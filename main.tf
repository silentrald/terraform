module "s3" {
    source = "./modules/S3"
    bucket_name = "${var.BUCKET_NAME}"
}