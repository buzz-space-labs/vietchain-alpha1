// Copyright 2025 Buzz Space Labs
//
// Adapted from the Aura project © 2023 Aura Contributors
// Licensed under the Apache License, Version 2.0
// See LICENSE file in the project root for full license text.

package keeper

import (
	"cosmossdk.io/math"
	custommint "github.com/anam-nw/anam/x/mint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
)

// Keeper of the mint store
type Keeper struct {
	mintkeeper.Keeper

	bankKeeper    custommint.BankKeeper
	stakingKeeper custommint.StakingKeeper
	anamKeeper    custommint.AnamKeeper
}

// NewKeeper creates a new mint Keeper instance
func NewKeeper(
	cdc codec.BinaryCodec, key storetypes.StoreKey,
	sk custommint.StakingKeeper, ak custommint.AccountKeeper, bk custommint.BankKeeper,
	anamKeeper custommint.AnamKeeper, feeCollectorName string, authority string,
) Keeper {
	return Keeper{
		Keeper:        mintkeeper.NewKeeper(cdc, key, sk, ak, bk, feeCollectorName, authority),
		bankKeeper:    bk,
		stakingKeeper: sk,
		anamKeeper:    anamKeeper,
	}
}

func (k Keeper) GetSupply(ctx sdk.Context, denom string) math.Int {
	return k.bankKeeper.GetSupply(ctx, denom).Amount
}

func (k Keeper) GetMaxSupply(ctx sdk.Context) string {
	return k.anamKeeper.GetMaxSupply(ctx)
}

//func (k Keeper) GetExcludeCirculatingAddr(ctx sdk.Context) []sdk.AccAddress {
//	return k.anamKeeper.GetExcludeCirculatingAddr(ctx)
//}

//func (k Keeper) GetExcludeCirculatingAmount(ctx sdk.Context, denom string) sdk.Coin {
//	excludeAddrs := k.anamKeeper.GetExcludeCirculatingAddr(ctx)
//	excludeAmount := sdk.NewInt64Coin(denom, 0)
//	for _, addr := range excludeAddrs {
//		k.Logger(ctx).Info("GetExcludeCirculatingAmount", "addr", addr.String())
//		amount := k.bankKeeper.GetBalance(ctx, addr, denom)
//		k.Logger(ctx).Info("GetExcludeCirculatingAmount", "amount", amount.Amount)
//		k.Logger(ctx).Info("GetExcludeCirculatingAmount", "amountString", amount.String())
//		excludeAmount = excludeAmount.Add(amount)
//		k.Logger(ctx).Info("GetExcludeCirculatingAmount", "excludeAmount", excludeAmount.String())
//	}
//	return excludeAmount
//}

func (k Keeper) GetExcludeCirculatingAmount(ctx sdk.Context, denom string) sdk.Coin {
	return k.bankKeeper.GetExcludeCirculatingAmount(ctx, denom)
}

// CustomStakingTokenSupply implements an alias call to the underlying staking keeper's
// CustomStakingTokenSupply to be used in BeginBlocker.
func (k Keeper) CustomStakingTokenSupply(ctx sdk.Context, excludeAmount math.Int) math.Int {
	return k.stakingKeeper.StakingTokenSupply(ctx).Sub(excludeAmount)
}

// CustomBondedRatio implements an alias call to the underlying staking keeper's
// CustomBondedRatio to be used in BeginBlocker.
func (k Keeper) CustomBondedRatio(ctx sdk.Context, excludeAmount math.Int) math.LegacyDec {

	stakeSupply := k.CustomStakingTokenSupply(ctx, excludeAmount)
	if stakeSupply.IsPositive() {
		totalBonded := k.stakingKeeper.TotalBondedTokens(ctx)
		return math.LegacyNewDecFromInt(totalBonded).QuoInt(stakeSupply)
	}

	return sdk.ZeroDec()
}
