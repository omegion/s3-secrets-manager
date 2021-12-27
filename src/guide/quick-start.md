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

## List Secret Versions

You can list the version of the secret if you enabled the versioning of your bucket. Each change will create a version
in the bucket that you can historically access them.

```shell
export S3SM_BUCKET=my-s3-bucket
s3sm secret versions --path secret/github/token/read-only-access
```

Output:

```shell
Order  Version ID                        Last Modified
----   ----                              ----
1      1mAx5J0P90m0mfFo_BnS8k9MUAHnGR5F  2021-12-27 22:35:37 +0000 UTC
2      Qm.a5lfhayBPO.btC8hkGF26dCGxWwaw  2021-12-27 22:35:34 +0000 UTC
3      PUV4vNVKq2NtP8JQ27kBADJKe79xDBvK  2021-12-27 22:35:29 +0000 UTC
4      VzRM7.138B0gt_vEtt77ST40WVyWXPzP  2021-12-27 22:35:27 +0000 UTC
5      OkRfb88ojsm2b_WViVH8PedmxpE4LxYU  2021-12-27 22:29:42 +0000 UTC
```

## Get Secret by Version

Once bucket versioning enabled, every change will be stored in a version. You can get old secret values by `version-id`.

```shell
export S3SM_BUCKET=my-s3-bucket
s3sm secret get --path secret/github/token/read-only-access --version-id <VERSION_ID>
```

Output:

```shell
Key    Value
----   ----
token  MYSECRETTOKEN2
```