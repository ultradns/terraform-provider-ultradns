---
subcategory: "CDN"
layout: "ultradns"
page_title: "ULTRADNS: ultradns_cdn"
description: |-
  Get a single CDN resource by FQDN from UltraDNS.
---

# Data Source: ultradns_cdn

Use this data source to get detailed information about a single CDN resource.

## Example Usage

```terraform
data "ultradns_cdn" "example" {
  account_name = "my-account"
  fqdn         = "cdn.example.com."
}
```

## Argument Reference

The following arguments are supported:

* `account_name` - (Required) (String) Account name.
* `fqdn` - (Required) (String) Fully qualified domain name for the CDN resource.

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

* `type` - (Computed) (String) CDN type, for example `BYOD` or `SYNTHETIC`.
* `resource_id` - (Computed) (Integer) Resource identifier.
* `name` - (Computed) (String) CDN resource name.
* `description` - (Computed) (String) CDN resource description.
