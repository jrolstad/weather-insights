resource "aws_s3_bucket" "observation_data" {
  bucket = local.observation_bucket_name

}

resource "aws_s3_bucket_acl" "observation_data" {
  bucket = aws_s3_bucket.observation_data.id
  acl    = "private"
}

resource "aws_s3_bucket_public_access_block" "observation_data" {
  bucket = aws_s3_bucket.observation_data.id

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

resource "aws_s3_bucket" "athena_resources" {
  bucket = local.athena_bucket_name
}

resource "aws_s3_bucket_acl" "athena_resources" {
  bucket = aws_s3_bucket.athena_resources.id
  acl    = "private"
}

resource "aws_s3_bucket_public_access_block" "athena_resources" {
  bucket = aws_s3_bucket.athena_resources.id

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}