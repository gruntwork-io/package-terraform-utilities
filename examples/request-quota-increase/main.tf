terraform {
  # This module is now only being tested with Terraform 0.14.x. However, to make upgrading easier, we are setting
  # 0.12.26 as the minimum version, as that version added support for required_providers with source URLs, making it
  # forwards compatible with 0.14.x code.
  required_version = ">= 0.12.26"
}

provider "aws" {
  region = var.aws_region
}

module "quota-increase" {
  source = "../../modules/request-quota-increase"

  resources_to_increase = {
    # In this example, to avoid opening a new request every time we run an automated test, we are setting the quotas
    # to their default values. In the real world, you'd want to set these quotes to higher values.
    vpc_active_vpc_peering_connections_per_vpc       = 50,
    vpc_egress_only_internet_gateways_per_region     = 10,
    vpc_gateway_vpc_endpoints_per_region             = 20,
    vpc_ipv4_cidr_blocks_per_vpc                     = 5,
    vpc_inbound_or_outbound_rules_per_security_group = 60,
    vpc_interface_vpc_endpoints_per_vpc              = 50,
    vpc_internet_gateways_per_region                 = 10,
    vpc_nat_gateways_per_availability_zone           = 35,
    vpc_network_acls_per_vpc                         = 200,
    vpc_network_interfaces_per_region                = 5000,
    vpc_outstanding_vpc_peering_connection_requests  = 25,
    vpc_participant_accounts_per_vpc                 = 100,
    vpc_route_tables_per_vpc                         = 200,
    vpc_routes_per_route_table                       = 50,
    vpc_rules_per_network_acl                        = 40,
    vpc_security_groups_per_network_interface        = 5,
    vpc_subnets_per_vpc                              = 200,
    vpc_subnets_that_can_be_shared_with_an_account   = 100,
    vpc_vpc_security_groups_per_region               = 2500,
    vpc_vpcs_per_region                              = 10,

    # IAM resources are only available in us-east-1, but it's a Global quota.
    iam_customer_managed_policies_per_account = 1500,
    iam_groups_per_account                    = 300,
    iam_instance_profiles_per_account         = 1000,
    iam_managed_policies_per_role             = 10,
    iam_managed_policies_per_user             = 10,
    iam_role_trust_policy_length              = 2048,
    iam_roles_per_account                     = 1000,
    iam_server_certificates_per_account       = 20
  }
}
