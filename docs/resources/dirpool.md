---
subcategory: "DIR-Pool"
layout: "ultradns"
page_title: "ULTRADNS: ultradns_dirpool"
description: |-
  Manages the Directional (DIR) pool records in UltraDNS.
---

# Resource: ultradns_dirpool

Use this resource to manage Directional (DIR) pool records in UltraDNS.

## Example Usage

### Create DIR pool record of type A (1)

```terraform
resource "ultradns_dirpool" "a" {
    zone_name = "example.com."
    owner_name = "a"
    record_type = "A"
    ignore_ecs = true
    conflict_resolve = "IP"
    rdata_info{
        rdata = "192.168.1.5"
        all_non_configured = true
        ttl = 800
    }
    rdata_info{
        rdata = "192.168.1.2"
        geo_group_name = "geo_group"
        geo_codes = ["NAM","EUR"]
        ip_group_name = "ip_group"
        ip{
            address = "200.1.1.1"
        }
        ip{
            start = "200.1.1.2"
            end = "200.1.1.5"
        }
        ip{
            cidr = "200.20.20.0/24"
        }
    }
    no_response{
        geo_group_name = "geo_response_group"
        geo_codes = ["AG"]
        ip_group_name = "ip_response_group"
        ip{
            address = "2.2.2.2"
        }
    }
}
```

### Create DIR pool record of Type PTR (12)

```terraform
resource "ultradns_dirpool" "ptr" {
    zone_name = "example.com."
    owner_name = "1"
    record_type = "PTR"
    rdata_info{
        rdata = "ns1.example.com."
        geo_group_name = "geo_group"
        geo_codes = ["NAM","EUR"]
    }
    no_response{
        all_non_configured = true
    }
}
```

### Create DIR pool record of Type MX (15)

```terraform
resource "ultradns_dirpool" "mx" {
    zone_name = "example.com."
    owner_name = "mx"
    record_type = "MX"
    rdata_info{
        rdata = "2 example.com."
        geo_group_name = "geo_group"
        geo_codes = ["NAM","EUR"]
    }
    no_response{
        all_non_configured = true
    }
}
```

### Create DIR pool record of Type TXT (16)

```terraform
resource "ultradns_dirpool" "txt" {
    zone_name = "example.com."
    owner_name = "txt.example.com."
    record_type = "TXT"
    rdata_info{
        rdata = "text data"
        geo_group_name = "geo_group"
        geo_codes = ["NAM","EUR"]
    }
    no_response{
        all_non_configured = true
    }
}
```

### Create DIR pool record of Type AAAA (28)

```terraform
resource "ultradns_dirpool" "aaaa" {
    zone_name = "example.com."
    owner_name = "aaaa"
    record_type = "AAAA"
    pool_description = "DIR Pool Resource of Type AAAA"
    ignore_ecs = true
    conflict_resolve = "IP"
    rdata_info{
        rdata = "aaaa:bbbb:cccc:dddd:eeee:ffff:1111:3333"
        geo_group_name = "geo_group"
        geo_codes = ["EUR"]
        ip_group_name = "ip_group"
        ip{
            start = "aaaa:bbbb:cccc:dddd:eeee:ffff:1111:4444"
            end = "aaaa:bbbb:cccc:dddd:eeee:ffff:1111:6666"
        }
    }
    no_response{
        geo_group_name = "geo_response_group"
        geo_codes = ["AI"]
        ip_group_name = "ip_response_group"
        ip{
            address = "aaaa:bbbb:cccc:dddd:eeee:ffff:3333:5555"
        }
    }
}
```

### Create DIR pool record of Type SRV (33)

```terraform
resource "ultradns_dirpool" "srv" {
    zone_name = "example.com."
    owner_name = "srv"
    record_type = "SRV"
    rdata_info{
        rdata = "5 6 7 example.com."
        geo_group_name = "geo_group"
        geo_codes = ["NAM","EUR"]
    }
    no_response{
        all_non_configured = true
    }
}
```

## Argument Reference

The following arguments are supported:

