# HTTP Probe Resources

## HTTP Probe Resource of TC Pool
resource "ultradns_probe_http" "http_tc" {
	zone_name = "${resource.ultradns_tcpool.rdpoola.zone_name}"
	owner_name = "${resource.ultradns_tcpool.rdpoola.owner_name}"
	interval = "HALF_MINUTE"
	agents = ["EUROPE_WEST", "SOUTH_AMERICA", "PALO_ALTO", "NEW_YORK"]
	threshold = 4
	total_limit{
		warning = "5"
		critical = "8"
		fail = "10"
	}
	transaction{
		method = "GET"
		protocol_version = "HTTP/1.0"
		url = "https://www.ultradns.com/"
		follow_redirects = false
		expected_response = "3XX"
		connect_limit{
			warning = "5"
			critical = "8"
			fail = "10"
		}
		run_limit{
			warning = "5"
			critical = "8"
			fail = "10"
		}
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
			fail = 11
		}
		run_limit{
			fail = 12
		}
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
			fail = 11
		}
		run_limit{
			fail = 12
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
