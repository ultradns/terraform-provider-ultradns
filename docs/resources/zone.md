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
    }
}
```

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

## Argument Reference

The following arguments are supported:

* `name` - (Required) (String) A Fully Qualified Domain Name should be provided.
* `account_name` - (Required) (String) 	Name of the account.
* `type` - (Required) (String) This is the type of the zone.Valid values are `PRIMARY`, `SECONDARY` or `ALIAS`.
* `change_comment` - (Optional) (String) This is used to provide comments on updates.
* `primary_create_info` - (Optional) (Block Set, Max: 1) Nested block describing the info of primary zone. The structure of this block is described below.
* `secondary_create_info` - (Optional) (Block Set, Max: 1)
Nested block describing the info of secondary zone. The structure of this block is described below.
* `alias_create_info` - (Optional) (Block Set, Max: 1)
Nested block describing the info of alias zone. The structure of this block is described below.

When `type` is "PRIMARY" `primary_create_info` should be provided.

Nested `primary_create_info` block have the following structure:

* `create_type` - (Required) (String) Indicates the method for creating the primary zone.Valid values are `NEW`, `COPY`, `TRANSFER`.
* `force_import` - (Optional) (Bool) Whether or not to move existing records from zones into this new zone. Default set to false.
* `original_zone_name` - (Optional) (String) The name of the zone being copied. The existing zone must be owned by the same account as the new zone. It is provided if `create_type` is "COPY".
* `inherit` - (Optional) (String) Defines whether this zone should inherit the zone transfer values from the Account, and also specifies which values to inherit.
* `name_server` - (Optional) (Block Set, Max: 1) Nested block describing Primary zone's name server.It is provided if `create_type` is "TRANSFER".
* `tsig` - (Optional) (Block Set, Max: 1) Nested block describing the TSIG information for the primary zone. The structure of this block is described below.
* `restrict_ip` - (Optional) (Block Set) Nested block describing the list of IP ranges that are allowed to transfer primary zones out using zone transfer protocol (AXFR/IXFR). The structure of this block is described below.
* `notify_addresses` - (Optional) (Block Set) Nested block describing the addresses that are notified when updates are made to the primary zone. The structure of this block is described below.

Nested `name_server` block have the following structure:

* `ip` - (Required) (String) The IP address of the primary name server for the source zone.
* `tsig_key` - (Optional) (String) If TSIG is enabled for this name server, the name of the TSIG key.
* `tsig_key_value` - (Optional) (String) If TSIG is enabled for this name server, the TSIG key's value.
* `tsig_algorithm` - (Optional) (String) The hash algorithm used to generate the TSIG key.Valid values are `hmac-md5`, `hmac-sha1`, `hmac-sha224`, `hmac-sha256`, `hmac-sha384`, `hmac-sha512`.

Nested `tsig` block have the following structure:

* `tsig_key_name` - (Required) (String) The name of the TSIG key for the zone.
* `tsig_key_value` - (Required) (String) The value of the TSIG key for the zone.
* `tsig_algorithm` - (Required) (String) The hash algorithm used to generate the TSIG key.Valid values are `hmac-md5`, `hmac-sha1`, `hmac-sha224`, `hmac-sha256`, `hmac-sha384`, `hmac-sha512`.
* `description` - (Optional) (String) A description of this key.


Nested `restrict_ip` block have the following structure:

* `start_ip` - (Optional) (String) The start of the IP range that is allowed to transfer this primary zone out using zone transfer protocol.
* `end_ip` - (Optional) (String) The start of the IP range that is allowed to transfer this primary zone out using zone transfer protocol.
* `cidr` - (Optional) (String) TThe start of the IP range that is allowed to transfer this primary zone out using zone transfer protocol.
* `single_ip` - (Optional) (String) The IP that is allowed to transfer this primary zone out using zone transfer protocol.
* `comment` - (Optional) (String) A description of this range of IP addresses.

Nested `notify_addresses` block have the following structure:

* `notify_address` - (Required) (String) The IP address that is notified when the primary zone is updated.
* `description` - (Optional) (String) A description of this address.


When `type` is "SECONDARY" `secondary_create_info` should be provided.

Nested `secondary_create_info` block have the following structure:

* `primary_name_server_1` - (Required) (Block Set) Nested block describing the primary name servers of the source zone for the secondary zone. The structure of this block is similar to the block `name_server` as described above.
* `primary_name_server_2` - (Optional) (Block Set) Nested block describing the primary name servers of the source zone for the secondary zone. The structure of this block is similar to the block `name_server` as described above.
* `primary_name_server_3` - (Optional) (Block Set) Nested block describing the primary name servers of the source zone for the secondary zone. The structure of this block is similar to the block `name_server` as described above.
* `notification_email_address` - (Optional) (String) The Notification Email for a secondary zone.

When `type` is "ALIAS" `alias_create_info` should be provided.

Nested `alias_create_info` block have the following structure:

* `original_zone_name` - (Required) (String) The name of the zone being aliased. The existing zone must be owned by the same account as the new zone.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `status` - (Computed) (String) Display the status of the zone.
* `dnssec_status` - (Computed) (String) Whether or not the zone is signed with DNSSEC. Valid values are `SIGNED` or `UNSIGNED`.
* `owner` - (Computed) (String) Name of the user that created the zone.
* `resource_record_count` - (Computed) (Integer) Number of records in the zone
* `last_modified_time` - (Computed) (String) The last date and time the zone was modified, represented in ISO8601 format.
* `registrar_info` - (Computed) (Block Set) Nested block describing information about the name server configuration for this zone. The structure of this block is described below.
* `transfer_status_details` - (Computed) (Block Set) Nested block describing the zone transfer details. The structure of this block is described below.


When `type` is "PRIMARY" `registrar_info` will be exported.

Nested `registrar_info` block have the following structure:

* `registrar` - (Computed) (String) The name of the domain registrar.
* `who_is_expiration` - (Computed) (String) The date when the domain name registration expires.
* `name_servers` - (Computed) (Block Set)  Nested block describing the name servers configuration of the zone. The structure of this block is described below.

Nested `registrar_info.name_servers` block have the following structure:

* `ok` - (Computed) (List String) List of UltraDNS name servers that are configured for this domain.
* `unknown` - (Computed) (List String) List of name servers that are configured for this domain, but are not UltraDNS-managed name servers.
* `missing` - (Computed) (List String) List of UltraDNS name servers that should be configured for this domain, but are not.
* `incorrect` - (Computed) (List String) List of any obsolete UltraDNS name servers that are still configured for this zone.

When `type` is "SECONDARY" `transfer_status_details` will be exported.

Nested `transfer_status_details` block have the following structure:

* `last_refresh` - (Computed) (String) Displays when the last transfer attempt or refresh was.
* `next_refresh` - (Computed) (String) Displays when the next transfer attempt or refresh is.
* `last_refresh_status` - (Computed) (String) Displays the status of the last transfer that was attempted. Valid values are `IN_PROGRESS`, `FAILED`, `SUCCESSFUL`
* `last_refresh_status_message` - (Computed) (String) Displays the last transfer’s status message. This is currently shown as failure reason.

## Import

Zones can be imported using their name(must be FQDN), e.g.,

```
$ terraform import ultradns_zone.example "example.com." 
```


