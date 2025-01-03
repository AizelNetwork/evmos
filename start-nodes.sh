#!/bin/bash
CHAINID="${CHAIN_ID:-aizel_9002-20151225}"
BASE_DENOM="aaizel"
TRACE=""
LOGLEVEL="info"
# Start the node1
aizeld start \
	--metrics "$TRACE" \
	--log_level $LOGLEVEL \
	--minimum-gas-prices=0.0001$BASE_DENOM \
	--json-rpc.api eth,txpool,personal,net,debug,web3 \
	--home "$EVMOSHOME/node1" \
	--chain-id "$CHAINID" \
	 > $EVMOSHOME/node1/node1.log 2>&1 &

# sleep 2

# Start the node2
aizeld start \
	--metrics "$TRACE" \
	--log_level $LOGLEVEL \
	--minimum-gas-prices=0.0001$BASE_DENOM \
	--json-rpc.api eth,txpool,personal,net,debug,web3 \
	--home "$EVMOSHOME/node2" \
	--chain-id "$CHAINID" \
	 > $EVMOSHOME/node2/node2.log 2>&1 &

# tmp
# aizeld query ibc client params --chain-id aizel_9002-20151225