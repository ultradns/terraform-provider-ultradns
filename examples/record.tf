# Record Resources

## Record Resource of Type A (1)
resource "ultradns_record" "a" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "a"
    record_type = "1"
    ttl = 800
    record_data = ["192.168.1.1"]
}

## Record Resource of Type NS (2)
resource "ultradns_record" "ns" {
		zone_name = "${resource.ultradns_zone.primary.id}"
		owner_name = "${resource.ultradns_zone.primary.id}"
		record_type = "NS"
		ttl = 800
		record_data = ["ns11.example.com.","ns12.example.com."]
}

## Record Resource of Type CNAME (5)
resource "ultradns_record" "cname" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "cname.${resource.ultradns_zone.primary.id}"
    record_type = "CNAME"
    ttl = 800
    record_data = ["example.com."]
}

## Record Resource of Type PTR (12)
resource "ultradns_record" "ptr" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "1"
    record_type = "12"
    ttl = 800
    record_data = ["ns1.example.com."]
}

## Record Resource of Type MX (15)
resource "ultradns_record" "mx" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "mx"
    record_type = "15"
    ttl = 800
    record_data = ["2 example.com."]
}

## Record Resource of Type TXT (16)
resource "ultradns_record" "txt" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "txt.${resource.ultradns_zone.primary.id}"
    record_type = "TXT"
    ttl = 800
    record_data = ["example.com."]
}

## Record Resource of Type AAAA (28)
resource "ultradns_record" "aaaa" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "aaaa"
    record_type = "AAAA"
    ttl = 800
    record_data = ["2001:db8:85a3:0:0:8a2e:370:7334"]
}

## Record Resource of Type SRV (33)
resource "ultradns_record" "srv" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "srv"
    record_type = "33"
    ttl = 800
    record_data = ["5 6 7 example.com."]
}

## Record Resource of Type SSHFP (44)
resource "ultradns_record" "sshfp" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "sshfp.${resource.ultradns_zone.primary.id}"
    record_type = "44"
    ttl = 800
    record_data = ["1 2 54B5E539EAF593AEA410F80737530B71CCDE8B6C3D241184A1372E98BC7EDB37"]
}

## Record Resource of Type APEXALIAS (65282)
resource "ultradns_record" "apex" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "${resource.ultradns_zone.primary.id}"
    record_type = "65282"
    ttl = 800
    record_data = ["example.com."]
}


# Record Datasource
data "ultradns_record" "record" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "${resource.ultradns_record.a.owner_name}"
    record_type = "${resource.ultradns_record.a.record_type}"
}