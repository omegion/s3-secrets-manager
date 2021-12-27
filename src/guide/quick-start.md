# Quick Start

It is assumed that you have AWS S3 bucket is created also you have AWS credentials are set.

## Set a Secret

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
s3sm secret get --field password --value MYSUPERSECRET --path  secret/database/my-db
```

Let's check `~/.ssh/keys` folder if our keys are added.

```shell
❯ ls -l ~/.ssh/keys/
-rw-------  1 X  staff   432 Mar 30 11:05 test
-rw-------  1 X  staff   112 Mar 30 11:05 test.pub
```

## Session Duration

After a login with Bitwarden CLI tool, it will return a `session key` that you will need to define it as environment
variable. Otherwise it will keep asking you to enter your credentials all the time. You can read for more info
at [Bitwarden documentation](https://bitwarden.com/help/article/cli/#environment-variable)
.