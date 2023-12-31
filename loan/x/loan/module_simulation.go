package loan

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"loan/testutil/sample"
	loansimulation "loan/x/loan/simulation"
	"loan/x/loan/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = loansimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgRequestLoan = "op_weight_msg_request_loan"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRequestLoan int = 100

	opWeightMsgApproveLoan = "op_weight_msg_approve_loan"
	// TODO: Determine the simulation weight value
	defaultWeightMsgApproveLoan int = 100

	opWeightMsgRepayLoan = "op_weight_msg_repay_loan"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRepayLoan int = 100

	opWeightMsgLiquidateLoan = "op_weight_msg_liquidate_loan"
	// TODO: Determine the simulation weight value
	defaultWeightMsgLiquidateLoan int = 100

	opWeightMsgCancelLoan = "op_weight_msg_cancel_loan"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCancelLoan int = 100

	opWeightMsgTokenMint = "op_weight_msg_token_mint"
	// TODO: Determine the simulation weight value
	defaultWeightMsgTokenMint int = 100

	opWeightMsgBurnToken = "op_weight_msg_burn_token"
	// TODO: Determine the simulation weight value
	defaultWeightMsgBurnToken int = 100

	opWeightMsgRedeem = "op_weight_msg_redeem"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRedeem int = 100

	opWeightMsgStake = "op_weight_msg_stake"
	// TODO: Determine the simulation weight value
	defaultWeightMsgStake int = 100

	opWeightMsgWithdrawStake = "op_weight_msg_withdraw_stake"
	// TODO: Determine the simulation weight value
	defaultWeightMsgWithdrawStake int = 100

	opWeightMsgAddCollateral = "op_weight_msg_add_collateral"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAddCollateral int = 100

	opWeightMsgWithdrawPartial = "op_weight_msg_withdraw_partial"
	// TODO: Determine the simulation weight value
	defaultWeightMsgWithdrawPartial int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	loanGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&loanGenesis)
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

	var weightMsgRequestLoan int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRequestLoan, &weightMsgRequestLoan, nil,
		func(_ *rand.Rand) {
			weightMsgRequestLoan = defaultWeightMsgRequestLoan
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRequestLoan,
		loansimulation.SimulateMsgRequestLoan(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgApproveLoan int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgApproveLoan, &weightMsgApproveLoan, nil,
		func(_ *rand.Rand) {
			weightMsgApproveLoan = defaultWeightMsgApproveLoan
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgApproveLoan,
		loansimulation.SimulateMsgApproveLoan(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRepayLoan int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRepayLoan, &weightMsgRepayLoan, nil,
		func(_ *rand.Rand) {
			weightMsgRepayLoan = defaultWeightMsgRepayLoan
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRepayLoan,
		loansimulation.SimulateMsgRepayLoan(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgLiquidateLoan int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgLiquidateLoan, &weightMsgLiquidateLoan, nil,
		func(_ *rand.Rand) {
			weightMsgLiquidateLoan = defaultWeightMsgLiquidateLoan
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgLiquidateLoan,
		loansimulation.SimulateMsgLiquidateLoan(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCancelLoan int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCancelLoan, &weightMsgCancelLoan, nil,
		func(_ *rand.Rand) {
			weightMsgCancelLoan = defaultWeightMsgCancelLoan
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCancelLoan,
		loansimulation.SimulateMsgCancelLoan(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgTokenMint int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgTokenMint, &weightMsgTokenMint, nil,
		func(_ *rand.Rand) {
			weightMsgTokenMint = defaultWeightMsgTokenMint
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgTokenMint,
		loansimulation.SimulateMsgTokenMint(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgBurnToken int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgBurnToken, &weightMsgBurnToken, nil,
		func(_ *rand.Rand) {
			weightMsgBurnToken = defaultWeightMsgBurnToken
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgBurnToken,
		loansimulation.SimulateMsgBurnToken(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRedeem int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRedeem, &weightMsgRedeem, nil,
		func(_ *rand.Rand) {
			weightMsgRedeem = defaultWeightMsgRedeem
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRedeem,
		loansimulation.SimulateMsgRedeem(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgStake int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgStake, &weightMsgStake, nil,
		func(_ *rand.Rand) {
			weightMsgStake = defaultWeightMsgStake
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgStake,
		loansimulation.SimulateMsgStake(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgWithdrawStake int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgWithdrawStake, &weightMsgWithdrawStake, nil,
		func(_ *rand.Rand) {
			weightMsgWithdrawStake = defaultWeightMsgWithdrawStake
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgWithdrawStake,
		loansimulation.SimulateMsgWithdrawStake(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAddCollateral int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgAddCollateral, &weightMsgAddCollateral, nil,
		func(_ *rand.Rand) {
			weightMsgAddCollateral = defaultWeightMsgAddCollateral
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAddCollateral,
		loansimulation.SimulateMsgAddCollateral(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgWithdrawPartial int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgWithdrawPartial, &weightMsgWithdrawPartial, nil,
		func(_ *rand.Rand) {
			weightMsgWithdrawPartial = defaultWeightMsgWithdrawPartial
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgWithdrawPartial,
		loansimulation.SimulateMsgWithdrawPartial(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgRequestLoan,
			defaultWeightMsgRequestLoan,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				loansimulation.SimulateMsgRequestLoan(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgApproveLoan,
			defaultWeightMsgApproveLoan,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				loansimulation.SimulateMsgApproveLoan(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgRepayLoan,
			defaultWeightMsgRepayLoan,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				loansimulation.SimulateMsgRepayLoan(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgLiquidateLoan,
			defaultWeightMsgLiquidateLoan,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				loansimulation.SimulateMsgLiquidateLoan(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCancelLoan,
			defaultWeightMsgCancelLoan,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				loansimulation.SimulateMsgCancelLoan(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgTokenMint,
			defaultWeightMsgTokenMint,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				loansimulation.SimulateMsgTokenMint(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgBurnToken,
			defaultWeightMsgBurnToken,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				loansimulation.SimulateMsgBurnToken(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgRedeem,
			defaultWeightMsgRedeem,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				loansimulation.SimulateMsgRedeem(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgStake,
			defaultWeightMsgStake,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				loansimulation.SimulateMsgStake(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgWithdrawStake,
			defaultWeightMsgWithdrawStake,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				loansimulation.SimulateMsgWithdrawStake(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgAddCollateral,
			defaultWeightMsgAddCollateral,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				loansimulation.SimulateMsgAddCollateral(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgWithdrawPartial,
			defaultWeightMsgWithdrawPartial,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				loansimulation.SimulateMsgWithdrawPartial(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
