## Installation

You can use `go` to build S3 Secrets Manager locally with:

```shell
go get -u github.com/omegion/s3-secrets-manager
```

Or, you can use the usual commands to install or upgrade:

On OS X
```console
$ curl -L https://github.com/docker/machine/releases/download/{{.Env.VERSION}}/docker-machine-`uname -s`-`uname -m` >/usr/local/bin/docker-machine && \
  chmod +x /usr/local/bin/docker-machine
```
On Linux
```console
$ curl -L https://github.com/docker/machine/releases/download/{{.Env.VERSION}}/docker-machine-`uname -s`-`uname -m` >/tmp/docker-machine &&
    chmod +x /tmp/docker-machine &&
    sudo cp /tmp/docker-machine /usr/local/bin/docker-machine
```

Otherwise, download one of the releases from the [release page](https://github.com/docker/machine/releases/) directly.

See the install [docs](https://docs.docker.com/machine/install-machine/) for more install options and instructions.

## Changelog

{{.Env.CHANGELOG}}