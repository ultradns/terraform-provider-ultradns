---
subcategory: "ZONE"
layout: "ultradns"
page_title: "ULTRADNS: ultradns_zone"
description: |-
  Get information of zones in UltraDNS.
---

# Data Source: ultradns_zone

Use this data source to get detailed information for your zones.

## Example Usage

```terraform
data "ultradns_zone" "zone" {
    name = "example.com."
}
```


## Argument Reference

The following arguments are supported:

* `name` - (Required) (String) A fully qualified domain name(FQDN) is required.


## Attributes Reference

In addition to all of the arguments above, the following attributes are exported:

* `account_name` - (Computed) (String) Name of the account.
* `type` - (Computed) (String) Type of zone. Valid values are PRIMARY, SECONDARY or ALIAS.
* `status` - (Computed) (String) Displays the status of the zone.
* `dnssec_status` - (Computed) (String) Whether or not the zone is signed with DNSSEC. Valid values are `SIGNED` or `UNSIGNED`.
* `owner` - (Computed) (String) Name of the user that created the zone.
* `resource_record_count` - (Computed) (Integer) Number of records in the zone.
* `last_modified_time` - (Computed) (String) The last date and time the zone was modified, represented in ISO8601 format(`yyyy-MM-ddTHH:mmZ`).<br/>
Example: `2021-12-07T11:25Z`.

#### When `type` is "PRIMARY" below attributes will be exported.

* `inherit` - (Computed) (String) Defines whether this zone should inherit the zone transfer values from the Account, and also specifies which values to inherit.
* `tsig` - (Computed) (Block Set, Max: 1) Nested block describing the TSIG information for the primary zone. The structure of this block is described below.
* `restrict_ip` - (Computed) (Block Set) Nested block describing the list of IP ranges that are allowed to transfer primary zones out using zone transfer protocol (AXFR/IXFR). The structure of this block is described below.
* `notify_addresses` - (Computed) (Block Set) Nested block describing the IP Addresses that are notified when updates are made to the primary zone. The structure of this block is described below.
* `registrar_info` - (Computed) (Block Set) Nested block describing information about the name server configuration for this zone. The structure of this block is described below.

#### When `type` is "SECONDARY" below attributes will be exported.

* `primary_name_server_1` - (Computed) (Block Set) The structure of this block is similar to the block <a href="#nested-name_server-block-has-the-following-structure">`name_server`</a> as described below. It is the info of primary name server.
* `primary_name_server_2` - (Computed) (Block Set) The structure of this block is similar to the block <a href="#nested-name_server-block-has-the-following-structure">`name_server`</a> as described below. It is the info of first backup primary name server.
* `primary_name_server_3` - (Computed) (Block Set) The structure of this block is similar to the block <a href="#nested-name_server-block-has-the-following-structure">`name_server`</a> as described below. It is the info of second backup primary name server.
* `notification_email_address` - (Computed) (String) The Notification Email for a secondary zone.
* `transfer_status_details` - (Computed) (Block Set) Nested block describing the zone transfer details. The structure of this block is described below.

#### When `type` is "ALIAS" below attributes will be exported.

* `original_zone_name` - (Computed) (String) The name of the zone being aliased. The existing zone must be owned by the same account as the new zone.

### Nested `name_server` block has the following structure:

* `ip` - (Required) (String) The IP address of the primary name server for the source zone.
* `tsig_key` - (Optional) (String) The name of the TSIG key, if TSIG is enabled for this name server.
* `tsig_key_value` - (Optional) (String) The TSIG key’s value, if TSIG is enabled for this name server.
* `tsig_algorithm` - (Optional) (String) The hash algorithm used to generate the TSIG key. Valid values are `hmac-md5`, `hmac-sha1`, `hmac-sha224`, `hmac-sha256`, `hmac-sha384`, `hmac-sha512`.

### Nested `tsig` block has the following structure:

* `tsig_key_name` - (Required) (String) The name of the TSIG key for the zone.
* `tsig_key_value` - (Required) (String) The value of the TSIG key for the zone.
* `tsig_algorithm` - (Required) (String) The hash algorithm used to generate the TSIG key. Valid values are `hmac-md5`, `hmac-sha1`, `hmac-sha224`, `hmac-sha256`, `hmac-sha384`, `hmac-sha512`.
* `description` - (Optional) (String) A description of this key.

### Nested `restrict_ip` block has the following structure:

* `start_ip` - (Optional) (String) The start of the IP range that is allowed to transfer this primary zone out using zone transfer protocol.
* `end_ip` - (Optional) (String) The end of the IP range that is allowed to transfer this primary zone out using zone transfer protocol.
* `cidr` - (Optional) (String) The IP Address ranges specified in CIDR.
* `single_ip` - (Optional) (String) The IP Address that is allowed to transfer this primary zone out using zone transfer protocol.
* `comment` - (Optional) (String) A description of this range of IP addresses.

### Nested `notify_addresses` block has the following structure:

* `notify_address` - (Required) (String) The IP Address that is notified when the primary zone is updated.
* `description` - (Optional) (String) A description of this address.

### Nested `registrar_info` block has the following structure:

* `registrar` - (Computed) (String) The name of the domain registrar.
* `who_is_expiration` - (Computed) (String) The date (`yyyy-MM-dd HH:mm:ss.S`) when the domain name registration expires.<br/>
Example: `2022-08-17 03:59:59.0`.
* `name_servers` - (Computed) (Block Set)  Nested block describing the name servers configuration of the zone. The structure of this block is described below.

### Nested `registrar_info.name_servers` block has the following structure:

* `ok` - (Computed) (List String) List of UltraDNS name servers that are configured for this domain.
* `unknown` - (Computed) (List String) List of name servers that are configured for this domain, but are not UltraDNS-managed name servers.
* `missing` - (Computed) (List String) List of UltraDNS name servers that should be configured for this domain, but are not.
* `incorrect` - (Computed) (List String) List of any obsolete UltraDNS name servers that are still configured for this zone.

### Nested `transfer_status_details` block has the following structure:

* `last_refresh` - (Computed) (String) Displays the date (`MM/dd/yy HH:mm:ss tt vvv`) when the last transfer attempt or refresh was.<br/>
Example: `03/18/21 10:20:34 AM GMT`.
* `next_refresh` - (Computed) (String) Displays the date (`MM/dd/yy HH:mm:ss tt vvv`) when the next transfer attempt or refresh is.<br/>
Example: `03/18/21 10:20:34 AM GMT`.
* `last_refresh_status` - (Computed) (String) Displays the status of the last transfer that was attempted. Valid values are `IN_PROGRESS`, `FAILED`, `SUCCESSFUL`
* `last_refresh_status_message` - (Computed) (String) Displays the last transfer’s status message. This is currently shown as failure reason.
