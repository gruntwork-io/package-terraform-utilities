# Request AWS Quota Increase

This module can be used to request a quota increase for an AWS Resource.

## Features

- Request a quota increase for VPC and IAM resources.

## Learn

### Core Concepts

- [AWS Service Quotas Documentation](https://docs.aws.amazon.com/servicequotas/?id=docs_gateway)


### Example code

See the [request-quota-increase example](/examples/request-quota-increase) for working sample code.



## Usage

Use the module in your Terraform code, replacing `<VERSION>` with the latest version from the [releases
page](https://github.com/gruntwork-io/terraform-aws-utilities/releases):

```hcl
module "request-quota-increase" {
  source = "git::git@github.com:gruntwork-io/terraform-aws-utilities.git//modules/request-quota-increase?ref=<VERSION>"

    request_quota_increase = {
      vpc_nat_gateways_per_availability_zone = 40,
      vpc_rules_per_network_acl              = 40,
    }
}
```

The only required argument is:

* `request_quota_increase`: A map with the desired resource and the new quota. The current supported services are VPC and IAM. Feel free to contribute to this module to add support for more services.


When you run `apply`, the `new_quotas` output variable will confirm to you that a quota request has been made!

```hcl
new_quotas = {
  "vpc_internet_gateways_per_region" = {
    "adjustable" = true
    "arn" = "arn:aws:servicequotas:us-east-1:<account-id>:vpc/L-FE5A380F"
    "default_value" = 5
    "id" = "vpc/L-FE5A380F"
    "quota_code" = "L-FE5A380F"
    "quota_name" = "NAT gateways per Availability Zone"
    "request_id" = "<request_id>"
    "request_status" = "PENDING"
    "service_code" = "vpc"
    "service_name" = "Amazon Virtual Private Cloud (Amazon VPC)"
    "value" = 40
  }
}
```

## Manage

You can see a full history of quota request changes using the [AWS
Console](https://console.aws.amazon.com/servicequotas/home#!/requests) or the AWS CLI:


```
aws service-quotas list-requested-service-quota-change-history --region <REGION>
```

### Request a new quota smaller than the current one

If the new value that you request is smaller than the current one, _nothing_ will happen. The
`terraform apply` output will contain the current quota. For example, if the NAT Gateway current
quota is 30 and you ask for a new quota of 25, this is the output:

```hcl
new_quotas = {
  "vpc_internet_gateways_per_region" = {
    "adjustable" = true
      "arn" = "arn:aws:servicequotas:us-east-1:<account-id>:vpc/L-FE5A380F"
      "default_value" = 5
      "id" = "vpc/L-FE5A380F"
      "quota_code" = "L-FE5A380F"
      "quota_name" = "NAT gateways per Availability Zone"
      "service_code" = "vpc"
      "service_name" = "Amazon Virtual Private Cloud (Amazon VPC)"
      "value" = 30   <------ Returned the current quota, not the requested one.
  }
}
```


### What happens when you run `destroy`


When you run `terraform destroy` on this module, it does not affect your current quotas or your
existing quota requests. In other words, you don't have to worry about quotas being reset to old
values; once they have been increased, they stay that way!

## How do I request quota increases for services this module doesn't support yet?

In order to request a Quota increase, AWS requires you to pass in special codes, like `L-2AEEBF1A`
and `L-2AEEBF1A`. Terraform has a data source [aws_servicequotas_service_quota](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/servicequotas_service_quota)
that returns this special code using the `quota_name` and `service_code` as parameters. Many of this
`quota_name` and `service_code` combinations are already defined in this module. To add others,
the [generate_code.rb](generate_code.rb) script can be used to generate the code for a given service.


You can check the available services with

```
aws service-quotas list-services --region <REGION> --output table
```

And use the [generate_code.rb](generate_code.rb) script to generate the necessary code to support
more resources.

Be aware that not all services are available in all regions, therefore if you are adding a resource
that is not available in certain region, it is necessary to filter it in the `locals.resources_code`
at [main.tf](main.tf)