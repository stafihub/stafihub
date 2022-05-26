# Settings on chain

## Setting examples for admin

### global:

```
stafihubd tx ledger set-protocol-fee-receiver stafi1ukq4mtq604prn5yxul7syh5ysvj0w5jrclvrvc --from admin --chain-id local-stafihub --keyring-backend file

stafihubd query ledger protocol-fee-receiver 
```

### add new rtoken:

```
# set rtoken metadata
stafihubd tx rbank add-denom cosmos cosmosvaloper ./metadata_example.json --chain-id local-stafihub --from admin --keyring-backend file

stafihubd query bank denom-metadata

stafihubd query rbank address-prefix uratom


# set relay fee receiver
stafihubd tx ledger set-relay-fee-receiver uratom stafi1mgjkpyfm00mxk0nmhvfvwhlr65067d538l6cxa --from admin --chain-id local-stafihub --keyring-backend file

stafihubd query ledger relay-fee-receiver uratom



# default 0.1
stafihubd tx ledger set-staking-reward-commission uratom 0.1 --from admin --chain-id local-stafihub --keyring-backend file

# default 0.002
stafihubd tx ledger set-unbond-commission uratom 0.002 --from admin --chain-id local-stafihub --keyring-backend file

# default 1000000ufis
stafihubd tx ledger set-unbond-relay-fee uratom 1000000ufis --from admin --chain-id local-stafihub --keyring-backend file



stafihubd tx relayers add-relayers ledger uratom stafi1ychj8z22pw0ruc65mx8nvdn7ca9qylpkauetvx:stafi1ukq4mtq604prn5yxul7syh5ysvj0w5jrclvrvc --keyring-backend file --from admin --chain-id local-stafihub

stafihubd query relayers relayers ledger uratom



stafihubd tx relayers set-threshold ledger uratom 1 --from admin --keyring-backend file --chain-id local-stafihub

stafihubd query relayers threshold ledger uratom


# set params used by relay
stafihubd tx ledger set-r-params uratom 0.00001stake 600 0 2 0stake cosmosvaloper129kf5egy80e8me93lg3h5lk54kp0tle7w9npre --from admin --keyring-backend file --chain-id local-stafihub

stafihubd query ledger r-params uratom



# this will init bonded pool, exchange rate, pipeline
stafihubd tx ledger set-pool-detail uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 cosmos1cad0efr25faywnjp8qp36l8zlqa2sgz0jwn0hl:cosmos13mwxtgrljf9d5r72sc28496ua4lsga0jvmqz8x 1 --from admin --chain-id local-stafihub --keyring-backend file

stafihubd query ledger pool-detail uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75

stafihubd query ledger bonded-pools uratom

stafihubd query ledger exchange-rate uratom

stafihubd query ledger bond-pipeline uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75
```




### bridge:

```
stafihubd tx bridge add-chain-id 1 --from admin --keyring-backend file --chain-id local-stafihub

stafihubd query bridge chain-ids



stafihubd tx relayers add-relayers bridge 1 stafi1ychj8z22pw0ruc65mx8nvdn7ca9qylpkauetvx --from admin --keyring-backend file --chain-id local-stafihub

stafihubd query relayers relayers bridge 1



stafihubd tx relayers set-threshold bridge 1 1 --from admin --keyring-backend file --chain-id local-stafihub

stafihubd query relayers threshold bridge 1



stafihubd tx bridge set-resourceid-to-denom  000000000000000000000000000000a9e0095b8965c01e6a09c97938f3860901 uratom --from admin --keyring-backend file --chain-id local-stafihub

stafihubd query bridge resourceid-to-denoms



stafihubd tx bridge set-relay-fee-receiver stafi1ychj8z22pw0ruc65mx8nvdn7ca9qylpkauetvx --from admin --keyring-backend file --chain-id local-stafihub

stafihubd query bridge relay-fee-receiver



stafihubd tx bridge set-relay-fee 1 1000000ufis --from admin --keyring-backend file --chain-id local-stafihub

stafihubd query bridge  relay-fee 1
```


### migrate rtoken (after adding new rtoken step):

```
stafihubd tx ledger migrate-init uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 100000000 150000000 200000000 300000000 1.23 --from admin --keyring-backend file --chain-id local-stafihub

stafihubd query bank  total 

stafihubd query ledger exchange-rate uratom

stafihubd query ledger bond-pipeline uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75



stafihubd tx ledger migrate-unbondings uratom --unbondings ./unbondings_example.json --from admin --keyring-backend file --chain-id local-stafihub

stafihubd query ledger pool-unbond uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 3



stafihubd tx bridge set-denom-type uratom  1 --from admin --keyring-backend file --chain-id local-stafihub

stafihubd query bridge denom-types
```


### rdex:

stafihubd tx rdex create-pool 10ufis 20uratom --from admin --chain-id local-stafihub --keyring-backend file

stafihubd tx rdex add-provider stafi1qzt0qajzr9df3en5sk06xlk26n30003c8uhdkg --from admin --chain-id local-stafihub --keyring-backend file

stafihubd tx rdex add-liquidity  100ufis 200uratom --from admin --chain-id local-stafihub --keyring-backend file

stafihubd tx rdex remove-liquidity 10 5 1uratom 1ufis ufis --from admin --chain-id local-stafihub --keyring-backend file

stafihubd tx rdex swap 2ufis 1uratom  --from admin --chain-id local-stafihub --keyring-backend file

### mining:

stafihubd tx mining add-mining-provider stafi1ychj8z22pw0ruc65mx8nvdn7ca9qylpkauetvx  --from admin --chain-id local-stafihub --keyring-backend file

stafihubd tx mining add-reward-token ufis 200 --from admin --chain-id local-stafihub --keyring-backend file




stafihubd tx mining add-stake-pool ufis ./add_stake_pool_example.json  --from relay1 --chain-id local-stafihub --keyring-backend file

stafihubd tx mining stake 0 10ufis 0 --from my-account --chain-id local-stafihub --keyring-backend file 

stafihubd tx mining claim-reward 0 0 --from my-account --chain-id local-stafihub --keyring-backend file

stafihubd tx mining add-reward 1 0 300 0 0 --from relay1 --chain-id local-stafihub --keyring-backend file

stafihubd tx mining withdraw 1 10ufis 0 --from test --chain-id local-stafihub --keyring-backend file

## Operate examples for user

### liquidity bond (gaiad example):

```
gaiad tx bank send userAccount cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 1000000stake --memo 1:stafi1ukq4mtq604prn5yxul7syh5ysvj0w5jrclvrvc --keyring-backend file --chain-id local-cosmos
```

### recover (gaiad example):

```
gaiad tx bank send userAccount cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 1stake --memo 2:stafi1ukq4mtq604prn5yxul7syh5ysvj0w5jrclvrvc:9A80F3E6A007E1144BE34F4A0AC35B9288C19641BCAD3464277168000AF5FC66 --keyring-backend file --chain-id local-cosmos
```

### liquidity unbond:

```
stafihubd tx ledger liquidity-unbond cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 100uratom cosmos1j9dues7ey2a39nes4ewfvyma96d3f5zrdhnfan --keyring-backend file --from user --home /Users/tpkeeper/gowork/stafi/rtoken-relay-core/keys/stafihub --chain-id local-stafihub
```

### deposit (transfer token to external chain):
```
stafihubd tx bridge deposit 1 uratom 800 dccf954570063847d73746afa0b0878f2c779d42089c5d9a107f2aca176e985f --from my-account --chain-id local-stafihub --keyring-backend file
```


