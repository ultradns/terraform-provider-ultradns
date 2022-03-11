# TC Pool Resources

## TC Pool Resource of Type A (1)
resource "ultradns_tcpool" "tcpoola" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "tcpoola"
    record_type = "A"
    ttl = 120
    pool_description = "TC Pool Resource of Type A"
    run_probes = true
    act_on_probes = true
    failure_threshold = 2
    max_to_lb = 1
    rdata_info{
        priority = 2
        threshold = 1
        rdata = "192.168.1.1"
        failover_delay = 2
        run_probes = true
        state = "ACTIVE"
        weight = 4
    }
    rdata_info{
        priority = 1
        threshold = 1
        rdata = "192.168.1.2"
        failover_delay = 1
        run_probes = false
        state = "NORMAL"
    }
    backup_record{
        rdata = "192.168.1.3"
        failover_delay = 1
    }
}

## TC Pool Datasource
data "ultradns_tcpool" "tcpool" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "${resource.ultradns_tcpool.tcpoola.owner_name}"
    record_type = "${resource.ultradns_tcpool.tcpoola.record_type}"
}