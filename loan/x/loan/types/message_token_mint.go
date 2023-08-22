package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgTokenMint = "token_mint"

var _ sdk.Msg = &MsgTokenMint{}

func NewMsgTokenMint(creator string, denom string, denomAmount int64) *MsgTokenMint {
	return &MsgTokenMint{
		Creator: creator,
		Denom: denom,
		DenomAmount: denomAmount,
	}
}

func (msg *MsgTokenMint) Route() string {
	return RouterKey
}

func (msg *MsgTokenMint) Type() string {
	return TypeMsgTokenMint
}

func (msg *MsgTokenMint) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgTokenMint) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTokenMint) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
