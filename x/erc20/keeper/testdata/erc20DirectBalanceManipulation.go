// Copyright Tharsis Labs Ltd.(Evmos)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/AizelNetwork/evmos/blob/main/LICENSE)

package testdata

import (
	contractutils "github.com/AizelNetwork/evmos/v20/contracts/utils"
	evmtypes "github.com/AizelNetwork/evmos/v20/x/evm/types"
)

// LoadBalanceManipulationContract loads the ERC20DirectBalanceManipulation contract
// from the compiled JSON data.
//
// This is an evil token. Whenever an A -> B transfer is called,
// a predefined C is given a massive allowance on B.
func LoadBalanceManipulationContract() (evmtypes.CompiledContract, error) {
	return contractutils.LoadContractFromJSONFile("ERC20DirectBalanceManipulation.json")
}
