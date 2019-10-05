package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformBasicExampleDefaultValues(t *testing.T) {
	//t.Parallel()

	expectedText := "example"
	expectedList := []string{expectedText}
	expectedMap := map[string]string{"expected": expectedText}
	expectedProvider := map[string]string{"profile":"my-default-profile", "region" : "us-west-2"}

	terraformOptions := &terraform.Options{
		TerraformDir: "../modules/terraform-basic-example",
		Vars: map[string]interface{}{
			"example": expectedText,
			"example_list": expectedList,
			"example_map":  expectedMap,
			"provider_set": expectedProvider,
		},
	}
	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	actualTextExample := terraform.Output(t, terraformOptions, "example")
	actualTextExample2 := terraform.Output(t, terraformOptions, "example2")
	actualExampleList := terraform.OutputList(t, terraformOptions, "example_list")
	actualExampleMap := terraform.OutputMap(t, terraformOptions, "example_map")
	actualExampleProvider := terraform.OutputMap(t, terraformOptions, "provider_output")


	assert.Equal(t, expectedText, actualTextExample)
	assert.Equal(t, "", actualTextExample2)
	assert.Equal(t, expectedList, actualExampleList)
	assert.Equal(t, expectedMap, actualExampleMap)
	assert.Equal(t, expectedProvider, actualExampleProvider)
}

func TestTerraformBasicExampleDevValues(t *testing.T) {
	//t.Parallel()

	expectedText := "example"
	expectedList := []string{expectedText}
	expectedMap := map[string]string{"expected": expectedText}
	expectedProvider := map[string]string{"profile":"dev-profile", "region" : "us-west-2"}

	terraformOptions := &terraform.Options{
		TerraformDir: "../modules/terraform-basic-example",
		Vars: map[string]interface{}{
			"example": expectedText,
			"example_list": expectedList,
			"example_map":  expectedMap,
			"provider_set": expectedProvider,
		},
		VarFiles: []string{"varfile.tfvars"},
	}
	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	actualTextExample := terraform.Output(t, terraformOptions, "example")
	actualTextExample2 := terraform.Output(t, terraformOptions, "example2")
	actualExampleList := terraform.OutputList(t, terraformOptions, "example_list")
	actualExampleMap := terraform.OutputMap(t, terraformOptions, "example_map")
	actualExampleProvider := terraform.OutputMap(t, terraformOptions, "provider_output")


	assert.Equal(t, expectedText, actualTextExample)
	assert.Equal(t, "test", actualTextExample2)
	assert.Equal(t, expectedList, actualExampleList)
	assert.Equal(t, expectedMap, actualExampleMap)
	assert.Equal(t, expectedProvider, actualExampleProvider)
}

func TestTerraformBasicExampleProdValues(t *testing.T) {
	//t.Parallel()

	expectedText := "example"
	expectedList := []string{expectedText}
	expectedMap := map[string]string{"expected": expectedText}
	expectedProvider := map[string]string{"profile":"prod-profile", "region" : "us-east-1"}

	terraformOptions := &terraform.Options{
		TerraformDir: "../modules/terraform-basic-example",
		Vars: map[string]interface{}{
			"example": expectedText,
			"example_list": expectedList,
			"example_map":  expectedMap,
			"provider_set": expectedProvider,
		},
		VarFiles: []string{"varfile-prod.tfvars"},
	}
	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	actualTextExample := terraform.Output(t, terraformOptions, "example")
	actualTextExample2 := terraform.Output(t, terraformOptions, "example2")
	actualExampleList := terraform.OutputList(t, terraformOptions, "example_list")
	actualExampleMap := terraform.OutputMap(t, terraformOptions, "example_map")
	actualExampleProvider := terraform.OutputMap(t, terraformOptions, "provider_output")


	assert.Equal(t, expectedText, actualTextExample)
	assert.Equal(t, "prod", actualTextExample2)
	assert.Equal(t, expectedList, actualExampleList)
	assert.Equal(t, expectedMap, actualExampleMap)
	assert.Equal(t, expectedProvider, actualExampleProvider)
}
