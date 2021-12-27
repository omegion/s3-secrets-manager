# Get Started

## Prerequisites

S3 Secrets Manager uses a bucket to store your secrets. In order to use the tool, you will need AWS S3 bucket created. You can use Terraform folder to create one for yourself.

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

Otherwise, download one of the releases from the [release page](https://github.com/omegion/s3-secrets-manager/releases/) directly.

------------
Let's verify that the binary has installed successfully.

```shell
❯ s3-secrets-manager version
INFO[27-12-2021 16:50:32] s3sm v0.3.0
```
