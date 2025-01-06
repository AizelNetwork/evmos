// Copyright Tharsis Labs Ltd.(Aizel)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/AizelNetwork/evmos/blob/main/LICENSE)

package grpc

import (
	"github.com/AizelNetwork/evmos/v20/testutil/integration/aizel/network"
	commongrpc "github.com/AizelNetwork/evmos/v20/testutil/integration/common/grpc"
	evmtypes "github.com/AizelNetwork/evmos/v20/x/evm/types"
	feemarkettypes "github.com/AizelNetwork/evmos/v20/x/feemarket/types"
	infltypes "github.com/AizelNetwork/evmos/v20/x/inflation/v1/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum/common"
)

// Handler is an interface that defines the methods that are used to query
// the network's modules via gRPC.
type Handler interface {
	commongrpc.Handler

	// EVM methods
	GetEvmAccount(address common.Address) (*evmtypes.QueryAccountResponse, error)
	EstimateGas(args []byte, GasCap uint64) (*evmtypes.EstimateGasResponse, error)
	GetEvmParams() (*evmtypes.QueryParamsResponse, error)
	GetEvmBaseFee() (*evmtypes.QueryBaseFeeResponse, error)
	GetBalanceFromEVM(address sdk.AccAddress) (*evmtypes.QueryBalanceResponse, error)

	// FeeMarket methods
	GetBaseFee() (*feemarkettypes.QueryBaseFeeResponse, error)
	GetFeeMarketParams() (*feemarkettypes.QueryParamsResponse, error)

	// Gov methods
	GetProposal(proposalID uint64) (*govtypes.QueryProposalResponse, error)
	GetGovParams(paramsType string) (*govtypes.QueryParamsResponse, error)

	// Inflation methods
	GetPeriod() (*infltypes.QueryPeriodResponse, error)
	GetEpochMintProvision() (*infltypes.QueryEpochMintProvisionResponse, error)
	GetSkippedEpochs() (*infltypes.QuerySkippedEpochsResponse, error)
	GetCirculatingSupply() (*infltypes.QueryCirculatingSupplyResponse, error)
	GetInflationRate() (*infltypes.QueryInflationRateResponse, error)
	GetInflationParams() (*infltypes.QueryParamsResponse, error)

	// Staking methods
	GetStakingParams() (*stakingtypes.QueryParamsResponse, error)
}

var _ Handler = (*IntegrationHandler)(nil)

// IntegrationHandler is a helper struct to query the network's modules
// via gRPC. This is to simulate the behavior of a real user and avoid querying
// the modules directly.
type IntegrationHandler struct {
	// We take the IntegrationHandler from common/grpc to get the common methods.
	*commongrpc.IntegrationHandler
	network network.Network
}

// NewIntegrationHandler creates a new IntegrationHandler instance.
func NewIntegrationHandler(network network.Network) Handler {
	return &IntegrationHandler{
		// Is there a better way to do this?
		IntegrationHandler: commongrpc.NewIntegrationHandler(network),
		network:            network,
	}
}