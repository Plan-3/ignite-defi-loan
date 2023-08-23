package keeper

import (
	"context"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
