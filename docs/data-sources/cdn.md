---
subcategory: "CDN"
layout: "ultradns"
page_title: "ULTRADNS: ultradns_cdn"
description: |-
  Get a single CDN resource by FQDN from UltraDNS.
---

# Data Source: ultradns_cdn

Use this data source to get detailed information of a CDN resource.

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

In addition to all of the arguments above, the following attributes are exported:

* `type` - (Computed) (String) CDN type, for example `BYOD` or `SYNTHETIC`.
* `resource_id` - (Computed) (Integer) Resource identifier.
* `name` - (Computed) (String) CDN resource name.
* `description` - (Computed) (String) CDN resource description.
* `ttl` - (Computed) (Integer) The time to live (in seconds) for the CDN resource.
* `content_type` - (Computed) (String) Content type for the CDN resource.
* `version` - (Computed) (String) Resource version.
* `last_updated` - (Computed) (String) Last update timestamp.
* `owner_name` - (Computed) (String) Resource owner.
* `cdn_providers` - (Computed) (Block List) CDN providers returned by UltraDNS. The structure of this block is described below.
* `config_properties` - (Computed) (Map String) JSON-encoded config properties returned by UltraDNS.
* `preference_properties` - (Computed) (Map String) JSON-encoded preference properties returned by UltraDNS.

### Nested `cdn_providers` block has the following structure:

* `client_cdn_id` - (Computed) (String) Client CDN identifier.
* `cdn_name` - (Computed) (String) Provider display name.
* `description` - (Computed) (String) Provider description.
* `fqdn` - (Computed) (String) Provider FQDN.
