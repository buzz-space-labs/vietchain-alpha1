// Copyright 2025 Buzz Space Labs
//
// Adapted from the Aura project © 2023 Aura Contributors
// Licensed under the Apache License, Version 2.0
// See LICENSE file in the project root for full license text.

package keeper

import (
	"github.com/anam-nw/anam/x/anam/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}
