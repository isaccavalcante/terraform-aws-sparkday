# ---------------------------------------------------------------------------------------------------------------------
# BASIC TERRAFORM EXAMPLE
# See test/terraform_aws_example.go for how to write automated tests for this code.
# ---------------------------------------------------------------------------------------------------------------------

provider "aws" {
  profile = "${var.provider_set.profile}"
  region = "${var.provider_set.region}"
}

data "template_file" "example" {
  template = "${var.example}"
}
