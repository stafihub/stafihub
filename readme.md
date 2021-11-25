# stafihub
**stafihub** is a blockchain built using Cosmos SDK and Tendermint and created with [Starport](https://github.com/tendermint/starport).

### Get started for dev

```
starport chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure for dev

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Starport docs](https://docs.starport.network).


### Running node for production
1. build the chain. `starport build --release -t linux:amd64 -t darwin:amd64 -t darwin:arm64`
2. unzip the release/stafihub_darwin_amd64.tar.gz file to get a execute file, like stafihubd
1. Setting up the keyring, commands:
    - `KEYPASSWD=123456789`
    - `./release/stafihubd config keyring-backend file`
    - `(echo $KEYPASSWD; echo $KEYPASSWD) | ./release/stafihubd keys add kael --keyring-backend file`
    - `MY_VALIDATOR_ADDRESS=$(./release/stafihubd show kael -a --keyring-backend file)`
2. Prepare genesis file and other config files
    - `./release/stafihubd init mynode -o`
    - `./release/stafihubd add-genesis-account $MY_VALIDATOR_ADDRESS 100000000000stake`
    - `./release/stafihubd gentx kael 100000000stake --keyring-backend file --chain-id stafihub`
    - `./release/stafihubd collect-gentxs`
3. Start node: `./release/stafihubd start`
4. There might be this kind of error:
`Error: couldn't get client config: open ~/.stafihub/config/client.toml: permission denied`, just add sudo to your command.
for example, start node: `sudo ./release/stafihubd start`

