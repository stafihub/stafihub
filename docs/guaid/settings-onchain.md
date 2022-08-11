# Settings on chain

## Setting examples for admin

### set protocol fee receiver

```bash
stafihubd tx ledger set-protocol-fee-receiver stafi1ukq4mtq604prn5yxul7syh5ysvj0w5jrclvrvc --from admin --chain-id local-stafihub --keyring-backend file

stafihubd query ledger protocol-fee-receiver 
```

### add new rtoken

```bash
# set rtoken metadata
stafihubd tx rbank add-denom cosmos cosmosvaloper ./metadata/metadata_ratom.json --chain-id local-stafihub --from admin --keyring-backend file

stafihubd query bank denom-metadata

stafihubd query rbank address-prefix uratom

# set relay fee receiver
stafihubd tx ledger set-relay-fee-receiver uratom stafi1mgjkpyfm00mxk0nmhvfvwhlr65067d538l6cxa --from admin --chain-id local-stafihub --keyring-backend file

stafihubd query ledger relay-fee-receiver uratom

# this will init bonded pool, exchange rate, pipeline
stafihubd tx ledger init-pool uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 --from admin --chain-id local-stafihub --keyring-backend file

stafihubd query ledger bonded-pools uratom

stafihubd query ledger exchange-rate uratom

stafihubd query ledger bond-pipeline uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75



# add relayers
stafihubd tx relayers add-relayers ledger uratom stafi1ychj8z22pw0ruc65mx8nvdn7ca9qylpkauetvx:stafi1ukq4mtq604prn5yxul7syh5ysvj0w5jrclvrvc --keyring-backend file --from admin --chain-id local-stafihub

stafihubd query relayers relayers ledger uratom

# set threshold
stafihubd tx relayers set-threshold ledger uratom 1 --from admin --keyring-backend file --chain-id local-stafihub

stafihubd query relayers threshold ledger uratom

# set params used by relay
stafihubd tx ledger set-r-params uratom 0.00001stake 600 0 2 0stake --from admin --keyring-backend file --chain-id local-stafihub

stafihubd query ledger r-params uratom

# set pool detail for multisig/ica pool
stafihubd tx ledger set-pool-detail uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 cosmos1cad0efr25faywnjp8qp36l8zlqa2sgz0jwn0hl:cosmos13mwxtgrljf9d5r72sc28496ua4lsga0jvmqz8x 1 --from admin --chain-id local-stafihub --keyring-backend file

stafihubd query ledger pool-detail uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75



# default 0.1
stafihubd tx ledger set-staking-reward-commission uratom 0.15 --from admin --chain-id local-stafihub --keyring-backend file

stafihubd q ledger staking-reward-commission uratom

# default 0.002
stafihubd tx ledger set-unbond-commission uratom 0.0025 --from admin --chain-id local-stafihub --keyring-backend file

stafihubd q ledger unbond-commission uratom

# default 1000000ufis
stafihubd tx ledger set-unbond-relay-fee uratom 1000005ufis --from admin --chain-id local-stafihub --keyring-backend file

stafihubd q ledger unbond-relay-fee uratom

```

### register ica pool
```
# register ica pool (need set rtoken metadata before this)
stafihubd tx ledger register-ica-pool uratom connection-0 --keyring-backend file --from admin --chain-id local-stafihub --gas 410000

stafihubd q ledger ica-pool-list uratom

# set withdrawal address
stafihubd tx ledger set-withdrawal-addr cosmos1gsth46z50w256p4kq36xquh4q90mfjq0t4lm9scln6zucg64epyqudzqzm --keyring-backend file --from admin --chain-id local-stafihub --gas 410000

```

### rvalidator

```bash
# add relayers
stafihubd tx relayers add-relayers rvalidator uratom stafi14z467aut40mcrt2ddyxf7e74fq99udul7kaf9g:stafi15lne70yk254s0pm2da6g59r82cjymzjqvvqxz7 --keyring-backend file --from admin --chain-id local-stafihub

stafihubd q relayers relayers rvalidator uratom

# set threshold
stafihubd tx relayers set-threshold rvalidator uratom 1 --from admin --keyring-backend file --chain-id local-stafihub

stafihubd q relayers threshold rvalidator uratom

# init rvalidator (should init target validators of pool before rtoken relay start)
stafihubd tx rvalidator init-r-validator uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 cosmosvaloper129kf5egy80e8me93lg3h5lk54kp0tle7w9npre --from admin --chain-id local-stafihub --keyring-backend file  

stafihubd q rvalidator r-validator-list uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75

# add rvalidator
stafihubd tx rvalidator add-r-validator uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 cosmosvaloper1cad0efr25faywnjp8qp36l8zlqa2sgz0h686mv  --chain-id local-stafihub --keyring-backend file --from admin

stafihubd q rvalidator r-validator-list uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75

# rm rvalidator
stafihubd tx rvalidator rm-r-validator uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 cosmosvaloper1cad0efr25faywnjp8qp36l8zlqa2sgz0h686mv cosmosvaloper129kf5egy80e8me93lg3h5lk54kp0tle7w9npre --from admin --chain-id local-stafihub --keyring-backend file
```



