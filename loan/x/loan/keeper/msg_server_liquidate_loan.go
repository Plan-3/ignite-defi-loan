package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"loan/x/loan/types"
)

// in future maybe add liquidator is not borrower also make a storage for accounts that get liquidated then create a check for bad actors in request loan
func (k msgServer) LiquidateLoan(goCtx context.Context, msg *types.MsgLiquidateLoan) (*types.MsgLiquidateLoanResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	loan, found := k.GetLoan(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrKeyNotFound, "key %d doesn't exist", msg.Id)
	}
	if loan.Lender != msg.Creator {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, "Cannot liquidate: not the lender")
	}
	if loan.State != "approved" {
		return nil, sdkerrors.Wrapf(types.ErrWrongLoanState, "%v", loan.State)
	}
	// should make it possible for anyone to liquidate by having msg.Creator == loan.Lender
	lender, _ := sdk.AccAddressFromBech32(loan.Lender) //loan.Lender
	collateral, _ := sdk.ParseCoinsNormalized(loan.Collateral)
	// convert deadline to int to compare to block height
	deadline, err := strconv.ParseInt(loan.Deadline, 10, 64)
	if err != nil {
		panic(err)
	}
	// check to make sure liquidation is happening after deadline and not before
	if ctx.BlockHeight() < deadline {
		return nil, sdkerrors.Wrap(types.ErrDeadline, "Cannot liquidate before deadline")
	}
	// send collateral to liquidator
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, lender, collateral)
	if err != nil {
		return nil, err
	}
	loan.State = "liquidated"
	k.SetLoan(ctx, loan)

	return &types.MsgLiquidateLoanResponse{}, nil
}
