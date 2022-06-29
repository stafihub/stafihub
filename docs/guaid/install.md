# Install

This guide will explain how to install the `stafihubd` entrypoint
onto your system. With these installed on a server, you can participate in the
mainnet as either a [Full Node](./join-mainnet.md) or a
[Validator](./validator-setup.md).

## Install build requirements

Install `make` and `gcc`.

On Ubuntu this can be done with the following:

```bash
sudo apt-get update

sudo apt-get install -y make gcc
```

## Install Go

Install `go` by following the [official docs](https://golang.org/doc/install).
Remember to set your `$PATH` environment variable, for example:

```bash
mkdir -p $HOME/go/bin
echo "export PATH=$PATH:$(go env GOPATH)/bin" >> ~/.bash_profile
source ~/.bash_profile
```

::: tip
**Go 1.17+** or later is required for the Cosmos SDK.
:::

## Install the binaries

Next, let's install the latest version of stafihubd. Make sure you `git checkout` the
correct [released version](https://github.com/stafihub/stafihub/releases).

```bash
git clone -b <latest-release-tag> https://github.com/stafihub/stafihub
cd stafihub && make install
```
