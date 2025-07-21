// Copyright 2025 Buzz Space Labs
//
// Adapted from the Aura project © 2023 Aura Contributors
// Licensed under the Apache License, Version 2.0
// See LICENSE file in the project root for full license text.

package smartaccount

import (
	"github.com/anam-nw/anam/x/smartaccount/keeper"
	"github.com/anam-nw/anam/x/smartaccount/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}

	k.SetNextAccountID(ctx, genState.GetSmartAccountId())
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	genesis.SmartAccountId = k.GetNextAccountID(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
