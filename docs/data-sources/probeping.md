---
subcategory: "PING-Probe"
layout: "ultradns"
page_title: "ULTRADNS: ultradns_probe_ping"
description: |-
  Get information of PING probe records in UltraDNS.
---

# Data Source: ultradns_probe_ping

Use this data source to get detailed information of PING probe records.

## Example Usage

### Get last created PING Probe

```terraform
data "ultradns_probe_ping" "ping" {
    zone_name = "example.com."
    owner_name = "www"
}
```

### Get PING Probe by guid

```terraform
data "ultradns_probe_ping" "ping" {
    zone_name = "example.com."
    owner_name = "www"
    guid = "06084A729D56C85C"
}
```

### Get PING Probe by filtering

```terraform
data "ultradns_probe_ping" "ping" {
    zone_name = "example.com."
    owner_name = "www"
    interval = "HALF_MINUTE"
	agents = [ "PALO_ALTO", "NEW_YORK"]
}
```

## Argument Reference

The following arguments are supported:

* `zone_name` - (Required) (String) Name of the zone.
* `owner_name` - (Required) (String) The domain name of the owner of the RRSet. Can be either a fully qualified domain name (FQDN) or a relative domain name. If provided as a FQDN, it must be contained within the zone name's FQDN.

The following arguments are used to filter the probes:

* `guid` - (Optional) (String) The internal id for this probe.

-> `guid` can be fetched using UltraDNS REST API.

* `interval` - (Optional) (String) Length of time between probes in minutes.
* `agents` - (Optional) (String List) Locations that will be used for probing. The exact list must be provided if probes need to be filtered using agents. Valid values are: `ASIA`, `CHINA`, `EUROPE_EAST`, `EUROPE_WEST`, `NORTH_AMERICA_CENTRAL`, `NORTH_AMERICA_EAST`, `NORTH_AMERICA_WEST`, `SOUTH_AMERICA`, `NEW_YORK`, `PALO_ALTO`, `DALLAS`, and `AMSTERDAM`.
* `threshold` - (Optional) (Integer) The probe threshold value.
* `pool_record` - (Optional) (String) The pool record associated with this probe.

->
1) If `guid` is provided, the probe with that guid is returned, and other filter options are not considered.</br>
2) If there is a conflict between probes due to filter options other than `guid`, the last created probe is returned.</br>
3) If no probe is found for the filter options, an error is returned.  

## Attributes Reference

In addition to all of the arguments above, the following attributes are exported:

* `packets` - (Computed) (String) Number of ICMP packets to send.
* `packet_size` - (Computed) (String) Size of packets in bytes.
* `loss_percent_limit` - (Computed) (Block Set, Max:1) Nested block describing the acceptable percentage of packets lost, which will in turn, generate either a warning or a failure. The structure of this block follows the same structure as the <a href="#nested-limit-block-has-the-following-structure">`limit`</a> block described below.
* `total_limit` - (Computed) (Block Set, Max:1) Nested block describing how long the probe should run in total for all pings. The structure of this block follows the same structure as the <a href="#nested-limit-block-has-the-following-structure">`limit`</a> block described below.
* `average_limit` - (Computed) (Block Set, Max:1) Nested block describing the mean (average) time to connect for the five most recent probes that have run on each agent. This is only used for Traffic Controller pools. The structure of this block follows the same structure as the <a href="#nested-limit-block-has-the-following-structure">`limit`</a> block described below.
* `run_limit` - (Computed) (Block Set, Max:1) Nested block describing how long the probe should run. The structure of this block follows the same structure as the <a href="#nested-limit-block-has-the-following-structure">`limit`</a> block described below.
* `avg_run_limit` - (Computed) (Block Set, Max:1) Nested block describing the mean (average) run-time for the five most recent probes that have run on each agent. This is only used for Traffic Controller pools. The structure of this block follows the same structure as the <a href="#nested-limit-block-has-the-following-structure">`limit`</a> block described below.

### Nested `limit` block has the following structure:

* `warning` - (Computed) (Integer) Indicates how long (in seconds, or by percentage value) the PING Probe should wait, before a warning is generated.
* `critical` - (Computed) (Integer) Indicates how long (in seconds, or by percentage value) the PING  Probe should wait, before a critical warning is generated.
* `fail` - (Computed) (Integer) Indicates how long (in seconds, or by percentage value) the PING Probe should wait, before causing the probe to fail.

-> `warning` and `critical` are only used for Traffic Controller pools.

