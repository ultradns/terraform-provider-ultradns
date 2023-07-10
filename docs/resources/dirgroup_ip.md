---
subcategory: "SourceIP"
layout: "ultradns"
page_title: "ULTRADNS: ultradns_dirgroup_ip"
description: |-
  Manages the account-level source IP groups in UltraDNS.
---

# Resource: ultradns_dirgroup_ip

Use this resource to manage account-level source IP groups in UltraDNS

## Example Usage

### Create Account-level sourceIP group

```terraform
resource "ultradns_dirgroup_ip" "lo" {
    name         = "loopback"
    account_name = "my_account"
	ip{
		address = "200.1.1.1"
	}
	ip{
		start = "200.1.1.2"
		end = "200.1.1.5"
	}
		ip{
			cidr = "200.20.20.0/24"
	}
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) (String) Name of the geoIP group.
* `account_name` - (Required) (String) 	Name of the account. It must be provided, but it can also be sourced from the `ULTRADNS_ACCOUNT` environment variable.
* `ip` - (Required) (Block Set) List of nested blocks describing the IP addresses and IP ranges this SourceIP group contains. The structure of this block is described below.
* `description` - (Optional) (String) </br>


### Nested `ip` block has the following structure:

* `start` - (Optional) (String) The starting IP address (IPv4 or IPv6). If the start value is present, the end value must be present as well. `cidr` and `address` cannot be present.
* `end` - (Optional) (String) The ending IP address (IPv4 or IPv6). If the end value is present, the start value must be present as well. `cidr` and `address` cannot be present.
* `cidr` - (Optional) (String) The CIDR format (IPv4 or IPv6) for an IP address range. If cidr is present, the `start`, `end`, and `address` cannot be present.
* `address` - (Optional) (String) A single IPv4 or IPv6 address. If address is present, the `start`, `end`, and `cidr` cannot be present.
## Import

GeoIP group can be imported by combining their `name` and `account_name`.<br/>
Example: `loopback:my_account`

Example:
```
$ terraform import ultradns_dirgrup_ip.lo "loopback:my_account"
```


