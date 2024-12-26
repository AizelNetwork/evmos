### 1. Initial node1 :

```bash
./prod_node1.sh
```

### 2. Update node1 config data :

#### 2.1 replace "localhost" with "0.0.0.0"

#### 2.2 replace "127.0.0.1" with "0.0.0.0"

#### 2.3 find "cors" and set it to true

### 3. Do with node2 

### 3.1 Clone node1 to node2 

```bash
sudo cp -r $EVMOSHOME/node1 $EVMOSHOME/node2
```

### 3.2 Customize node2 config

```bash
name "node1" => "node2"
port 8545 => 8555
port 8546 => 8566
```
### 3.3 Run `gentx` in node2 folders
```bash
./prod_node2.sh
```
### 3.4 the `gentx-*` folders under  `$EVMOSHOME/node2/config/gentx/` folders into the original `$EVMOSHOME/node1/config/gentx/`

```bash
sudo cp -r $EVMOSHOME/node2/config/gentx/*  $EVMOSHOME/node1/config/gentx/
```

### 4. Back with node2 
#### 4.1 Collect genesis tx 

```bash
./collect-gentxs.sh
```
#### 4.2 Vadliate genisis file

```bash
evmosd validate-genesis --home "$EVMOSHOME/node1"
```

#### 4.2 Start the nodes

```bash
./start-nodes.sh
```
