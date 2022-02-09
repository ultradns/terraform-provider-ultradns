---
subcategory: "SF-Pool"
layout: "ultradns"
page_title: "ULTRADNS: ultradns_sfpool"
description: |-
  Manages the Simple Monitor/Failover (SF) pool records in UltraDNS.
---

# Resource: ultradns_sfpool

Use this resource to manage Simple Monitor/Failover (SF) pool records in UltraDNS

## Example Usage

### Create SF pool record of type A (1)

```terraform
resource "ultradns_sfpool" "a" {
    zone_name = "example.com."
    owner_name = "a"
    record_type = "A"
    ttl = 120
    record_data = ["192.1.1.3"]
    region_failure_sensitivity = "HIGH"
    live_record_state = "NOT_FORCED"
    live_record_description = "Maintainance"
    pool_description = "SF Pool Resource of Type A"
    monitor{
        url = "https://example.com"
        method = "GET"
    }
    backup_record{
        rdata = "192.1.1.4"
        description = "Backup record"
    }
}
```

### Create SF pool record of type AAAA (28)

```terraform
resource "ultradns_sfpool" "aaaa" {
    zone_name = "example.com."
    owner_name = "aaaa"
    record_type = "AAAA"
    ttl = 120
    record_data = ["2001:db8:85a3:0:0:8a2e:370:7334"]
    region_failure_sensitivity = "LOW"
    monitor{
        url = "https://example.com"
        method = "POST"
    }
    backup_record{
        rdata = "2001:db8:85a3:0:0:8a2e:370:7324"
        description = "Backup record"
    }
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
* `ttl` - (Optional) (Integer) The time to live (in seconds) for the record. Must be a value between 0 and 2147483647, inclusive.
* `region_failure_sensitivity` - (Required) (String) Specifies the sensitivity to the number of regions that need to fail for the backup record to be made active. Valid values are `LOW`, `HIGH`.
* `live_record_state` - (Required) (String) Whether or not the live record is currently active. Valid values are:</br> 
`FORCED_INACTIVE` – the backup record should always be returned by a DNS query.</br> 
`NOT_FORCED` – the monitor should determine if the live record or the backup record is returned by a DNS query.
* `live_record_description` - (Optional) (String) An optional description of the live record.
* `pool_description` - (Optional) (String) An optional description of the Simple Failover field.
* `monitor` - (Required) (Block Set) Nested block describing the information for the monitor. The structure of this block is described below.
* `backup_record` - (Optional) (Block Set) Nested block describing the information for the backup record. The structure of this block is described below.

### Nested `monitor` block has the following structure:

* `url` - (Required) (String) Monitored URL. A full URL including: protocol, host, and URI. Required.
Valid protocols are HTTP and HTTPS.
* `method` - (Required) (String) HTTP method used to connect to the monitored URL. Valid values are `GET`, `POST`.
* `transmitted_data` - (Optional) (String) If a monitor is sending a POST, the data that is sent as the body of the request.
* `search_string` - (Optional) (String) If supplied, a string that is checked against the returned data from the request. The monitor fails if the searchString is not present.

### Nested `backup_record` block has the following structure:

* `rdata` - (Required) (String) An IPv4 address or IPv6 address as a backup record.
* `description` - (Optional) (String) An optional description for the backup record.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `status` - (Computed) (String) Current status of the serving record. Valid values are:</br>
`OK` – Live record is being served.</br>
`CRITICAL` – The backup record is being served due to the monitor detecting a failure.</br>
`MANUAL` – The backup record is being served due to user forcing the live record to be inactive.

## Import

Simple Monitor/Failover (SF) pool records can be imported by combining their `owner_name`, `zone_name`, `record_type` using colon.<br/>
Example : `www.example.com.:example.com.:A (1)`.


-> For import, the `owner_name` and `zone_name` must be a FQDN and `record_type` should have the type as well as the corresponding number as shown in the example above.

e.g.,
```
$ terraform import ultradns_sfpool.example "www.example.com.:example.com.:A (1)" 
```