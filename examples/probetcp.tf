# TCP Probe Resources

## TCP Probe Resource of SB Pool
resource "ultradns_probe_tcp" "tcp_sb" {
	zone_name = "${resource.ultradns_sbpool.sbpoola.zone_name}"
	owner_name = "sbpoola"
	interval = "HALF_MINUTE"
	agents = ["EUROPE_WEST", "SOUTH_AMERICA", "PALO_ALTO", "NEW_YORK"]
	threshold = 4
	port = 443
	control_ip = "192.168.1.1"
	connect_limit{
		fail = 11
	}
}

## TCP Probe Resource of TC Pool
resource "ultradns_probe_tcp" "tcp_tc" {
	zone_name = "${resource.ultradns_tcpool.tcpoolaaaa.zone_name}"
	owner_name = "tcpoolaaaa"
	pool_type = "AAAA"
	interval = "HALF_MINUTE"
	agents = ["EUROPE_WEST", "SOUTH_AMERICA", "PALO_ALTO", "NEW_YORK"]
	threshold = 4
	port = 443
	connect_limit{
		warning = 6
		critical = 9
		fail = 11
	}
	avg_connect_limit{
		warning = 5
		critical = 8
		fail = 10
	}
}

## PING Probe Datasource
data "ultradns_probe_tcp" "probetcp" {
    zone_name = "${resource.ultradns_probe_tcp.tcp_tc.zone_name}"
    owner_name = "${resource.ultradns_probe_tcp.tcp_tc.owner_name}"
	interval = "${resource.ultradns_probe_tcp.tcp_tc.interval}"
	threshold = "${resource.ultradns_probe_tcp.tcp_tc.threshold}"
	agents = "${resource.ultradns_probe_tcp.tcp_tc.agents}"
	pool_type = "AAAA"
}