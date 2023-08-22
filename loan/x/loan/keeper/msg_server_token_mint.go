package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	//sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	sdkmath "cosmossdk.io/math"
	"loan/x/loan/types"
)

func (k msgServer) TokenMint(goCtx context.Context, msg *types.MsgTokenMint) (*types.MsgTokenMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	recipientAddr, _ := sdk.AccAddressFromBech32(msg.Creator)
	// TODO: Handling the message
	err := k.MintTokens(ctx, recipientAddr, sdk.NewCoin(msg.Denom, sdkmath.NewInt(int64(msg.DenomAmount))))
	if err != nil {
		return nil, err
	}

	return &types.MsgTokenMintResponse{}, nil
}
