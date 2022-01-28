# Terraform and AWS S3 w/ Terratest Automation

Requirements:

- Terraform
- Golang (for Terratest)

## Terraform

The terraform scripts will create a bucket containing 2 files, `text1.txt` and `text2.txt`,
which both contain a timestamp of their creation.

```bash
path/to/terraform $ terraform init
path/to/terraform $ terraform plan -no-color -var AWS_ACCESS_KEY="<access_key>" -var AWS_SECRET_KEY="<secret_key>"
path/to/terraform $ terraform apply -auto-approve -var AWS_ACCESS_KEY="<access_key>" -var AWS_SECRET_KEY="<secret_key>"
```

To destroy the created files (both on local and on the bucket)

```bash
terraform destroy -auto-approve -var AWS_ACCESS_KEY="<access_key>" -var AWS_SECRET_KEY="<secret_key>"
```

## Terratest

```bash
path/to/terratest $ go mod tidy
path/to/terratest $ go test s3_test.go
```
