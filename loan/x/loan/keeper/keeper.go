package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"loan/x/loan/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace

		bankKeeper types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,

	bankKeeper types.BankKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,

		bankKeeper: bankKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) MintTokens(ctx sdk.Context, receiver sdk.AccAddress, tokens sdk.Coin) error {
	// mint new tokens if the source of the transfer is the same chain
	if err := k.bankKeeper.MintCoins(
		ctx, types.ModuleName, sdk.NewCoins(tokens),
	); err != nil {
		return err
	}

	// send to receiver
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx, types.ModuleName, receiver, sdk.NewCoins(tokens),
	); err != nil {
		panic(fmt.Sprintf("unable to send coins from module to account despite previously minting coins to module account: %v", err))
	}

	return nil
}

func (k Keeper) BurnTokens(ctx sdk.Context, receiver sdk.AccAddress, tokens sdk.Coin) error {

	// send to receiver
	if err := k.bankKeeper.SendCoinsFromAccountToModule(
		ctx, receiver, types.ModuleName, sdk.NewCoins(tokens),
	); err != nil {
		panic(fmt.Sprintf("unable to send coins from account to module: %v", err))
	}

	// mint new tokens if the source of the transfer is the same chain
	if err := k.bankKeeper.BurnCoins(
		ctx, types.ModuleName, sdk.NewCoins(tokens),
	); err != nil {
		return err
	}

	return nil
}

func (k Keeper) TypedLoan(ctx sdk.Context, token sdk.Coins) *types.TokenPrice {
	// set up pointer to TokenPrice collateral price
	collateralPrice := &types.TokenPrice{}

	// switch on denom string to set parsed coin t a type TokenPrice{sdk.Coin, int}
	switch token[0].Denom {
	case "ctz":
		collateralPrice.Denom = token[0]
		collateralPrice.Price = 1800
		break
	case "cqt":
		collateralPrice.Denom = token[0]
		collateralPrice.Price = 100
		break
	default:
		break
	}
	return collateralPrice
}

/*
add more keeper methods need one for
calculating risk,
one for calculating collateral in terms of cwei,
and one for calculating interest
*/
