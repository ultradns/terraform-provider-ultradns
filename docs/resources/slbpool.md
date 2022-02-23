---
subcategory: "SLB-Pool"
layout: "ultradns"
page_title: "ULTRADNS: ultradns_slbpool"
description: |-
  Manages the Simple Load Balancing (SLB) pool records in UltraDNS.
---

# Resource: ultradns_slbpool

Use this resource to manage Simple Load Balancing (SLB) pool records in UltraDNS.

## Example Usage

### Create SLB pool record of type A (1)

```terraform
resource "ultradns_slbpool" "a" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "a"
    record_type = "A"
    ttl = 120
    rdata_info{
        description = "first"
        rdata = "192.201.127.33"
        probing_enabled = false
    }
    rdata_info{
        description = "second"
        rdata = "192.168.1.2"
        probing_enabled = true
    }
    region_failure_sensitivity = "HIGH"
    serving_preference = "AUTO_SELECT"
    response_method = "ROUND_ROBIN"
    monitor{
        url = "https://example.com"
        method = "POST"
    }
    all_fail_record{
        rdata = "192.127.127.33"
    }
}
```

### Create SLB pool record of type AAAA (28)

```terraform
resource "ultradns_slbpool" "aaaa" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "aaaa"
    record_type = "28"
    ttl = 120
    rdata_info{
        description = "first"
        rdata = "2001:db8:85a3:0:0:8a2e:370:7334"
        probing_enabled = false
    }
    region_failure_sensitivity = "LOW"
    serving_preference = "AUTO_SELECT"
    response_method = "ROUND_ROBIN"
    monitor{
        url = "https://example.com"
        method = "GET"
    }
    all_fail_record{
        rdata = "2001:db8:85a3:0:0:8a2e:370:7324"
    }
}
```

## Argument Reference

The following arguments are supported:

* `zone_name` - (Required) (String) Name of the zone.
* `owner_name` - (Required) (String) The domain name of the owner of the RRSet. Can be either a fully qualified domain name (FQDN) or a relative domain name. If provided as a FQDN, it must be contained within the zone name's FQDN.
* `record_type` - (Required) (String) Must be formatted as a well-known resource record type (A or AAAA), or the corresponding number for the type (1 or 28).<br/>
Below are the supported resource record types with the corresponding number:<br/>
`A (1)`
`AAAA (28)`
* `ttl` - (Optional) (Integer) The time to live (in seconds) for the record. Must be a value between 0 and 2147483647, inclusive.
* `region_failure_sensitivity` - (Required) (String) Specifies the sensitivity to the number of regions that need to fail for the backup record to be made active. Valid values are `LOW`, and `HIGH`.
* `response_method` - (Required) (String) The method used to select which record is returned from the primary record pool. Valid values are:</br>
`PRIORITY_HUNT` – The sequence of the records listed in the primary record pool determines the priority. The first record listed is the highest priority record. Once a record starts being served, it will always be served until the probe detects a failure on this record, or, the record is set to FORCED_INACTIVE.</br>
`RANDOM` – A random record is returned from the following set of primary records.</br>
`ROUND_ROBIN` -A record is returned in (a round robin fashion) rotation, based upon the priority of the following active set of records.


-> The top priority record is always returned amongst the set of primary records, when the following conditions are met:</br></br>
1)  Pool Probe is determined to be passing successfully for the record (based upon Threshold configuration), along with the record forced_state is **NOT_FORCED** and probing_enabled at this record level is set to true.</br></br>OR</br></br>
2)  Record forced_state is set to **FORCED_ACTIVE**.

* `serving_preference` - (Required) (String) It determines if records will be selected from the Primary Records pool or from the All Fail Record. Valid values are:</br>
`AUTO_SELECT`: Serving method switches from serving Primary Records, to All Fail Record based upon probe results, and the Forced State of the Primary Records.</br>
`SERVE_PRIMARY`: Only the Primary Records are served based upon the probe results and the Forced State of the Primary Records.</br>
`SERVE_ALL_FAIL`: Only the All Fail Record will be served, ignoring the probe results and the Forced State of the Primary Records.


-> Please be aware that it may take up to 15 seconds to see the updated value, after switching between Auto Select/Serve Primary and Serve All Fail.


* `pool_description` - (Optional) (String) An optional description of the Simple Load Balancing (SLB) field.
* `monitor` - (Required) (Block Set) Nested block describing the information for the monitor. The structure of this block is described below.
* `all_fail_record` - (Required) (Block Set) Nested block describing the information for the backup record. The structure of this block is described below.
* `rdata_info` - (Required) (Block Set, Max: 5) Nested block describing the pool records. The structure of this block is described below.


### Nested `monitor` block has the following structure:

* `url` - (Required) (String) Monitored URL. A full URL including the protocol, host, and URI. Valid protocols are HTTP and HTTPS.
* `method` - (Required) (String) HTTP method used to connect to the monitored URL. Valid values are `GET`, and `POST`.
* `transmitted_data` - (Optional) (String) If a monitor is sending a POST, this is the data sent as the body of the request.
* `search_string` - (Optional) (String) A string that is checked against the returned data from the request. The monitor fails if the search string is not present.

### Nested `all_fail_record` block has the following structure:

* `rdata` - (Required) (String) An IPv4 or IPv6 address as a backup record.
* `description` - (Optional) (String) An optional description for the backup record.
* `serving` - (Computed) (Boolean) Serving status of the AllFail Record.

### Nested `rdata_info` block has the following structure:

* `rdata` - (Required) (String) An IPv4 address or IPv6 address.
* `description` - (Optional) (String) An optional description of the record in the live pool.
* `forced_state` - (Optional) (String) The Forced State of the record that indicates whether the record needs to be: force served, forced to be inactive, or the force status not being considered (monitoring result decides the record state). Valid values are `FORCED_ACTIVE`, `FORCED_INACTIVE`, or `NOT_FORCED`. Default set to `NOT_FORCED`.
* `probing_enabled` - (Optional) (Boolean) Can be set at the record level to indicate whether probing is required (true) or not (false) for the given record. Default set to true.
* `available_to_serve` - (Computed) (Boolean) Indicates whether the record is available to be served (true) or not (false), based upon the probe results or the forced state of the record.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `status` - (Computed) (String)  Current status of the serving record. Valid values are:</br>
`OK`- Priority record(s) are being served.</br>
`WARNING` – One of the priority records is not being served due to the monitor detecting a failure, but there is still a priority record to be served.</br>
`CRITICAL` – The backup All Fail record is being served due to the monitor detecting a failure.

## Import

Simple Load Balancing (SLB) pool records can be imported by combining their `owner_name`, `zone_name`, and `record_type`, separated by a colon.<br/>
Example : `www.example.com.:example.com.:A (1)`.


-> For import, the `owner_name` and `zone_name` must be a FQDN, and `record_type` should have the type, as well as the corresponding number, as shown in the example below.

Example:
```terraform
$ terraform import ultradns_slbpool.example "www.example.com.:example.com.:A (1)" 
```