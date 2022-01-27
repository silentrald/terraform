resource "random_id" "id" {
    byte_length = 8
}

resource "aws_s3_bucket" "bucket-resource" {
    bucket = "${var.bucket_name}-${random_id.id.hex}"
    acl = "${var.acl_value}"
}

resource "local_file" "text1" {
    content = "${timestamp()}"
    filename = "text1.txt"
}

resource "local_file" "text2" {
    content = "${timestamp()}"
    filename = "text2.txt"
}

resource "aws_s3_bucket_object" "text1-send" {
    bucket = aws_s3_bucket.bucket-resource.id
    key    = local_file.text1.filename
    acl    = "private"
    source = local_file.text1.filename
}

resource "aws_s3_bucket_object" "text2-send" {
    bucket = aws_s3_bucket.bucket-resource.id
    key    = local_file.text2.filename
    acl    = "private"
    source = local_file.text2.filename
}