---
subcategory: "DNS-Probe"
layout: "ultradns"
page_title: "ULTRADNS: ultradns_probe_dns"
description: |-
  Manages the DNS probe records in UltraDNS.
---

# Resource: ultradns_probe_dns

Use this resource to manage the DNS probe records in UltraDNS.

## Example Usage

### Create DNS Probe for SB Pool

```terraform
resource "ultradns_probe_dns" "dns_sb" {
	zone_name = "example.com."
	owner_name = "sb.example.com."
	interval = "HALF_MINUTE"
	agents = ["EUROPE_WEST", "SOUTH_AMERICA", "PALO_ALTO", "NEW_YORK"]
	threshold = 4
	port = 43
	tcp_only = false
	type = "SOA"
	query_name = "${resource.ultradns_sbpool.sbpoola.zone_name}"
	response{
		fail = "fail"
	}
	run_limit{
		fail = 12
	}
}
```

### Create DNS Probe for TC Pool

```terraform
resource "ultradns_probe_dns" "dns_tc" {
	zone_name = "example.com."
	owner_name = "tc.example.com."
	interval = "HALF_MINUTE"
	agents = ["EUROPE_WEST", "SOUTH_AMERICA", "PALO_ALTO", "NEW_YORK"]
	threshold = 4
	port = 43
	tcp_only = false
	type = "SOA"
	query_name = "${resource.ultradns_tcpool.tcpoola.zone_name}"
	response{
		warning = "warning" 
		critical = "critical"
		fail = "fail"
	}
	run_limit{
		warning = 7 
		critical = 10
		fail = 12
	}
	avg_run_limit{
		warning = 7 
		critical = 10
		fail = 12
	}
}
```

## Argument Reference

The following arguments are supported:

* `zone_name` - (Required) (String) Name of the zone.
* `owner_name` - (Required) (String) The domain name of the owner of the RRSet. Can be either a fully qualified domain name (FQDN) or a relative domain name. If provided as a FQDN, it must be contained within the zone name's FQDN.
* `interval` - (Optional) (String) Length of time between probes in minutes. Valid values are `HALF_MINUTE`, `ONE_MINUTE`, `TWO_MINUTES`, `FIVE_MINUTES`, `TEN_MINUTES`, and `FIFTEEN_MINUTES`.</br>Default value set to `FIVE_MINUTES`.
* `agents` - (Required) (String List) Locations that will be used for probing. Multiple values can be comma separated. Valid values are: `ASIA`, `CHINA`, `EUROPE_EAST`, `EUROPE_WEST`, `NORTH_AMERICA_CENTRAL`, `NORTH_AMERICA_EAST`, `NORTH_AMERICA_WEST`, `SOUTH_AMERICA`, `NEW_YORK`, `PALO_ALTO`, `DALLAS`, and `AMSTERDAM`.
* `threshold` - (Required) (Integer) Number of agents that must agree for a probe state to be changed.
* `pool_record` - (Optional) (String) The pool record associated with this probe. Specified when creating a record-level probe.
* `port` - (Optional) (String) The Port that should be used for DNS lookup. Default value set to 53.
* `type` - (Optional) (String) Select the record type that the probe will check for. Valid values are `NULL`, `AXFR`, or any Resource Record Type. Default value set to `NULL`.
* `query_name` - (Optional) (String) The name that should be queried.
* `tcp_only` - (Optional) (Boolean) Indicates whether or not the probe should use TCP only, or first UDP then TCP. Default value set to false.
* `response` - (Optional) (Block Set, Max:1) Nested block describing the strings to match the response that will generate a warning or failure. The structure of this block follows the same structure as the <a href="#nested-limit-block-has-the-following-structure">`limit`</a> block described below.
* `run_limit` - (Optional) (Block Set, Max:1) Nested block describing how long the probe should run. The structure of this block follows the same structure as the <a href="#nested-limit-block-has-the-following-structure">`limit`</a> block described below.
* `avg_run_limit` - (Optional) (Block Set, Max:1) Nested block describing the mean (average) run-time for the five most recent probes that have run on each agent. This is only used for Traffic Controller pools. The structure of this block follows the same structure as the <a href="#nested-limit-block-has-the-following-structure">`limit`</a> block described below.

### Nested `limit` block has the following structure:

* `warning` - (Optional) (Integer) Indicates how long (in seconds) the DNS Probe should wait, before a warning is generated.
* `critical` - (Optional) (Integer) Indicates how long (in seconds) the DNS  Probe should wait, before a critical warning is generated.
* `fail` - (Optional) (Integer) Indicates how long (in seconds) the DNS Probe should wait, before causing the probe to fail.

### Nested `response` block has the following structure:

* `warning` - (Optional) (String) Will match exactly for records containing a single field response (i.e., A, CNAME, DNAME, NS, MB, MD, MF, MG, MR, PTR), and partial matches for record types with multiple field response. Fields separated by spaces will be combined with the matches to trigger a warning.
* `critical` - (Optional) (String) Will match exactly for records containing a single field response (i.e., A, CNAME, DNAME, NS, MB, MD, MF, MG, MR, PTR), and partial matches for record types with multiple field response. Fields separated by spaces will be combined with the matches to trigger a critical warning.
* `fail` - (Optional) (String) Will match exactly for records containing a single field response (i.e., A, CNAME, DNAME, NS, MB, MD, MF, MG, MR, PTR), and partial matches for record types with multiple field response. Fields separated by spaces will be combined with the matches to trigger a failure.

-> `warning` and `critical` are only used for Traffic Controller pools.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `guid` - (Computed) (String) The internal id for this probe.


## Import

DNS probe records can be imported by combining their `owner_name`, `zone_name`, `record_type`, and `guid`, separated by a colon.<br/>
Example : `www.example.com.:example.com.:A (1):06084A729D56C85C`.


-> For import, the `owner_name` and `zone_name` must be a FQDN, and `record_type` should have the type, as well as the corresponding number, as shown in the example below.

Example:
```terraform
$ terraform import ultradns_probe_dns.example "www.example.com.:example.com.:A (1):06084A729D56C85C" 
```