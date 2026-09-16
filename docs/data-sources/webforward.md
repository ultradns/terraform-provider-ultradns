---
subcategory: "Web Forwarding"
layout: "ultradns"
page_title: "ULTRADNS: ultradns_webforward"
description: |-
  Get a web forward from UltraDNS.
---

# Data Source: ultradns_webforward

Gets a web forward by zone and either its UltraDNS-generated GUID or its
source hostname.

## Example Usage

```terraform
data "ultradns_webforward" "by_guid" {
  zone_name = "example.com."
  guid      = "ABC123"
}

data "ultradns_webforward" "by_request_to" {
  zone_name  = "example.com."
  request_to = "www.example.com"
}
```

## Argument Reference

* `zone_name` - (Required) Zone containing the web forward.
* `guid` - (Optional) UltraDNS-generated web forward identifier. Exactly one of `guid` or `request_to` must be set.
* `request_to` - (Optional) Source hostname of the web forward. Exactly one of `guid` or `request_to` must be set.

## Attributes Reference

All web forward resource fields are exported, including `request_to`,
`default_redirect_to`, `default_forward_type`, `relative_forward_type`, and
advanced `record` blocks.
