---
subcategory: "Web Forwarding"
layout: "ultradns"
page_title: "ULTRADNS: ultradns_webforward"
description: |-
  Manages a web forward in UltraDNS.
---

# Resource: ultradns_webforward

Manages a web forward in an UltraDNS zone.

## Example Usage

```terraform
resource "ultradns_webforward" "www" {
  zone_name             = "example.com."
  request_to            = "www.example.com"
  default_redirect_to   = "https://example.com/"
  default_forward_type  = "HTTP_301_REDIRECT"
  relative_forward_type = "PARAMETER_AND_PATH"
}
```

## Argument Reference

* `zone_name` - (Required) Zone containing the web forward. Changing this value forces a new resource.
* `request_to` - (Required) Hostname to forward. Changing this value forces a new resource.
* `default_redirect_to` - (Optional) Default redirect target.
* `default_forward_type` - (Optional) Redirect type: `HTTP_301_REDIRECT`, `HTTP_302_REDIRECT`, `HTTP_303_REDIRECT`, `HTTP_307_REDIRECT`, or `Framed`.
* `relative_forward_type` - (Optional) Request content preserved in the redirect: `PATH`, `PARAMETER`, or `PARAMETER_AND_PATH`.
* `default_redirect_type` - (Optional) Default redirect type returned or accepted by UltraDNS.
* `certificate_id` - (Optional) Certificate identifier for HTTPS web forwarding.
* `certificate_managed_type` - (Optional) Certificate management type for HTTPS web forwarding.
* `certificate_name` - (Optional) Certificate name for HTTPS web forwarding.
* `expiration_days` - (Optional) Certificate expiration period.
* `record` - (Optional) Advanced forwarding rules. The structure is described below.

### Nested `record` block

* `redirect_to` - (Optional) Redirect target for requests matching the rules.
* `forward_type` - (Optional) Redirect type.
* `priority` - (Optional) Rule priority.
* `rule` - (Optional) Header matching rules. The structure is described below.

### Nested `rule` block

* `header` - (Optional) HTTP request header to match.
* `match_criteria` - (Optional) Match operation.
* `value` - (Optional) Value to match.
* `case_insensitive` - (Optional) Whether matching ignores case.

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

* `guid` - UltraDNS-generated web forward identifier.
* `state` - Web forward state.
* `error_description` - Error detail returned by UltraDNS.

## Import

Import a web forward with its zone name and GUID separated by a colon:

```shell
terraform import ultradns_webforward.example "example.com.:ABC123"
```
