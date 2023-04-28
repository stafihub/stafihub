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
cd $HOME
wget -O go1.20.3.linux-amd64.tar.gz https://go.dev/dl/go1.20.3.linux-amd64.tar.gz
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.20.3.linux-amd64.tar.gz && rm go1.20.3.linux-amd64.tar.gz
echo 'export GOROOT=/usr/local/go' >> $HOME/.bashrc
echo 'export GOPATH=$HOME/go' >> $HOME/.bashrc
echo 'export GO111MODULE=on' >> $HOME/.bashrc
echo 'export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin' >> $HOME/.bashrc && . $HOME/.bashrc
go version
```

::: tip
**Go 1.18+** or later is required for the Cosmos SDK.
:::

## Install the binaries

Next, let's install the latest version of stafihubd. Make sure you `git checkout` the
correct [released version](https://github.com/stafihub/stafihub/releases).

```bash
git clone -b <latest-release-tag> https://github.com/stafihub/stafihub
cd stafihub && make install
```
