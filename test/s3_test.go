package test

import (
	"fmt"
	"strings"
	"testing"
	"os"

	"github.com/joho/godotenv"
	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// An example of how to test the Terraform module in examples/terraform-aws-s3-example using Terratest.
func TestTerraformAwsS3(t *testing.T) {
	t.Parallel()

	// Load env
	godotenv.Load(".env")

	// Give this S3 Bucket a unique ID for a name tag so we can distinguish it from any other Buckets provisioned
	// in your AWS account
	expectedName := fmt.Sprintf("terratest-aws-s3-example-%s", strings.ToLower(random.UniqueId()))

	// Pick a random AWS region to test in. This helps ensure your code works in all regions.
	awsRegion := aws.GetRandomStableRegion(t, nil, nil)

	// Construct the terraform options with default retryable errors to handle the most common retryable errors in
	// terraform testing.
	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../",

		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"AWS_ACCESS_KEY": os.Getenv("AWS_ACCESS_KEY"),
			"AWS_SECRET_KEY": os.Getenv("AWS_SECRET_KEY"),
			"REGION": awsRegion,
			"BUCKET_NAME": expectedName,
		},
	})

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the value of an output variable
	bucketId := terraform.Output(t, terraformOptions, "bucket_id")
	timestamp := terraform.Output(t, terraformOptions, "timestamp")

	// Check if the bucket exist
	aws.AssertS3BucketExists(t, awsRegion, bucketId)
	t.Logf("Bucket Id: %s", bucketId)

	// Check for the files if they exist
	text1 := aws.GetS3ObjectContents(t, awsRegion, bucketId, "text1.txt")
	text2 := aws.GetS3ObjectContents(t, awsRegion, bucketId, "text2.txt")
	
	assert.NotEmpty(t, text1, "text1.txt not sent to the bucket")
	assert.NotEmpty(t, text2, "text2.txt not sent to the bucket")
	t.Logf("text1.txt: %s", text1)
	t.Logf("text2.txt: %s", text2)

	// Check if the contents are the same with the timestamp log
	assert.Equal(t, text1, timestamp)
	assert.Equal(t, text2, timestamp)
}
