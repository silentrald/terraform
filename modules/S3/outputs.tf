output "bucket_id" {
  value = aws_s3_bucket.bucket-resource.id
}

output "content" {
  value = local.content
}