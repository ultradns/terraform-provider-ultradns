---
subcategory: "SLB-Pool"
layout: "ultradns"
page_title: "ULTRADNS: ultradns_slbpool"
description: |-
  Get information of Simple Load Balancing (SLB) pool records in UltraDNS.
---

# Data Source: ultradns_slbpool

Use this data source to get detailed information of Simple Load Balancing (SLB) pool records.

## Example Usage

```terraform
data "ultradns_slbpool" "slbpool" {
    zone_name = "example.com."
    owner_name = "www"
    record_type = "A"
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


## Attributes Reference

In addition to all of the arguments above, the following attributes are exported:

* `ttl` - (Computed) (Integer) The time to live (in seconds) for the record. Must be a value between 0 and 2147483647, inclusive.
* `region_failure_sensitivity` - (Computed) (String) Specifies the sensitivity to the number of regions that need to fail for the backup record to be made active.
* `response_method` - (Required) (String) The method used to select which record is returned from the primary record pool. valid values are:</br>
`PRIORITY_HUNT` – The sequence of the records listed in the primary record pool determines the priority. The first record listed is the highest priority record. Once a record starts being served, it will always be served until the probe detects a failure on this record, or, the record is set to FORCED_INACTIVE.</br>
`RANDOM` – A random record is returned from the following set of primary records.</br>
`ROUND_ROBIN` - A record is returned in (a round robin fashion) rotation, based upon the priority of the following active set of records.
* `serving_preference` - (Required) (String) It determines if records will be selected from the Primary Records pool , or, from the All Fail Record. Valid values are:</br>
`AUTO_SELECT`: Serving method switches from serving Primary Records, to All Fail Record based upon probe results, and the Forced State of the Primary Records.</br>
`SERVE_PRIMARY`: Only the Primary Records are served based upon the probe results and the Forced State of the Primary Records.</br>
`SERVE_ALL_FAIL`: Only the All Fail Record will be served, ignoring the probe results and the Forced State of the Primary Records.
* `pool_description` - (Computed) (String) An optional description of the Simple Load Balancing (SLB) field.
* `monitor` - (Computed) (Block Set) Nested block describing the information for the monitor. The structure of this block is described below.
* `all_fail_record` - (Required) (Block Set) Nested block describing the information for the backup record. The structure of this block is described below.
* `rdata_info` - (Required) (Block Set, Max: 5) Nested block describing the pool records. The structure of this block is described below.
* `status` - (Computed) (String)  Current status of the serving record. Valid values are:</br>
`OK`- Priority record(s) are being served.</br>
`WARNING` – One of the priority records is not being served due to the monitor detecting a failure, but there is still a priority record to be served.</br>
`CRITICAL` – The backup All Fail record is being served due to the monitor detecting a failure.

### Nested `monitor` block has the following structure:

* `url` - (Computed) (String) Monitored URL. A full URL including the protocol, host, and URI. Valid protocols are HTTP and HTTPS.
* `method` - (Computed) (String) HTTP method used to connect to the monitored URL.
* `transmitted_data` - (Computed) (String) If a monitor is sending a POST, the data that is sent as the body of the request.
* `search_string` - (Computed) (String) A string that is checked against the returned data from the request. 

### Nested `all_fail_record` block has the following structure:

* `rdata` - (Computed) (String) An IPv4 address or IPv6 address as a backup record.
* `description` - (Computed) (String) An optional description for the backup record.
* `serving` - (Computed) (Boolean) Serving status of the AllFail Record.

### Nested `rdata_info` block has the following structure:

* `rdata` - (Computed) (String) An IPv4 or IPv6 address.
* `description` - (Computed) (String) An optional description of the record in the live pool.
* `forced_state` - (Computed) (String) The Forced State of the record that indicates whether the record needs to be: force served, forced to be inactive, or the force status not being considered (monitoring result decides the record state).
* `probing_enabled` - (Computed) (Boolean) Can be set at the record level to indicate whether probing is required (true) or not (false) for the given record.
* `available_to_serve` - (Computed) (Boolean) Indicates whether the record is available to be served (true) or not (false), based upon the probe results or the forced state of the record.