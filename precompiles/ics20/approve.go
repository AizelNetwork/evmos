// Copyright Tharsis Labs Ltd.(Aizel)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/AizelNetwork/evmos/blob/main/LICENSE)

package ics20

import (
	"github.com/AizelNetwork/evmos/v20/precompiles/authorization"
	"github.com/AizelNetwork/evmos/v20/x/evm/core/vm"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

// Approve implements the ICS20 approve transactions.
func (p Precompile) Approve(
	ctx sdk.Context,
	origin common.Address,
	stateDB vm.StateDB,
	method *abi.Method,
	args []interface{},
) ([]byte, error) {
	grantee, transferAuthz, err := NewTransferAuthorization(method, args)
	if err != nil {
		return nil, err
	}

	// Approve from ICS20 common module
	if err := Approve(
		ctx,
		p.AuthzKeeper,
		p.channelKeeper,
		p.Address(),
		grantee,
		origin,
		p.ApprovalExpiration,
		transferAuthz,
		p.ABI.Events[authorization.EventTypeIBCTransferAuthorization],
		stateDB,
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

// Revoke implements the ICS20 authorization revoke transactions.
func (p Precompile) Revoke(
	ctx sdk.Context,
	origin common.Address,
	stateDB vm.StateDB,
	method *abi.Method,
	args []interface{},
) ([]byte, error) {
	grantee, err := checkRevokeArgs(args)
	if err != nil {
		return nil, err
	}

	// Revoke from ICS20 common module
	if err := Revoke(
		ctx,
		p.AuthzKeeper,
		p.Address(),
		grantee,
		origin,
		p.ABI.Events[authorization.EventTypeIBCTransferAuthorization],
		stateDB,
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

// IncreaseAllowance implements the ICS20 increase allowance transactions.
func (p Precompile) IncreaseAllowance(
	ctx sdk.Context,
	origin common.Address,
	stateDB vm.StateDB,
	method *abi.Method,
	args []interface{},
) ([]byte, error) {
	grantee, sourcePort, sourceChannel, denom, amount, err := checkAllowanceArgs(args)
	if err != nil {
		return nil, err
	}

	// IncreaseAllowance from ICS20 common module
	if err := IncreaseAllowance(
		ctx,
		p.AuthzKeeper,
		p.Address(),
		grantee,
		origin,
		sourcePort,
		sourceChannel,
		denom,
		amount,
		p.ABI.Events[authorization.EventTypeIBCTransferAuthorization],
		stateDB,
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}

// DecreaseAllowance implements the ICS20 decrease allowance transactions.
func (p Precompile) DecreaseAllowance(
	ctx sdk.Context,
	origin common.Address,
	stateDB vm.StateDB,
	method *abi.Method,
	args []interface{},
) ([]byte, error) {
	grantee, sourcePort, sourceChannel, denom, amount, err := checkAllowanceArgs(args)
	if err != nil {
		return nil, err
	}

	// DecreaseAllowance from ICS20 common module
	if err := DecreaseAllowance(
		ctx,
		p.AuthzKeeper,
		p.Address(),
		grantee,
		origin,
		sourcePort,
		sourceChannel,
		denom,
		amount,
		p.ABI.Events[authorization.EventTypeIBCTransferAuthorization],
		stateDB,
	); err != nil {
		return nil, err
	}

	return method.Outputs.Pack(true)
}
