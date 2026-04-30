---
subcategory: "CDN"
layout: "ultradns"
page_title: "ULTRADNS: ultradns_cdn"
description: |-
  Manages CDN resources in UltraDNS.
---

# Resource: ultradns_cdn

Use this resource to manage CDN resources in UltraDNS

## Example Usage

### Create BYOD CDN

```terraform
resource "ultradns_cdn" "byod" {
  account_name = "my-account"
  fqdn         = "cdn.example.com."
  type         = "BYOD"
  name         = "example-byod"
  description  = "BYOD CDN resource"
  ttl          = 300
  content_type = "dynamic"

  cdn_providers {
    client_cdn_id = "TEST-PROVIDER-1"
    cdn_name      = "BYOD Provider"
    fqdn          = "byod-cdn-provider.example.com."
  }

  config_properties = {
    cdnEnablementMap = jsonencode({
      worldDefault = ["TEST-PROVIDER-1"]
      asnOverrides = {}
      continents   = {}
    })
    trafficDistribution = jsonencode({
      worldDefault = {
        options = [{
          name        = "performance_based"
          description = "Performance-based routing"
          equalWeight = true
          distribution = [{
            id = "TEST-PROVIDER-1"
          }]
        }]
      }
    })
  }

  preference_properties = {
    availabilityThresholds = jsonencode({
      world      = 85
      continents = {}
    })
    performanceFiltering = jsonencode({
      world      = { mode = "absolute" }
      continents = {}
    })
    enabledSubdivisionCountries = jsonencode({
      continents = {}
    })
  }
}
```

### Create Synthetic CDN

```terraform
resource "ultradns_cdn" "synthetic" {
  account_name = "my-account"
  fqdn         = "synthetic.example.com."
  type         = "SYNTHETIC"
  name         = "example-synthetic"
  description  = "Synthetic CDN resource"
  ttl          = 300
  content_type = "dynamic"

  cdn_providers {
    client_cdn_id = "TEST-PROVIDER-1"
    cdn_name      = "Synthetic Provider"
    fqdn          = "synthetic-cdn-provider.example.com."
  }

  config_properties = {
    checksConfig = jsonencode({
      protocol = "HTTP"
      path     = "/"
    })
    cdnEnablementMap = jsonencode({
      worldDefault = ["ultradns"]
      asnOverrides = {}
      continents   = {}
    })
    trafficDistribution = jsonencode({
      worldDefault = {
        options = [{
          name        = "performance_based"
          description = "Performance-based routing"
          equalWeight = true
          distribution = [{
            id = "ultradns"
          }]
        }]
      }
    })
  }

  preference_properties = {
    availabilityThresholds = jsonencode({
      world      = 85
      continents = {}
    })
    performanceFiltering = jsonencode({
      world      = { mode = "absolute" }
      continents = {}
    })
    enabledSubdivisionCountries = jsonencode({
      continents = {}
    })
  }
}
```

## Argument Reference

The following arguments are supported:

* `account_name` - (Required) (String) Account name. Changing this value forces a new resource.
* `fqdn` - (Required) (String) Fully qualified domain name for the CDN resource. Changing this value forces a new resource.
* `type` - (Required) (String) CDN type. Valid values are `BYOD` and `SYNTHETIC`.
* `name` - (Required) (String) CDN resource name.
* `description` - (Optional) (String) CDN resource description.
* `ttl` - (Optional) (Integer) The time to live (in seconds) for the CDN resource.
* `content_type` - (Optional) (String) Content type for the CDN resource.
* `cdn_providers` - (Required) (Block List, Min: 1) CDN providers for this CDN resource. The structure of this block is described below.
* `config_properties` - (Required) (Map String) JSON-encoded config properties sent to UltraDNS.
* `preference_properties` - (Required) (Map String) JSON-encoded preference properties sent to UltraDNS.

### Nested `cdn_providers` block has the following structure:

* `client_cdn_id` - (Required) (String) Client CDN identifier. Must match `^[A-Za-z0-9\-_]{1,64}$`.
* `cdn_name` - (Optional) (String) Provider display name.
* `description` - (Optional) (String) Provider description.
* `fqdn` - (Optional) (String) Provider FQDN.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `resource_id` - (Computed) (Integer) Resource identifier.
* `version` - (Computed) (String) Resource version.
* `last_updated` - (Computed) (String) Last update timestamp.
* `owner_name` - (Computed) (String) Resource owner.

## Import

CDN resources can be imported using `account_name` and `fqdn` separated by a colon.

Example:

```terraform
$ terraform import ultradns_cdn.example "my-account:cdn.example.com."
```