// Copyright Tharsis Labs Ltd.(Aizel)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/aizel/aizel/blob/main/LICENSE)
package keeper

import (
	v8 "github.com/AizelNetwork/evmos/v20/x/evm/migrations/v8"
	"github.com/AizelNetwork/evmos/v20/x/evm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper         Keeper
	legacySubspace types.Subspace
}

// NewMigrator returns a new Migrator instance.
func NewMigrator(keeper Keeper, legacySubspace types.Subspace) Migrator {
	return Migrator{
		keeper:         keeper,
		legacySubspace: legacySubspace,
	}
}

// Migrate7to8 migrates the store from consensus version 6 to 7.
func (m Migrator) Migrate7to8(ctx sdk.Context) error {
	return v8.MigrateStore(ctx, m.keeper.storeKey, m.keeper.cdc)
}
