#!/bin/bash

CHAINID="${CHAIN_ID:-aizel_2015-3333}"
BASE_DENOM="aaizel"
MONIKER="node2"
# Remember to change to other types of keyring like 'file' in-case exposing to outside world,
# otherwise your balance will be wiped quickly
# The keyring test does not require private key to steal tokens from you
KEYRING="file"
KEYALGO="eth_secp256k1"
LOGLEVEL="info"
# Set dedicated home directory for the aizeld instance
HOMEDIR="$AIZELHOME/node2"
# to trace evm
#TRACE="--trace"
TRACE=""

# feemarket params basefee
BASEFEE=1000000000

# Path variables
CONFIG=$HOMEDIR/config/config.toml
APP_TOML=$HOMEDIR/config/app.toml
GENESIS=$HOMEDIR/config/genesis.json
TMP_GENESIS=$HOMEDIR/config/tmp_genesis.json


VAL2_KEY="validator2"
VAL2_MNEMONIC="surround soft tragic ensure accuse tooth soul attack ahead cheese few taxi hope priority globe crater enable still silly glad hire correct brush defense"
VAL2_KEY_PASS="20151225"

# Sign genesis transaction
aizeld gentx "$VAL2_KEY" 1000000000000000000000$BASE_DENOM --gas-prices ${BASEFEE}$BASE_DENOM --keyring-backend "$KEYRING" --chain-id "$CHAINID" --home "$HOMEDIR"