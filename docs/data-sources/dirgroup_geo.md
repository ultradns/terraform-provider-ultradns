---
subcategory: "Geo"
layout: "ultradns"
page_title: "ULTRADNS: ultradns_dirgroup_geo"
description: |-
  Get information of Account-level GeoIP Groups in UltraDNS.
---

# Data Source: ultradns_dirgroup_geo

Use this data source to get detailed information for your zones.

## Example Usage

```terraform
data "ultradns_dirgroup_geo" "eu_group" {
    name = "EU Group"
}
```


## Argument Reference

The following arguments are supported:

* `name` - (Required) (String) Name of the geoIP group.


## Attributes Reference

In addition to the arguments above, the following attributes are exported:

* `account_name` - (Computed) (String) Name of the account.
* `codes` - (Computed) (String List) The codes for the geographical territories that make up this group.
* `description` - (Computed) (String) The description for the group.

