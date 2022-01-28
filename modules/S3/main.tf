resource "random_id" "id" {
    byte_length = 8
}

resource "aws_s3_bucket" "bucket-resource" {
    bucket = "${var.bucket_name}-${random_id.id.hex}"
    acl = "${var.acl_value}"
}

resource "aws_s3_bucket_object" "text1" {
    bucket = aws_s3_bucket.bucket-resource.id
    key    = "text1.txt"
    acl    = "private"
    
    content = local.content
}

resource "aws_s3_bucket_object" "text2" {
    bucket = aws_s3_bucket.bucket-resource.id
    key    = "text2.txt"
    acl    = "private"

    content = local.content
}