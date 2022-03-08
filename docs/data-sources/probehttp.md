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
* `interval` - (Optional) (String) Length of time between probes in minutes. Valid values are `HALF_MINUTE`, `ONE_MINUTE`, `TWO_MINUTES`, `FIVE_MINUTES`, `TEN_MINUTES`, and `FIFTEEN_MINUTES`.
* `agents` - (Required) (String List) Locations that will be used for probing. Exact list must be provided if probes needed to be filtered using agents.
* `threshold` - (Optional) (Integer) The probe threshold.
* `pool_record` - (Optional) (String) The pool record associated with this probe.

->
1) If `guid` is provided, the probe with that guid is returned.</br>
2) If there is a conflict between probes due to filter options other than `guid`, last created probe is returned.</br>
3) If no probe found for the filter options it returns error.  

## Attributes Reference

In addition to all of the arguments above, the following attributes are exported:

* `transaction` - (Computed) (Block Set) List of nested blocks describing the http requests sent for a single probe. The structure of this block is described below.
* `total_limit` - (Computed) (Block Set) Nested block describing the total amount of time spent on all http transactions. The structure of this block follows the same structure as the <a href="#nested-limit-block-has-the-following-structure">`limit`</a> block described below.

### Nested `transaction` block has the following structure:

* `method` - (Computed) (String) HTTP method. Valid values are `GET` or `POST`.
* `protocol_version` - (Computed) (String) HTTP protocol version. Valid values are `HTTP/1.0`, `HTTP/1.1`, `HTTP/2`.
* `url` - (Computed) (String) URL to probe.
* `transmitted_data` - (Computed) (String) Data to send to URL.
* `follow_redirects` - (Computed) (Boolean) Indicates whether or not to follow redirects. Default set to false.
* `expected_response` - (Computed) (String) The Expected Response code for probes to be returned as Successful.
* `search_string` - (Computed) (Block Set) Nested block describing the strings need to be search on probes successful response. It does not search the status line and headers. The structure of this block is described below.
* `connect_limit` - (Computed) (Block Set) Nested block describing how long the probe stays connected to the resource. The structure of this block follows the same structure as the <a href="#nested-limit-block-has-the-following-structure">`limit`</a> block described below.
* `avg_connect_limit` - (Computed) (Block Set) Nested block describing the mean connect time over the five most recent probes run on each agent. Only used for Traffic Controller Pools. The structure of this block follows the same structure as the <a href="#nested-limit-block-has-the-following-structure">`limit`</a> block described below.
* `run_limit` - (Computed) (Block Set) Nested block describing how long the probe should run. The structure of this block follows the same structure as the <a href="#nested-limit-block-has-the-following-structure">`limit`</a> block described below.
* `avg_run_limit` - (Computed) (Block Set) Nested block describing the mean run time over the five most recent probes run on each agent. Only used for Traffic Controller Pools. The structure of this block follows the same structure as the <a href="#nested-limit-block-has-the-following-structure">`limit`</a> block described below.

### Nested `limit` block has the following structure:

* `warning` - (Computed) (Integer) Time for the HTTP transactional probe for a warning to be generated.
* `critical` - (Computed) (Integer)  Time for the HTTP transactional probe for a critical warning to be generated.
* `fail` - (Computed) (Integer)  Time for the HTTP transactional probe for the probe to fail.

-> `warning` and `critical` are only used for Traffic Controller Pools.

### Nested `search_string` block has the following structure:

* `warning` - (Computed) (String) If the probe does not find the string within the response, or does not match it as a regular expression, a warning will be generated. 
* `critical` - (Computed) (String) If the probe does not find the string within the response, or does not match it as a regular expression, a critical warning will be generated.
* `fail` - (Computed) (String) If the probe does not find the string within the response, or does not match it as a regular expression, the probe will fail.

