---
subcategory: "TC-Pool"
layout: "ultradns"
page_title: "ULTRADNS: ultradns_tcpool"
description: |-
  Manages the Traffic Controller (TC) pool records in UltraDNS.
---

# Resource: ultradns_tcpool

Use this resource to manage Traffic Controller (TC) pool records in UltraDNS.

## Example Usage

### Create TC pool record of type A (1)

```terraform
resource "ultradns_tcpool" "a" {
    zone_name = "example.com."
    owner_name = "a"
    record_type = "A"
    ttl = 120
    pool_description = "TC Pool Resource of Type A"
    run_probes = true
    act_on_probes = true
    failure_threshold = 2
    max_to_lb = 1
    rdata_info{
        priority = 2
        threshold = 1
        rdata = "192.168.1.1"
        failover_delay = 2
        run_probes = true
        state = "ACTIVE"
        weight = 4
    }
    rdata_info{
        priority = 1
        threshold = 1
        rdata = "192.168.1.2"
        failover_delay = 1
        run_probes = false
        state = "NORMAL"
    }
    backup_record{
        rdata = "192.168.1.3"
        failover_delay = 1
    }
}
```

## Argument Reference

The following arguments are supported:

* `zone_name` - (Required) (String) Name of the zone.
* `owner_name` - (Required) (String) The domain name of the owner of the RRSet. Can be either a fully qualified domain name (FQDN) or a relative domain name. If provided as a FQDN, it must be contained within the zone name's FQDN.
* `record_type` - (Required) (String) Must be formatted as a well-known resource record type (A), or the corresponding number for the type (1).<br/>
Below are the supported resource record types with the corresponding number:<br/>
`A (1)`
* `ttl` - (Optional) (Integer) The time to live (in seconds) for the record. Must be a value between 0 and 2147483647, inclusive.
* `pool_description` - (Optional) (String) An optional description of the Traffic Controller (TC) field.
* `run_probes` - (Optional) (Boolean) Indicates whether or not the probes are run for this pool. Default value set to true.
* `act_on_probes` - (Optional) (Boolean) Indicates whether or not pool records will be enabled (true) or disabled (false) when probes are run. Default value set to true.
* `failure_threshold` - (Optional) (Integer) The minimum number of records that must fail for a pool to be labeled 'FAILED'. If the number of failed records in the pool is greater than or equal to the 'Failure Threshold' value, the pool will be labeled FAILED.<br/>
For example, a pool with six priority records, one all-fail record, and the Failure Threshold value is set to four (4). If four or more priority records are not available to serve, the pool will be labeled FAILED, and the all-fail record will be served.<br/>
Valid value between 0 and the number of priority records in the pool.
* `max_to_lb` - (Optional) (Integer) Specifies the maximum number of active servers in a pool. The maximum value is the number of pool records. Default value set to 0.
* `backup_record` - (Required) (Block Set, Max: 1) Nested block describing the information of the backup record for the Traffic Controller pool. The backup record is served if all other records fail. There can be one A record, or a single CNAME record. The structure of this block is described below.
* `rdata_info` - (Required) (Block Set) List of nested blocks describing the pool records. The structure of this block is described below.

### Nested `backup_record` block has the following structure:

* `rdata` - (Required) (String) The IPv4 address or CNAME for the backup record.
* `failover_delay` - (Optional) (Integer) Specifies the time, between 0 – 30 minutes, that the Traffic Controller waits after detecting that the pool record has failed, prior to activating the primary records. Default value set to 0.
* `available_to_serve` - (Computed) (Boolean) Indicates whether the pool's backup record is active and available to serve records.

### Nested `rdata_info` block has the following structure:

* `rdata` - (Required) (String) The IPv4 address or CNAME.
* `priority` - (Required) (Integer) Indicates the serving preference for this pool record.
* `threshold` - (Optional) (Integer) Specifies how many probes must agree before the record state is changed. Default value set to 1.
* `weight` - (Optional) (Integer) Determines the traffic load to send to each server in the Traffic Controller pool. Even integers from 2 to 100. Default value set to 2.
* `failover_delay` - (Optional) (Integer) Specifies the time, between 0 – 30 minutes, that the Traffic Controller waits after detecting that the pool record has failed, prior to activating the secondary records. Default value set to 0.
* `state` - (Optional) (String) The current state of the pool record. Valid values are `NORMAL`, `ACTIVE`, and `INACTIVE`. Default value set to `NORMAL`.
* `run_probes` - (Optional) (Boolean) Indicates whether or not probes are run for this pool record. Default value set to true.
* `available_to_serve` - (Computed) (Boolean) Indicates whether the pool record is active and available to serve records.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `status` - (Computed) (String)  Current status of the serving record. Valid values are:</br>
`OK`- If the number of records serving is equal to the Max Active value, and all the active records are top priority records.</br>
For example, if a pool has a Max Active of 1 and the Priority 1 record is serving.</br>
`WARNING` – If the number of records serving is equal to the Max Active value, and the active records are not top priority records.</br>
For example, if a pool has a Max Active of 1, and the Priority 1 record is not serving and the Priority 2 record is serving.</br>
`CRITICAL` – If the number of records serving is less than the Max Active value, or the All Fail record is being served.</br>
For example, if a pool has a Max Active of 2, and only one record is serving.</br>
`FAILED` - If the FailureThreshold value is 0 or null, and no records are serving, and there is no All Fail record configured.</br>OR</br>If the number of priority records not available to serve equals or exceeds the FailureThreshold’s value.</br>
For example, if the Failure Threshold value is 3, and there are 3 or more Priority Records that are not available to serve.

## Import

Traffic Controller (TC) pool records can be imported by combining their `owner_name`, `zone_name`, and `record_type`, separated by a colon.<br/>
Example : `www.example.com.:example.com.:A (1)`.


-> For import, the `owner_name` and `zone_name` must be a FQDN, and `record_type` should have the type, as well as the corresponding number, as shown in the example below.

Example:
```terraform
$ terraform import ultradns_tcpool.example "www.example.com.:example.com.:A (1)" 
```