package acctest

import "fmt"

func TestAccResourceZonePrimary(resourceName, zoneName string) string {
	return fmt.Sprintf(`
	resource "ultradns_zone" "%s" {
		name        = "%s"
		account_name = "%s"
		type        = "PRIMARY"
		primary_create_info {
			create_type = "NEW"
			notify_addresses {
				notify_address = "192.168.1.1"
			}
			notify_addresses {
				notify_address = "192.168.1.2"
			}
			notify_addresses {
				notify_address = "192.168.1.3"
			}
			restrict_ip {
				single_ip = "192.168.1.1"
			}
			restrict_ip {
				single_ip = "192.168.1.2"
			}
			restrict_ip {
				single_ip = "192.168.1.3"
			}
		}
	}
	`, resourceName, zoneName, TestAccount)
}

func TestAccResourceSBPool(zoneName, ownerName string) string {
	return fmt.Sprintf(`
	%s
	resource "ultradns_sbpool" "a" {
		zone_name = "${resource.ultradns_zone.primary_sbpool.id}"
		owner_name = "%s.${resource.ultradns_zone.primary_sbpool.id}"
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
			failover_delay = 1
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
	}
	`, TestAccResourceZonePrimary("primary_sbpool", zoneName), ownerName)
}

func TestAccResourceTCPool(zoneName, ownerName string) string {
	return fmt.Sprintf(`
	%s
	resource "ultradns_tcpool" "a" {
		zone_name = "${resource.ultradns_zone.primary_tcpool.id}"
		owner_name = "%s.${resource.ultradns_zone.primary_tcpool.id}"
		record_type = "A"
		ttl = 120
    	run_probes = true
    	act_on_probes = true
    	failure_threshold = 2
    	max_to_lb = 2
    	rdata_info{
			priority = 2
			threshold = 1
			rdata = "192.168.1.1"
			failover_delay = 1
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
	}
	`, TestAccResourceZonePrimary("primary_tcpool", zoneName), ownerName)
}
