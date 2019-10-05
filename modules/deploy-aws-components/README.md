# deploy-aws-components

Module that deploy single ec2 instance and a S3 bucket:

## To deploy

```bash
terraform init
terraform apply
```

## To destroy

```bash
terraform destroy
```

## To run tests:

```bash
cd ../test/
dep ensure
go test
```
