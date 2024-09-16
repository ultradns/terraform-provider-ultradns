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

### Create DNS record of type NS (2)

```terraform
resource "ultradns_record" "ns" {
    zone_name = "example.com."
    owner_name = "example.com."
    record_type = "NS"
    ttl = 120
    record_data = ["ns11.sample.com.","ns12.sample.com."]
}
```

### Create DNS record of type CNAME (5)

```terraform
resource "ultradns_record" "cname" {
    zone_name = "example.com."
    owner_name = "cname"
    record_type = "CNAME"
    ttl = 120
    record_data = ["host.sample.com."]
}
```

### Create DNS record of type PTR (12)

```terraform
resource "ultradns_record" "ptr" {
    zone_name = "70.154.156.in-addr.arpa."
    owner_name = "1"
    record_type = "12"
    ttl = 120
    record_data = ["ns1.example.com."]
}
```

### Create DNS record of type MX (15)

```terraform
resource "ultradns_record" "mx" {
    zone_name = "example.com."
    owner_name = "mx"
    record_type = "15"
    ttl = 120
    record_data = ["2 example.com."]
}
```

### Create DNS record of type TXT (16)

```terraform
resource "ultradns_record" "txt" {
    zone_name = "example.com."
    owner_name = "txt"
    record_type = "TXT"
    ttl = 120
    record_data = ["text data"]
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
    record_data = ["5 6 7 example.com."]
}
```

### Create DNS record of type DS (43)

```terraform
resource "ultradns_record" "ds" {
    zone_name = "example.com."
    owner_name = "example.com."
    record_type = "DS"
    ttl = 800
    record_data = ["25286 1 1 340437DC66C3DFAD0B3E849740D2CF1A4151671D"]
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

### Create DNS record of type SVCB (64)

```terraform
resource "ultradns_record" "svcb" {
    zone_name = "example.com."
    owner_name = "svcb"
    record_type = "SVCB"
    ttl = 800
    record_data = ["0 www.ultradns.com."]
}
```

### Create DNS record of type HTTPS (65)

```terraform
resource "ultradns_record" "https" {
    zone_name = "example.com."
    owner_name = "https"
    record_type = "HTTPS"
    ttl = 800
    record_data = ["1 www.ultradns.com. ech=dGVzdA== mandatory=alpn,key65444 no-default-alpn port=8080 ipv4hint=1.2.3.4,9.8.7.6 key65444=privateKeyTesting ipv6hint=2001:db8:3333:4444:5555:6666:7777:8888,2001:db8:3333:4444:cccc:dddd:eeee:ffff alpn=h3,h3-29,h2"]
}
```

### Create DNS record of type CAA (257)

```terraform
resource "ultradns_record" "caa" {
    zone_name = "example.com."
    owner_name = "example.com."
    record_type = "CAA"
    ttl = 800
    record_data = ["0 issue ultradns"]
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

### Manage the fields of a SOA record
UltraDNS does not allow the creation or deletion of a zone SOA record. The resource can be used to manage all of the fields of the SOA _except for_ the zone serial number which is managed by UltraDNS.

The record data for SOA is a string with space separated fields for `mname`, `rname`, `refresh`, `retry`, `expire`, `minimum`

```terraform
resource "ultradns_record" "soa" {
    zone_name = "example.com."
    owner_name = "example.com."
    record_type = "SOA"
    ttl = 86400
    record_data = ["ns.example.com admin@example.com 7200 3600 1209600 36000"]
}
```


## Argument Reference

The following arguments are supported:

* `zone_name` - (Required) (String) Name of the zone.
* `owner_name` - (Required) (String) The domain name of the owner of the RRSet. Can be either a fully qualified domain name (FQDN) or a relative domain name. If provided as a FQDN, it must be contained within the zone name's FQDN.
* `record_type` - (Required) (String) Must be formatted as the well-known resource record type (A, AAAA, TXT, etc.) or the corresponding number for the type; between 1 and 65535.<br/>
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
* `ttl` - (Optional) (Integer) The time to live (in seconds) for the record. Must be a value between 0 and 2147483647, inclusive.
* `record_data` - (Required) (String List) The data for the record displayed as the BIND presentation format for the specified RRTYPE.<br/>
Example : For a SRV record, the format of data is ["priority weight port target"] (["2 2 523 example.com."])<br/>
Additionally for MX, CNAME, and PTR record types, the data value must be a FQDN, as it cannot be relative to the zone name.<br/>

## Import

Records can be imported by combining their `owner_name`, `zone_name`, and `record_type`, separated by a colon.<br/>
Example : `www.example.com.:example.com.:A (1)`.


-> For import, the `owner_name` and `zone_name` must be a FQDN, and `record_type` must have the type as well as the corresponding number as shown in the example below.

Example:
```
$ terraform import ultradns_record.example "www.example.com.:example.com.:A (1)" 
```
