provider "aws" {
	profile = "tiago-melo-developer"
	region = "us-east-1"
}

resource "aws_s3_bucket" "my-bucket" {
  bucket = "tiago-melo-terraform-s3"
  acl    = "private"

    tags  = {
    Name = "Tiago Melo S3 bucket",
    owner = "Tiago Melo"
    purpose = "Check Tags"
  }
}

resource "aws_instance" "example" {
  ami           = "ami-2757f631"
  instance_type = "t2.micro"
  tags  = {
    Name = "Tiago Melo EC2 instance",
    owner = "Tiago Melo"
    purpose = "Check Tags"
  }
}
	