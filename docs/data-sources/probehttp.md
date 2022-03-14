---
subcategory: "HTTP-Probe"
layout: "ultradns"
page_title: "ULTRADNS: ultradns_probe_http"
description: |-
  Get information of HTTP probe records in UltraDNS.
---

# Data Source: ultradns_probe_http

Use this data source to get detailed information of HTTP probe records.

## Example Usage

### Get last created HTTP Probe

```terraform
data "ultradns_probe_http" "http" {
    zone_name = "example.com."
    owner_name = "www"
}
```

### Get HTTP Probe by guid

```terraform
data "ultradns_probe_http" "http" {
    zone_name = "example.com."
    owner_name = "www"
    guid = "06084A729D56C85C"
}
```

### Get HTTP Probe by filtering

```terraform
data "ultradns_probe_http" "http" {
    zone_name = "example.com."
    owner_name = "www"
    interval = "HALF_MINUTE"
	agents = [ "PALO_ALTO", "NEW_YORK"]
}
```

## Argument Reference

The following arguments are supported:

* `zone_name` - (Required) (String) Name of the zone.
* `owner_name` - (Required) (String) The domain name of the owner of the RRSet. Can be either a fully qualified domain name (FQDN) or a relative domain name. If provided as a FQDN, it must be contained within the zone name's FQDN.

The following arguments are used to filter the probes:

* `guid` - (Optional) (String) The internal id for this probe.
* `interval` - (Optional) (String) Length of time between probes in minutes.
* `agents` - (Optional) (String List) Locations that will be used for probing. The exact list must be provided if probes need to be filtered using agents.
* `threshold` - (Optional) (Integer) The probe threshold value.
* `pool_record` - (Optional) (String) The pool record associated with this probe.

->
1) If `guid` is provided, the probe with that guid is returned, and other filter options are not considered.</br>
2) If there is a conflict between probes due to filter options other than `guid`,the last created probe is returned.</br>
3) If no probe is found for the filter options, an error is returned.  

## Attributes Reference

In addition to all of the arguments above, the following attributes are exported:

* `transaction` - (Computed) (Block Set) List of nested blocks describing the http requests sent for a single probe. The structure of this block is described below.
* `total_limit` - (Computed) (Block Set) Nested block describing the total amount of time spent on all http transactions. The structure of this block follows the same structure as the <a href="#nested-limit-block-has-the-following-structure">`limit`</a> block described below.

### Nested `transaction` block has the following structure:

* `method` - (Computed) (String) HTTP method.
* `protocol_version` - (Computed) (String) HTTP protocol version.
* `url` - (Computed) (String) The URL that will be probed.
* `transmitted_data` - (Computed) (String) The data to send to the URL.
* `follow_redirects` - (Computed) (Boolean) Indicates whether or not to follow redirects.
* `expected_response` - (Computed) (String) The Expected Response code for probes to be returned as Successful.
* `search_string` - (Computed) (Block Set) Nested block describing the strings required to be searched for a probe’s successful response. This does not search the status line or headers. The structure of this block is described below.
* `connect_limit` - (Computed) (Block Set) Nested block describing how long the probe stays connected to the resource. The structure of this block follows the same structure as the <a href="#nested-limit-block-has-the-following-structure">`limit`</a> block described below.
* `avg_connect_limit` - (Computed) (Block Set) Nested block describing the mean (average) time to connect for the five most recent probes that have run on each agent. This is only used for Traffic Controller pools. The structure of this block follows the same structure as the <a href="#nested-limit-block-has-the-following-structure">`limit`</a> block described below.
* `run_limit` - (Computed) (Block Set) Nested block describing how long the probe should run. The structure of this block follows the same structure as the <a href="#nested-limit-block-has-the-following-structure">`limit`</a> block described below.
* `avg_run_limit` - (Computed) (Block Set) Nested block describing the mean (average) run-time for the five most recent probes that have run on each agent. This is only used for Traffic Controller pools. The structure of this block follows the same structure as the <a href="#nested-limit-block-has-the-following-structure">`limit`</a> block described below.

### Nested `limit` block has the following structure:

* `warning` - (Computed) (Integer) Indicates how long the HTTP Transactional Probe should wait before a warning is generated.
* `critical` - (Computed) (Integer) Indicates how long the HTTP Transactional Probe should wait before a critical warning is generated.
* `fail` - (Computed) (Integer) Indicates how long the HTTP Transactional Probe should wait before it make the probe to fail.

### Nested `search_string` block has the following structure:

* `warning` - (Computed) (String) If the probe does not find the search string within the response, or does not match it as a regular expression, a warning will be generated. 
* `critical` - (Computed) (String) If the probe does not find the search string within the response, or does not match it as a regular expression, a critical warning will be generated.
* `fail` - (Computed) (String) If the probe does not find the search string within the response, or does not match it as a regular expression, the probe will fail.

-> `warning` and `critical` are only used for Traffic Controller pools

