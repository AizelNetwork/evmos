// Copyright Tharsis Labs Ltd.(Aizel)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/aizel/aizel/blob/main/LICENSE)
package testutil

import (
	"math/big"
	"testing"

	storetypes "cosmossdk.io/store/types"
	"github.com/AizelNetwork/evmos/v20/x/evm/core/vm"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

// NewPrecompileContract creates a new precompile contract and sets the gas meter.
func NewPrecompileContract(t *testing.T, ctx sdk.Context, caller common.Address, precompile vm.ContractRef, gas uint64) (*vm.Contract, sdk.Context) {
	contract := vm.NewContract(vm.AccountRef(caller), precompile, big.NewInt(0), gas)
	ctx = ctx.WithGasMeter(storetypes.NewInfiniteGasMeter())
	initialGas := ctx.GasMeter().GasConsumed()
	require.Equal(t, uint64(0), initialGas)
	return contract, ctx
}
