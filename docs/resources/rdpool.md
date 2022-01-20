---
subcategory: "RD-Pool"
layout: "ultradns"
page_title: "ULTRADNS: ultradns_rdpool"
description: |-
  Manages the Resource Distribution (RD) pool records in UltraDNS.
---

# Resource: ultradns_rdpool

Use this resource to manage resource distribution pool records in UltraDNS

## Example Usage

### Create RD pool record of type A (1)

```terraform
resource "ultradns_rdpool" "a" {
    zone_name = "example.com."
    owner_name = "a"
    record_type = "1"
    ttl = 120
    record_data = ["192.168.1.1"]
    order = "RANDOM"
}
```

### Create RD pool record of type AAAA (28)

```terraform
resource "ultradns_rdpool" "aaaa" {
    zone_name = "example.com."
    owner_name = "aaaa"
    record_type = "AAAA"
    ttl = 120
    record_data = ["2001:db8:85a3:0:0:8a2e:370:7334"]
    order = "ROUND_ROBIN"
}
```

## Argument Reference

The following arguments are supported:

* `zone_name` - (Required) (String) Name of the zone.
* `owner_name` - (Required) (String) The domain name of the owner of the RRSet. Can be either a fully qualified domain name (FQDN) or a relative domain name. If provided as a FQDN, it must be contained within the zone name's FQDN.
* `record_type` - (Required) (String) Must be formatted as the well-known resource record type (A or AAAA) or the corresponding number for the type (1 or 28).<br/>
Below are the supported resource record type with its corresponding number:<br/>
`A (1)`
`AAAA (28)`
* `record_data` - (Required) (String List) The list of IPv4 or IPv6 addresses.
* `order` - (Required) (String) The order of the records will be returned in. Valid values are `FIXED`, `RANDOM`, `ROUND_ROBIN`.
* `ttl` - (Optional) (Integer) The time to live (in seconds) for the record. Must be a value between 0 and 2147483647, inclusive.
* `description` - (Optional) (String) An optional description of the RD pool.

## Import

Resource Distribution (RD) pool records can be imported by combining their `owner_name`, `zone_name`, `record_type` using colon.<br/>
Example : `www.example.com.:example.com.:A (1)`.


-> For import, the `owner_name` and `zone_name` must be a FQDN and `record_type` should have the type as well as the corresponding number as shown in the example above.

e.g.,
```
$ terraform import ultradns_rdpool.example "www.example.com.:example.com.:A (1)" 
```