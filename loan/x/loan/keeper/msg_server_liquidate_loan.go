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

	// convert deadline to int to compare to block height
	deadline, err := strconv.ParseInt(loan.Deadline, 10, 64)
	if err != nil {
		panic(err)
	}
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrKeyNotFound, "key %d doesn't exist", msg.Id)
	}
	if (loan.Timestamp + deadline) > ctx.BlockHeight() {
		return nil, sdkerrors.Wrap(types.ErrUnauthorized, "Cannot liquidate: not past deadline")
	}
	if loan.State != "approved" {
		return nil, sdkerrors.Wrapf(types.ErrWrongLoanState, "%v", loan.State)
	}

	borrower, _ := sdk.AccAddressFromBech32(loan.Borrower)
	collateral, _ := sdk.ParseCoinsNormalized(loan.Collateral)
	amount, _ := sdk.ParseCoinsNormalized(loan.Amount)
	
	// burn 99% of collateral
	collateralPrice := k.TypedLoan(ctx, collateral)
	collateralAmount := collateral[0].Amount.Mul(sdk.NewInt(int64(collateralPrice.Price)))
	collateralTotal := types.Cwei.Mul(collateralAmount)
	collateralBurn := collateralTotal.MulRaw(99).QuoRaw(100)
	collateralLiquidatedToPool := collateralTotal.Sub(collateralBurn)
	// convert to sdk.Coin to send to burn coins
	burnCoin := sdk.NewCoin(collateral[0].Denom, collateralBurn)
	liquidatedCollateral := sdk.NewCoin(collateral[0].Denom, collateralLiquidatedToPool)
	burnZusd := sdk.NewCoin("zusd", amount[0].Amount)
	errB := k.bankKeeper.BurnCoins(ctx, types.Nbtp, sdk.NewCoins(burnCoin))
	if errB != nil {
		return nil, errB
	}
	// for time being force zusd back from borrower to loan module to burn
	errB1 := k.bankKeeper.SendCoinsFromAccountToModule(ctx, borrower, types.ModuleName, sdk.NewCoins(burnZusd))
	if errB1 != nil {
		return nil, errB1
	}
	errB2 := k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(burnZusd))
	if errB2 != nil {
		return nil, errB2
	}
	errS := k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.Nbtp, types.ModuleName, sdk.NewCoins(liquidatedCollateral))
	if errS != nil {
		return nil, errS
	}
	
	loan.State = "liquidated"
	k.SetLoan(ctx, loan)

	return &types.MsgLiquidateLoanResponse{}, nil
}
