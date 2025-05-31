package keeper

import (
	"github.com/annam-nw/annam/x/smartaccount/types"
)

var _ types.QueryServer = Keeper{}
