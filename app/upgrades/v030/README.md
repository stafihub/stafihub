# StaFiHub v0.3.0 Upgrade (Ledger support)

The Upgrade is scheduled for block `2626590`. A countdown clock is [here](https://www.mintscan.io/stafi/blocks/2626590)

This guide assumes that you use cosmovisor to manage upgrades.

## Changelog

1. Ledger support ([a1dd186](https://github.com/stafihub/stafihub/commit/a1dd1863f558f7f811708b1646b1bfd78f6adc06), [f139982](https://github.com/stafihub/stafihub/commit/f1399828f39069f00a1922c2308865bed69bafac)). 
2. SDK bump up ([bc57f81](https://github.com/stafihub/stafihub/commit/bc57f81496bab9b1bc13934c2e2e070ee6487a78)). cosmos-sdk (0.45.9 -> 0.45.11), ibc-go (3.1.1 to 3.4.0).
3. Lint warns fix ([90f4029](https://github.com/stafihub/stafihub/commit/90f40296e84a921a5cc51a44c8b5ed5fb559c6fc)).

## Install

```bash
cd stafihub
git pull
git checkout v0.3.0
make install
```

## Check the version

```bash
# should be 0.3.0
stafihubd version
# Should be commit fd0c05441ce8e6bddd681fa6cb4090697db84c53
stafihubd version --long
```

## Make new directory and copy binary

```bash
mkdir -p $HOME/.stafihub/cosmovisor/upgrades/v030/bin
cp $HOME/go/bin/stafihubd $HOME/.stafihub/cosmovisor/upgrades/v030/bin
```

## Check the version again

```bash
# should be 0.3.0
$HOME/.stafihub/cosmovisor/upgrades/v030/bin/stafihubd version
```






