// Copyright Tharsis Labs Ltd.(Aizel)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/aizel/aizel/blob/main/LICENSE)

package tests

import (
	aizeltypes "github.com/AizelNetwork/evmos/v20/types"
	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
)

var (
	UosmoDenomtrace = transfertypes.DenomTrace{
		Path:      "transfer/channel-0",
		BaseDenom: "uosmo",
	}
	UosmoIbcdenom = UosmoDenomtrace.IBCDenom()

	UatomDenomtrace = transfertypes.DenomTrace{
		Path:      "transfer/channel-1",
		BaseDenom: "uatom",
	}
	UatomIbcdenom = UatomDenomtrace.IBCDenom()

	UaizelDenomtrace = transfertypes.DenomTrace{
		Path:      "transfer/channel-0",
		BaseDenom: aizeltypes.BaseDenom,
	}
	UaizelIbcdenom = UaizelDenomtrace.IBCDenom()

	UatomOsmoDenomtrace = transfertypes.DenomTrace{
		Path:      "transfer/channel-0/transfer/channel-1",
		BaseDenom: "uatom",
	}
	UatomOsmoIbcdenom = UatomOsmoDenomtrace.IBCDenom()

	AaizelDenomtrace = transfertypes.DenomTrace{
		Path:      "transfer/channel-0",
		BaseDenom: aizeltypes.BaseDenom,
	}
	AaizelIbcdenom = AaizelDenomtrace.IBCDenom()
)
