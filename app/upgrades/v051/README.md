# StaFiHub v0.5.1 Upgrade

The Upgrade is scheduled for block ``. A countdown clock is [here](https://www.mintscan.io/stafi/blocks/)

This guide assumes that you use cosmovisor to manage upgrades.

## Changelog

([v0.5.0...v0.5.1](https://github.com/stafihub/stafihub/compare/v0.5.0...v0.5.1))

1. Add more encoding methods
2. fix pool status when reopen channel

## Install

```bash
cd stafihub
git pull
git checkout v0.5.1
make install
```

## Check the version

```bash
# should be 0.5.1
stafihubd version
# Should be commit 
stafihubd version --long
```

## Make new directory and copy binary

```bash
mkdir -p $HOME/.stafihub/cosmovisor/upgrades/v051/bin
cp $HOME/go/bin/stafihubd $HOME/.stafihub/cosmovisor/upgrades/v051/bin
```

## Check the version again

```bash
# should be 0.5.1
$HOME/.stafihub/cosmovisor/upgrades/v051/bin/stafihubd version
```
