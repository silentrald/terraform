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

func TestTerraformAwsS3(t *testing.T) {
	t.Parallel()

	// Load env
	godotenv.Load(".env")

	expectedName := fmt.Sprintf("terratest-aws-s3-example-%s", strings.ToLower(random.UniqueId()))
	awsRegion := aws.GetRandomStableRegion(t, nil, nil)

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../",
		Vars: map[string]interface{}{
			"AWS_ACCESS_KEY": os.Getenv("AWS_ACCESS_KEY"),
			"AWS_SECRET_KEY": os.Getenv("AWS_SECRET_KEY"),
			"REGION": awsRegion,
			"BUCKET_NAME": expectedName,
		},
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	bucketId := terraform.Output(t, terraformOptions, "bucket_id")
	timestamp := terraform.Output(t, terraformOptions, "timestamp")

	// Check if the bucket exist
	aws.AssertS3BucketExists(t, awsRegion, bucketId)
	t.Logf("Bucket Id: %s", bucketId)

	// Get the files from the S3 bucket
	text1 := aws.GetS3ObjectContents(t, awsRegion, bucketId, "text1.txt")
	text2 := aws.GetS3ObjectContents(t, awsRegion, bucketId, "text2.txt")

	// Check if the contents are the same with the timestamp log
	assert.Equal(t, text1, timestamp)
	assert.Equal(t, text2, timestamp)
}
