# SF Pool Resources

## SF Pool Resource of Type A (1)
resource "ultradns_sfpool" "sfpoola" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "sfpoola"
    record_type = "A"
    ttl = 120
    record_data = ["192.1.1.3"]
    region_failure_sensitivity = "HIGH"
    live_record_state = "FORCED_INACTIVE"
    live_record_description = "Maintainance"
    pool_description = "SF Pool Resource of Type A"
    monitor{
        url = "${var.ultradns_host_url}"
        method = "GET"
        search_string = "test"
    }
    backup_record{
        rdata = "192.1.1.4"
        description = "Backup record"
    }
}

## SF Pool Resource of Type AAAA (28)
resource "ultradns_sfpool" "sfpoolaaaa" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "sfpoolaaaa"
    record_type = "AAAA"
    ttl = 120
    record_data = ["2001:db8:85a3:0:0:8a2e:370:7334"]
    region_failure_sensitivity = "LOW"
    monitor{
        url = "${var.ultradns_host_url}"
        method = "POST"
    }
    backup_record{
        rdata = "2001:db8:85a3:0:0:8a2e:370:7324"
        description = "Backup record"
    }
}


# SF Pool Datasource
data "ultradns_sfpool" "sfpool" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "${resource.ultradns_sfpool.sfpoola.owner_name}"
    record_type = "${resource.ultradns_sfpool.sfpoola.record_type}"
}