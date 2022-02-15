---
subcategory: "RD-Pool"
layout: "ultradns"
page_title: "ULTRADNS: ultradns_rdpool"
description: |-
  Get information of Resource Distribution (RD) pool records in UltraDNS.
---

# Data Source: ultradns_rdpool

Use this data source to get detailed information of Resource Distribution (RD) pool records.

## Example Usage

```terraform
data "ultradns_rdpool" "rdpool" {
    zone_name = "example.com."
    owner_name = "www"
    record_type = "A"
}
```


## Argument Reference

The following arguments are supported:

* `zone_name` - (Required) (String) Name of the zone.
* `owner_name` - (Required) (String) The domain name of the owner of the RRSet. Can be either a fully qualified domain name (FQDN) or a relative domain name. If provided as a FQDN, it must be contained within the zone name's FQDN.
* `record_type` - (Required) (String) Must be formatted as the well-known resource record type (A or AAAA) or the corresponding number for the type (1 or 28).<br/>
Below are the supported resource record types with the corresponding number:<br/>
`A (1)`
`AAAA (28)`


## Attributes Reference

In addition to all of the arguments above, the following attributes are exported:

* `ttl` - (Computed) (Integer) The time to live (in seconds) for the record. Must be a value between 0 and 2147483647, inclusive.
* `record_data` - (Computed) (String List) The list of IPv4 or IPv6 addresses.
* `order` - (Computed) (String) The order of the records will be returned in. Valid values are `FIXED`, `RANDOM`, `ROUND_ROBIN`.
* `description` - (Computed) (String) An optional description of the RD pool.