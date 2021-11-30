---
subcategory: "ZONE"
layout: "ultradns"
page_title: "ULTRADNS: ultradns_zone"
description: |-
  Get information of zones in UltraDNS.
---

# Data Source: ultradns_zone

Use this data source to get the information of zones.

## Example Usage

```terraform
data "ultradns_zone" "zone" {
     limit = 1
     query = "name:example.com."
}
```


## Argument Reference

* `query` - (Optional) (String) The query used to construct the list.
* `sort` - (Optional) (String) The sort column used to order the list.
* `reverse` - (Optional) (Bool) Whether the list is ascending (false) or descending (true).
* `limit` - (Optional) (Integer) The maximum number of rows requested.
* `offset` - (Optional) (Integer) The position in the list for the first returned element (0 based).
* `total_count` - (Computed) (Integer) Count of all zones in the system for the specified query.
* `returned_count` - (Computed) (Integer) The number of zones returned.
* `zones` - (Computed) (List) List of the returned zones. Each item in the list matches the structure below.

Structure of each item in list `zones`:

* `name` - (Computed) (String)	Name of the zone, with trailing periods.
* `account_name` - (Computed) (String) Name of the account.
* `type` - (Computed) (String) Type of zone. Valid values are PRIMARY, SECONDARY or ALIAS.

