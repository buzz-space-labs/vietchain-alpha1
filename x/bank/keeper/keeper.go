package keeper

import (
	"github.com/annam-nw/annam/x/bank/types"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank/keeper"
)

type BaseKeeper struct {
	keeper.BaseKeeper

	annamKeeper types.AnnamKeeper
}

func NewBaseKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	ak types.AccountKeeper,
	blockedAddrs map[string]bool,
	authority string,
	annamKeeper types.AnnamKeeper,
) BaseKeeper {
	return BaseKeeper{
		BaseKeeper: keeper.NewBaseKeeper(cdc, storeKey, ak, blockedAddrs, authority),
		annamKeeper: annamKeeper,
	}
}

func (k BaseKeeper) GetExcludeCirculatingAmount(ctx sdk.Context, denom string) sdk.Coin {
	excludeAddrs := k.annamKeeper.GetExcludeCirculatingAddr(ctx)
	excludeAmount := sdk.NewInt64Coin(denom, 0)
	for _, addr := range excludeAddrs {
		amount := k.BaseKeeper.GetBalance(ctx, addr, denom)
		excludeAmount = excludeAmount.Add(amount)
	}

	return excludeAmount
}
