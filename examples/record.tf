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

## Record Resource of Type SOA (6)
resource "ultradns_record" "soa" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "${resource.ultradns_zone.primary.id}"
    record_type = "SOA"
    ttl = 86400
    record_data = ["ns.example.com. admin@example.com. 7200 3600 1209600 36000"]
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

## Record Resource of Type DS (43)
resource "ultradns_record" "ds" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "${resource.ultradns_zone.primary.id}"
    record_type = "DS"
    ttl = 800
    record_data = ["25286 1 1 340437DC66C3DFAD0B3E849740D2CF1A4151671D"]
}

## Record Resource of Type SSHFP (44)
resource "ultradns_record" "sshfp" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "sshfp.${resource.ultradns_zone.primary.id}"
    record_type = "44"
    ttl = 800
    record_data = ["1 2 54B5E539EAF593AEA410F80737530B71CCDE8B6C3D241184A1372E98BC7EDB37"]
}

## Record Resource of Type SVCB (64)
resource "ultradns_record" "svcb" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "svcb"
    record_type = "SVCB"
    ttl = 800
    record_data = ["0 www.ultradns.com."]
}

## Record Resource of Type HTTPS (65)
resource "ultradns_record" "https" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "https"
    record_type = "HTTPS"
    ttl = 800
    record_data = ["1 www.ultradns.com. ech=dGVzdA== mandatory=alpn,key65444 no-default-alpn port=8080 ipv4hint=1.2.3.4,9.8.7.6 key65444=privateKeyTesting ipv6hint=2001:db8:3333:4444:5555:6666:7777:8888,2001:db8:3333:4444:cccc:dddd:eeee:ffff alpn=h3,h3-29,h2"]
}

## Record Resource of Type CAA (257)
resource "ultradns_record" "caa" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "${resource.ultradns_zone.primary.id}"
    record_type = "CAA"
    ttl = 800
    record_data = ["0 issue ultradns"]
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