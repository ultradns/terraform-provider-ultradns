---
subcategory: "CDN"
layout: "ultradns"
page_title: "ULTRADNS: ultradns_cdn_list"
description: |-
  List CDN resources for an UltraDNS account.
---

# Data Source: ultradns_cdn_list

Use this data source to list CDN resources in UltraDNS.

## Example Usage

```terraform
data "ultradns_cdn_list" "all" {
  account_name = "my-account"
  page         = 1
  size         = 100
}
```

## Argument Reference

The following arguments are supported:

* `account_name` - (Required) (String) Account name.
* `page` - (Optional) (Integer) Page number. Default: `1`.
* `size` - (Optional) (Integer) Page size. Default: `100`.

## Attributes Reference

In addition to all of the arguments above, the following attributes are exported:

* `total_pages` - (Computed) (Integer) Total number of pages.
* `total_elements` - (Computed) (Integer) Total number of CDN resources.
* `cdns` - (Computed) (Block List) List of CDN resources. The structure of this block is described below.

### Nested `cdns` block has the following structure:

* `fqdn` - (Computed) (String) Fully qualified domain name.
* `type` - (Computed) (String) CDN type, for example `BYOD` or `SYNTHETIC`.
* `resource_id` - (Computed) (Integer) Resource identifier.
* `name` - (Computed) (String) CDN resource name.
