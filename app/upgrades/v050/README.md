# StaFiHub v0.5.0 Upgrade

The Upgrade is scheduled for block `5868372`. A countdown clock is [here](https://www.mintscan.io/stafi/blocks/5868372)

This guide assumes that you use cosmovisor to manage upgrades.

## Changelog

1. LSM bond support ([v0.4.3...v0.5.0](https://github.com/stafihub/stafihub/compare/v0.4.3...v0.5.0)).

## Install

```bash
cd stafihub
git pull
git checkout v0.5.0
make install
```

## Check the version

```bash
# should be 0.5.0
stafihubd version
# Should be commit 920922cd686a94664f736bd604093a305475d0dc
stafihubd version --long
```

## Make new directory and copy binary

```bash
mkdir -p $HOME/.stafihub/cosmovisor/upgrades/v050/bin
cp $HOME/go/bin/stafihubd $HOME/.stafihub/cosmovisor/upgrades/v050/bin
```

## Check the version again

```bash
# should be 0.5.0
$HOME/.stafihub/cosmovisor/upgrades/v050/bin/stafihubd version
```
