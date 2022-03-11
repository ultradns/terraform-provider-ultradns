---
subcategory: "TC-Pool"
layout: "ultradns"
page_title: "ULTRADNS: ultradns_tcpool"
description: |-
  Get information of Traffic Controller (TC) pool records in UltraDNS.
---

# Data Source: ultradns_tcpool

Use this data source to get detailed information of Traffic Controller (TC) pool records.

## Example Usage

```terraform
data "ultradns_tcpool" "tcpool" {
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

* `ttl` - (Computed) (Integer) The time to live (in seconds) for the record. Must be a value between 0 and 2147483647, inclusive.
* `pool_description` - (Computed) (String) An optional description of the Traffic Controller (TC) field.
* `run_probes` - (Computed) (Boolean) Indicates whether or not the probes are run for this pool.
* `act_on_probes` - (Computed) (Boolean) Indicates whether or not pool records will be enabled (true) or disabled (false) when probes are run.
* `failure_threshold` - (Computed) (Integer) The minimum number of records that must fail for a pool to be labeled 'FAILED'. If the number of failed records in the pool is greater than or equal to the 'Failure Threshold' value, the pool will be labeled FAILED.<br/>
For example, a pool with six priority records, one all-fail record, and the Failure Threshold value is set to four (4). If four or more priority records are not available to serve, the pool will be labeled FAILED, and the all-fail record will be served.
* `max_to_lb` - (Computed) (Integer) Specifies the maximum number of active servers in a pool. The maximum value is the number of pool records.
* `backup_record` - (Computed) (Block Set, Max: 1) Nested block describing the information of the backup record for the Traffic Controller pool. The backup record is served if all other records fail. There can be one or more A records used as backup records, or a single CNAME record. The structure of this block is described below.
* `rdata_info` - (Computed) (Block Set) List of nested blocks describing the pool records. The structure of this block is described below.
* `status` - (Computed) (String)  Current status of the serving record. Valid values are:</br>
`OK`- If the number of records serving is equal to the Max Active value, and all the active records are top priority records.</br>
For example, if a pool has a Max Active of 1 and the Priority 1 record is serving.</br>
`WARNING` – If the number of records serving is equal to the Max Active value, and the active records are not top priority records.</br>
For example, if a pool has a Max Active of 1 and the Priority 1 records is not serving and the Priority 2 record is serving.</br>
`CRITICAL` – If the number of records serving is less than the Max Active value, or the All Fail record is being served.</br>
For example, if a pool has a Max Active of 2, and only one record is serving.</br>
`FAILED` - If the FailureThreshold value is 0 or null, and no records are serving, and there is no All Fail record configured.</br>OR</br>If the number of priority records not available to serve equals or exceeds the FailureThreshold’s value.</br>
For example, if the Failure Threshold value is 3, and there are 3 or more Priority Records that are not available to serve.

### Nested `backup_record` block has the following structure:

* `rdata` - (Computed) (String) The IPv4 address or CNAME for the backup record.
* `failover_delay` - (Computed) (Integer) Specifies the time, between 0 – 30 minutes, that the Traffic Controller waits after detecting that the pool record has failed, prior to activating the primary records. Default value set to 0.
* `available_to_serve` - (Computed) (Boolean) Indicates whether the pool's backup record is active and available to serve records.

### Nested `rdata_info` block has the following structure:

* `rdata` - (Computed) (String) The IPv4 address or CNAME.
* `priority` - (Computed) (Integer) Indicates the serving preference for this pool record.
* `threshold` - (Computed) (Integer) Specifies how many probes must agree before the record state is changed.
* `weight` - (Optional) (Integer) Determines the traffic load to send to each server in the Traffic Controller pool. Even integers from 2 to 100.
* `failover_delay` - (Computed) (Integer) Specifies the time, between 0 – 30 minutes, that the Traffic Controller waits after detecting that the pool record has failed, prior to activating the secondary records. Default value set to 0.
* `state` - (Computed) (String) The current state of the pool record.
* `run_probes` - (Computed) (Boolean) Indicates whether or not probes are run for this pool record.
* `available_to_serve` - (Computed) (Boolean) Indicates whether the pool record is active and available to serve records.