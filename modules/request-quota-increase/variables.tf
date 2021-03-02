variable "resources_to_increase" {
  description = "A map where the key is the resource and the value is the desired quota. The only services supported at the moment are VPC and IAM. You can also use the `aws_servicequotas_service_quota` resource directly!"
  type        = map(number)

  # Example:
  # {
  # vpc_ipv4_cidr_blocks_per_vpc     = 10,
  # vpc_internet_gateways_per_region = 15,
  # }
}
