package keeper

import (
	// this line is used by starport scaffolding
	"github.com/earth2378/logistic/x/logistic/types"
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewQuerier creates a new querier for logistic clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		// this line is used by starport scaffolding # 2
		// Step 9: register list and get deal to main querier

		case types.QueryDeal:
			return getDeal(ctx, path[1:], k)
		case types.QueryListDeal:
			return getListDeal(ctx, k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown logistic query endpoint")
		}
	}
}
