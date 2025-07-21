// Copyright 2025 Buzz Space Labs
//
// Adapted from the Aura project © 2023 Aura Contributors
// Licensed under the Apache License, Version 2.0
// See LICENSE file in the project root for full license text.

package feegrant

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
)

// RegisterInterfaces registers the interfaces types with the interface registry
func RegisterInterfaces(registry types.InterfaceRegistry) {

	registry.RegisterImplementations((*sdk.Msg)(nil),
		&feegrant.MsgGrantAllowance{},
		&feegrant.MsgRevokeAllowance{},
	)

	registry.RegisterInterface(
		"cosmos.feegrant.v1beta1.FeeAllowanceI",
		(*feegrant.FeeAllowanceI)(nil),
		&feegrant.BasicAllowance{},
		&feegrant.PeriodicAllowance{},
		&feegrant.AllowedMsgAllowance{},
		&AllowedContractAllowance{},
	)

	feegrant.RegisterInterfaces(registry)
}
