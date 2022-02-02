# SLB Pool Resources

## SLB Pool Resource of Type A (1)
resource "ultradns_slbpool" "slbpoola" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "slbpoola"
    record_type = "A"
    ttl = 120
    rdata_info{
        description = "first"
        rdata = "192.201.127.33"
        probing_enabled = false
    }
    rdata_info{
        description       = "second"
        rdata = "192.168.1.2"
        probing_enabled = true
    }
    region_failure_sensitivity = "HIGH"
    serving_preference = "AUTO_SELECT"
    response_method = "ROUND_ROBIN"
    monitor{
        url = "${var.ultradns_host_url}"
        method = "POST"
        search_string = "test"
    }
    all_fail_record{
        rdata = "192.127.127.33"
    }
}

## SLB Pool Resource of Type AAAA (28)
resource "ultradns_slbpool" "slbpoolaaaa" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "slbpoolaaaa"
    record_type = "AAAA"
    ttl = 120
    rdata_info{
        description = "first"
        rdata = "2001:db8:85a3:0:0:8a2e:370:7334"
        probing_enabled = false
    }
    region_failure_sensitivity = "LOW"
    serving_preference = "AUTO_SELECT"
    response_method = "ROUND_ROBIN"
    monitor{
        url = "${var.ultradns_host_url}"
        method = "GET"
    }
    all_fail_record{
        rdata = "2001:db8:85a3:0:0:8a2e:370:7324"
    }
}


# SLB Pool Datasource
data "ultradns_slbpool" "slbpool" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "${resource.ultradns_slbpool.slbpoola.owner_name}"
    record_type = "${resource.ultradns_slbpool.slbpoola.record_type}"
}