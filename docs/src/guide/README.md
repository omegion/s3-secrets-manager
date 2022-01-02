# Introduction

AWS S3 is widely used storage system secured very well. S3 Secrets Manager uses S3 to store your secrets in a structured way that you can access anywhere any time.

You can set your S3 bucket to store your objects with KSM key to add extra security. In addition, you can create IAM role to limit the tools operations.

```shell
S3 Secrets Management for AWS S3.

Usage:
  s3sm [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  secret      Secret operations.
  version     Print the version/build number

Flags:
      --config string      config file (default is $HOME/.s3sm/config.yaml)
  -h, --help               help for s3sm
      --interactive        Set the interactivity (default true)
      --logFormat string   Set the logging format. One of: text|json (default "text") (default "text")
      --logLevel string    Set the logging level. One of: debug|info|warn|error (default "info")

Use "s3sm [command] --help" for more information about a command.
```
