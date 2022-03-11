# HTTP Probe Resources

## HTTP Probe Resource of SB Pool
resource "ultradns_probe_http" "http_sb" {
	zone_name = "${resource.ultradns_sbpool.sbpoola.zone_name}"
	owner_name = "sbpoola"
	interval = "HALF_MINUTE"
	agents = ["EUROPE_WEST", "SOUTH_AMERICA", "PALO_ALTO", "NEW_YORK"]
	threshold = 4
	total_limit{
		fail = 20
	}
	transaction{
		method = "POST"
		protocol_version = "HTTP/1.0"
		url = "https://www.ultradns.com/"
		transmitted_data = "foo=bar"
		follow_redirects = true
		expected_response = "3XX"
		search_string {
			fail = "Failure"
		}
		connect_limit{ 
			fail = 10
		}
		run_limit{
			fail = 10
		}
	}
}

## HTTP Probe Resource of TC Pool
resource "ultradns_probe_http" "http_tc" {
	zone_name = "${resource.ultradns_tcpool.tcpoola.zone_name}"
	owner_name = "tcpoola"
	interval = "HALF_MINUTE"
	agents = ["EUROPE_WEST", "SOUTH_AMERICA", "PALO_ALTO", "NEW_YORK"]
	threshold = 4
	total_limit{
		warning = 10 
		critical = 15 
		fail = 20
	}
	transaction{
		method = "POST"
		protocol_version = "HTTP/1.0"
		url = "https://www.ultradns.com/"
		transmitted_data = "foo=bar"
		follow_redirects = true
		expected_response = "3XX"
		search_string {
			warning = "Warning"
			critical = "Critical"
			fail = "Failure"
		}
		connect_limit{
			warning = 5 
			critical = 8 
			fail = 10
		}
		avg_connect_limit{
			warning = 5 
			critical = 8 
			fail = 10
		}
		run_limit{
			warning = 5 
			critical = 8 
			fail = 10
		}
		avg_run_limit{
			warning = 5 
			critical = 8 
			fail = 10
		}
	}
}

## HTTP Probe Datasource
data "ultradns_probe_http" "probehttp" {
    zone_name = "${resource.ultradns_probe_http.http_tc.zone_name}"
    owner_name = "${resource.ultradns_probe_http.http_tc.owner_name}"
	interval = "${resource.ultradns_probe_http.http_tc.interval}"
	threshold = "${resource.ultradns_probe_http.http_tc.threshold}"
	agents = "${resource.ultradns_probe_http.http_tc.agents}"
}
