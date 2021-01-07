terraform {
  required_version = ">= 0.12.26"
}

resource "aws_servicequotas_service_quota" "increase_quotas" {
  for_each = var.resources_to_increase

  quota_code = local.codes[each.key].quota_code
  service_code = local.codes[each.key].service_code
  value = each.value
}

locals {
  codes = {
    nacl_rules = {
      quota_code = "L-2AEEBF1A"
      service_code = "vpc"
    }
  }
}
