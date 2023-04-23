# StaFiHub v0.4.0 Upgrade

The Upgrade is scheduled for block `3766566`. A countdown clock is [here](https://www.mintscan.io/stafi/blocks/3766566)

This guide assumes that you use cosmovisor to manage upgrades.

## Changelog

1. SDK bump up ([08fd361](https://github.com/stafihub/stafihub/commit/08fd36175df8249421b56f95d948e289f500f41a)). cosmos-sdk (0.45.11 -> 0.46.12), ibc-go (3.4.0 to 5.2.0).
2. Fix duplicate events issue ([93de08f](https://github.com/stafihub/stafihub/commit/93de08f2172b2ee9f25f57ab63689c85ef450878)).

## Install

```bash
cd stafihub
git pull
git checkout v0.4.0
make install
```

## Check the version

```bash
# should be 0.4.0
stafihubd version
# Should be commit d989bf1c03397b33750ce6db1cd94a1ca3042d02
stafihubd version --long
```

## Make new directory and copy binary

```bash
mkdir -p $HOME/.stafihub/cosmovisor/upgrades/v040/bin
cp $HOME/go/bin/stafihubd $HOME/.stafihub/cosmovisor/upgrades/v040/bin
```

## Check the version again

```bash
# should be 0.4.0
$HOME/.stafihub/cosmovisor/upgrades/v040/bin/stafihubd version
```
