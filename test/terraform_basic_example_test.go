package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformBasicExampleDefaultValues(t *testing.T) {
	//t.Parallel()

	expectedText := "default-value"
	expectedProvider := map[string]string{"profile":"my-default-profile", "region" : "us-west-2"}

	terraformOptions := &terraform.Options{
		TerraformDir: "../modules/terraform-basic-example",
		Vars: map[string]interface{}{
			"example": expectedText,
			"provider_set": expectedProvider,
		},
	}
	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	actualTextExample := terraform.Output(t, terraformOptions, "example")
	actualExampleProvider := terraform.OutputMap(t, terraformOptions, "provider_output")

	assert.Equal(t, expectedText, actualTextExample)
	assert.Equal(t, expectedProvider, actualExampleProvider)
}

func TestTerraformBasicExampleDevValues(t *testing.T) {
	//t.Parallel()

	expectedText := "dev-environment"
	expectedProvider := map[string]string{"profile":"dev-profile", "region" : "us-west-2"}

	terraformOptions := &terraform.Options{
		TerraformDir: "../modules/terraform-basic-example",
		Vars: map[string]interface{}{
			"example": expectedText,
			"provider_set": expectedProvider,
		},
		VarFiles: []string{"varfile.tfvars"},
	}
	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	actualTextExample := terraform.Output(t, terraformOptions, "example")
	actualExampleProvider := terraform.OutputMap(t, terraformOptions, "provider_output")


	assert.Equal(t, expectedText, actualTextExample)
	assert.Equal(t, expectedProvider, actualExampleProvider)
}

func TestTerraformBasicExampleProdValues(t *testing.T) {
	//t.Parallel()

	expectedText := "prod-environment"
	expectedProvider := map[string]string{"profile":"prod-profile", "region" : "us-east-1"}

	terraformOptions := &terraform.Options{
		TerraformDir: "../modules/terraform-basic-example",
		Vars: map[string]interface{}{
			"example": expectedText,
			"provider_set": expectedProvider,
		},
		VarFiles: []string{"varfile-prod.tfvars"},
	}
	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	actualTextExample := terraform.Output(t, terraformOptions, "example")
	actualExampleProvider := terraform.OutputMap(t, terraformOptions, "provider_output")


	assert.Equal(t, expectedText, actualTextExample)
	assert.Equal(t, expectedProvider, actualExampleProvider)
}
