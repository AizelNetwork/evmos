// Copyright Tharsis Labs Ltd.(Aizel)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/AizelNetwork/evmos/blob/main/LICENSE)

package utils

import (
	"context"
	"fmt"

	"cosmossdk.io/math"
	"github.com/AizelNetwork/evmos/v20/testutil/integration/aizel/keyring"
	cmnfactory "github.com/AizelNetwork/evmos/v20/testutil/integration/common/factory"
	cmnnet "github.com/AizelNetwork/evmos/v20/testutil/integration/common/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// FundAccountWithBaseDenom funds the given account with the given amount of the network's
// base denomination.
func FundAccountWithBaseDenom(tf cmnfactory.CoreTxFactory, nw cmnnet.Network, sender keyring.Key, receiver sdk.AccAddress, amount math.Int) error {
	return tf.FundAccount(sender, receiver, sdk.NewCoins(sdk.NewCoin(nw.GetBaseDenom(), amount)))
}

// CheckBalances checks that the given accounts have the expected balances and
// returns an error if that is not the case.
func CheckBalances(ctx context.Context, client banktypes.QueryClient, balances []banktypes.Balance) error {
	for _, balance := range balances {
		addr := balance.GetAddress()
		for _, coin := range balance.GetCoins() {
			balance, err := client.Balance(ctx, &banktypes.QueryBalanceRequest{Address: addr, Denom: coin.Denom})
			if err != nil {
				return err
			}

			if !balance.Balance.IsEqual(coin) {
				return fmt.Errorf(
					"expected balance %s, got %s for address %s",
					coin, balance.Balance, addr,
				)
			}
		}
	}

	return nil
}