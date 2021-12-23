terraform {
    required_providers {
        ultradns = {
            source = "ultradns.com/ultradns/ultradns"
        }
    }
}

provider "ultradns" {
   username = "${var.ultradns_username}"
   password = "${var.ultradns_password}"
   hosturl = "${var.ultradns_host_url}"
}

resource "ultradns_zone" "primary" {
    name        = "${var.ultradns_primary_zone_name}"
    account_name = "${var.ultradns_username}"
    type        = "PRIMARY"
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

resource "ultradns_zone" "secondary" {
    name        = "${var.ultradns_secondary_zone_name}"
    account_name = "${var.ultradns_username}"
    type        = "SECONDARY"
    secondary_create_info {
        primary_name_server_1 {
            ip = "${var.ultradns_primary_name_server}"
        } 
        notification_email_address = "${var.ultradns_notification_email_address}"
    }
}

resource "ultradns_zone" "alias" {
    name        = "${var.ultradns_alias_zone_name}"
    account_name = "${var.ultradns_username}"
    type        = "ALIAS"
    alias_create_info {
        original_zone_name = "${resource.ultradns_zone.primary.id}"
  }
}

resource "ultradns_record" "a" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "a"
    record_type = "A"
    ttl = 120
    record_data = ["192.168.1.1"]
}

resource "ultradns_record" "aaaa" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "aaaa"
    record_type = "AAAA"
    ttl = 120
    record_data = ["2001:db8:85a3:0:0:8a2e:370:7334"]
}

resource "ultradns_record" "cname" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "cname"
    record_type = "CNAME"
    ttl = 120
    record_data = ["google.com."]
}

resource "ultradns_record" "mx" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "mx"
    record_type = "MX"
    ttl = 120
    record_data = ["2 google.com."]
}

resource "ultradns_record" "srv" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "srv"
    record_type = "SRV"
    ttl = 120
    record_data = ["5 6 7 google.com."]
}

resource "ultradns_record" "txt" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "txt"
    record_type = "TXT"
    ttl = 120
    record_data = ["google.com."]
}

resource "ultradns_record" "ptr" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "192.168.1.1"
    record_type = "PTR"
    ttl = 120
    record_data = ["google.com."]
}

data "ultradns_zone" "zone" {
    name = "${resource.ultradns_zone.primary.id}"
}

data "ultradns_record" "record_a" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "${resource.ultradns_record.a.owner_name}"
    record_type = "${resource.ultradns_record.a.record_type}"
}
