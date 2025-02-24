#!/bin/bash
CHAINID="${CHAIN_ID:-evmos_9002-20151225}"
BASE_DENOM="aevmos"
VAL1_KEY="validator1"
# Query Proposal
evmosd query gov proposal 1 \
  --home "$EVMOSHOME/node1" \
  --chain-id "$CHAINID" \
  --node "tcp://localhost:26657" \
  > "$EVMOSHOME/query-proposal.log"