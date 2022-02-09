# Zone Resources

## Primary Zone
resource "ultradns_zone" "primary" {
    name        = "${var.ultradns_primary_zone_name}"
    account_name = "${var.ultradns_account}"
    type        = "PRIMARY"
    change_comment = "New Primary zone"
    primary_create_info {
        create_type = "NEW"
        force_import = false
        restrict_ip{
            single_ip = "192.168.1.11"
        }
        restrict_ip{
            single_ip = "192.168.1.5"
        }
        notify_addresses {
            notify_address = "192.168.1.15"
        }
        notify_addresses {
            notify_address = "192.168.1.2"
        }
    }
}

## Secondary Zone
resource "ultradns_zone" "secondary" {
    name        = "${var.ultradns_secondary_zone_name}"
    account_name = "${var.ultradns_account}"
    type        = "SECONDARY"
    secondary_create_info {
        primary_name_server_1 {
            ip = "${var.ultradns_primary_name_server}"
        } 
        notification_email_address = "${var.ultradns_notification_email_address}"
    }
}

## Alias Zone
resource "ultradns_zone" "alias" {
    name        = "${var.ultradns_alias_zone_name}"
    account_name = "${var.ultradns_account}"
    type        = "ALIAS"
    change_comment = "New alias zone"
    alias_create_info {
        original_zone_name = "${resource.ultradns_zone.primary.id}"
  }
}

# Zone Datasource
data "ultradns_zone" "zone" {
    name = "${resource.ultradns_zone.primary.id}"
}