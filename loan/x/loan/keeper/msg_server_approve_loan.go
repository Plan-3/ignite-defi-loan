package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"loan/x/loan/types"
)

func (k msgServer) ApproveLoan(goCtx context.Context, msg *types.MsgApproveLoan) (*types.MsgApproveLoanResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// retrieve loan to approve k is the msgServer object getLoan is a method of a keeper
	loan, found := k.GetLoan(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "loan %d doesn't exist", msg.Id)
	}

	// make sure loan is in requested state
	if loan.State != "requested" {
		return nil, sdkerrors.Wrapf(types.ErrWrongLoanState, "loan %d is not in requested state", msg.Id)
	}

	// set up lender account from account calling func
	lender, _ := sdk.AccAddressFromBech32(msg.Creator)
	// get borrower account
	borrower, _ := sdk.AccAddressFromBech32(loan.Borrower)
	// get loan amount
	amount, err := sdk.ParseCoinsNormalized(loan.Amount)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidRequest, "Can't parse loan amount")
	}
	// send coins to the module from the borrower not to the lender
	err = k.bankKeeper.SendCoins(ctx, lender, borrower, amount)
	if err != nil {
		return nil, err
	}
	loan.Lender = msg.Creator
	loan.State = "approved"
	// store updated loan values
	k.SetLoan(ctx, loan)

	return &types.MsgApproveLoanResponse{}, nil
}
