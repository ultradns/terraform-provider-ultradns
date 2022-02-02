# Terraform configuration for UltraDNS provider
terraform {
    required_providers {
        ultradns = {
            source = "ultradns.com/ultradns/ultradns"
        }
    }
}

# UltraDNS provider configuration
provider "ultradns" {
   username = "${var.ultradns_username}"
   password = "${var.ultradns_password}"
   hosturl = "${var.ultradns_host_url}"
}
