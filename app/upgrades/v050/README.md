# StaFiHub v0.5.0 Upgrade

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
# Should be commit fbeedb2441f614c9936e63a39cdec29abc70503f
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
