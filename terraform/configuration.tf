variable "AWS_ACCESS_KEY" {
  description = "AWS_ACCESS_KEY"
  type        = string
  sensitive   = true
}

variable "AWS_SECRET_KEY" {
  description = "AWS_SECRET_KEY"
  type        = string
  sensitive   = true
}

variable "REGION" {
  default = "ap-northeast-1"
}

variable "BUCKET_NAME" {
  default = "sample"
}
