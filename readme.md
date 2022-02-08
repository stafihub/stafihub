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

### Sudo operate example

common:

```

./stafihubd tx ledger set-receiver fis1mgjkpyfm00mxk0nmhvfvwhlr65067d53cd49zm --from my-account --chain-id my-chain --keyring-backend file

./stafihubd tx ledger set-commission 0.1 --from my-account --chain-id my-chain --keyring-backend file

```

when add new denom:

```

./stafihubd tx sudo add-denom  --metadata ./metadata.json --chain-id my-chain --from my-account --keyring-backend file

./stafihubd tx ledger set-unbond-fee ratom4 10stake  --from my-account --chain-id my-chain --keyring-backend file

./stafihubd tx ledger set-chain-bonding-duration ratom4 2 --chain-id my-chain --from my-account --keyring-backend file 

./stafihubd tx ledger set-least-bond ratom4 2 --from my-account --chain-id my-chain --keyring-backend file

./stafihubd tx ledger add-new-pool ratom4 cosmos1rp7lvszhm2c4724fzah5qa5ekjq4ckfv226p3n --from my-account --chain-id my-chain --keyring-backend file

./stafihubd tx ledger set-pool-detail ratom4 cosmos1rp7lvszhm2c4724fzah5qa5ekjq4ckfv226p3n cosmos1cad0efr25faywnjp8qp36l8zlqa2sgz0jwn0hl+cosmos1u8pqvzscpp24x8jnaq3l3qtks0l3segy2pvm4t 1 --from my-account --chain-id my-chain --keyring-backend file

./stafihubd tx ledger set-init-bond cosmos1rp7lvszhm2c4724fzah5qa5ekjq4ckfv226p3n 1ratom4 fis1mgjkpyfm00mxk0nmhvfvwhlr65067d53cd49zm --from my-account --chain-id my-chain --keyring-backend file

./stafihubd tx relayers create-relayer ratom4 fis1ychj8z22pw0ruc65mx8nvdn7ca9qylpkzwkkgq --keyring-backend file --from my-account --chain-id my-chain

./stafihubd tx relayers update-threshold ratom4 1 --from my-account --keyring-backend file --chain-id my-chain

```