### bridge

```bash
stafihubd tx bridge add-chain-id 1 --from admin --keyring-backend file --chain-id local-stafihub

stafihubd query bridge chain-ids



stafihubd tx relayers add-relayers bridge 1 stafi1ychj8z22pw0ruc65mx8nvdn7ca9qylpkauetvx --from admin --keyring-backend file --chain-id local-stafihub

stafihubd query relayers relayers bridge 1



stafihubd tx relayers set-threshold bridge 1 1 --from admin --keyring-backend file --chain-id local-stafihub

stafihubd query relayers threshold bridge 1



stafihubd tx bridge set-resourceid-to-denom  000000000000000000000000000000a9e0095b8965c01e6a09c97938f3860901 uratom NATIVE --from admin --keyring-backend file --chain-id local-stafihub

stafihubd query bridge resourceid-to-denoms



stafihubd tx bridge set-relay-fee-receiver stafi1ychj8z22pw0ruc65mx8nvdn7ca9qylpkauetvx --from admin --keyring-backend file --chain-id local-stafihub

stafihubd query bridge relay-fee-receiver



stafihubd tx bridge set-relay-fee 1 1000000ufis --from admin --keyring-backend file --chain-id local-stafihub

stafihubd query bridge  relay-fee 1


stafihubd tx bridge add-banned-denom 1 uratom --from admin --keyring-backend file --chain-id local-stafihub

stafihubd q bridge banned-denom-list
```

### migrate rtoken (after adding new rtoken step)

```bash
stafihubd tx ledger migrate-init uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 100000000 150000000 200000000 300000000 1.23 --from admin --keyring-backend file --chain-id local-stafihub

stafihubd query bank  total 

stafihubd query ledger exchange-rate uratom

stafihubd query ledger bond-pipeline uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75



stafihubd tx ledger migrate-unbondings uratom --unbondings ./unbondings_example.json --from admin --keyring-backend file --chain-id local-stafihub

stafihubd query ledger pool-unbond uratom cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 3



stafihubd tx bridge set-denom-type uratom  1 --from admin --keyring-backend file --chain-id local-stafihub

stafihubd query bridge denom-types
```

### rdex

```bash
stafihubd tx rdex create-pool 10ufis 20uratom --from admin --chain-id local-stafihub --keyring-backend file

stafihubd tx rdex add-provider stafi1qzt0qajzr9df3en5sk06xlk26n30003c8uhdkg --from admin --chain-id local-stafihub --keyring-backend file

stafihubd tx rdex add-liquidity  100ufis 200uratom --from admin --chain-id local-stafihub --keyring-backend file

stafihubd tx rdex remove-liquidity 10 5 1uratom 1ufis ufis --from admin --chain-id local-stafihub --keyring-backend file

stafihubd tx rdex swap 2ufis 1uratom  --from admin --chain-id local-stafihub --keyring-backend file
```

### mining

```bash
stafihubd tx mining add-mining-provider stafi1ychj8z22pw0ruc65mx8nvdn7ca9qylpkauetvx  --from admin --chain-id local-stafihub --keyring-backend file

stafihubd tx mining add-reward-token ufis 200 --from admin --chain-id local-stafihub --keyring-backend file


stafihubd tx mining add-stake-pool ufis ./add_stake_pool_example.json  --from relay1 --chain-id local-stafihub --keyring-backend file

stafihubd tx mining stake 0 10ufis 0 --from my-account --chain-id local-stafihub --keyring-backend file 

stafihubd tx mining claim-reward 0 0 --from my-account --chain-id local-stafihub --keyring-backend file

stafihubd tx mining add-reward 1 0 300 0 0 --from relay1 --chain-id local-stafihub --keyring-backend file

stafihubd tx mining withdraw 1 10ufis 0 --from test --chain-id local-stafihub --keyring-backend file
```



## Operate examples for user

### liquidity bond (gaiad example)

```bash
gaiad tx bank send userAccount cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 1000000stake --note 1:stafi1ukq4mtq604prn5yxul7syh5ysvj0w5jrclvrvc --keyring-backend file --chain-id local-cosmos
```

### recover (gaiad example)

```bash
gaiad tx bank send userAccount cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 1stake --note 2:stafi1ukq4mtq604prn5yxul7syh5ysvj0w5jrclvrvc:9A80F3E6A007E1144BE34F4A0AC35B9288C19641BCAD3464277168000AF5FC66 --keyring-backend file --chain-id local-cosmos
```

### liquidity unbond

```bash
stafihubd tx ledger liquidity-unbond cosmos13jd2vn5wt8h6slj0gcv05lasgpkwpm26n04y75 100uratom cosmos1j9dues7ey2a39nes4ewfvyma96d3f5zrdhnfan --keyring-backend file --from user --home /Users/tpkeeper/gowork/stafi/rtoken-relay-core/keys/stafihub --chain-id local-stafihub
```

### deposit (transfer token to external chain)

```bash
stafihubd tx bridge deposit 1 uratom 800 dccf954570063847d73746afa0b0878f2c779d42089c5d9a107f2aca176e985f --from my-account --chain-id local-stafihub --keyring-backend file
```
