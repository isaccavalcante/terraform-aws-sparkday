# terraform-variables-output-testing

Module with two variables and outputs:
- example: Single text content
- provider-set: map with a provider name and aws-region

It has two terraform variale files: `varfile-dev.tvars` and `varfile-prod.tvars` used on `terraform-variables-output-testing_test.go`.

To run tests:

```bash
cd ../test/
dep ensure
go test
```
