package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/loan module sentinel errors
var (
	ErrWrongLoanState = sdkerrors.Register(ModuleName, 1, "loan is not in the correct state for this action")
	ErrKeyNotFound    = sdkerrors.Register(ModuleName, 2, "key not found")
	ErrInvalidRequest = sdkerrors.Register(ModuleName, 3, "invalid request")
	ErrUnauthorized   = sdkerrors.Register(ModuleName, 4, "unauthorized for this request")
	ErrDeadline       = sdkerrors.Register(ModuleName, 5, "deadline not reached")
)
