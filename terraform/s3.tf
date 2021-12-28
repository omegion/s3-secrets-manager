resource "aws_s3_bucket" "this" {
  bucket = var.bucket_name
  acl    = "private"

  versioning {
    enabled = var.enable_versioning
  }

  tags = merge(
  var.tags,
  {
    Name = var.bucket_name
  },
  )
}