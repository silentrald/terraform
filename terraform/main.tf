provider "aws" {
    access_key = "${var.AWS_ACCESS_KEY}"
    secret_key = "${var.AWS_SECRET_KEY}"
    region = "${var.REGION}"
}

module "s3" {
    source = "./S3"
    bucket_name = "${var.BUCKET_NAME}"
}