* `zone_name` - (Required) (String) Name of the zone.
* `owner_name` - (Required) (String) The domain name of the owner of the RRSet. Can be either a fully qualified domain name (FQDN) or a relative domain name. If provided as a FQDN, it must be contained within the zone name's FQDN.
* `record_type` - (Required) (String) Must be formatted as a well-known resource record type (A, AAAA, TXT, etc.), or the corresponding number for the type; between 1 and 33.<br/>
Below are the supported resource record types with the corresponding number:<br/>
`A (1)`
`PTR (12)`
`MX (15)`
`TXT (16)`
`AAAA (28)`
`SRV (33)`
* `pool_description` - (Optional) (String) An optional description of the Directional (DIR) field.
* `conflict_resolve` - (Optional) (String) When there is a conflict between a matching GeoIP group and a matching SourceIP group, this will determine which should take precedence. This only applies to a mixed pool (contains both GeoIP and SourceIP data). Valid values are `GEO` and `IP`. Default value set to `GEO`
* `ignore_ecs` - (Optional) (Boolean) Whether to ignore the EDNSO (which is an extended label type allowing for greater DNS message size) Client Subnet data when available in the DNS request.</br>
`false`= EDNSO data will be used for IP directional routing.</br>
`true` = EDNSO data will not be used and IP directional routing decisions will always use the IP address of the recursive server.</br>
Default value set to false.
* `no_response` - (Optional) (Block Set, Max: 1) Nested block describing certain geographical territories and IP addresses that will get no response if they try to access the directional pool. The structure of this block is described below.
* `rdata_info` - (Required) (Block Set, Min: 1) List of nested blocks describing the pool records. The structure of this block is described below.

### Nested `rdata_info` block has the following structure:

