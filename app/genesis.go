// Copyright 2025 Buzz Space Labs
//
// Adapted from the Aura project © 2023 Aura Contributors
// Licensed under the Apache License, Version 2.0
// See LICENSE file in the project root for full license text.

package app

import (
	"encoding/json"

	//"github.com/CosmWasm/wasmd/x/wasm"
	//wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/codec"
)

// The genesis state of the blockchain is represented here as a map of raw json
// messages key'd by a identifier string.
// The identifier is used to determine which module genesis information belongs
// to so it may be appropriately routed during init chain.
// Within this application default genesis information is retrieved from
// the ModuleBasicManager which populates json from each BasicModule
// object provided to it during init.
type GenesisState map[string]json.RawMessage

// const (
// 	// DefaultMaxWasmCodeSize limit max bytes read to prevent gzip bombs

// 	DefaultMaxWasmCodeSize = 1000 * 1024 * 2

// 	DefaultVotingPeriod  = 1000000000 * 60 * 60 * 12 // nano seconds = 12h
// 	DefaultDepositPeriod = 1000000000 * 60 * 60 * 12 // nano seconds = 12h
// )

// NewDefaultGenesisState generates the default state for the application.
func NewDefaultGenesisState(cdc codec.JSONCodec) GenesisState {
	genesis := ModuleBasics.DefaultGenesis(cdc)
	// wasmGen := wasm.GenesisState{
	// 	Params: wasmtypes.Params{
	// 		CodeUploadAccess:             wasmtypes.AllowNobody,
	// 		InstantiateDefaultPermission: wasmtypes.AccessTypeEverybody,
	// 		MaxWasmCodeSize:              DefaultMaxWasmCodeSize,
	// 	},
	// }
	// genesis[wasm.ModuleName] = cdc.MustMarshalJSON(&wasmGen)
	return genesis
}
