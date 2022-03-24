---
subcategory: "DIR-Pool"
layout: "ultradns"
page_title: "ULTRADNS: ultradns_dirpool"
description: |-
  Get information of Directional (DIR) pool records in UltraDNS.
---

# Data Source: ultradns_dirpool

Use this data source to get detailed information of Directional (DIR) pool records.

## Example Usage

```terraform
data "ultradns_dirpool" "dirpool" {
    zone_name = "example.com."
    owner_name = "www"
    record_type = "A"
}
```


## Argument Reference

The following arguments are supported:

* `zone_name` - (Required) (String) Name of the zone.
* `owner_name` - (Required) (String) The domain name of the owner of the RRSet. Can be either a fully qualified domain name (FQDN) or a relative domain name. If provided as a FQDN, it must be contained within the zone name's FQDN.
* `record_type` - (Required) (String) Must be formatted as a well-known resource record type (A), or the corresponding number for the type (1).<br/>
Below are the supported resource record types with the corresponding number:<br/>
`A (1)`


## Attributes Reference

In addition to all of the arguments above, the following attributes are exported:

* `pool_description` - (Computed) (String) An optional description of the Directional (DIR) field.
* `conflict_resolve` - (Computed) (String) When there is a conflict between a matching GeoIP group and a matching SourceIP group, this will determine which should take precedence. This only applies to a mixed pool (contains both GeoIP and SourceIP data).
* `ignore_ecs` - (Computed) (Boolean) Whether to ignore the EDNSO (which is an extended label type allowing for greater DNS message size) Client Subnet data when available in the DNS request.</br>
`false`= EDNSO data will be used for IP directional routing.</br>
`true` = EDNSO data will not be used and IP directional routing decisions will always use the IP address of the recursive server.
* `no_response` - (Computed) (Block Set, Max: 1) Nested block describing certain geographical territories and IP addresses that will get no response if they try to access the directional pool. The structure of this block is described below.
* `rdata_info` - (Computed) (Block Set, Min: 1) List of nested blocks describing the pool records. The structure of this block is described below.

### Nested `rdata_info` block has the following structure:

* `rdata` - (Computed) (String) The IPv4/IPv6 address, CNAME, MX, TXT, or SRV format data.
* `type` - (Computed) (String) The type of pool record.
* `ttl` - (Computed) (Integer) The time to live (in seconds) for the corresponding record in rdata. Must be a value between 0 and 2147483647, inclusive.
* `all_non_configured` - (Computed) (Boolean) Indicates whether or not the associated rdata is used for all non-configured geographical territories and SourceIP ranges. At most, one entry in rdataInfo can have this set to true.
* `geo_group_name` - (Computed) (String) The name of the GeoIP group.
* `geo_codes` - (Computed) (String List) The codes for the geographical territories that make up this group.
* `ip_group_name` - (Computed) (String) The name of the SourceIP group.
* `ip` - (Computed) (Block Set) List of nested blocks describing the IP addresses and IP ranges this SourceIP group contains. The structure of this block is described below.

### Nested `no_response` block has the following structure:

* `all_non_configured` - (Computed) (Boolean) Indicates whether or not “no response” is returned for all non-configured geographical territories and IP ranges. This can only be set to true if there is no entry in rdataInfo with allNonConfigured set to true.
* `geo_group_name` - (Computed) (String) The name for the “no response” GeoIP group.
* `geo_codes` - (Computed) (String List) The codes for the geographical territories that make up the “no response” group.
* `ip_group_name` - (Computed) (String) The name of the “no response” SourceIP group.
* `ip` - (Computed) (Block Set) List of nested blocks describing the IP addresses and IP range for the “no response” SourceIP group. The structure of this block is described below.

### Nested `ip` block has the following structure:

* `start` - (Computed) (String) The starting IP address (v4 or v6) for a SourceIP range.
* `end` - (Computed) (String) The ending IP address (v4 or v6) for a SourceIP range.
* `cidr` - (Computed) (String) The CIDR format (IPv4 or IPv6) for an IP address range.
* `address` - (Computed) (String) A single IPv4 address.
