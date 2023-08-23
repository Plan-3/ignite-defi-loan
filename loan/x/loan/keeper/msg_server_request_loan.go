package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	sdkmath "cosmossdk.io/math"
	"loan/x/loan/types"
)

func (k msgServer) RequestLoan(goCtx context.Context, msg *types.MsgRequestLoan) (*types.MsgRequestLoanResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// first create loan
	var loan = types.Loan{
		Amount:     msg.Amount,
		Fee:        msg.Fee,
		Collateral: msg.Collateral,
		Deadline:   msg.Deadline,
		State:      "requested",
		Borrower:   msg.Creator,
	}

	// get borrower account
	borrower, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}

	// parse collateral and amount string to sdk.Coin
	collateral, err := sdk.ParseCoinsNormalized(loan.Collateral)
	if err != nil {
		panic(err)
	}

	amount, err := sdk.ParseCoinsNormalized(loan.Amount)
	if err != nil {
		panic(err)
	}

	// set up pointer to TokenPrice collateral price
	collateralPrice := &types.TokenPrice{}


	// switch on denom string to set parsed coin t a type TokenPrice{sdk.Coin, int}
	switch collateral[0].Denom {
		case "ctz":
			collateralPrice.Denom = collateral[0];
			collateralPrice.Price = 1800;
			break;
		case "cqt":
			collateralPrice.Denom = collateral[0];
			collateralPrice.Price = 100;
			break;
		default:
			break;
	}

	// no switch needed here all loan amounts are paid out in zusd
	amountPrice := &types.TokenPrice{amount[0], 1};

	// need to use sdkmath.Float64 since numbers are sdk.Int takes
	// Float64 is a method on LegacyDec type 
	// can use sdkmath ToLegacyDec
	// turn prices into floats for risk check
	collateralFloat, _ := sdkmath.LegacyDec(collateral[0].Amount).Float64()
	amountFloat, _ := sdkmath.LegacyDec(amount[0].Amount).Float64()

	collateralPriceFloat := collateralFloat * float64(collateralPrice.Price)
	amountPriceFloat := amountFloat * float64(amountPrice.Price)

	// calculate risk using ratio collateral price / amount price > .909090909
	risk := collateralPriceFloat / amountPriceFloat
	
	if risk < .909090909 {
	// send collateral from borrower to loan module account
	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, borrower, types.ModuleName, collateral)
	if sdkError != nil {
		return nil, sdkError
	}

	// append loan to store
	k.AppendLoan(ctx, loan)
	return &types.MsgRequestLoanResponse{}, nil

	} else {
		return nil, sdkerrors.Wrap(types.ErrInvalidRequest, "Loan risk too high")
	}

}
