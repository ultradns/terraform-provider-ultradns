---
subcategory: "ZONE"
layout: "ultradns"
page_title: "ULTRADNS: ultradns_zone"
description: |-
  Manges the zones in UltraDNS.
---

# Resource: ultradns_zone

Use this resource to manage zones in UltraDNS

## Example Usage

### Create Primary Zone

```terraform
resource "ultradns_zone" "primary" {
    name        = "example.com."
    account_name = "account"
    type        = "PRIMARY"
    primary_create_info {
        create_type = "NEW"
        restrict_ip{
            single_ip = "192.168.1.1"
        }
        restrict_ip{
            single_ip = "192.168.1.2"
        }
        notify_addresses {
            notify_address = "192.168.1.3"
        }
        notify_addresses {
            notify_address = "192.168.1.4"
        }
        tsig {
            tsig_key_name = "example.com.0.325349282751.key."
            tsig_key_value = "ZWFlY2U1MTBlRmM2Y0NGQ5MTlmYTdmZTE0Njc="
            tsig_algorithm  = "hmac-md5"
            description = "description"
        }
    }
}
```

### Create Secondary Zone

```terraform
resource "ultradns_zone" "secondary" {
    name        = "example.com."
    account_name = "account"
    type        = "SECONDARY"
    secondary_create_info {
        primary_name_server_1 {
            ip = "192.168.1.1"
        } 
        notification_email_address = "test@example.com"
    }
}
```

### Create Alias Zone

```terraform
resource "ultradns_zone" "alias" {
    name        = "example.com."
    account_name = "account"
    type        = "ALIAS"
    alias_create_info {
        original_zone_name = "ultradns.com."
  }
}
```


## Argument Reference

The following arguments are supported:

* `name` - (Required) (String) Name of the zone.
* `account_name` - (Required) (String) 	Name of the account. It must be provided, but it can also be sourced from the `ULTRADNS_ACCOUNT` environment variable.
* `type` - (Required) (String) This is the type of the zone. Valid values are `PRIMARY`, `SECONDARY` or `ALIAS`.
* `change_comment` - (Optional) (String) This is used to provide comments on updates.
* `primary_create_info` - (Optional) (Block Set, Max: 1) Nested block describing the info of primary zone. The structure of this block is described below.
* `secondary_create_info` - (Optional) (Block Set, Max: 1)
Nested block describing the info of secondary zone. The structure of this block is described below.
* `alias_create_info` - (Optional) (Block Set, Max: 1)
Nested block describing the info of alias zone. The structure of this block is described below.

#### When `type` is "PRIMARY" <a href="#nested-primary_create_info-block-has-the-following-structure">`primary_create_info`</a> needs to be provided.
#### When `type` is "SECONDARY" <a href="#nested-secondary_create_info-block-has-the-following-structure">`secondary_create_info`</a> needs to be provided.
#### When `type` is "ALIAS" <a href="#nested-alias_create_info-block-has-the-following-structure">`alias_create_info`</a> needs to be provided.

### Nested `primary_create_info` block has the following structure:

* `create_type` - (Required) (String) Indicates the method for creating the primary zone. Valid values are `NEW`, `COPY`, `TRANSFER`.
* `force_import` - (Optional) (Boolean) Indicates whether or not to move existing records from zones into this new zone. Default set to false.
* `original_zone_name` - (Optional) (String) The name of the zone being copied. The existing zone must be owned by the same account as the new zone. It needs to be provided if <a href="#create_type">`create_type`</a> is `COPY`.
* `inherit` - (Optional) (String) Defines whether this zone should inherit the zone transfer values from the Account, and also specifies which values to inherit. Valid values are `ALL`, `NONE`, any combination of `IP_RANGE`, `NOTIFY_IP`, `TSIG`. Separate multiple values with a comma.<br/>
Example: `IP_RANGE, NOTIFY_IP`
* `name_server` - (Optional) (Block Set, Max: 1) Nested block describing the Primary zone's name server. It needs to be provided if <a href="#create_type">`create_type`</a> is `TRANSFER`.
* `tsig` - (Optional) (Block Set, Max: 1) Nested block describing the TSIG information for the primary zone. The structure of this block is described below.
* `restrict_ip` - (Optional) (Block Set) Nested block describing the list of IP ranges that are allowed to transfer primary zones out using zone transfer protocol (AXFR/IXFR). The structure of this block is described below.
* `notify_addresses` - (Optional) (Block Set) Nested block describing the addresses that are notified when updates are made to the primary zone. The structure of this block is described below.

### Nested `name_server` block has the following structure:

* `ip` - (Required) (String) The IPv4 or IPv6 address of the primary name server for the source zone.
* `tsig_key` - (Optional) (String) The name of the TSIG key, if TSIG is enabled for this name server.
* `tsig_key_value` - (Optional) (String) The TSIG key’s value, if TSIG is enabled for this name server.
* `tsig_algorithm` - (Optional) (String) The hash algorithm used to generate the TSIG key. Valid values are `hmac-md5`, `hmac-sha1`, `hmac-sha224`, `hmac-sha256`, `hmac-sha384`, `hmac-sha512`.

### Nested `tsig` block has the following structure:

The following tsig values are required if TSIG is enabled for the zone.

