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

The following arguments are used to filter the probes:

* `guid` - (Optional) (String) The internal id for this probe.
* `interval` - (Optional) (String) Length of time between probes in minutes.
* `agents` - (Optional) (String List) Locations that will be used for probing. The exact list must be provided if probes need to be filtered using agents.
* `threshold` - (Optional) (Integer) The probe threshold value.
* `pool_record` - (Optional) (String) The pool record associated with this probe.

->
1) If `guid` is provided, the probe with that guid is returned, and other filter options are not considered.</br>
2) If there is a conflict between probes due to filter options other than `guid`,the last created probe is returned.</br>
3) If no probe is found for the filter options, an error is returned.  

## Attributes Reference

In addition to all of the arguments above, the following attributes are exported:

* `port` - (Optional) (String) The Port that should be used for DNS lookup.
* `type` - (Optional) (String) Select which kind of record should be checked for. Valid values are `NULL`, `AXFR`, or any Resource Record Type.
* `query_name` - (Optional) (String) The name that should be queried.
* `tcp_only` - (Optional) (Boolean) Indicates whether or not the probe should use TCP only, or first UDP then TCP.
* `response` - (Optional) (Block Set, Max:1) Nested block describing the strings to match the response that will generates a warning or failure. The structure of this block follows the same structure as the <a href="#nested-limit-block-has-the-following-structure">`limit`</a> block described below.
* `run_limit` - (Optional) (Block Set, Max:1) Nested block describing how long the probe should run. The structure of this block follows the same structure as the <a href="#nested-limit-block-has-the-following-structure">`limit`</a> block described below.
* `avg_run_limit` - (Optional) (Block Set, Max:1) Nested block describing the mean (average) run-time for the five most recent probes that have run on each agent. This is only used for Traffic Controller pools. The structure of this block follows the same structure as the <a href="#nested-limit-block-has-the-following-structure">`limit`</a> block described below.

### Nested `limit` block has the following structure:

* `warning` - (Optional) (Integer) Indicates how long the DNS Probe should wait before a warning is generated.
* `critical` - (Optional) (Integer) Indicates how long the DNS  Probe should wait before a critical warning is generated.
* `fail` - (Optional) (Integer) Indicates how long the DNS Probe should wait before it make the probe to fail.

### Nested `response` block has the following structure:

* `warning` - (Optional) (String) Match exactly for records with single field responses (that is: A, CNAME, DNAME, NS, MB, MD, MF, MG, MR, PTR), and match partially for types with multiple field responses, and join all fields separated by spaces and match to trigger a warning.
* `critical` - (Optional) (String) Match exactly for records with single field responses (that is: A, CNAME, DNAME, NS, MB, MD, MF, MG, MR, PTR), and match partially for types with multiple field responses, and join all fields separated by spaces and match to trigger a critical warning.
* `fail` - (Optional) (String) Match exactly for records with single field responses (that is: A, CNAME, DNAME, NS, MB, MD, MF, MG, MR, PTR), and match partially for types with multiple field responses, and join all fields separated by spaces and match to trigger a failure.

-> `warning` and `critical` are only used for Traffic Controller pools