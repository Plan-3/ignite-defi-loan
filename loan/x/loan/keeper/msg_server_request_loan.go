package keeper

import (
	"context"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"loan/x/loan/types"
)

var (
	moduleAccountLoan = "cosmos1gu4m79yj8ch8em7c22vzt3qparg69ymm75qf6l"
	blackList         = make(map[string]bool)
)

func (k msgServer) RequestLoan(goCtx context.Context, msg *types.MsgRequestLoan) (*types.MsgRequestLoanResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	blackList["cosmos1gxrdcutv2plpdqcm8ldg4frafy7tms0qk9lcn6"] = true
	// first create loan
	var loan = types.Loan{
		Amount:     msg.Amount,
		Fee:        msg.Fee,
		Collateral: msg.Collateral,
		Deadline:   msg.Deadline,
		State:      "requested",
		Borrower:   msg.Creator,
		Lender:     "",
		Timestamp:  ctx.BlockHeight(),
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

	fee, err := sdk.ParseCoinsNormalized(loan.Fee)
	if err != nil {
		panic(err)
	}

	collateralPrice := k.TypedLoan(ctx, collateral)
	// first times collateral price by collateral[0].amount
	dollarAmountC := collateral[0].Amount.MulRaw(int64(collateralPrice.Price))
	// times dollar amount by 1 billion to get the amount needed in ctz to send
	requiredCollateral := types.Cwei.Mul(dollarAmountC)

	sdkError2 := k.bankKeeper.SendCoinsFromAccountToModule(ctx, borrower, types.ModuleName, fee)
	if sdkError2 != nil {
		return nil, sdkError2
	}
	// no switch needed here all loan amounts are paid out in zusd amount = Coin{denom:zusd, amount:1}
	amountPrice := &types.TokenPrice{amount[0], 1}

	// need to use sdkmath.Float64 since numbers are sdk.Int takes
	// Float64 is a method on LegacyDec type
	// can use sdkmath ToLegacyDec
	// turn prices into floats for risk check
	collateralFloat, _ := sdkmath.LegacyDec(collateral[0].Amount).Float64()
	amountFloat, _ := sdkmath.LegacyDec(amount[0].Amount).Float64()

	collateralPriceFloat := collateralFloat * float64(collateralPrice.Price)
	amountPriceFloat := amountFloat * float64(amountPrice.Price)

	// calculate risk using ratio collateral price / amount price > .909090909
	risk := amountPriceFloat / collateralPriceFloat

	if risk < .909090909 && !blackList[msg.Creator] {
		err := k.MintTokens(ctx, borrower, sdk.NewCoin("zusd", amount[0].Amount))
		if err != nil {
			return nil, sdkerrors.Wrap(types.ErrInvalidRequest, "Error minting tokens")
		}
		// send collateral from borrower to loan module account
		// can't send type coin needs type coins
		// make a coin from collateral[0].Denom and requiredCollateral
		cCoin := sdk.NewCoin(collateral[0].Denom, requiredCollateral)
		// can now pass cCoin as type coins
		sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, borrower, types.ModuleName, sdk.NewCoins(cCoin))
		if sdkError != nil {
			return nil, sdkError
		}

		// append loan to store
		loan.Lender = moduleAccountLoan
		loan.State = "approved"
		k.AppendLoan(ctx, loan)
		return &types.MsgRequestLoanResponse{}, nil

	} else {
		return nil, sdkerrors.Wrap(types.ErrInvalidRequest, "Loan risk too high")
	}

}