* `rdata` - (Required) (String) The IPv4/IPv6 address, CNAME, MX, TXT, or SRV format data.
* `type` - (Computed) (String) The type of pool record.
* `ttl` - (Optional) (Integer) The time to live (in seconds) for the corresponding record in rdata. Must be a value between 0 and 2147483647, inclusive.
* `all_non_configured` - (Optional) (Boolean) Indicates whether or not the associated rdata is used for all non-configured geographical territories and SourceIP ranges. At most, one entry in rdataInfo can have this set to true. If this is set to true, then geoInfo and ipInfo are ignored. Default value set to false.
* `geo_group_name` - (Optional) (String) The name of the GeoIP group.
* `geo_codes` - (Optional) (String List) The codes for the geographical territories that make up this group. [Valid GEO codes](#valid-geo-codes).
* `ip_group_name` - (Optional) (String) The name of the SourceIP group.
* `ip` - (Optional) (Block Set) List of nested blocks describing the IP addresses and IP ranges this SourceIP group contains. The structure of this block is described below.

### Nested `no_response` block has the following structure:

* `all_non_configured` - (Optional) (Boolean) Indicates whether or not “no response” is returned for all non-configured geographical territories and IP ranges. This can only be set to true if there is no entry in rdataInfo with allNonConfigured set to true. If this is set to true, then geoInfo and ipInfo are ignored. Default value set to false.
* `geo_group_name` - (Optional) (String) The name for the “no response” GeoIP group.
* `geo_codes` - (Optional) (String List) The codes for the geographical territories that make up the “no response” group. [Valid GEO codes](#valid-geo-codes).
* `ip_group_name` - (Optional) (String) The name of the “no response” SourceIP group.
* `ip` - (Optional) (Block Set) List of nested blocks describing the IP addresses and IP range for the “no response” SourceIP group. The structure of this block is described below.

### Nested `ip` block has the following structure:

* `start` - (Optional) (String) The starting IP address (v4 or v6) for a SourceIP range. If start is present, end must be present as well. `cidr` and `address` must not be present.
* `end` - (Optional) (String) The ending IP address (v4 or v6) for a SourceIP range. If end is present, start must be present as well. `cidr` and `address` must not be present.
* `cidr` - (Optional) (String) The CIDR format (IPv4 or IPv6) for an IP address range. If CIDR is present, the `start`, `end`, and `address` must not be present.
* `address` - (Optional) (String) A single IPv4 address. If address is present, the `start`, `end`, and `CIDR` must not be present.


## Import

Directional (DIR) pool records can be imported by combining their `owner_name`, `zone_name`, and `record_type`, separated by a colon.<br/>
Example : `www.example.com.:example.com.:A (1)`.


-> For import, the `owner_name` and `zone_name` must be a FQDN, and `record_type` should have the type, as well as the corresponding number, as shown in the example below.

Example:
```terraform
$ terraform import ultradns_dirpool.example "www.example.com.:example.com.:A (1)" 
```


## Valid GEO Codes:

| Code 	| Meaning | Equivalent ISO codes |
| :--- 	| :----: | :--- |
|_________________________|__________________________________________________|__________________________________________________|
| `A1`  | Anonymous Proxy | None |
|_________________________|__________________________________________________|__________________________________________________|
| `A2`	| Satellite Provider | None |
|_________________________|__________________________________________________|__________________________________________________|
| `A3`	| Unknown / Uncategorized IPs | None |
|_________________________|__________________________________________________|__________________________________________________|
| `NAM`	| North America (including Central America and the Caribbean) | `AG`,`AI`,`AN`,`AW`,`BB`,`BL`,`BM`,</br>`BQ`,`BS`,`BZ`,`CA`,`CR`,`CU`,`CW`,</br>`DM`,`DO`,`GD`,`GL`,`GP`,`GT`,`HN`,</br>`HT`,`JM`,`KN`,`KY`,`LC`,`MF`,`MQ`,</br>`MS`,`MX`,`NI`,`PA`,`PM`,`PR`,`SV`,</br>`SX`,`TC`,`TT`,`U3`,`US`,`VC`,`VG`,</br>`VI` |
|_________________________|__________________________________________________|__________________________________________________|
| `SAM`	| South America | `AR`,`BO`,`BR`,`CL`,`CO`,`EC`,`FK`,</br>`GF`,`GS`,`GY`,`PE`,`PY`,`SR`,`U4`,</br>`UY`,`VE` |
|_________________________|__________________________________________________|__________________________________________________|
| `EUR`	| Europe | `AD`,`AL`,`AM`,`AT`,`AX`,`AZ`,`BA`,</br>`BE`,`BG`,`BY`,`CH`,`CZ`,`DE`,`DK`,</br>`EE`,`ES`,`FI`,`FO`,`FR`,`GB`,`GE`,</br>`GG`,`GI`,`GR`,`HR`,`HU`,`IE`,`IM`,</br>`IS`,`IT`,`JE`,`LI`,`LT`,`LU`,`LV`,</br>`MC`,`MD`,`ME`,`MK`,`MT`,`NL`,`NO`,</br>`PL`,`PT`,`RO`,`RS`,`SE`,`SI`,`SJ`,</br>`SK`,`SM`,`U5`,`UA`,`VA` |
|_________________________|__________________________________________________|__________________________________________________|
| `AFR`	| Africa | `AO`,`BF`,`BI`,`BJ`,`BW`,`CD`,`CF`,</br>`CG`,`CI`,`CM`,`CV`,`DJ`,`DZ`,`EG`,</br>`EH`,`ER`,`ET`,`GA`,`GH`,`GM`,`GN`,</br>`GQ`,`GW`,`KE`,`KM`,`LR`,`LS`,`LY`,</br>`MA`,`MG`,`ML`,`MR`,`MU`,`MW`,`MZ`,</br>`NA`,`NE`,`NG`,`RE`,`RW`,`SC`,`SD`,</br>`SH`,`SL`,`SN`,`SO`,`SS`,`ST`,`SZ`,</br>`TD`,`TG`,`TN`,`TZ`,`U7`,`UG`,`YT`,</br>`ZA`,`ZM`,`ZW` |
|_________________________|__________________________________________________|__________________________________________________|
| `ASI`	| Asia (including Middle East and the Russian Federation) | `AE`,`AF`,`BD`,`BH`,`BN`,`BT`,`CN`,</br>`CY`,`HK`,`ID`,`IL`,`IN`,`IO`,`IQ`,</br>`IR`,`JO`,`JP`,`KG`,`KH`,`KP`,`KR`,</br>`KW`,`KZ`,`LA`,`LB`,`LK`,`MM`,`MN`,</br>`MO`,`MV`,`MY`,`NP`,`OM`,`PH`,`PK`,</br>`PS`,`QA`,`RU`,`SA`,`SG`,`SY`,`TH`,</br>`TJ`,`TL`,`TM`,`TR`,`TW`,`U6`,`U8`,</br>`UZ`,`VN`,`YE` |
|_________________________|__________________________________________________|__________________________________________________|
| `OCN`	| Australia / Oceania | `AS`,`AU`,`CC`,`CK`,`CX`,`FJ`,`FM`,</br>`GU`,`HM`,`KI`,`MH`,`MP`,`NC`,`NF`,</br>`NR`,`NU`,`NZ`,`PF`,`PG`,`PN`,`PW`,</br>`SB`,`TK`,`TO`,`TV`,`U9`,`UM`,`VU`,</br>`WF`,`WS` |
|_________________________|__________________________________________________|__________________________________________________|
| `ANT`	| Antarctica | `AQ`, `TF`, `BV` |
|_________________________|__________________________________________________|__________________________________________________|