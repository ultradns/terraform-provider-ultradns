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

data "ultradns_record" "all" {
    depends_on = [resource.ultradns_record.ptr]
    zone_name = "${resource.ultradns_zone.primary.id}"
    # owner_name = "www"
    # record_type = "1"
}

data "ultradns_zone" "all" {
    # query = "name:${var.ultradns_primary_zone_name}"
    # cursor = "fjpMQVNU"
    limit = 1
}
