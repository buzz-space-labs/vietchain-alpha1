package tests

import (
	"time"

	"github.com/anam-nw/anam/app"
	tmtypes "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
)

type KeeperTestHelper struct {
	suite.Suite

	App *app.App
	Ctx sdk.Context
}

func (s *KeeperTestHelper) Setup() {
	s.App = Setup(false)
	s.Ctx = s.App.BaseApp.NewContext(false, tmtypes.Header{Height: 1, ChainID: "anam-1", Time: time.Now().UTC()})
}
