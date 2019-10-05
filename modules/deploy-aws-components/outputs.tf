########################### S3 OUTPUTS ###########################

output "bucket_name" {
  value = "${aws_s3_bucket.my-bucket.bucket}"
}

output "bucket_arn" {
  value = "${aws_s3_bucket.my-bucket.arn}"
}

output "bucket_tags" {
  value = "${aws_s3_bucket.my-bucket.tags}"
}

########################### EC2 OUTPUTS ###########################

output "instance_arn" {
  value = "${aws_instance.example.arn}"
}

output "instance_private_id"{
    value = "${aws_instance.example.id}"
}

output "instance_private_ip"{
    value = "${aws_instance.example.private_ip}"
}

output "instance_private_dns"{
    value = "${aws_instance.example.private_dns}"
}

output "instance_security_groups"{
    value = "${aws_instance.example.security_groups}"
}

output "instance_tags" {
  value = "${aws_instance.example.tags}"
}

output "instance_type" {
  value = "${aws_instance.example.instance_type}"
}

output "instance_ami" {
  value = "${aws_instance.example.ami}"
}
