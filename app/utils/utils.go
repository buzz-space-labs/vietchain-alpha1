package utils

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strings"
)

const (
	annamExponent = 6
	BaseCoinUnit = "uannam"
)

var (
	// DevnetChainID defines the annam chain ID for devnet
	DevnetChainID = "annam-testnet"

	// SerenityChainID defines the annam chain ID for serenity testnet
	SerenityChainID = "serenity-testnet"

	// EuphoriaChainID defines the annam chain ID for euphoria testnet
	EuphoriaChainID = "euphoria"
)

// IsDevnet returns true if the chain-id has the annam devnet chain prefix.
func IsDevnet(chainID string) bool {
	return strings.HasPrefix(chainID, DevnetChainID)
}

// IsSerenity returns true if the chain-id has the annam serenity network chain prefix.
func IsSerenity(chainID string) bool {
	return strings.HasPrefix(chainID, SerenityChainID)
}

// IsEuphoria returns true if the chain-id has the annam euphoria network chain prefix.
func IsEuphoria(chainID string) bool {
	return strings.HasPrefix(chainID, EuphoriaChainID)
}

// RegisterDenoms registers token denoms.
func RegisterDenoms() {
	err := sdk.RegisterDenom(BaseCoinUnit, sdk.NewDecWithPrec(1, annamExponent))
	if err != nil {
		panic(err)
	}
}
