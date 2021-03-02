package test

import (
	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
	"testing"
)

type (
	QuotaAndServiceName struct {
		QuotaName   string `json:"quota_name"`
		ServiceName string `json:"service_name"`
	}
	QuotaIncreaseOutput struct {
		VpcActiveVpcPeeringConnectionsPerVpc       QuotaAndServiceName `json:"vpc_active_vpc_peering_connections_per_vpc"`
		VpcEgressOnlyInternetGatewaysPerRegion     QuotaAndServiceName `json:"vpc_egress-only_internet_gateways_per_region"`
		VpcGatewayVpcEndpointsPerRegion            QuotaAndServiceName `json:"vpc_gateway_vpc_endpoints_per_region"`
		VpcIpv4CidrBlocksPerVpc                    QuotaAndServiceName `json:"vpc_ipv4_cidr_blocks_per_vpc"`
		VpcInboundOrOutboundRulesPerSecurityGroup  QuotaAndServiceName `json:"vpc_inbound_or_outbound_rules_per_security_group"`
		VpcInterfaceVpcEndpointsPerVpc             QuotaAndServiceName `json:"vpc_interface_vpc_endpoints_per_vpc"`
		VpcInternetGatewaysPerRegion               QuotaAndServiceName `json:"vpc_internet_gateways_per_region"`
		VpcNatGatewaysPerAvailabilityZone          QuotaAndServiceName `json:"vpc_nat_gateways_per_availability_zone"`
		VpcNetworkAclsPerVpc                       QuotaAndServiceName `json:"vpc_network_acls_per_vpc"`
		VpcNetworkInterfacesPerRegion              QuotaAndServiceName `json:"vpc_network_interfaces_per_region"`
		VpcOutstandingVpcPeeringConnectionRequests QuotaAndServiceName `json:"vpc_outstanding_vpc_peering_connection_requests"`
		VpcParticipantAccountsPerVpc               QuotaAndServiceName `json:"vpc_participant_accounts_per_vpc"`
		VpcRouteTablesPerVpc                       QuotaAndServiceName `json:"vpc_route_tables_per_vpc"`
		VpcRoutesPerRouteTable                     QuotaAndServiceName `json:"vpc_routes_per_route_table"`
		VpcRulesPerNetworkAcl                      QuotaAndServiceName `json:"vpc_rules_per_network_acl"`
		VpcSecurityGroupsPerNetworkInterface       QuotaAndServiceName `json:"vpc_security_groups_per_network_interface"`
		VpcSubnetsPerVpc                           QuotaAndServiceName `json:"vpc_subnets_per_vpc"`
		VpcSubnetsThatCanBeSharedWithAnAccount     QuotaAndServiceName `json:"vpc_subnets_that_can_be_shared_with_an_account"`
		VpcVpcSecurityGroupsPerRegion              QuotaAndServiceName `json:"vpc_vpc_security_groups_per_region"`
		VpcVpcsPerRegion                           QuotaAndServiceName `json:"vpc_vpcs_per_region"`

		IamCustomerManagedPoliciesPerAccount QuotaAndServiceName `json:"iam_customer_managed_policies_per_account"`
		IamGroupsPerAccount                  QuotaAndServiceName `json:"iam_groups_per_account"`
		IamInstanceProfilesPerAccount        QuotaAndServiceName `json:"iam_instance_profiles_per_account"`
		IamManagedPoliciesPerRole            QuotaAndServiceName `json:"iam_managed_policies_per_role"`
		IamManagedPoliciesPerUser            QuotaAndServiceName `json:"iam_managed_policies_per_user"`
		IamRoleTrustPolicyLength             QuotaAndServiceName `json:"iam_role_trust_policy_length"`
		IamRolesPerAccount                   QuotaAndServiceName `json:"iam_roles_per_account"`
		IamServerCertificatesPerAccount      QuotaAndServiceName `json:"iam_server_certificates_per_account"`
	}
)

