---
subcategory: "DNS-Probe"
layout: "ultradns"
page_title: "ULTRADNS: ultradns_probe_dns"
description: |-
  Get information of DNS probe records in UltraDNS.
---

# Data Source: ultradns_probe_dns

Use this data source to get detailed information of DNS probe records.

## Example Usage

### Get last created DNS Probe

```terraform
data "ultradns_probe_dns" "dns" {
    zone_name = "example.com."
    owner_name = "www"
}
```

### Get last created DNS Probe of AAAA Pool

```terraform
data "ultradns_probe_dns" "dns" {
    zone_name = "example.com."
    owner_name = "www"
    pool_type = "AAAA"
}
```

### Get DNS Probe by guid

```terraform
data "ultradns_probe_dns" "dns" {
    zone_name = "example.com."
    owner_name = "www"
    guid = "06084A729D56C85C"
}
```

### Get DNS Probe by filtering

```terraform
data "ultradns_probe_dns" "dns" {
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
* `pool_type` - (Optional) (String) Pool type of the probe. Valid values are `A`, `AAAA`.</br>Default value set to `A`.

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

* `port` - (Computed) (String) The Port that should be used for DNS lookup.
* `type` - (Computed) (String) Select the record type that the probe will check for. Valid values are `NULL`, `AXFR`, or any Resource Record Type.
* `query_name` - (Computed) (String) The name that should be queried.
* `tcp_only` - (Computed) (Boolean) Indicates whether or not the probe should use TCP only, or first UDP then TCP.
* `response` - (Computed) (Block Set, Max:1) Nested block describing the strings to match the response that will generate a warning or failure. The structure of this block is described below.
* `run_limit` - (Computed) (Block Set, Max:1) Nested block describing how long the probe should run. The structure of this block follows the same structure as the [`limit`](#nested-limit-block-has-the-following-structure) block described below.
* `avg_run_limit` - (Computed) (Block Set, Max:1) Nested block describing the mean (average) run-time for the five most recent probes that have run on each agent. This is only used for Traffic Controller pools. The structure of this block follows the same structure as the [`limit`](#nested-limit-block-has-the-following-structure) block described below.

### Nested `limit` block has the following structure:

* `warning` - (Computed) (Integer) Indicates how long (in seconds) the DNS Probe should wait, before a warning is generated.
* `critical` - (Computed) (Integer) Indicates how long (in seconds) the DNS  Probe should wait, before a critical warning is generated.
* `fail` - (Computed) (Integer) Indicates how long (in seconds) the DNS Probe should wait, before causing the probe to fail.

### Nested `response` block has the following structure:

* `warning` - (Computed) (String) Will match exactly for records containing a single field response (i.e., A, CNAME, DNAME, NS, MB, MD, MF, MG, MR, PTR), and partial matches for record types with multiple field response. Fields separated by spaces will be combined with the matches to trigger a warning.
* `critical` - (Computed) (String) Will match exactly for records containing a single field response (i.e., A, CNAME, DNAME, NS, MB, MD, MF, MG, MR, PTR), and partial matches for record types with multiple field response. Fields separated by spaces will be combined with the matches to trigger a critical warning.
* `fail` - (Computed) (String) Will match exactly for records containing a single field response (i.e., A, CNAME, DNAME, NS, MB, MD, MF, MG, MR, PTR), and partial matches for record types with multiple field response. Fields separated by spaces will be combined with the matches to trigger a failure.

-> `warning` and `critical` are only used for Traffic Controller pools.