variable "bucket_name" {}

variable "acl_value" {
    default = "private"
}

locals {
    content = "${timestamp()}"
}
