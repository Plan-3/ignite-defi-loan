package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"loan/x/loan/types"
)

func (k msgServer) RepayLoan(goCtx context.Context, msg *types.MsgRepayLoan) (*types.MsgRepayLoanResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// get loan
	loan, found := k.GetLoan(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrKeyNotFound, "key %d not found", msg.Id)
	}
	// check if loan is in the correct state, approved is only state allowing repayment
	if loan.State != "approved" {
		return nil, sdkerrors.Wrapf(types.ErrWrongLoanState, "loan is not in the correct state for this action, loan is in %s state. Needs to be in approved", loan.State)
	}

	lender, _ := sdk.AccAddressFromBech32(loan.Lender)
	borrower, _ := sdk.AccAddressFromBech32(loan.Borrower)
	if msg.Creator != loan.Borrower {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, "Cannot repay: not the borrower")
	}

	// grab necessary coins from borrower and send to lender
	amount, _ := sdk.ParseCoinsNormalized(loan.Amount)
	fee, _ := sdk.ParseCoinsNormalized(loan.Fee)
	collateral, _ := sdk.ParseCoinsNormalized(loan.Collateral)

	// send coins out to and from appropriate accounts
	err := k.bankKeeper.SendCoins(ctx, borrower, lender, amount)
	if err != nil {
		return nil, err
	}
	err = k.bankKeeper.SendCoins(ctx, borrower, lender, fee)
	if err != nil {
		return nil, err
	}
	// collateral is sent back to the borrower from the module not the lender
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, borrower, collateral)
	if err != nil {
		return nil, err
	}
	loan.State = "repayed"
	k.SetLoan(ctx, loan)
	return &types.MsgRepayLoanResponse{}, nil
}
