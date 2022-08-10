# Join the StaFiHub Mainnet

## Quickstart

This Quickstart tutorial completes the following actions:

* Ensure that you have [compilation prerequisites](./install.md)
* Compile stafihub
* Give your node a moniker and configure it
* Download compressed genesis state
* Put the `genesis.json` file in the correct location

```bash
git clone -b vx.x.x https://github.com/stafihub/stafihub
cd stafihub
make install
stafihubd init chooseanicehandle
wget https://github.com/stafihub/network/blob/main/mainnets/stafihub-1/genesis.json
mv genesis.json ~/.stafihub/config/genesis.json
```

**Go**
Starts stafihub

```bash
stafihubd start  --p2p.seeds bf8328b66dceb4987e5cd94430af66045e59899f@xxx:26656,cfd785a4224c7940e9a10f6c1ab24c343e923bec@xxxx:26656,d72b3011ed46d783e369fdf8ae2055b99a1e5074@xxxx:26656
```

To save those seeds to your settings, put the comma-separated list format seeds in `~/.stafihub/config/config.toml` in the p2p section under seeds.

## Manual Setup of a new Node

These instructions are for setting up a brand new full node from scratch.

Make sure to have the [latest stafihub version installed](./join-mainnet.md).
First, initialize the node.

```bash
stafihubd init <your_custom_moniker>
```

**Note**
Monikers can contain only ASCII characters. Using Unicode characters is not supported and renders your node unreachable.

By default, the `init` command creates your `~/.stafihub` directory with subfolders `config` and `data`.
In the `config` directory, the most important files for configuration are `app.toml` and `config.toml`.

You can edit the `moniker` in the `~/.stafihub/config/config.toml` file:

```toml
# A custom human readable name for this node
moniker = "<your_custom_moniker>"
```

For optimized node performance, edit the `~/.stafihub/config/app.toml` file to enable the anti-spam mechanism and reject incoming transactions with less than the minimum gas prices:

```
# This is a TOML config file.
# For more information, see https://github.com/toml-lang/toml

##### main base config options #####

# The minimum gas prices a validator is willing to accept for processing a
# transaction. A transaction's fees must meet the minimum of any denomination
# specified in this config (for example, 10ufis).

minimum-gas-prices = "0.01ufis"
```

Your full node has been initialized!

## Genesis & Seeds

### Copy the Genesis File

Fetch the mainnet's `genesis.json` file into `stafihubd`'s config directory.

```bash
mkdir -p $HOME/.stafihub/config
wget https://github.com/stafihub/network/blob/main/mainnets/stafihub-1/genesis.json
mv genesis.json $HOME/.stafihub/config
```

If you want to connect to the public testnet instead, click [here](./join-testnet.md)

To verify the correctness of the configuration run:

```bash
stafihubd start
```

### Add Seed Nodes

Your node needs to know how to find peers. You'll need to add healthy seed nodes to `$HOME/.stafihub/config/config.toml`. The [`mainnets`](https://github.com/stafihub/network/tree/main/mainnets) repo contains links to some seed nodes.

If those seeds aren't working, you can find more seeds and persistent peers on a StaFiHub explorer (a list can be found on the [mainnets](https://github.com/stafihub/network/tree/main/mainnets)).

## A Note on Gas and Fees

On StaFiHub mainnet, the accepted denom is `ufis`, where `1fis = 1.000.000ufis`

Transactions on the StaFiHub network need to include a transaction fee in order to be processed. This fee pays for the gas required to run the transaction. The formula is the following:

```
fees = ceil(gas * gasPrices)
```

The `gas` is dependent on the transaction. Different transaction require different amount of `gas`. The `gas` amount for a transaction is calculated as it is being processed, but there is a way to estimate it beforehand by using the `auto` value for the `gas` flag. Of course, this only gives an estimate. You can adjust this estimate with the flag `--gas-adjustment` (default `1.0`) if you want to be sure you provide enough `gas` for the transaction.

The `gasPrice` is the price of each unit of `gas`. Each validator sets a `min-gas-price` value, and will only include transactions that have a `gasPrice` greater than their `min-gas-price`.

The transaction `fees` are the product of `gas` and `gasPrice`. As a user, you have to input 2 out of 3. The higher the `gasPrice`/`fees`, the higher the chance that your transaction will get included in a block.

For mainnet, the recommended `gas-prices` is `0.0025ufis`.

## Set `minimum-gas-prices`

Your full-node keeps unconfirmed transactions in its mempool. In order to protect it from spam, it is better to set a `minimum-gas-prices` that the transaction must meet in order to be accepted in your node's mempool. This parameter can be set in the following file `~/.stafihub/config/app.toml`.

The initial recommended `min-gas-prices` is `0.0025ufis`, but you might want to change it later.

## Pruning of State

There are four strategies for pruning state. These strategies apply only to state and do not apply to block storage.
To set pruning, adjust the `pruning` parameter in the `~/.stafihub/config/app.toml` file.
The following pruning state settings are available:

