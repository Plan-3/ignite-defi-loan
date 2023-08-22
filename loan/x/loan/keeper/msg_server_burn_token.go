package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkmath "cosmossdk.io/math"
	"loan/x/loan/types"
)

func (k msgServer) BurnToken(goCtx context.Context, msg *types.MsgBurnToken) (*types.MsgBurnTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	senderAddr, _ := sdk.AccAddressFromBech32(msg.Creator)
	// TODO: Handling the message
	err := k.BurnTokens(ctx, senderAddr, sdk.NewCoin(msg.Denom, sdkmath.NewInt(int64(msg.Amount))))
	if err != nil {
		return nil, err
	}

	return &types.MsgBurnTokenResponse{}, nil
}
