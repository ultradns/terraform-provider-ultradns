---
subcategory: "Record"
layout: "ultradns"
page_title: "ULTRADNS: ultradns_record"
description: |-
  Manages the standard DNS records in UltraDNS.
---

# Resource: ultradns_record

Use this resource to manage standard DNS records in UltraDNS

## Example Usage

### Create DNS record of type A (1)

```terraform
resource "ultradns_record" "a" {
    zone_name = "example.com."
    owner_name = "a"
    record_type = "1"
    ttl = 120
    record_data = ["192.168.1.1"]
}
```

### Create DNS record of type CNAME (5)

```terraform
resource "ultradns_record" "cname" {
    zone_name = "example.com."
    owner_name = "cname"
    record_type = "CNAME"
    ttl = 120
    record_data = ["sample.com."]
}
```

### Create DNS record of type PTR (12)

```terraform
resource "ultradns_record" "ptr" {
    zone_name = "example.com."
    owner_name = "192.168.1.1"
    record_type = "12"
    ttl = 120
    record_data = ["sample.com."]
}
```

### Create DNS record of type HINFO (13)

```terraform
resource "ultradns_record" "hinfo" {
    zone_name = "example.com."
    owner_name = "hinfo"
    record_type = "HINFO"
    ttl = 120
    record_data = ["\"PC\" \"Linux\"","\"Laptop\" \"Windows\""]
}
```

### Create DNS record of type MX (15)

```terraform
resource "ultradns_record" "mx" {
    zone_name = "example.com."
    owner_name = "mx"
    record_type = "15"
    ttl = 120
    record_data = ["2 sample.com."]
}
```

### Create DNS record of type TXT (16)

```terraform
resource "ultradns_record" "txt" {
    zone_name = "example.com."
    owner_name = "txt"
    record_type = "TXT"
    ttl = 120
    record_data = ["example.com."]
}
```

### Create DNS record of type RP (17)

```terraform
resource "ultradns_record" "rp" {
    zone_name = "example.com."
    owner_name = "example.com."
    record_type = "17"
    ttl = 120
    record_data = ["test.sample.com. sample.128/134.123.178.178.in-addr.arpa."]
}
```

### Create DNS record of type AAAA (28)

```terraform
resource "ultradns_record" "aaaa" {
    zone_name = "example.com."
    owner_name = "aaaa"
    record_type = "AAAA"
    ttl = 120
    record_data = ["2001:db8:85a3:0:0:8a2e:370:7334"]
}
```

### Create DNS record of type SRV (33)

```terraform
resource "ultradns_record" "srv" {
    zone_name = "example.com."
    owner_name = "srv"
    record_type = "33"
    ttl = 120
    record_data = ["5 6 7 sample.com."]
}
```

### Create DNS record of type NAPTR (35)

```terraform
resource "ultradns_record" "naptr" {
    zone_name = "example.com."
    owner_name = "naptr"
    record_type = "NAPTR"
    ttl = 120
    record_data = ["1 2 \"3\" \"test\" \"\" test.com."]
}
```

### Create DNS record of type SSHFP (44)

```terraform
resource "ultradns_record" "sshfp" {
    zone_name = "example.com."
    owner_name = "sshfp"
    record_type = "SSHFP"
    ttl = 120
    record_data = ["1 2 54B5E539EAF593AEA410F80737530B71CCDE8B6C3D241184A1372E98BC7EDB37"]
}
```

### Create DNS record of type TLSA (52)

```terraform
resource "ultradns_record" "tlsa" {
    zone_name = "example.com."
    owner_name = "_23._tcp.tlsatest"
    record_type = "52"
    ttl = 120
    record_data = ["0 0 0 aaaaaaaa"]
}
```

### Create DNS record of type SPF (99)

```terraform
resource "ultradns_record" "spf" {
    zone_name = "example.com."
    owner_name = "example.com."
    record_type = "SPF"
    ttl = 120
    record_data = ["v=spf1 ip4:1.2.3.4 ~all"]
}
```

### Create DNS record of type CAA (257)

```terraform
resource "ultradns_record" "caa" {
    zone_name = "example.com."
    owner_name = "caa"
    record_type = "257"
    ttl = 120
    record_data = ["1 issue \"asdfsadf\""]
}
```

### Create DNS record of type APEXALIAS (65282)

```terraform
resource "ultradns_record" "apex" {
    zone_name = "example.com."
    owner_name = "example.com."
    record_type = "APEXALIAS"
    ttl = 120
    record_data = ["sample.com."]
}
```

## Argument Reference

The following arguments are supported:

* `zone_name` - (Required) (String) Name of the zone.
* `owner_name` - (Required) (String) The domain name of the owner of the RRSet. Can be either a fully qualified domain name (FQDN) or a relative domain name. If provided as a FQDN, it must be contained within the zone name's FQDN.
* `record_type` - (Required) (String) Must be formatted as the well-known resource record type (A, AAAA, TXT, etc.) or the corresponding number for the type; between 1 and 65535.<br/>
Below are the supported resource record type with its corresponding number:<br/>
`A (1)`
`CNAME (5)`
`PTR (12)`
`HINFO (13)`
`MX (15)`
`TXT (16)`
`RP (17)`
`AAAA (28)`
`SRV (33)`
`NAPTR (35)`
`SSHFP (44)`
`TLSA (52)`
`SPF (99)`
`CAA (257)`
`APEXALIAS (65282)`
* `ttl` - (Optional) (Integer) The time to live (in seconds) for the record. Must be a value between 0 and 2147483647, inclusive.
* `record_data` - (Required) (String List) The data for the record displayed as the BIND presentation format for the specified RRTYPE.<br/>
Example : For a SRV record, the format of data is ["priority weight port target"] (["2 2 523 example.com."])<br/>
Additionally for MX, CNAME, and PTR record types, the data value must be a FQDN, as it cannot be relative to the zone name.<br/>

## Import

Records can be imported by combining their `owner_name`, `zone_name`, `record_type` using colon.<br/>
Example : `www.example.com.:example.com.:A (1)`.


-> For import, the `owner_name` and `zone_name` must be a FQDN and `record_type` should have the type as well as the corresponding number as shown in the example above.

e.g.,
```
$ terraform import ultradns_record.example "www.example.com.:example.com.:A (1)" 
```