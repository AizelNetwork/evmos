### 1. Initial node1 :

#### 1.1. Generate bench32 format address from evm cold wallet address.

```bash
aizeld debug addr [evm-address]
```

apply the above generated bench32 address for the USER1, USER2, USER3, USER4 in file prod_node1.sh


#### 1.2. Initial node1 by running script :

```bash
./prod_node1.sh
```

### 2. Do with node2 

### 2.1 Clone node1 to node2 

```bash
sudo cp -r $AIZELHOME/node1 $AIZELHOME/node2
```

### 2.2 Customize node2 config

```bash
name "node1" => "node2"
port 8545 => 8555
port 8546 => 8566
port 26657 => 26667
# port 9090 => 9090 # keep it as it is
port 26658 => 26668
port 26656 => 26666
port 6060 => 6070

pruning="nothing" => pruning="custom"

# get boot node id by 
aizeld tendermint show-node-id --home $AIZELHOME/node1
# set the persistent_peers in node2/config/config.toml file
persistent_peers = "<node1_id>@<node1_ip>:26656"
```

### 2.3 Generate node_key and validator_key for node2

```bash
# inital node in other folder
aizeld init node3 --chain-id=evmos_9002-20151225 --home=$AIZELHOME/node3

# delete key files in node2
sudo rm $AIZELHOME/node2/config/node_key.json
sudo rm $AIZELHOME/node2/config/priv_validator_key.json

# copy node_key and validator_key to node2 folder
sudo cp $AIZELHOME/node3/config/node_key.json  $AIZELHOME/node2/config/
sudo cp $AIZELHOME/node3/config/priv_validator_key.json  $AIZELHOME/node2/config/

# delete node3 folder
sudo rm -rf $AIZELHOME/node3
```

### 2.4 Run `gentx` in node2 folders
```bash
./prod_node2.sh
```
### 2.5 copy the `gentx-*` folders under  `$AIZELHOME/node2/config/gentx/` folders into the original `$AIZELHOME/node1/config/gentx/`

```bash
sudo cp -r $AIZELHOME/node2/config/gentx/*  $AIZELHOME/node1/config/gentx/
```

### 3. Back with node2 
#### 3.1 Collect genesis tx 

```bash
./collect-gentxs.sh
```

#### 3.2 Copy genisis file from node1 to node2

```bash
sudo rm $AIZELHOME/node2/config/genesis.json

sudo cp $AIZELHOME/node1/config/genesis.json $AIZELHOME/node2/config/
```

#### 3.3 Start the nodes

```bash
./start-nodes.sh
```

### 3.4 Get block number to test the chain

```bash
curl -X POST \
  -H "Content-Type: application/json" \
  --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}' \
  http://localhost:18545
```

## FAQ
###If you meet this issue when run "make install" : flag provided but not defined: -L/opt/homebrew/opt/openblas/lib, Just need to do :

```bash
unset LDFLAGS
unset CFLAGS
```
