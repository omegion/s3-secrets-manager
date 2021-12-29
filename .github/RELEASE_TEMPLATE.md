## Installation

You can use `go` to build S3 Secrets Manager locally with:

```shell
go get -u github.com/omegion/s3-secrets-manager
```

Or, you can use the usual commands to install or upgrade:

On OS X

```shell
$ curl -L https://github.com/omegion/s3-secrets-manager/releases/download/{{.Env.VERSION}}/s3sm-darwin-amd64 >/usr/local/bin/s3sm && \
  chmod +x /usr/local/bin/s3sm
```

On Linux

```shell
$ curl -L https://github.com/omegion/s3-secrets-manager/releases/download/{{.Env.VERSION}}/s3sm-linux-amd64 >/usr/local/bin/s3sm && \
    chmod +x /tmp/s3sm && \
    sudo cp /tmp/s3sm /usr/local/bin/s3sm
```

Otherwise, download one of the releases from the [release page](https://github.com/omegion/s3-secrets-manager/releases/)
directly.

See the install [docs](https://s3-secrets-manager.omegion.dev) for more install options and instructions.

## Changelog

