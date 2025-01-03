// Copyright Tharsis Labs Ltd.(Aizel)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/aizel/aizel/blob/main/LICENSE)

package keeper

import (
	"github.com/AizelNetwork/evmos/v20/x/inflation/v1/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetParams returns the total set of inflation parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.ParamsKey)
	if len(bz) == 0 {
		return params
	}

	k.cdc.MustUnmarshal(bz, &params)
	return params
}

// SetParams sets the inflation params in a single key
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) error {
	store := ctx.KVStore(k.storeKey)
	bz, err := k.cdc.Marshal(&params)
	if err != nil {
		return err
	}

	store.Set(types.ParamsKey, bz)

	return nil
}
