# SB Pool Resources

## SB Pool Resource of Type A (1)
resource "ultradns_sbpool" "sbpoola" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "sbpoola"
    record_type = "A"
    ttl = 120
    pool_description = "SB Pool Resource of Type A"
    run_probes = true
    act_on_probes = true
    order = "ROUND_ROBIN"
    failure_threshold = 2
    max_active = 1
    max_served = 1
    rdata_info{
        priority = 2
        threshold = 1
        rdata = "192.168.1.1"
        failover_delay = 2
        run_probes = true
        state = "ACTIVE"
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
    backup_record{
        rdata = "192.168.1.4"
        failover_delay = 1
    }
}

## SB Pool Resource of Type AAAA (28)
resource "ultradns_sbpool" "sbpoolaaaa" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "sbpoolaaaa"
    record_type = "AAAA"
    ttl = 120
    pool_description = "SB Pool Resource of Type AAAA"
    run_probes = true
    act_on_probes = true
    order = "ROUND_ROBIN"
    failure_threshold = 2
    max_active = 1
    max_served = 1
    rdata_info{
        priority = 2
        threshold = 1
        rdata = "2001:db8:85a3:0:0:8a2e:370:7335"
        failover_delay = 2
        run_probes = true
        state = "ACTIVE"
    }
    rdata_info{
        priority = 1
        threshold = 1
        rdata = "2001:db8:85a3:0:0:8a2e:370:7336"
        failover_delay = 1
        run_probes = false
        state = "NORMAL"
    }
    backup_record{
        rdata = "2001:db8:85a3:0:0:8a2e:370:7337"
        failover_delay = 1
    }
    backup_record{
        rdata = "2001:db8:85a3:0:0:8a2e:370:7338"
        failover_delay = 1
    }
}

## SB Pool Datasource
data "ultradns_sbpool" "sbpool" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "${resource.ultradns_sbpool.sbpoola.owner_name}"
    record_type = "${resource.ultradns_sbpool.sbpoola.record_type}"
}