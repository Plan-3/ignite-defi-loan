package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"loan/x/loan/types"
)

func (k msgServer) AddCollateral(goCtx context.Context, msg *types.MsgAddCollateral) (*types.MsgAddCollateralResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	// retrieve loan to approve k is the msgServer object getLoan is a method of a keeper
	loan, found := k.GetLoan(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "loan %d doesn't exist", msg.Id)
	}

	if loan.State != "approved" {
		return nil, sdkerrors.Wrapf(types.ErrWrongLoanState, "loan %d is not in requested state", msg.Id)
	}

	// get borrower account
	borrower, _ := sdk.AccAddressFromBech32(loan.Borrower)
	if loan.Borrower != msg.Creator {
		return nil, sdkerrors.Wrap(types.ErrNotBorrower, "You are not borrower")
	}
	// get loan amount
	amount, err := sdk.ParseCoinsNormalized(msg.Amount)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidRequest, "Can't parse loan amount")
	}
	// get collateral
	collateral, err := sdk.ParseCoinsNormalized(loan.Collateral)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidRequest, "Can't parse collateral")
	}

	addition := collateral[0].Amount.Add(amount[0].Amount)

	// updated collateral
	newCollateral := sdk.NewCoin(collateral[0].Denom, addition)

	getCwei := amount[0].Amount.MulRaw(int64(1000000000))

	cCoin := sdk.NewCoin(amount[0].Denom, getCwei)

	// send coins to the module account that holds collateral
	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, borrower, types.Nbtp, sdk.NewCoins(cCoin))
	if sdkError != nil {
		return nil, sdkError
	}

	// update values
	loan.Collateral = newCollateral.String()
	// store updated loan values
	k.SetLoan(ctx, loan)

	return &types.MsgAddCollateralResponse{}, nil
}
