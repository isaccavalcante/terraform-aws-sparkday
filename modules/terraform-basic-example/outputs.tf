output "example" {
  value = "${data.template_file.example.rendered}"
}

output "provider_output" {
    value = "${var.provider_set}"
}