#!/bin/bash
CHAINID="${CHAIN_ID:-evmos_9002-20151225}"
BASE_DENOM="aevmos"
VAL1_KEY="validator1"
# Submit Proposal
evmosd tx gov submit-proposal "$EVMOSHOME/evm-v9-upgrade-proposal.json" \
  --fees 4430808153$BASE_DENOM \
  --home "$EVMOSHOME/node1" \
  --chain-id "$CHAINID" \
  --from "$VAL1_KEY" \
  --yes \
  > "$EVMOSHOME/submit-proposal.log"


# evmosd tx gov draft-proposal \
#   --home "$EVMOSHOME/node1" \
#   --chain-id "evmos_9002-20151225" \
#   --from "validator1" \