## 1. Initial nodes :

```bash
evmosd init node1 --chain-id=evmosix_9000-2024 --home=$EVMOSHOME/node1

evmosd init node2 --chain-id=evmosix_9000-2024 --home=$EVMOSHOME/node2

evmosd init node3 --chain-id=evmosix_9000-2024 --home=$EVMOSHOME/node3
```
## 2.0 add keys :

### 2.0.1 For keys validator1, only put in node1:
```bash 
# Replace <your-key-name> with a name for your key that you will remember, such as validator1, validator2, validator3
evmosd keys add validator1 --home=$EVMOSHOME/node1
# or import it from the exported private key 
# evmosd keys import validator1 $EVMOSHOME/validator1.pem --home=$EVMOSHOME/node1
# To see a list of wallets on your node
evmosd keys list --home=$EVMOSHOME/node1
```

## 2.1 Add Genesis Accounts : Add initial funds to the accounts in the genesis file:

```bash
# pre fund ukava
evmosd add-genesis-account validator1  100000000000stake  --home=$EVMOSHOME/node1

```

## 2.2 stake and set validators :
### 2.2.1 Manually add balance and total_supply for stake , each account one by one
```bash
      "balances": [
        {
          "address": "kava1qqkdj2vvzdgq7e8yzmncjq5nfl0lhpy4rt6e2x",
          "coins": [
            {
              "denom": "stake",
              "amount": "1374397274592"
            },
            {
              "denom": "ukava",
              "amount": "1000000000"
            }
          ]
        },
        {
          "address": "kava18p4et3um5aez9cnmxwld7yft6xu2xtvnv5jwqm",
          "coins": [
            {
              "denom": "stake",
              "amount": "1374397274592"
            },
            {
              "denom": "ukava",
              "amount": "1000000000"
            }
          ]
        },
        {
          "address": "kava1m57dmeyqqa3ujv3zpy353em8fdram97e49789t",
          "coins": [
            {
              "denom": "stake",
              "amount": "1374397274592"
            },
            {
              "denom": "ukava",
              "amount": "1000000000"
            }
          ]
        }
      ],
      "supply": [
        {
          "denom": "stake",
          "amount": "4123191823776"
        },
        {
          "denom": "ukava",
          "amount": "3000000000"
        }
      ]
```
### 2.2.2 Gentx for validator1
```bash
# for validator1
evmosd gentx validator1 100000000000stake --chain-id aizel_2024-12 --commission-rate 0.10 --commission-max-rate 0.20 --commission-max-change-rate 0.01 --min-self-delegation 1 --home=$EVMOSHOME/node1
```

### 2.2.3 collect gents
```bash
evmosd collect-gentxs --home=$EVMOSHOME/node1
```

## 2.3 pre-fund evm addresses :

### 2.3.1 using the following tool to convert evm address to bech32 format address :

```golang
package main

import (
	"encoding/hex"
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/bech32"
)

func evmToBech32(evmAddr, prefix string) (string, error) {
	// Remove the "0x" prefix and decode the hex address
	addressBytes, err := hex.DecodeString(evmAddr[2:])
	if err != nil {
		return "", err
	}

	// Convert to Bech32 format
	bech32Addr, err := bech32.ConvertAndEncode(prefix, addressBytes)
	if err != nil {
		return "", err
	}

	return bech32Addr, nil
}

func main() {
	evmAddr := "0xaaafB3972B05630fCceE866eC69CdADd9baC2771"
	prefix := "evmos"

	bech32Addr, err := evmToBech32(evmAddr, prefix)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Bech32 Address:", bech32Addr)
}
```

### 2.3.2 add the converted evm address to genesis account and pre-fund with aphoton

```json
## add new account to accounts
{
          "@type": "/ethermint.types.v1.EthAccount",
          "base_account": {
            "address": "kava142hm89etq43sln8wsehvd8x6mkd6cfm3yugv6c",
            "pub_key": null,
            "account_number": "3",
            "sequence": "1"
          },
          "code_hash": "0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470"
}
## add prefund to balances
{
          "address": "kava142hm89etq43sln8wsehvd8x6mkd6cfm3yugv6c",
          "coins": [
            {
              "denom": "aphoton",
              "amount": "100000000000000000000000000"
            }
          ]
}
## add new coins to total_supply
{
          "denom": "aphoton",
          "amount": "100000000000000000000000000"
}
```

### 2.3.3 update genesis eve account :

```bash
evmosd debug addr AAAFB3972B05630FCCEE866EC69CDADD9BAC2771  --home=$EVMOSHOME/node1
```

## 2.4 Other settings

### 2.4.1 update genesis file's consensus_params to follow :

```json
 "block": {
      "max_bytes": "142020096",
      "max_gas": "3000000000"
    },
    "evidence": {
      "max_age_num_blocks": "100000",
      "max_age_duration": "172800000000000",
      "max_bytes": "1448576"
    },
```

### 2.4.2 Validate genesis file :

```bash
evmosd validate-genesis --home=$EVMOSHOME/node1
```
### 2.4.3 Copy the final validated genesis file to other nodes's config directory
```bash
cp $EVMOSHOME/node1/config/genesis.json $EVMOSHOME/node2/config/
cp $EVMOSHOME/node1/config/genesis.json $EVMOSHOME/node3/config/
```

### 2.4.4. Start nodes :
```bash
nohup evmosd start --home=$EVMOSHOME/node1 > $EVMOSHOME/node1/node1.log 2>&1 &
nohup evmosd start --home=$EVMOSHOME/node2 > $EVMOSHOME/node2/node2.log 2>&1 &
nohup evmosd start --home=$EVMOSHOME/node3 > $EVMOSHOME/node3/node3.log 2>&1 &
```

## 3 FQA :

### 3.1 make install issue

When run "make install" and occurs the following errors :

```bash
flag provided but not defined: -L/opt/homebrew/opt/openblas/lib
```

Just run the following CLI before make install

```bash
unset LDFLAGS
unset CFLAGS
```
