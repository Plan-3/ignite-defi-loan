package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"loan/x/loan/types"
)

func (k msgServer) Stake(goCtx context.Context, msg *types.MsgStake) (*types.MsgStakeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	amount, _ := sdk.ParseCoinsNormalized(msg.Amount)

	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, amount)
	if err != nil {
		return nil, err
	}

	// check keeper.go for the function
	ctzPrice, cqtPrice, zusdTotalAtTimeOfDeposit, _ := k.ModuleStakingAmounts(ctx)

	// create coins based on the amounts from loop
	zusdPlaceHolder := sdk.NewCoin("zPh", zusdTotalAtTimeOfDeposit)
	collateralPlaceHolder := sdk.NewCoin("cPh", ctzPrice.Add(cqtPrice))
	positionCoin := sdk.NewCoin("posi", amount[0].Amount)

	errZ := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(zusdPlaceHolder))
	if errZ != nil {
		return nil, errZ
	}
	errC := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(collateralPlaceHolder))
	if errC != nil {
		return nil, errC
	}
	errP := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(positionCoin))
	if errP != nil {
		return nil, errP
	}

	errSZ := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creator, sdk.NewCoins(zusdPlaceHolder))
	if errSZ != nil {
		return nil, errSZ
	}
	errSC := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creator, sdk.NewCoins(collateralPlaceHolder))
	if errSC != nil {
		return nil, errSC
	}
	errSP := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creator, sdk.NewCoins(positionCoin))
	if errSP != nil {
		return nil, errSP
	}

	return &types.MsgStakeResponse{}, nil
}
