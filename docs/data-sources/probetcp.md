---
subcategory: "TCP-Probe"
layout: "ultradns"
page_title: "ULTRADNS: ultradns_probe_tcp"
description: |-
  Get information of TCP probe records in UltraDNS.
---

# Data Source: ultradns_probe_tcp

Use this data source to get detailed information of TCP probe records.

## Example Usage

### Get last created TCP Probe

```terraform
data "ultradns_probe_tcp" "tcp" {
    zone_name = "example.com."
    owner_name = "www"
}
```

### Get TCP Probe by guid

```terraform
data "ultradns_probe_tcp" "tcp" {
    zone_name = "example.com."
    owner_name = "www"
    guid = "06084A729D56C85C"
}
```

### Get TCP Probe by filtering

```terraform
data "ultradns_probe_tcp" "tcp" {
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

* `port` - (Computed) (Integer) TCP port number to connect to. Default value set to 80.
* `control_ip` - (Computed) (String) The target IP address of the TCP probe.
* `connect_limit` - (Computed) (Block Set, Max:1) Nested block describing how long the probe stays connected to the resource. The structure of this block follows the same structure as the [`limit`](#nested-limit-block-has-the-following-structure) block described below.

### Nested `limit` block has the following structure:

* `warning` - (Computed) (Integer) Indicates how long (in seconds, or by percentage value) the TCP Probe should wait, before a warning is generated.
* `critical` - (Computed) (Integer) Indicates how long (in seconds, or by percentage value) the TCP  Probe should wait, before a critical warning is generated.
* `fail` - (Computed) (Integer) Indicates how long (in seconds, or by percentage value) the TCP Probe should wait, before causing the probe to fail.

-> `warning` and `critical` are only used for Traffic Controller pools.

