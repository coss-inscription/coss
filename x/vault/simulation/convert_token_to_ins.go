package simulation

import (
	"math/rand"

	"coss/x/vault/keeper"
	"coss/x/vault/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgConvertTokenToIns(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgConvertTokenToIns{
			Sender: simAccount.Address.String(),
		}

		// TODO: Handling the ConvertTokenToIns simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "ConvertTokenToIns simulation not implemented"), nil, nil
	}
}
