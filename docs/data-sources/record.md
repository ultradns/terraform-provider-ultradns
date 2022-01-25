---
subcategory: "Record"
layout: "ultradns"
page_title: "ULTRADNS: ultradns_record"
description: |-
  Get information of standard DNS records in UltraDNS.
---

## Example Usage

```terraform
data "ultradns_record" "record" {
    zone_name = "example.com."
    owner_name = "www"
    record_type = "A"
}
```


## Argument Reference

The following arguments are supported:

* `zone_name` - (Required) (String) Name of the zone.
* `owner_name` - (Required) (String) The domain name of the owner of the RRSet. Can be either a fully qualified domain name (FQDN) or a relative domain name. If provided as a FQDN, it must be contained within the zone name's FQDN.
* `record_type` - (Required) (String) Must be formatted as the well-known resource record type (A, AAAA, TXT, etc.) or the corresponding number for the type, between 1 and 65535.<br/>
Below are the supported resource record type with its corresponding number:<br/>
`A (1)`
`CNAME (5)`
`PTR (12)`
`MX (15)`
`TXT (16)`
`AAAA (28)`
`SRV (33)`
`SSHFP (44)`
`APEXALIAS (65282)`


## Attributes Reference

In addition to all of the arguments above, the following attributes are exported:

* `ttl` - (Computed) (Integer) The time to live (in seconds) for the record. Must be a value between 0 and 2147483647, inclusive.
* `record_data` - (Computed) (String List) The data for the record displayed as the BIND presentation format for the specified RRTYPE.<br/>
Example : For a SRV record, the format of data is ["priority weight port target"] (["2 2 523 example.com."]).