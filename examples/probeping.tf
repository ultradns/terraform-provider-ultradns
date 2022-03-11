# PING Probe Resources

## PING Probe Resource of SB Pool
resource "ultradns_probe_ping" "ping_sb" {
	zone_name = "${resource.ultradns_sbpool.sbpoola.zone_name}"
	owner_name = "sbpoola"
	interval = "HALF_MINUTE"
	agents = ["EUROPE_WEST", "SOUTH_AMERICA", "PALO_ALTO", "NEW_YORK"]
	threshold = 4
	packets = 6
	packet_size = 106
	loss_percent_limit{
		fail = 11
	}
	total_limit{
		fail = 11
	}
	run_limit{
		fail = 11
	}
}

## PING Probe Resource of TC Pool
resource "ultradns_probe_ping" "ping_tc" {
	zone_name = "${resource.ultradns_tcpool.tcpoola.zone_name}"
	owner_name = "tcpoola"
	interval = "HALF_MINUTE"
	agents = ["EUROPE_WEST", "SOUTH_AMERICA", "PALO_ALTO", "NEW_YORK"]
	threshold = 4
	packets = 6
	packet_size = 106
	loss_percent_limit{
		warning = 6 
		critical = 9
		fail = 11
	}
	total_limit{
		warning = 6 
		critical = 9
		fail = 11
	}
	average_limit{
		warning = 6 
		critical = 9
		fail = 11
	}
	run_limit{
		warning = 6 
		critical = 9
		fail = 11
	}
	avg_run_limit{
		warning = 6 
		critical = 9
		fail = 11
	}
}

## PING Probe Datasource
data "ultradns_probe_ping" "probeping" {
    zone_name = "${resource.ultradns_probe_ping.ping_tc.zone_name}"
    owner_name = "${resource.ultradns_probe_ping.ping_tc.owner_name}"
	interval = "${resource.ultradns_probe_ping.ping_tc.interval}"
	threshold = "${resource.ultradns_probe_ping.ping_tc.threshold}"
	agents = "${resource.ultradns_probe_ping.ping_tc.agents}"
}
