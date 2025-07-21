// Copyright 2025 Buzz Space Labs
//
// Adapted from the Aura project © 2023 Aura Contributors
// Licensed under the Apache License, Version 2.0
// See LICENSE file in the project root for full license text.

package keeper

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/cometbft/cometbft/libs/log"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"

	"github.com/anam-nw/anam/x/anam/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramSpace paramtypes.Subspace
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	paramSpace paramtypes.Subspace,

) Keeper {
	// set KeyTable if it has not already been set
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}
	return Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramSpace: paramSpace,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// GetParams returns the total set of anam parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the total set of anam parameters.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

// GetMaxSupply return max supply of anam coin
func (k Keeper) GetMaxSupply(ctx sdk.Context) string {
	params := k.GetParams(ctx)
	return params.MaxSupply
}

// GetExcludeCirculatingAddr return list exclude address do not calculator for circulating
func (k Keeper) GetExcludeCirculatingAddr(ctx sdk.Context) []sdk.AccAddress {
	params := k.GetParams(ctx)
	excludeAddr := make([]sdk.AccAddress, 0, len(params.ExcludeCirculatingAddr))
	for _, addrBech32 := range params.ExcludeCirculatingAddr {
		addr, err := sdk.AccAddressFromBech32(addrBech32)
		if err != nil {
			panic(err)
		}
		excludeAddr = append(excludeAddr, addr)
	}

	return excludeAddr
}
