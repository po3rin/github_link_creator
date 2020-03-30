resource "aws_s3_bucket" "bucket" {
  bucket = "github-link-card-lambda-function"
  acl    = "private"

  versioning {
    enabled = true
  }

  # 暗号化を有効
  server_side_encryption_configuration {
    rule {
      apply_server_side_encryption_by_default {
        sse_algorithm = "AES256"
      }
    }
  }

  tags = {
    Name        = "ghcard"
  }
}

resource "aws_s3_bucket_public_access_block" "private" {
  bucket                  = aws_s3_bucket.bucket.id
  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}
