// Copyright 2025 Buzz Space Labs
//
// Adapted from the Aura project © 2023 Aura Contributors
// Licensed under the Apache License, Version 2.0
// See LICENSE file in the project root for full license text.

package utils

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strings"
)

const (
	anamExponent = 6
	BaseCoinUnit = "uanam"
)

var (
	// DevnetChainID defines the anam chain ID for devnet
	DevnetChainID = "anam-testnet"

	// SerenityChainID defines the anam chain ID for serenity testnet
	SerenityChainID = "serenity-testnet"

	// EuphoriaChainID defines the anam chain ID for euphoria testnet
	EuphoriaChainID = "euphoria"
)

// IsDevnet returns true if the chain-id has the anam devnet chain prefix.
func IsDevnet(chainID string) bool {
	return strings.HasPrefix(chainID, DevnetChainID)
}

// IsSerenity returns true if the chain-id has the anam serenity network chain prefix.
func IsSerenity(chainID string) bool {
	return strings.HasPrefix(chainID, SerenityChainID)
}

// IsEuphoria returns true if the chain-id has the anam euphoria network chain prefix.
func IsEuphoria(chainID string) bool {
	return strings.HasPrefix(chainID, EuphoriaChainID)
}

// RegisterDenoms registers token denoms.
func RegisterDenoms() {
	err := sdk.RegisterDenom(BaseCoinUnit, sdk.NewDecWithPrec(1, anamExponent))
	if err != nil {
		panic(err)
	}
}
