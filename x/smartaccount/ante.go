// Copyright 2025 Buzz Space Labs
//
// Adapted from the Aura project © 2023 Aura Contributors
// Licensed under the Apache License, Version 2.0
// See LICENSE file in the project root for full license text.

package smartaccount

import (
	"encoding/json"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	sakeeper "github.com/anam-nw/anam/x/smartaccount/keeper"
	"github.com/anam-nw/anam/x/smartaccount/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

func GetSmartAccountTxSigner(ctx sdk.Context, sigTx authsigning.SigVerifiableTx, saKeeper sakeeper.Keeper) (*types.SmartAccount, error) {
	sigs, err := sigTx.GetSignaturesV2()
	if err != nil {
		return nil, err
	}

	signerAddrs := sigTx.GetSigners()

	// signer of smartaccount tx must stand alone
	if len(signerAddrs) != 1 || len(sigs) != 1 {
		return nil, nil
	}

	saAcc, err := saKeeper.GetSmartAccountByAddress(ctx, signerAddrs[0])
	if err != nil {
		return nil, err
	}

	return saAcc, nil
}

func GetValidActivateAccountMessage(sigTx authsigning.SigVerifiableTx) (*types.MsgActivateAccount, error) {
	msgs := sigTx.GetMsgs()

	if len(msgs) != 1 {
		// smart account activation message must stand alone
		for _, msg := range msgs {
			if _, ok := msg.(*types.MsgActivateAccount); ok {
				return nil, errorsmod.Wrap(types.ErrInvalidTx, "smart account activation message must stand alone")
			}
		}

		return nil, nil
	}

	activateMsg, ok := msgs[0].(*types.MsgActivateAccount)
	if !ok {
		return nil, nil
	}

	sigs, err := sigTx.GetSignaturesV2()
	if err != nil {
		return nil, err
	}

	signer := sigTx.GetSigners()

	// do not allow multi signer and signature
	if len(signer) != 1 || len(sigs) != 1 {
		return nil, errorsmod.Wrap(types.ErrInvalidTx, "smart-account activation tx does not allow multiple signers")
	}

	return activateMsg, nil
}

// ------------------------- SmartAccount Decorator ------------------------- \\

type SmartAccountDecorator struct {
	SaKeeper sakeeper.Keeper
}

func NewSmartAccountDecorator(saKeeper sakeeper.Keeper) *SmartAccountDecorator {
	return &SmartAccountDecorator{
		SaKeeper: saKeeper,
	}
}

// AnteHandle is used for performing basic validity checks on a transaction such that it can be thrown out of the mempool.
func (d *SmartAccountDecorator) AnteHandle(
	ctx sdk.Context,
	tx sdk.Tx,
	simulate bool,
	next sdk.AnteHandler,
) (newCtx sdk.Context, err error) {

	sigTx, ok := tx.(authsigning.SigVerifiableTx)
	if !ok {
		return ctx, errorsmod.Wrap(types.ErrInvalidTx, "not a SigVerifiableTx")
	}

	feeTx, ok := tx.(sdk.FeeTx)
	if !ok {
		return ctx, errorsmod.Wrap(types.ErrInvalidTx, "not a FeeTx")
	}

	activateMsg, err := GetValidActivateAccountMessage(sigTx)
	if err != nil {
		return ctx, err
	}

	if activateMsg == nil {
		err = handleSmartAccountTx(ctx, d.SaKeeper, sigTx, feeTx, simulate)
		if err != nil {
			return ctx, err
		}
	} else {
		err = handleSmartAccountActivate(ctx, d.SaKeeper, activateMsg, simulate)
		if err != nil {
			return ctx, err
		}
	}

	return next(ctx, tx, simulate)
}

func handleSmartAccountTx(
	ctx sdk.Context,
	saKeeper sakeeper.Keeper,
	sigTx authsigning.SigVerifiableTx,
	feeTx sdk.FeeTx,
	simulate bool,
) error {

	signerAcc, err := GetSmartAccountTxSigner(ctx, sigTx, saKeeper)
	if err != nil {
		return err
	}

	// if is not smartaccount tx type
	if signerAcc == nil {
		// do some thing
		return nil
	}

	// save the account address to the module store. we will need it in the
	// posthandler
	//
	// TODO: a question is that instead of writing to store, can we just put this
	// in memory instead. in practice however, the address is deleted in the post
	// handler, so it's never actually written to disk, meaning the difference in
	// gas consumption should be really small. still worth investigating tho.
	saKeeper.SetSignerAddress(ctx, signerAcc.GetAddress())

	// not support smartaccount tx simulation yet
	if simulate {
		return errorsmod.Wrap(types.ErrNotSupported, "Simulation of SmartAccount txs isn't supported yet")
	}

	msgs := sigTx.GetMsgs()

	err = saKeeper.CheckAllowedMsgs(ctx, msgs)
	if err != nil {
		return err
	}

	msgsData, err := types.ParseMessagesString(msgs)
	if err != nil {
		return err
	}

	txInfor := types.TxInfor{
		Gas:     feeTx.GetGas(),
		Fee:     feeTx.GetFee(),
		Payer:   feeTx.FeePayer().String(),
		Granter: feeTx.FeeGranter().String(),
	}

	preExecuteMessage, err := json.Marshal(&types.AccountMsg{
		PreExecuteTx: &types.PreExecuteTx{
			Msgs:    msgsData,
			TxInfor: txInfor,
		},
	})
	if err != nil {
		return err
	}

	params := saKeeper.GetParams(ctx)

	// execute SA contract for pre-execute transaction with limit gas
	err = sudoWithGasLimit(ctx, saKeeper.ContractKeeper, signerAcc.GetAddress(), preExecuteMessage, params.MaxGasExecute)
	if err != nil {
		return err
	}

	return nil
}

func handleSmartAccountActivate(
	ctx sdk.Context,
	saKeeper sakeeper.Keeper,
	activateMsg *types.MsgActivateAccount,
	simulate bool,
) error {
	// in ReCheckTx mode, below check may not be necessary

	// get signer of smart account activation message
	signer := activateMsg.GetSigners()[0]

	// decode string to pubkey
	pubKey, err := types.PubKeyDecode(activateMsg.PubKey)
	if err != nil {
		return err
	}

	// generate predictable address using Instantiate2's PredicableAddressGenerator
	predicAddr, err := types.Instantiate2Address(
		ctx,
		saKeeper.WasmKeeper,
		activateMsg.CodeID,
		activateMsg.InitMsg,
		activateMsg.Salt,
		pubKey,
	)
	if err != nil {
		return err
	}

	// the signer of the activation message must be the address generated by Instantiate2's PredicableAddressGenerator
	if !signer.Equals(predicAddr) {
		return errorsmod.Wrap(types.ErrInvalidAddress, "not the same as predicted")
	}

	// if in delivery mode, remove temporary pubkey from account
	if !ctx.IsCheckTx() && !ctx.IsReCheckTx() && !simulate {
		// get smart contract account by address
		sAccount := saKeeper.AccountKeeper.GetAccount(ctx, signer)
		_, isBase := sAccount.(*authtypes.BaseAccount)
		_, isSa := sAccount.(*types.SmartAccount)
		if !isBase && !isSa {
			return errorsmod.Wrap(types.ErrAccountNotFoundForAddress, signer.String())
		}

		// remove temporary pubkey for account
		err = saKeeper.UpdateAccountPubKey(ctx, sAccount, nil)
		if err != nil {
			return err
		}
	}

	return nil
}

// Call a contract's sudo with a gas limit
// referenced from Osmosis' protorev posthandler:
// https://github.com/osmosis-labs/osmosis/blob/98025f185ab2ee1b060511ed22679112abcc08fa/x/protorev/keeper/posthandler.go#L42-L43
func sudoWithGasLimit(
	ctx sdk.Context, contractKeeper *wasmkeeper.PermissionedKeeper,
	contractAddr sdk.AccAddress, msg []byte, maxGas sdk.Gas,
) error {
	cacheCtx, write := ctx.CacheContext()
	cacheCtx = cacheCtx.WithGasMeter(sdk.NewGasMeter(maxGas))

	if _, err := contractKeeper.Sudo(
		cacheCtx,
		contractAddr, // contract address
		msg,
	); err != nil {
		return err
	}

	write()
	ctx.EventManager().EmitEvents(cacheCtx.EventManager().Events())

	return nil
}

// ------------------------- SetPubKey Decorator ------------------------- \\

type SetPubKeyDecorator struct {
	saKeeper sakeeper.Keeper
}

func NewSetPubKeyDecorator(saKeeper sakeeper.Keeper) *SetPubKeyDecorator {
	return &SetPubKeyDecorator{
		saKeeper: saKeeper,
	}
}

func (d *SetPubKeyDecorator) AnteHandle(
	ctx sdk.Context,
	tx sdk.Tx,
	simulate bool,
	next sdk.AnteHandler,
) (newCtx sdk.Context, err error) {

	sigTx, ok := tx.(authsigning.SigVerifiableTx)
	if !ok {
		return ctx, errorsmod.Wrap(types.ErrInvalidTx, "not a SigVerifiableTx")
	}

	activateMsg, err := GetValidActivateAccountMessage(sigTx)
	if err != nil {
		return ctx, err
	}

	// if is smart account activation message
	if activateMsg != nil {
		// get message signer
		signer := activateMsg.GetSigners()[0]

		// get smart contract account by address, account must be inactivate smart account
		sAccount, err := d.saKeeper.IsInactiveAccount(ctx, signer)
		if err != nil {
			return ctx, err
		}

		// decode any to pubkey
		pubKey, err := types.PubKeyDecode(activateMsg.PubKey)
		if err != nil {
			return ctx, err
		}

		// set temporary pubkey for account
		// need this for the next ante signature checks
		err = d.saKeeper.UpdateAccountPubKey(ctx, sAccount, pubKey)
		if err != nil {
			return ctx, err
		}

		return next(ctx, tx, simulate)
	}

	signerAcc, err := GetSmartAccountTxSigner(ctx, sigTx, d.saKeeper)
	if err != nil {
		return ctx, err
	}

	// if is smart account tx skip authante NewSetPubKeyDecorator
	// need this to avoid pubkey and address equal check of above decorator
	if signerAcc != nil {
		// if this is smart account tx, check if pubkey is set
		if signerAcc.GetPubKey() == nil {
			return ctx, errorsmod.Wrap(types.ErrNilPubkey, signerAcc.String())
		}

		// Also emit the following events, so that txs can be indexed by these
		var events sdk.Events
		events = append(events, sdk.NewEvent(types.EventTypeSmartAccountTx,
			sdk.NewAttribute(sdk.AttributeKeyAccountSequence, fmt.Sprintf("%s/%d", signerAcc.GetAddress(), signerAcc.GetSequence())),
		))

		// maybe need add more event here

		ctx.EventManager().EmitEvents(events)

		return next(ctx, tx, simulate)
	}

	// default authant SetPubKeyDecorator
	svd := authante.NewSetPubKeyDecorator(d.saKeeper.AccountKeeper)
	return svd.AnteHandle(ctx, tx, simulate, next)
}

// AfterTx
type AfterTxDecorator struct {
	saKeeper sakeeper.Keeper
}

func NewAfterTxDecorator(saKeeper sakeeper.Keeper) AfterTxDecorator {
	return AfterTxDecorator{
		saKeeper: saKeeper,
	}
}

// referenced from Larry0x' abstractaccount posthandler:
// https://github.com/larry0x/abstract-account/blob/b3c6432e593d450e7c58dae94cdf2a95930f8159/x/abstractaccount/ante.go#L152-L185
func (d AfterTxDecorator) PostHandle(ctx sdk.Context, tx sdk.Tx, simulate, success bool, next sdk.PostHandler) (newCtx sdk.Context, err error) {
	// load the signer address, which we determined during the AnteHandler
	//
	// if not found, it means this tx is simply not an AA tx. we skip
	signerAddr := d.saKeeper.GetSignerAddress(ctx)
	if signerAddr == nil {
		return next(ctx, tx, simulate, success)
	}

	d.saKeeper.DeleteSignerAddress(ctx)

	afterExecuteMessage, err := json.Marshal(&types.AccountMsg{
		AfterExecuteTx: &types.AfterExecuteTx{},
	})
	if err != nil {
		return ctx, err
	}

	params := d.saKeeper.GetParams(ctx)

	// execute SA contract for after-execute transaction with limit gas
	err = sudoWithGasLimit(ctx, d.saKeeper.ContractKeeper, signerAddr, afterExecuteMessage, params.MaxGasExecute)
	if err != nil {
		return ctx, err
	}

	return next(ctx, tx, simulate, success)
}
