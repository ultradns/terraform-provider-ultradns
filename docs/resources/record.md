---
subcategory: "RECORD"
layout: "ultradns"
page_title: "ULTRADNS: ultradns_record"
description: |-
  Manges the standard DNS records in UltraDNS.
---

# Resource: ultradns_record

Use this resource to manage standard DNS records in UltraDNS

## Example Usage

### Create DNS record of type A (1)

```terraform
resource "ultradns_record" "www" {
    zone_name = "example.com."
    owner_name = "www"
    record_type = "A"
    ttl = 120
    record_data = ["192.168.1.1"]
}
```

### Create DNS record of type AAAA (28)

```terraform
resource "ultradns_record" "aaaa" {
    zone_name = "example.com."
    owner_name = "aaaa"
    record_type = "28"
    ttl = 120
    record_data = ["2001:db8:85a3:0:0:8a2e:370:7334"]
}
```

### Create DNS record of type CNAME (5)

```terraform
resource "ultradns_record" "cname" {
    zone_name = "example.com."
    owner_name = "cname"
    record_type = "CNAME"
    ttl = 120
    record_data = ["google.com."]
}
```

### Create DNS record of type MX (15)

```terraform
resource "ultradns_record" "mx" {
    zone_name = "example.com."
    owner_name = "mx"
    record_type = "15"
    ttl = 120
    record_data = ["2 google.com."]
}
```

### Create DNS record of type SRV (33)

```terraform
resource "ultradns_record" "srv" {
    zone_name = "example.com."
    owner_name = "srv"
    record_type = "SRV"
    ttl = 120
    record_data = ["5 6 7 google.com."]
}
```

### Create DNS record of type TXT (16)

```terraform
resource "ultradns_record" "txt" {
    zone_name = "example.com."
    owner_name = "txt"
    record_type = "16"
    ttl = 120
    record_data = ["google.com."]
}
```

### Create DNS record of type PTR (12)

```terraform
resource "ultradns_record" "ptr" {
    zone_name = "example.com."
    owner_name = "192.168.1.1"
    record_type = "PTR"
    ttl = 120
    record_data = ["google.com."]
}
```

## Argument Reference

The following arguments are supported:

* `zone_name` - (Required) (String) Name of the zone.
* `owner_name` - (Required) (String) The domain name of the owner of the RRSet. Can be either fully qualified domain name (FQDN) or relative domain name. If a FQDN, it must be contained within the zone name FQDN.
* `record_type` - (Required) (String) Must be formatted as the well-known resource record type (A, AAAA, TXT, etc.) ot the corresponding number for the type, between 1 and 65535.<br/>
Below are suported resource record type with its corresponding number:<br/>
`A (1)`
`AAAA (28)`
`CNAME (5)`
`MX (15)`
`SRV (33)`
`TXT (16)`
`PTR (12)`
* `ttl` - (Optional) (Integer) The time to live (in seconds) for for the record. Must be a value between 0 and 2147483647, inclusive.
* `record_data` - (Required) (String List) The data for the record. Must use the BIND presentation format for the specified rrtype.<br/>
Example : For SRV record, the format of data is ["priority weight port target"] (["2 2 523 example.com."])<br/>
Also for MX, NS, CNAME, PTR, and APEXALIAS record types, the data value cannot be relative to the zone name. It must be a FQDN.<br/>

## Import

Records can be imported by combining thier `owner_name`, `zone_name`, `record_type` using colen.<br/>
Example : www.example.com.:example.com.:A (1).


-> For import, the `owner_name`, `zone_name` must be FQDN and `record_type` should have the type with corresponding number as shown in the above example.

e.g.,
```
$ terraform import ultradns_record.example "www.example.com.:example.com.:A (1)" 
```