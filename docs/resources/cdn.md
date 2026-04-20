---
subcategory: "CDN"
layout: "ultradns"
page_title: "ULTRADNS: ultradns_cdn"
description: |-
  Manages CDN resources in UltraDNS.
---

# Resource: ultradns_cdn

Use this resource to create and manage a CDN resource.

## Example Usage

```terraform
resource "ultradns_cdn" "byod" {
  account_name = "my-account"
  fqdn         = "cdn.example.com."
  type         = "BYOD"
  name         = "example-byod"
  description  = "BYOD CDN resource"
}

resource "ultradns_cdn" "synthetic" {
  account_name = "my-account"
  fqdn         = "synthetic.example.com."
  type         = "SYNTHETIC"
  name         = "example-synthetic"
}
```

## Argument Reference

The following arguments are supported:

* `account_name` - (Required) (String) Account name. Changing this value forces a new resource.
* `fqdn` - (Required) (String) Fully qualified domain name for the CDN resource. Changing this value forces a new resource.
* `type` - (Required) (String) CDN type. Valid values are `BYOD` and `SYNTHETIC`.
* `name` - (Required) (String) CDN resource name.
* `description` - (Optional) (String) CDN resource description.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `resource_id` - (Computed) (Integer) Resource identifier.

## Import

CDN resources can be imported by combining `account_name` and `fqdn`, separated by a colon.

Example:

```terraform
$ terraform import ultradns_cdn.example "my-account:cdn.example.com."
```