1. `everything`: Prune all saved states other than the current state.
2. `nothing`: Save all states and delete nothing.
3. `default`: Save the last 100 states and the state of every 10,000th block.
4. `custom`: Specify pruning settings with the `pruning-keep-recent`, `pruning-keep-every`, and `pruning-interval` parameters.

By default, every node is in `default` mode which is the recommended setting for most environments.
If you would like to change your nodes pruning strategy then you must do so when the node is initialized. Passing a flag when starting `stafihub` will always override settings in the `app.toml` file, if you would like to change your node to the `everything` mode then you can pass the `---pruning everything` flag when you call `stafihubd start`.

> Note: When you are pruning state you will not be able to query the heights that are not in your store.

## Run a Full Node

Start the full node with this command:

```bash
stafihubd start
```

Check that everything is running smoothly:

```bash
stafihubd status
```

View the status of the network with the [StaFiHub Explorer](https://stafihub.io).

## Enable the REST API

By default, the REST API is disabled. To enable the REST API, edit the `~/.stafihub/config/app.toml` file, and set `enable` to `true` in the `[api]` section.

```toml
###############################################################################
###                           API Configuration                             ###
###############################################################################
[api]
# Enable defines if the API server should be enabled.
enable = true
# Swagger defines if swagger documentation should automatically be registered.
swagger = false
# Address defines the API server to listen on.
address = "tcp://0.0.0.0:1317"
```

Optionally, you can activate swagger by setting `swagger` to `true` or change the port of the REST API in the parameter `address`.
After restarting your application, you can access the REST API on `YOURNODEIP:1317`.

## GRPC Configuration

By default, gRPC is enabled on port `9090`. In the `~/.stafihub/config/app.toml` file, you can make changes in the gRPC section. To disable the gRPC endpoint, set `enable` to `false`. To change the port, use the `address` parameter.

```toml
###############################################################################
###                           gRPC Configuration                            ###
###############################################################################
[grpc]
# Enable defines if the gRPC server should be enabled.
enable = true
# Address defines the gRPC server address to bind to.
address = "0.0.0.0:9090"
```

## Upgrades

To be best prepared for eventual upgrades, it is recommended to setup [Cosmovisor](https://docs.cosmos.network/master/run-node/cosmovisor.html), a small process manager,  which can swap in new `stafihubd` binaries.

## Background Process

To run the node in a background process with automatic restarts, you can use a service manager like `systemd`. To set this up run the following:

```bash
sudo tee /etc/systemd/system/stafihubd.service > /dev/null <<EOF  
[Unit]
Description=stafihub Daemon
After=network-online.target

[Service]
User=$USER
ExecStart=$(which stafihubd) start
Restart=always
RestartSec=3
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target
EOF
```

If you're using Cosmovisor you want to add

```bash
Environment="DAEMON_HOME=$HOME/.stafihub"
Environment="DAEMON_NAME=stafihubd"
Environment="DAEMON_ALLOW_DOWNLOAD_BINARIES=false"
Environment="DAEMON_RESTART_AFTER_UPGRADE=true"
```

after the `LimitNOFILE` line and replace `$(which stafihubd)` with `$(which cosmovisor)`.

Then setup the daemon

```bash
sudo -S systemctl daemon-reload
sudo -S systemctl enable stafihubd
```

We can then start the process and confirm that it is running

```bash
sudo -S systemctl start stafihubd

sudo service stafihubd status
```

## Export State

stafihub can dump the entire application state into a JSON file. This application state dump is useful for manual analysis and can also be used as the genesis file of a new network.

Export state with:

```bash
stafihubd export > [filename].json
```

You can also export state from a particular height (at the end of processing the block of that height):

```bash
stafihubd export --height [height] > [filename].json
```

If you plan to start a new network from the exported state, export with the `--for-zero-height` flag:

```bash
stafihubd export --height [height] --for-zero-height > [filename].json
```

## Verify Mainnet

Help to prevent a catastrophe by running invariants on each block on your full
node. In essence, by running invariants you ensure that the state of mainnet is
the correct expected state. One vital invariant check is that no atoms are
being created or destroyed outside of expected protocol, however there are many
other invariant checks each unique to their respective module. Because invariant checks
are computationally expensive, they are not enabled by default. To run a node with
these checks start your node with the assert-invariants-blockly flag:

```bash
stafihubd start --assert-invariants-blockly
```

If an invariant is broken on your node, your node will panic and prompt you to send
a transaction which will halt mainnet. For example the provided message may look like:

```bash
invariant broken:
    loose token invariance:
        pool.NotBondedTokens: 100
        sum of account tokens: 101
    CRITICAL please submit the following transaction:
        stafihubd tx crisis invariant-broken staking supply

```

When submitting a invariant-broken transaction, transaction fee tokens are not
deducted as the blockchain will halt (invariant-broken transactions are free transactions).
