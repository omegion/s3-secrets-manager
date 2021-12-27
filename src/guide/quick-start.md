# Quick Start

It is assumed that you have AWS S3 bucket is created also you have AWS credentials are set.

## Set Secret

Let's create a secret into S3. You can either define environment variable
`S3SM_BUCKET`, or set `--bucket` flag.

```shell
export S3SM_BUCKET=my-s3-bucket
s3sm secret set --name password --value MYSUPERSECRET --path  secret/database/my-db
```

Output:

```shell
Key       Value
----      ----
password  MYSUPERSECRET
```

## Get Secret

Once we put a secret to S3. Let's get it back.

```shell
export S3SM_BUCKET=my-s3-bucket
s3sm secret get --path  secret/database/my-db
```

Output:

```shell
Key       Value
----      ----
host      example.com
password  MYSUPERSECRET
username  root
```

You can get only a field with `--field` flag.

```shell
export S3SM_BUCKET=my-s3-bucket
s3sm secret get --field host --path  secret/database/my-db
```

Output:

```shell
example.com
```

## List Secret

You can list all secrets with given path.

```shell
export S3SM_BUCKET=my-s3-bucket
s3sm secret list --path  secret
```

Output:

```shell
Secret                                Last Modified
----                                  ----
secret/aws/rds-1                      2021-12-27 22:19:48 +0000 UTC
secret/database/my-db                 2021-12-27 22:15:54 +0000 UTC
secret/github/token/read-only-access  2021-12-27 22:29:42 +0000 UTC
secret/github/token/write-read-access  2021-12-27 22:30:35 +0000 UTC
```

You can narrow down the secret list by defining specific path:

```shell
export S3SM_BUCKET=my-s3-bucket
s3sm secret list --path  secret/github
```

Output:

```shell
Secret                                 Last Modified
----                                   ----
secret/github/token/read-only-access   2021-12-27 22:29:42 +0000 UTC
secret/github/token/write-read-access  2021-12-27 22:30:35 +0000 UTC
```