// Copyright Tharsis Labs Ltd.(Aizel)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/AizelNetwork/evmos/blob/main/LICENSE)

package contracts

import (
	contractutils "github.com/AizelNetwork/evmos/v20/contracts/utils"
	evmtypes "github.com/AizelNetwork/evmos/v20/x/evm/types"
)

func LoadFlashLoanContract() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("FlashLoan.json")
}
