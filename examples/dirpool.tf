# DIR Pool Resources

## DIR Pool Resource of Type A (1)
resource "ultradns_dirpool" "a" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "dirpool_a"
    record_type = "A"
    ignore_ecs = true
    conflict_resolve = "IP"
    rdata_info{
        rdata = "192.168.1.5"
        all_non_configured = true
        ttl = 800
    }
    rdata_info{
        rdata = "192.168.1.2"
        geo_group_name = "geo_group"
        geo_codes = ["NAM","EUR"]
        ip_group_name = "ip_group"
        ip{
            address = "200.1.1.1"
        }
        ip{
            start = "200.1.1.2"
            end = "200.1.1.5"
        }
        ip{
            cidr = "200.20.20.0/24"
        }
    }
    no_response{
        geo_group_name = "geo_response_group"
        geo_codes = ["AG"]
        ip_group_name = "ip_response_group"
        ip{
            address = "2.2.2.2"
        }
    }
}

## DIR Pool Resource of Type PTR (12)
resource "ultradns_dirpool" "ptr" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "dirpool_12"
    record_type = "PTR"
    rdata_info{
        rdata = "ns1.example.com."
        geo_group_name = "geo_group"
        geo_codes = ["NAM","EUR"]
    }
    no_response{
        all_non_configured = true
    }
}

## DIR Pool Resource of Type MX (15)
resource "ultradns_dirpool" "mx" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "dirpool_mx"
    record_type = "MX"
    rdata_info{
        rdata = "2 example.com."
        geo_group_name = "geo_group"
        geo_codes = ["NAM","EUR"]
    }
    no_response{
        all_non_configured = true
    }
}

## DIR Pool Resource of Type TXT (16)
resource "ultradns_dirpool" "txt" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "dirpool_txt.${resource.ultradns_zone.primary.id}"
    record_type = "TXT"
    rdata_info{
        rdata = "text data"
        geo_group_name = "geo_group"
        geo_codes = ["NAM","EUR"]
    }
    no_response{
        all_non_configured = true
    }
}

## DIR Pool Resource of Type AAAA (28)
resource "ultradns_dirpool" "aaaa" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "dirpool_aaaa"
    record_type = "AAAA"
    pool_description = "DIR Pool Resource of Type AAAA"
    ignore_ecs = true
    conflict_resolve = "IP"
    rdata_info{
        rdata = "aaaa:bbbb:cccc:dddd:eeee:ffff:1111:3333"
        geo_group_name = "geo_group"
        geo_codes = ["EUR"]
        ip_group_name = "ip_group"
        ip{
            start = "aaaa:bbbb:cccc:dddd:eeee:ffff:1111:4444"
            end = "aaaa:bbbb:cccc:dddd:eeee:ffff:1111:6666"
        }
    }
    no_response{
        geo_group_name = "geo_response_group"
        geo_codes = ["AI"]
        ip_group_name = "ip_response_group"
        ip{
            address = "aaaa:bbbb:cccc:dddd:eeee:ffff:3333:5555"
        }
    }
}

## DIR Pool Resource of Type SRV (33)
resource "ultradns_dirpool" "srv" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "dirpool_srv"
    record_type = "SRV"
    rdata_info{
        rdata = "5 6 7 example.com."
        geo_group_name = "geo_group"
        geo_codes = ["NAM","EUR"]
    }
    no_response{
        all_non_configured = true
    }
}


# Record Datasource
data "ultradns_dirpool" "dirpool" {
    zone_name = "${resource.ultradns_zone.primary.id}"
    owner_name = "${resource.ultradns_dirpool.a.owner_name}"
    record_type = "${resource.ultradns_dirpool.a.record_type}"
}