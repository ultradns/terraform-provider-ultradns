---
subcategory: "SF-Pool"
layout: "ultradns"
page_title: "ULTRADNS: ultradns_sfpool"
description: |-
  Get information of Simple Monitor/Failover (SF) pool records in UltraDNS.
---

## Example Usage

```terraform
data "ultradns_sfpool" "sfpool" {
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
Below are the supported resource record type with its corresponding number:<br/>
`A (1)`
`AAAA (28)`


## Attributes Reference

In addition to all of the arguments above, the following attributes are exported:

* `ttl` - (Computed) (Integer) The time to live (in seconds) for the record. Must be a value between 0 and 2147483647, inclusive.
* `record_data` - (Computed) (String List) The list of IPv4 or IPv6 addresses.
* `status` - (Computed) (String) Current status of the serving record. Valid values are:</br>
`OK` – Live record is being served.</br>
`CRITICAL` – The backup record is being served due to the monitor detecting a failure.</br>
`MANUAL` – The backup record is being served due to user forcing the live record to be inactive.
* `region_failure_sensitivity` - (Computed) (String) Specifies the sensitivity to the number of regions that need to fail for the backup record to be made active. Valid values are `LOW`, `HIGH`.
* `live_record_description` - (Computed) (String) An optional description of the live record.
* `pool_description` - (Computed) (String) An optional description of the Simple Failover field.
* `monitor` - (Computed) (Block Set) Nested block describing the information for the monitor. The structure of this block is described below.
* `backup_record` - (Computed) (Block Set) Nested block describing the information for the backup record. The structure of this block is described below.

### Nested `monitor` block has the following structure:

* `url` - (Computed) (String) Monitored URL. A full URL including: protocol, host, and URI. Required.
Valid protocols are HTTP and HTTPS.
* `method` - (Computed) (String) HTTP method used to connect to the monitored URL. Valid values are `GET`, `POST`.
* `transmitted_data` - (Computed) (String) If a monitor is sending a POST, the data that is sent as the body of the request.
* `search_string` - (Computed) (String) A string that is checked against the returned data from the request. 

### Nested `backup_record` block has the following structure:

* `rdata` - (Computed) (String) An IPv4 address or IPv6 address as a backup record.
* `description` - (Computed) (String) An optional description for the backup record.