// Copyright Tharsis Labs Ltd.(Aizel)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/aizel/aizel/blob/main/LICENSE)

package transfer

import (
	"github.com/AizelNetwork/evmos/v20/x/ibc/transfer/keeper"
	ibctransfer "github.com/cosmos/ibc-go/v8/modules/apps/transfer"
	porttypes "github.com/cosmos/ibc-go/v8/modules/core/05-port/types"
)

var _ porttypes.IBCModule = IBCModule{}

// IBCModule implements the ICS26 interface for transfer given the transfer keeper.
type IBCModule struct {
	*ibctransfer.IBCModule
}

// NewIBCModule creates a new IBCModule given the keeper
func NewIBCModule(k keeper.Keeper) IBCModule {
	transferModule := ibctransfer.NewIBCModule(*k.Keeper)
	return IBCModule{
		IBCModule: &transferModule,
	}
}