func TestRequestQuotaIncrease(t *testing.T) {
	t.Parallel()

	awsRegion := aws.GetRandomRegion(t, nil, nil)

	terraformOptions := &terraform.Options{
		TerraformDir: "../examples/request-quota-increase",
		Vars: map[string]interface{}{
			"aws_region": awsRegion,
		},
	}

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	output := QuotaIncreaseOutput{}
	terraform.OutputStruct(t, terraformOptions, "new_quotas", &output)

	assert.Equal(t, output.VpcActiveVpcPeeringConnectionsPerVpc.QuotaName, "Active VPC peering connections per VPC")
	assert.Equal(t, output.VpcActiveVpcPeeringConnectionsPerVpc.ServiceName, "Amazon Virtual Private Cloud (Amazon VPC)")

	assert.Equal(t, output.VpcEgressOnlyInternetGatewaysPerRegion.QuotaName, "Egress-only internet gateways per Region")
	assert.Equal(t, output.VpcEgressOnlyInternetGatewaysPerRegion.ServiceName, "Amazon Virtual Private Cloud (Amazon VPC)")

	assert.Equal(t, output.VpcGatewayVpcEndpointsPerRegion.QuotaName, "Gateway VPC endpoints per Region")
	assert.Equal(t, output.VpcGatewayVpcEndpointsPerRegion.ServiceName, "Amazon Virtual Private Cloud (Amazon VPC)")

	assert.Equal(t, output.VpcIpv4CidrBlocksPerVpc.QuotaName, "IPv4 CIDR blocks per VPC")
	assert.Equal(t, output.VpcIpv4CidrBlocksPerVpc.ServiceName, "Amazon Virtual Private Cloud (Amazon VPC)")

	assert.Equal(t, output.VpcInboundOrOutboundRulesPerSecurityGroup.QuotaName, "Inbound or outbound rules per security group")
	assert.Equal(t, output.VpcInboundOrOutboundRulesPerSecurityGroup.ServiceName, "Amazon Virtual Private Cloud (Amazon VPC)")

	assert.Equal(t, output.VpcInterfaceVpcEndpointsPerVpc.QuotaName, "Interface VPC endpoints per VPC")
	assert.Equal(t, output.VpcInterfaceVpcEndpointsPerVpc.ServiceName, "Amazon Virtual Private Cloud (Amazon VPC)")

	assert.Equal(t, output.VpcInternetGatewaysPerRegion.QuotaName, "Internet gateways per Region")
	assert.Equal(t, output.VpcInternetGatewaysPerRegion.ServiceName, "Amazon Virtual Private Cloud (Amazon VPC)")

	assert.Equal(t, output.VpcNatGatewaysPerAvailabilityZone.QuotaName, "NAT gateways per Availability Zone")
	assert.Equal(t, output.VpcNatGatewaysPerAvailabilityZone.ServiceName, "Amazon Virtual Private Cloud (Amazon VPC)")

	assert.Equal(t, output.VpcNetworkAclsPerVpc.QuotaName, "Network ACLs per VPC")
	assert.Equal(t, output.VpcNetworkAclsPerVpc.ServiceName, "Amazon Virtual Private Cloud (Amazon VPC)")

	assert.Equal(t, output.VpcNetworkInterfacesPerRegion.QuotaName, "Network interfaces per Region")
	assert.Equal(t, output.VpcNetworkInterfacesPerRegion.ServiceName, "Amazon Virtual Private Cloud (Amazon VPC)")

	assert.Equal(t, output.VpcOutstandingVpcPeeringConnectionRequests.QuotaName, "Outstanding VPC peering connection requests")
	assert.Equal(t, output.VpcOutstandingVpcPeeringConnectionRequests.ServiceName, "Amazon Virtual Private Cloud (Amazon VPC)")

	assert.Equal(t, output.VpcParticipantAccountsPerVpc.QuotaName, "Participant accounts per VPC")
	assert.Equal(t, output.VpcParticipantAccountsPerVpc.ServiceName, "Amazon Virtual Private Cloud (Amazon VPC)")

	assert.Equal(t, output.VpcRouteTablesPerVpc.QuotaName, "Route tables per VPC")
	assert.Equal(t, output.VpcRouteTablesPerVpc.ServiceName, "Amazon Virtual Private Cloud (Amazon VPC)")

	assert.Equal(t, output.VpcRoutesPerRouteTable.QuotaName, "Routes per route table")
	assert.Equal(t, output.VpcRoutesPerRouteTable.ServiceName, "Amazon Virtual Private Cloud (Amazon VPC)")

	assert.Equal(t, output.VpcRulesPerNetworkAcl.QuotaName, "Rules per network ACL")
	assert.Equal(t, output.VpcRulesPerNetworkAcl.ServiceName, "Amazon Virtual Private Cloud (Amazon VPC)")

	assert.Equal(t, output.VpcSecurityGroupsPerNetworkInterface.QuotaName, "Security groups per network interface")
	assert.Equal(t, output.VpcSecurityGroupsPerNetworkInterface.ServiceName, "Amazon Virtual Private Cloud (Amazon VPC)")

	assert.Equal(t, output.VpcSubnetsPerVpc.QuotaName, "Subnets per VPC")
	assert.Equal(t, output.VpcSubnetsPerVpc.ServiceName, "Amazon Virtual Private Cloud (Amazon VPC)")

	assert.Equal(t, output.VpcSubnetsThatCanBeSharedWithAnAccount.QuotaName, "Subnets that can be shared with an account")
	assert.Equal(t, output.VpcSubnetsThatCanBeSharedWithAnAccount.ServiceName, "Amazon Virtual Private Cloud (Amazon VPC)")

	assert.Equal(t, output.VpcVpcSecurityGroupsPerRegion.QuotaName, "VPC security groups per Region")
	assert.Equal(t, output.VpcVpcSecurityGroupsPerRegion.ServiceName, "Amazon Virtual Private Cloud (Amazon VPC)")

	assert.Equal(t, output.VpcVpcsPerRegion.QuotaName, "VPCs per Region")
	assert.Equal(t, output.VpcVpcsPerRegion.ServiceName, "Amazon Virtual Private Cloud (Amazon VPC)")

	assert.Equal(t, output.IamCustomerManagedPoliciesPerAccount.QuotaName, "Customer managed policies per account")
	assert.Equal(t, output.IamCustomerManagedPoliciesPerAccount.ServiceName, "AWS Identity and Access Management (IAM)")

	assert.Equal(t, output.IamGroupsPerAccount.QuotaName, "Groups per account")
	assert.Equal(t, output.IamGroupsPerAccount.ServiceName, "AWS Identity and Access Management (IAM)")

	assert.Equal(t, output.IamInstanceProfilesPerAccount.QuotaName, "Instance profiles per account")
	assert.Equal(t, output.IamInstanceProfilesPerAccount.ServiceName, "AWS Identity and Access Management (IAM)")

	assert.Equal(t, output.IamManagedPoliciesPerRole.QuotaName, "Managed policies per role")
	assert.Equal(t, output.IamManagedPoliciesPerRole.ServiceName, "AWS Identity and Access Management (IAM)")

	assert.Equal(t, output.IamManagedPoliciesPerUser.QuotaName, "Managed policies per user")
	assert.Equal(t, output.IamManagedPoliciesPerUser.ServiceName, "AWS Identity and Access Management (IAM)")

	assert.Equal(t, output.IamRoleTrustPolicyLength.QuotaName, "Role trust policy length")
	assert.Equal(t, output.IamRoleTrustPolicyLength.ServiceName, "AWS Identity and Access Management (IAM)")

	assert.Equal(t, output.IamRolesPerAccount.QuotaName, "Roles per account")
	assert.Equal(t, output.IamRolesPerAccount.ServiceName, "AWS Identity and Access Management (IAM)")

	assert.Equal(t, output.IamServerCertificatesPerAccount.QuotaName, "Server certificates per account")
	assert.Equal(t, output.IamServerCertificatesPerAccount.ServiceName, "AWS Identity and Access Management (IAM)")
}
