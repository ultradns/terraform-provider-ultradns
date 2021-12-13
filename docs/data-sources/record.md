---
subcategory: "RECORD"
layout: "ultradns"
page_title: "ULTRADNS: ultradns_record"
description: |-
  Get information of standard DNS records in UltraDNS.
---

## Example Usage

```terraform
data "ultradns_record" "record" {
    zone_name = "example.com."
    owner_name = "www"
    record_type = "A"
}
```


## Argument Reference

The following arguments are supported:

* `zone_name` - (Optional) (String) The query used to construct the list.
* `sort` - (Optional) (String) The sort column name used to order the list.
* `reverse` - (Optional) (Boolean) Whether the list is ascending (false) or descending (true).
* `query` - (Optional) (String) The query used to construct the list.
* `sort` - (Optional) (String) The sort column name used to order the list.
* `reverse` - (Optional) (Boolean) Whether the list is ascending (false) or descending (true).
* `limit` - (Optional) (Integer) The maximum number of rows requested.
* `offset` - (Optional) (Integer) The position in the list for the first returned element (0 based).

## Attributes Reference

In addition to all of the arguments above, the following attributes are exported:

* `total_count` - (Computed) (Integer) Count of all zones in the system for the specified query.
* `returned_count` - (Computed) (Integer) The number of zones returned.
* `record_sets` - (Computed) (List) List of the returned rrsets. Each item in the list matches the structure below.