---
subcategory: "IP"
layout: "ultradns"
page_title: "ULTRADNS: ultradns_dirgroup_ip"
description: |-
  Get information of Account-level source IP Groups in UltraDNS.
---

# Data Source: ultradns_dirgroup_ip

Use this data source to get detailed information for your zones.

## Example Usage

```terraform
data "ultradns_dirgroup_ip" "lo" {
    name = "loopback"
}
```


## Argument Reference

The following arguments are supported:

* `name` - (Required) (String) Name of the sourceIP group.


## Attributes Reference

In addition to the arguments above, the following attributes are exported:

* `account_name` - (Computed) (String) Name of the account.
* `ip` - (Computed) (Block Set) List of nested blocks describing the IP addresses and IP ranges this SourceIP group contains. The structure of this block is described below.
* `description` - (Computed) (String) The description for the group.

### Nested `ip` block has the following structure:

* `start` - (Computed) (String) The starting IP address (IPv4 or IPv6).
* `end` - (Computed) (String) The ending IP address (IPv4 or IPv6).
* `cidr` - (Computed) (String) The CIDR format (IPv4 or IPv6) for an IP address range. 
* `address` - (Computed) (String) A single IPv4 or IPv6 address.
