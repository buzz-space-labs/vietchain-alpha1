// Copyright 2025 Buzz Space Labs
//
// Adapted from the Aura project © 2023 Aura Contributors
// Licensed under the Apache License, Version 2.0
// See LICENSE file in the project root for full license text.

package keeper

import (
	"context"

	"github.com/anam-nw/anam/x/smartaccount/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GenerateAccount(goCtx context.Context, req *types.QueryGenerateAccountRequest) (*types.QueryGenerateAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	pubKey, err := types.PubKeyDecode(req.PubKey)
	if err != nil {
		return nil, err
	}

	contractAddress, err := types.Instantiate2Address(ctx, k.WasmKeeper, req.CodeID, req.InitMsg, req.Salt, pubKey)
	if err != nil {
		return nil, err
	}

	return &types.QueryGenerateAccountResponse{
		Address: contractAddress.String(),
	}, nil
}
