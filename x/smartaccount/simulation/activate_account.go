package simulation

import (
	"math/rand"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	"github.com/anam-nw/anam/x/smartaccount/keeper"
	"github.com/anam-nw/anam/x/smartaccount/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgActivateAccount(
	wk *wasmkeeper.PermissionedKeeper,
	ak types.AccountKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgActivateAccount{
			AccountAddress: simAccount.Address.String(),
		}

		// TODO: Handling the ActivateAccount simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "ActivateAccount simulation not implemented"), nil, nil
	}
}
