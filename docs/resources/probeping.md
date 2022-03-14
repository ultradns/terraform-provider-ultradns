---
subcategory: "PING-Probe"
layout: "ultradns"
page_title: "ULTRADNS: ultradns_probe_ping"
description: |-
  Manages the PING probe records in UltraDNS.
---

# Resource: ultradns_probe_ping

Use this resource to manage the PING probe records in UltraDNS.

## Example Usage

### Create PING Probe for SB Pool

```terraform
resource "ultradns_probe_ping" "ping_sb" {
	zone_name = "example.com."
	owner_name = "sb.example.com."
	interval = "HALF_MINUTE"
	agents = ["EUROPE_WEST", "SOUTH_AMERICA", "PALO_ALTO", "NEW_YORK"]
	threshold = 4
	packets = 6
	packet_size = 106
	loss_percent_limit{
		fail = 11
	}
	total_limit{
		fail = 11
	}
	run_limit{
		fail = 11
	}
}
```

### Create PING Probe for TC Pool

```terraform
resource "ultradns_probe_ping" "ping_tc" {
	zone_name = "example.com."
	owner_name = "tc.example.com."
	interval = "HALF_MINUTE"
	agents = ["EUROPE_WEST", "SOUTH_AMERICA", "PALO_ALTO", "NEW_YORK"]
	threshold = 4
	packets = 6
	packet_size = 106
	loss_percent_limit{
		warning = 6 
		critical = 9
		fail = 11
	}
	total_limit{
		warning = 6 
		critical = 9
		fail = 11
	}
	average_limit{
		warning = 6 
		critical = 9
		fail = 11
	}
	run_limit{
		warning = 6 
		critical = 9
		fail = 11
	}
	avg_run_limit{
		warning = 6 
		critical = 9
		fail = 11
	}
}
```

## Argument Reference

The following arguments are supported:

* `zone_name` - (Required) (String) Name of the zone.
* `owner_name` - (Required) (String) The domain name of the owner of the RRSet. Can be either a fully qualified domain name (FQDN) or a relative domain name. If provided as a FQDN, it must be contained within the zone name's FQDN.
* `interval` - (Optional) (String) Length of time between probes in minutes. Valid values are `HALF_MINUTE`, `ONE_MINUTE`, `TWO_MINUTES`, `FIVE_MINUTES`, `TEN_MINUTES`, and `FIFTEEN_MINUTES`.</br>Default value set to `FIVE_MINUTES`.
* `agents` - (Required) (String List) Locations that will be used for probing. One or more values must be specified.
Valid values are `ASIA`, `CHINA`, `EUROPE_EAST`, `EUROPE_WEST`, `NORTH_AMERICA_CENTRAL`, `NORTH_AMERICA_EAST`, `NORTH_AMERICA_WEST`, `SOUTH_AMERICA`, `NEW_YORK`, `PALO_ALTO`, `DALLAS`, and `AMSTERDAM`.
* `threshold` - (Required) (Integer) Number of agents that must agree for a probe state to be changed.
* `pool_record` - (Optional) (String) The pool record associated with this probe. Specified when creating a record-level probe.
* `packets` - (Optional) (String) Number of ICMP packets to send. Default value set to 3.
* `packet_size` - (Optional) (String) Size of packets in bytes. Default value set to 56.
* `loss_percent_limit` - (Optional) (Block Set, Max:1) Nested block describing the percentage of packets lost will be acceptable and beyond that it will generates warning or failure. The structure of this block follows the same structure as the <a href="#nested-limit-block-has-the-following-structure">`limit`</a> block described below.
* `total_limit` - (Optional) (Block Set, Max:1) Nested block describing how long the probe should run in total for all pings. The structure of this block follows the same structure as the <a href="#nested-limit-block-has-the-following-structure">`limit`</a> block described below.
* `average_limit` - (Optional) (Block Set, Max:1) Nested block describing the mean (average) time to connect for the five most recent probes that have run on each agent. This is only used for Traffic Controller pools. The structure of this block follows the same structure as the <a href="#nested-limit-block-has-the-following-structure">`limit`</a> block described below.
* `run_limit` - (Optional) (Block Set, Max:1) Nested block describing how long the probe should run. The structure of this block follows the same structure as the <a href="#nested-limit-block-has-the-following-structure">`limit`</a> block described below.
* `avg_run_limit` - (Optional) (Block Set, Max:1) Nested block describing the mean (average) run-time for the five most recent probes that have run on each agent. This is only used for Traffic Controller pools. The structure of this block follows the same structure as the <a href="#nested-limit-block-has-the-following-structure">`limit`</a> block described below.

### Nested `limit` block has the following structure:

* `warning` - (Optional) (Integer) Indicates how long/percent the PING Probe should wait before a warning is generated.
* `critical` - (Optional) (Integer) Indicates how long/percent the PING  Probe should wait before a critical warning is generated.
* `fail` - (Optional) (Integer) Indicates how long/percent the PING Probe should wait before it make the probe to fail.

-> `warning` and `critical` are only used for Traffic Controller pools


## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `guid` - (Computed) (String) The internal id for this probe.


## Import

PING probe records can be imported by combining their `owner_name`, `zone_name`, `record_type`, and `guid`, separated by a colon.<br/>
Example : `www.example.com.:example.com.:A (1):06084A729D56C85C`.


-> For import, the `owner_name` and `zone_name` must be a FQDN, and `record_type` should have the type, as well as the corresponding number, as shown in the example below.

Example:
```terraform
$ terraform import ultradns_probe_ping.example "www.example.com.:example.com.:A (1):06084A729D56C85C" 
```