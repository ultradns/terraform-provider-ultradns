# Terraform configuration for UltraDNS provider
terraform {
    required_providers {
        ultradns = {
            source = "ultradns.com/ultradns/ultradns"
        }
    }
}

# UltraDNS provider configuration
provider "ultradns" {
   username = "${var.ultradns_username}"
   password = "${var.ultradns_password}"
   hosturl = "${var.ultradns_host_url}"
}

# Zone Resources

## Primary Zone
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

## Secondary Zone
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

## Alias Zone
resource "ultradns_zone" "alias" {
    name        = "${var.ultradns_alias_zone_name}"
    account_name = "${var.ultradns_username}"
    type        = "ALIAS"
    alias_create_info {
        original_zone_name = "${resource.ultradns_zone.primary.id}"
  }
}

# Record Resources

## Record Resource of Type A (1)
resource "ultradns_record" "a" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "a"
    record_type = "1"
    ttl = 120
    record_data = ["192.168.1.1"]
}

## Record Resource of Type CNAME (5)
resource "ultradns_record" "cname" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "cname"
    record_type = "CNAME"
    ttl = 120
    record_data = ["example.com."]
}

## Record Resource of Type PTR (12)
resource "ultradns_record" "ptr" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "192.168.1.1"
    record_type = "12"
    ttl = 120
    record_data = ["example.com."]
}

## Record Resource of Type HINFO (13)
resource "ultradns_record" "hinfo" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "hinfo"
    record_type = "HINFO"
    ttl = 120
    record_data = ["\"PC\" \"Linux\"","\"Laptop\" \"Windows\""]
}

## Record Resource of Type MX (15)
resource "ultradns_record" "mx" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "mx"
    record_type = "15"
    ttl = 120
    record_data = ["2 example.com."]
}

## Record Resource of Type TXT (16)
resource "ultradns_record" "txt" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "txt"
    record_type = "TXT"
    ttl = 120
    record_data = ["example.com."]
}

## Record Resource of Type RP (17)
resource "ultradns_record" "rp" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "${resource.ultradns_zone.primary.id}"
    record_type = "17"
    ttl = 120
    record_data = ["test.example.com. example.128/134.123.178.178.in-addr.arpa."]
}

## Record Resource of Type AAAA (28)
resource "ultradns_record" "aaaa" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "aaaa"
    record_type = "AAAA"
    ttl = 120
    record_data = ["2001:db8:85a3:0:0:8a2e:370:7334"]
}

## Record Resource of Type SRV (33)
resource "ultradns_record" "srv" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "srv"
    record_type = "33"
    ttl = 120
    record_data = ["5 6 7 example.com."]
}

## Record Resource of Type NAPTR (35)
resource "ultradns_record" "naptr" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "naptr"
    record_type = "NAPTR"
    ttl = 120
    record_data = ["1 2 \"3\" \"test\" \"\" test.com."]
}

## Record Resource of Type SSHFP (44)
resource "ultradns_record" "sshfp" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "sshfp"
    record_type = "44"
    ttl = 120
    record_data = ["1 2 54B5E539EAF593AEA410F80737530B71CCDE8B6C3D241184A1372E98BC7EDB37"]
}

## Record Resource of Type TLSA (52)
resource "ultradns_record" "tlsa" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "_23._tcp.tlsatest"
    record_type = "TLSA"
    ttl = 120
    record_data = ["0 0 0 aaaaaaaa"]
}

## Record Resource of Type SPF (99)
resource "ultradns_record" "spf" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "${resource.ultradns_zone.primary.id}"
    record_type = "99"
    ttl = 120
    record_data = ["v=spf1 ip4:1.2.3.4 ~all"]
}

## Record Resource of Type CAA (257)
resource "ultradns_record" "caa" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "caa"
    record_type = "CAA"
    ttl = 120
    record_data = ["1 issue \"asdfsadf\""]
}

## Record Resource of Type APEXALIAS (65282)
resource "ultradns_record" "apex" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "${resource.ultradns_zone.primary.id}"
    record_type = "65282"
    ttl = 120
    record_data = ["example.com."]
}

# RD Pool Resources

## RD Pool Resource of Type A (1)
resource "ultradns_rdpool" "rdpoola" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "rdpoola"
    record_type = "A"
    order = "ROUND_ROBIN"
    ttl = 120
    record_data = ["192.1.1.1","192.168.1.2"]
}

## RD Pool Resource of Type AAAA (28)
resource "ultradns_rdpool" "rdpoolaaaa" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "rdpoolaaaa"
    record_type = "AAAA"
    order = "RANDOM"
    description = "description 123"
    ttl = 120
    record_data = ["2001:db8:85a3:0:0:8a2e:370:7334"]
}

# Zone Datasource
data "ultradns_zone" "zone" {
    name = "0-0-0-0-0antony.com."
}

# Record Datasource
data "ultradns_record" "record" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "${resource.ultradns_record.a.owner_name}"
    record_type = "${resource.ultradns_record.a.record_type}"
}

# RD Pool Datasource
data "ultradns_rdpool" "rdpool" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "${resource.ultradns_rdpool.rdpoola.owner_name}"
    record_type = "${resource.ultradns_rdpool.rdpoola.record_type}"
}