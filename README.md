<h1 align="center">
Cheapest Secure Secret Management on AWS S3.
</h1>

<p align="center">
  <a href="https://omegion.dev" target="_blank">
    <img width="180" src="https://s3-secrets-manager.omegion.dev/img/logo.svg" alt="logo">
  </a>
</p>

<p align="center">
    <img src="https://img.shields.io/github/workflow/status/omegion/s3-secrets-manager/Test" alt="Test"></a>
    <img src="https://coveralls.io/repos/github/omegion/s3-secrets-manager/badge.svg?branch=master" alt="Coverall"></a>
    <img src="https://goreportcard.com/badge/github.com/omegion/s3-secrets-manager" alt="Report"></a>
    <a href="http://pkg.go.dev/github.com/omegion/s3-secrets-manager"><img src="https://img.shields.io/badge/pkg.go.dev-doc-blue" alt="Doc"></a>
    <a href="https://github.com/omegion/s3-secrets-manager/blob/master/LICENSE"><img src="https://img.shields.io/github/license/omegion/s3-secrets-manager" alt="License"></a>
</p>

```shell
S3 Secrets Management for AWS S3.

Usage:
  s3sm [command]

Available Commands:
  completion  generate the autocompletion script for the specified shell
  help        Help about any command
  secret      Adds two numbers
  version     Print the version/build number

Flags:
  -h, --help               help for s3sm
      --interactive        Set the interactivity (default true)
      --logFormat string   Set the logging format. One of: text|json (default "text") (default "text")
      --logLevel string    Set the logging level. One of: debug|info|warn|error (default "info")

Use "s3sm [command] --help" for more information about a command.
```

## Requirements

* AWS CLI

## What does it do?

S3 Secrets Management CLI tool allows you to manage your secrets on S3 in cheaper way.

## Installation

You can use `go` to build S3 Secrets Manager locally with:

```shell
go get -u github.com/omegion/s3-secrets-manager
```

This will install `bw-ssh` binary to your `GOPATH`.

Or, you can use the usual commands to install or upgrade:

On OS X

```shell
VERSION=v0.3.0
$ curl -L https://github.com/omegion/s3-secrets-manager/releases/download/$VERSION/s3-secrets-manager-darwin-amd64 >/usr/local/bin/s3sm && \
  chmod +x /usr/local/bin/s3sm
```

On Linux

```shell
VERSION=v0.3.0
$ curl -L https://github.com/omegion/s3-secrets-manager/releases/download/$VERSION/s3-secrets-manager-linux/amd64 >/usr/local/bin/s3sm && \
    chmod +x /tmp/s3sm &&
    sudo cp /tmp/s3sm /usr/local/bin/s3sm
```

Otherwise, download one of the releases from the [release page](https://github.com/omegion/s3-secrets-manager/releases/)
directly.

## Improvements to be made

* 100% test coverage.
* Kubernetes Operator for sync.
* Better covering for other features.

