---
subcategory: "Record"
layout: "ultradns"
page_title: "ULTRADNS: ultradns_record"
description: |-
  Get information of standard DNS records in UltraDNS.
---

# Data Source: ultradns_record

Use this data source to get detailed information of standard DNS records.

## Example Usage

```terraform
data "ultradns_record" "record" {
    zone_name = "example.com."
    owner_name = "www"
    record_type = "A"
}
```

```terraform
data "ultradns_record" "soarecord" {
    zone_name = "example.com."
    owner_name = "example.com."
    record_type = "SOA"
}
```

## Argument Reference

The following arguments are supported:

* `zone_name` - (Required) (String) Name of the zone.
* `owner_name` - (Required) (String) The domain name of the owner of the RRSet. Can be either a fully qualified domain name (FQDN) or a relative domain name. If provided as a FQDN, it must be contained within the zone name's FQDN.
* `record_type` - (Required) (String) Must be formatted as the well-known resource record type (A, AAAA, TXT, etc.) or the corresponding number for the type, between 1 and 65535.<br/>
Below are the supported resource record types with the corresponding number:<br/>
`A (1)`
`NS (2)`
`CNAME (5)`
`SOA (6)`
`PTR (12)`
`MX (15)`
`TXT (16)`
`AAAA (28)`
`SRV (33)`
`DS (43)`
`SSHFP (44)`
`SVCB (64)`
`HTTPS (65)`
`CAA (257)`
`APEXALIAS (65282)`


## Attributes Reference

In addition to all of the arguments above, the following attributes are exported:

* `ttl` - (Computed) (Integer) The time to live (in seconds) for the record. Must be a value between 0 and 2147483647, inclusive.
* `record_data` - (Computed) (String List) The data for the record displayed as the BIND presentation format for the specified RRTYPE.<br/>

__Note__

In SOA records the serial number is ignored and removed from the data: `["mname rname refresh retry expire minimum"]`
```["ns.example.com admin@example.com 7200 3600 1209600 36000"]```

For a SRV record, the format of data is `["priority weight port target"]`
```["2 2 523 example.com."]```
