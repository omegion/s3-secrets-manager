locals {
  role_name = "s3-secrets-manager-${var.bucket_name}"
}

data "aws_caller_identity" "current" {}

data "aws_iam_policy_document" "bucket_assume_role_policy" {
  statement {
    principals {
      type        = "Service"
      identifiers = ["s3.amazonaws.com"]
    }

    principals {
      type        = "AWS"
      identifiers = [data.aws_caller_identity.current.account_id]
    }

    effect = "Allow"

    actions = [
      "sts:AssumeRole",
    ]
  }
}

data "aws_iam_policy_document" "bucket_policy" {
  statement {
    effect = "Allow"

    actions = [
      "s3:ListBucket",
      "s3:GetBucketLocation",
      "s3:ListBucketVersions",
      "s3:GetBucketVersioning"
    ]

    resources = [
      aws_s3_bucket.this.arn,
    ]
  }

  statement {
    effect = "Allow"

    actions = [
      "s3:*"
    ]

    resources = [
      "${aws_s3_bucket.this.arn}/*",
    ]
  }
}

resource "aws_iam_role" "this" {
  name               = local.role_name
  assume_role_policy = data.aws_iam_policy_document.bucket_assume_role_policy.json
  tags               = merge(
  var.tags,
  {
    Name = local.role_name
  },
  )
}

resource "aws_iam_policy" "this" {
  name   = local.role_name
  policy = data.aws_iam_policy_document.bucket_policy.json
  tags   = merge(
  var.tags,
  {
    Name = local.role_name
  },
  )
}

resource "aws_iam_policy_attachment" "this" {
  name       = local.role_name
  roles      = [aws_iam_role.this.name]
  policy_arn = aws_iam_policy.this.arn
}
