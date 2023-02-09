---
subcategory: "TCP-Probe"
layout: "ultradns"
page_title: "ULTRADNS: ultradns_probe_tcp"
description: |-
  Manages the TCP probe records in UltraDNS.
---

# Resource: ultradns_probe_tcp

Use this resource to manage the TCP probe records in UltraDNS.

## Example Usage

### Create TCP Probe for SB Pool

```terraform
resource "ultradns_probe_tcp" "tcp_sb" {
	zone_name = "example.com."
	owner_name = "sb.example.com."
	interval = "HALF_MINUTE"
	agents = ["EUROPE_WEST", "SOUTH_AMERICA", "PALO_ALTO", "NEW_YORK"]
	threshold = 4
	port = 443
	control_ip = "192.168.1.1"
	connect_limit{
		fail = 11
	}
}
```

### Create TCP Probe for TC Pool

```terraform
resource "ultradns_probe_tcp" "tcp_tc" {
	zone_name = "example.com."
	owner_name = "tc.example.com."
	interval = "HALF_MINUTE"
	agents = ["EUROPE_WEST", "SOUTH_AMERICA", "PALO_ALTO", "NEW_YORK"]
	threshold = 4
	port = 6
	control_ip = "192.168.1.1"
	connect_limit{
		warning = 6
		critical = 9
		fail = 11
	}
	avg_connect_limit{
		warning = 5
		critical = 8
		fail = 10
	}
}
```

## Argument Reference

The following arguments are supported:

* `zone_name` - (Required) (String) Name of the zone.
* `owner_name` - (Required) (String) The domain name of the owner of the RRSet. Can be either a fully qualified domain name (FQDN) or a relative domain name. If provided as a FQDN, it must be contained within the zone name's FQDN.
* `interval` - (Optional) (String) Length of time between probes in minutes. Valid values are `HALF_MINUTE`, `ONE_MINUTE`, `TWO_MINUTES`, `FIVE_MINUTES`, `TEN_MINUTES`, and `FIFTEEN_MINUTES`.</br>Default value set to `FIVE_MINUTES`.
* `agents` - (Required) (String List) Locations that will be used for probing. Multiple values can be comma separated. Valid values are:  `ASIA`, `CHINA`, `EUROPE_EAST`, `EUROPE_WEST`, `NORTH_AMERICA_CENTRAL`, `NORTH_AMERICA_EAST`, `NORTH_AMERICA_WEST`, `SOUTH_AMERICA`, `NEW_YORK`, `PALO_ALTO`, `DALLAS`, and `AMSTERDAM`.
* `threshold` - (Required) (Integer) Number of agents that must agree for a probe state to be changed.
* `pool_record` - (Optional) (String) The pool record associated with this probe. Specified when creating a record-level probe.
* `port` - (Optional) (Integer) TCP port number to connect to. Default value set to 80.
* `control_ip` - (Optional) (String) The target IP address of the TCP probe.
* `connect_limit` - (Required) (Block Set, Max:1) Nested block describing how long the probe stays connected to the resource. The structure of this block follows the same structure as the [`limit`](#nested-limit-block-has-the-following-structure) block described below.

### Nested `limit` block has the following structure:

* `warning` - (Optional) (Integer) Indicates how long (in seconds, or by percentage value) the TCP Probe should wait, before a warning is generated.
* `critical` - (Optional) (Integer) Indicates how long (in seconds, or by percentage value) the TCP  Probe should wait, before a critical warning is generated.
* `fail` - (Optional) (Integer) Indicates how long (in seconds, or by percentage value) the TCP Probe should wait, before causing the probe to fail.

-> `warning` and `critical` are only used for Traffic Controller pools.


## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `guid` - (Computed) (String) The internal id for this probe.


## Import

TCP probe records can be imported by combining their `owner_name`, `zone_name`, `record_type`, and `guid`, separated by a colon.<br/>
Example : `www.example.com.:example.com.:A (1):06084A729D56C85C`.


-> For import, the `owner_name` and `zone_name` must be a FQDN, and `record_type` should have the type, as well as the corresponding number, as shown in the example below.

Example:
```terraform
$ terraform import ultradns_probe_tcp.example "www.example.com.:example.com.:A (1):06084A729D56C85D"
```
