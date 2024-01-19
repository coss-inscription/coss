package vault

import (
	"math/rand"

	"coss/testutil/sample"
	vaultsimulation "coss/x/vault/simulation"
	"coss/x/vault/types"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = vaultsimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgCreateTokenAdmin = "op_weight_msg_token_admin"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateTokenAdmin int = 100

	opWeightMsgUpdateTokenAdmin = "op_weight_msg_token_admin"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateTokenAdmin int = 100

	opWeightMsgDeleteTokenAdmin = "op_weight_msg_token_admin"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteTokenAdmin int = 100

	opWeightMsgCreateToken = "op_weight_msg_token"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateToken int = 100

	opWeightMsgUpdateToken = "op_weight_msg_token"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateToken int = 100

	opWeightMsgDeleteToken = "op_weight_msg_token"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteToken int = 100

	opWeightMsgAuditToken = "op_weight_msg_audit_token"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAuditToken int = 100

	opWeightMsgConvertTokenToIns = "op_weight_msg_convert_token_to_ins"
	// TODO: Determine the simulation weight value
	defaultWeightMsgConvertTokenToIns int = 100

	opWeightMsgConvertInsToToken = "op_weight_msg_convert_ins_to_token"
	// TODO: Determine the simulation weight value
	defaultWeightMsgConvertInsToToken int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	vaultGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		TokenList: []types.Token{
			{
				Owner: sample.AccAddress(),
				Denom: "0",
			},
			{
				Owner: sample.AccAddress(),
				Denom: "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&vaultGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateTokenAdmin int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateTokenAdmin, &weightMsgCreateTokenAdmin, nil,
		func(_ *rand.Rand) {
			weightMsgCreateTokenAdmin = defaultWeightMsgCreateTokenAdmin
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateTokenAdmin,
		vaultsimulation.SimulateMsgCreateTokenAdmin(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateTokenAdmin int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateTokenAdmin, &weightMsgUpdateTokenAdmin, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateTokenAdmin = defaultWeightMsgUpdateTokenAdmin
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateTokenAdmin,
		vaultsimulation.SimulateMsgUpdateTokenAdmin(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateToken int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateToken, &weightMsgCreateToken, nil,
		func(_ *rand.Rand) {
			weightMsgCreateToken = defaultWeightMsgCreateToken
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateToken,
		vaultsimulation.SimulateMsgCreateToken(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateToken int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateToken, &weightMsgUpdateToken, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateToken = defaultWeightMsgUpdateToken
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateToken,
		vaultsimulation.SimulateMsgUpdateToken(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAuditToken int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAuditToken, &weightMsgAuditToken, nil,
		func(_ *rand.Rand) {
			weightMsgAuditToken = defaultWeightMsgAuditToken
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAuditToken,
		vaultsimulation.SimulateMsgAuditToken(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgConvertTokenToIns int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgConvertTokenToIns, &weightMsgConvertTokenToIns, nil,
		func(_ *rand.Rand) {
			weightMsgConvertTokenToIns = defaultWeightMsgConvertTokenToIns
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgConvertTokenToIns,
		vaultsimulation.SimulateMsgConvertTokenToIns(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgConvertInsToToken int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgConvertInsToToken, &weightMsgConvertInsToToken, nil,
		func(_ *rand.Rand) {
			weightMsgConvertInsToToken = defaultWeightMsgConvertInsToToken
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgConvertInsToToken,
		vaultsimulation.SimulateMsgConvertInsToToken(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateTokenAdmin,
			defaultWeightMsgCreateTokenAdmin,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				vaultsimulation.SimulateMsgCreateTokenAdmin(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateTokenAdmin,
			defaultWeightMsgUpdateTokenAdmin,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				vaultsimulation.SimulateMsgUpdateTokenAdmin(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateToken,
			defaultWeightMsgCreateToken,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				vaultsimulation.SimulateMsgCreateToken(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateToken,
			defaultWeightMsgUpdateToken,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				vaultsimulation.SimulateMsgUpdateToken(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgAuditToken,
			defaultWeightMsgAuditToken,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				vaultsimulation.SimulateMsgAuditToken(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgConvertTokenToIns,
			defaultWeightMsgConvertTokenToIns,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				vaultsimulation.SimulateMsgConvertTokenToIns(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgConvertInsToToken,
			defaultWeightMsgConvertInsToToken,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				vaultsimulation.SimulateMsgConvertInsToToken(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
