// Copyright Tharsis Labs Ltd.(Aizel)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/AizelNetwork/evmos/blob/main/LICENSE)

package utils

import (
	"encoding/json"
	"math/big"

	"github.com/AizelNetwork/evmos/v20/contracts"
	"github.com/AizelNetwork/evmos/v20/testutil/integration/aizel/network"
	evmtypes "github.com/AizelNetwork/evmos/v20/x/evm/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// GetERC20Balance returns the token balance of a given account address for
// an ERC-20 token at the given contract address.
func GetERC20Balance(nw network.Network, tokenAddress, accountAddress common.Address) (*big.Int, error) {
	input, err := contracts.ERC20MinterBurnerDecimalsContract.ABI.Pack(
		"balanceOf",
		accountAddress,
	)
	if err != nil {
		return nil, err
	}

	callData, err := json.Marshal(evmtypes.TransactionArgs{
		To:    &tokenAddress,
		Input: (*hexutil.Bytes)(&input),
	})
	if err != nil {
		return nil, err
	}

	ethRes, err := nw.GetEvmClient().EthCall(
		nw.GetContext(),
		&evmtypes.EthCallRequest{
			Args: callData,
		},
	)
	if err != nil {
		return nil, err
	}

	var balance *big.Int
	err = contracts.ERC20MinterBurnerDecimalsContract.ABI.UnpackIntoInterface(&balance, "balanceOf", ethRes.Ret)
	if err != nil {
		return nil, err
	}

	return balance, nil
}