* `tsig_key_name` - (Required) (String) The name of the TSIG key for the zone.
* `tsig_key_value` - (Required) (String) The value of the TSIG key for the zone.
* `tsig_algorithm` - (Required) (String) The hash algorithm used to generate the TSIG key. Valid values are `hmac-md5`, `hmac-sha1`, `hmac-sha224`, `hmac-sha256`, `hmac-sha384`, `hmac-sha512`.
* `description` - (Optional) (String) A description of this key.

### Nested `restrict_ip` block has the following structure:

* `start_ip` - (Optional) (String) The start of the IPv4 or IPv6 range that is allowed to transfer this primary zone out using zone transfer protocol.
* `end_ip` - (Optional) (String) The end of the IPv4 or IPv6 range that is allowed to transfer this primary zone out using zone transfer protocol.
* `cidr` - (Optional) (String) The IP Address ranges specified in CIDR.
* `single_ip` - (Optional) (String) The IP Address that is allowed to transfer this primary zone out using zone transfer protocol.
* `comment` - (Optional) (String) A description of this range of IP addresses.

### Nested `notify_addresses` block has the following structure:

* `notify_address` - (Required) (String) The IP Address that is notified when the primary zone is updated.
* `description` - (Optional) (String) A description of this IP Address.

### Nested `secondary_create_info` block has the following structure:

* `primary_name_server_1` - (Required) (Block Set) The structure of this block follows the same structure as the <a href="#nested-name_server-block-has-the-following-structure">`name_server`</a> block described above. It is the info of primary name server.
* `primary_name_server_2` - (Optional) (Block Set) The structure of this block follows the same structure as the <a href="#nested-name_server-block-has-the-following-structure">`name_server`</a> block described above. It is the info of first backup primary name server.
* `primary_name_server_3` - (Optional) (Block Set) The structure of this block follows the same structure as the <a href="#nested-name_server-block-has-the-following-structure">`name_server`</a> block described above. It is the info of second backup primary name server.
* `notification_email_address` - (Optional) (String) The Notification Email for a secondary zone.

### Nested `alias_create_info` block has the following structure:

* `original_zone_name` - (Required) (String) The name of the zone being aliased. The existing zone must be owned by the same account as the new zone.


## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `status` - (Computed) (String) Display the status of the zone.
* `dnssec_status` - (Computed) (String) Whether or not the zone is signed with DNSSEC. Valid values are `SIGNED` or `UNSIGNED`.
* `owner` - (Computed) (String) Name of the user that created the zone.
* `resource_record_count` - (Computed) (Integer) Number of records in the zone.
* `last_modified_time` - (Computed) (String) The last date and time the zone was modified, represented in ISO8601 format (`yyyy-MM-ddTHH:mmZ`).<br/>
Example: `2021-12-07T11:25Z`.
* `registrar_info` - (Computed) (Block Set) Nested block describing information about the name server configuration for this zone. The structure of this block is described below.
* `transfer_status_details` - (Computed) (Block Set) Nested block describing the zone transfer details. The structure of this block is described below.

#### When `type` is "PRIMARY" <a href="#nested-registrar_info-block-has-the-following-structure">`registrar_info`</a> will be exported.
#### When `type` is "SECONDARY" <a href="#nested-transfer_status_details-block-has-the-following-structure">`transfer_status_details`</a> will be exported.

### Nested `registrar_info` block has the following structure:

* `registrar` - (Computed) (String) The name of the domain registrar.
* `who_is_expiration` - (Computed) (String) The date  (`yyyy-MM-dd HH:mm:ss.S`) when the domain name registration expires.<br/>
Example: `2022-08-17 03:59:59.0`.
* `name_servers` - (Computed) (Block Set)  Nested block describing the name servers configuration of the zone. The structure of this block is described below.

### Nested `registrar_info.name_servers` block has the following structure:

* `ok` - (Computed) (List String) List of UltraDNS name servers that are configured for this domain.
* `unknown` - (Computed) (List String) List of name servers that are configured for this domain, but are not UltraDNS managed name servers.
* `missing` - (Computed) (List String) List of UltraDNS name servers that should be configured for this domain, but are not.
* `incorrect` - (Computed) (List String) List of any obsolete UltraDNS name servers that are still configured for this zone.

### Nested `transfer_status_details` block has the following structure:

* `last_refresh` - (Computed) (String) Displays the date (`MM/dd/yy HH:mm:ss tt vvv`) when the last transfer attempt or refresh was.<br/>
Example: `03/18/21 10:20:34 AM GMT`.
* `next_refresh` - (Computed) (String) Displays the date (`MM/dd/yy HH:mm:ss tt vvv`) when the next transfer attempt or refresh is.<br/>
Example: `03/18/21 10:20:34 AM GMT`.
* `last_refresh_status` - (Computed) (String) Displays the status of the last transfer that was attempted. Valid values are `IN_PROGRESS`, `FAILED`, `SUCCESSFUL`.
* `last_refresh_status_message` - (Computed) (String) Displays the last transfer’s status message. This is currently shown as failure reason.

## Import

Zones can be imported using their name (must be a FQDN), e.g.,

```
$ terraform import ultradns_zone.example "example.com." 
```