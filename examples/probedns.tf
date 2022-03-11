# DNS Probe Resources

## DNS Probe Resource of SB Pool
resource "ultradns_probe_dns" "dns_sb" {
	zone_name = "${resource.ultradns_sbpool.sbpoola.zone_name}"
	owner_name = "sbpoola"
	interval = "HALF_MINUTE"
	agents = ["EUROPE_WEST", "SOUTH_AMERICA", "PALO_ALTO", "NEW_YORK"]
	threshold = 4
	port = 43
	tcp_only = false
	type = "SOA"
	query_name = "${resource.ultradns_sbpool.sbpoola.zone_name}"
	response{
		fail = "fail"
	}
	run_limit{
		fail = 12
	}
}

## DNS Probe Resource of TC Pool
resource "ultradns_probe_dns" "dns_tc" {
	zone_name = "${resource.ultradns_tcpool.tcpoola.zone_name}"
	owner_name = "tcpoola"
	interval = "HALF_MINUTE"
	agents = ["EUROPE_WEST", "SOUTH_AMERICA", "PALO_ALTO", "NEW_YORK"]
	threshold = 4
	port = 43
	tcp_only = false
	type = "SOA"
	query_name = "${resource.ultradns_tcpool.tcpoola.zone_name}"
	response{
		warning = "warning" 
		critical = "critical"
		fail = "fail"
	}
	run_limit{
		warning = 7 
		critical = 10
		fail = 12
	}
	avg_run_limit{
		warning = 7 
		critical = 10
		fail = 12
	}
}

## DNS Probe Datasource
data "ultradns_probe_dns" "probedns" {
    zone_name = "${resource.ultradns_probe_dns.dns_tc.zone_name}"
    owner_name = "${resource.ultradns_probe_dns.dns_tc.owner_name}"
	interval = "${resource.ultradns_probe_dns.dns_tc.interval}"
	threshold = "${resource.ultradns_probe_dns.dns_tc.threshold}"
	agents = "${resource.ultradns_probe_dns.dns_tc.agents}"
}
