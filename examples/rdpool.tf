# RD Pool Resources

## RD Pool Resource of Type A (1)
resource "ultradns_rdpool" "rdpoola" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "rdpoola"
    record_type = "1"
    order = "ROUND_ROBIN"
    ttl = 800
    record_data = ["192.1.1.1","192.168.1.2"]
}

## RD Pool Resource of Type AAAA (28)
resource "ultradns_rdpool" "rdpoolaaaa" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "rdpoolaaaa"
    record_type = "AAAA"
    order = "RANDOM"
    description = "Record AAAA RD pool"
    ttl = 800
    record_data = ["2001:db8:85a3:0:0:8a2e:370:7334"]
}

## RD Pool Datasource
data "ultradns_rdpool" "rdpool" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "${resource.ultradns_rdpool.rdpoola.owner_name}"
    record_type = "${resource.ultradns_rdpool.rdpoola.record_type}"
}