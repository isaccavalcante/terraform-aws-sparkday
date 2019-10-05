package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestS3Variables(t *testing.T) {
	t.Parallel()

	expectedBucketName := "tiago-melo-terraform-s3"
	expectedBucketARN := "arn:aws:s3:::tiago-melo-terraform-s3"
	expectedBucketTAGS := map[string]string{"Name":"Tiago Melo S3 bucket", "owner" : "Tiago Melo", "purpose" : "Check Tags"}

	terraformOptions := &terraform.Options{
		TerraformDir: "../modules/deploy-aws-components",
	}
	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	actualBucketName := terraform.Output(t, terraformOptions, "bucket_name")
	actualBucketARN := terraform.Output(t, terraformOptions, "bucket_arn")
	actualBucketTags := terraform.OutputMap(t, terraformOptions, "bucket_tags")

	assert.Equal(t, expectedBucketName, actualBucketName)
	assert.Equal(t, expectedBucketARN, actualBucketARN)
	assert.Equal(t, expectedBucketTAGS, actualBucketTags)
}

func TestEC2Variables(t *testing.T) {
	t.Parallel()

	expectedInstanceType := "t2.micro"
	expectedInstanceAMI := "ami-2757f631"
	expectedInstanceTAGS := map[string]string{"Name":"Tiago Melo EC2 instance", "owner" : "Tiago Melo", "purpose" : "Check Tags"}

	terraformOptions := &terraform.Options{
		TerraformDir: "../modules/deploy-aws-components",
	}
	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	actualInstanceType := terraform.Output(t, terraformOptions, "instance_type")
	actualInstanceAMI := terraform.Output(t, terraformOptions, "instance_ami")
	actualInstanceTags := terraform.OutputMap(t, terraformOptions, "instance_tags")

	assert.Equal(t, expectedInstanceType, actualInstanceType)
	assert.Equal(t, expectedInstanceAMI, actualInstanceAMI)
	assert.Equal(t, expectedInstanceTAGS, actualInstanceTags)
}
