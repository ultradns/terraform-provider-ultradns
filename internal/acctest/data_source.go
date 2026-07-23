package acctest

import (
	"fmt"
)

func TestAccDataSourceRRSet(dataSourceType, datasourceName, zoneName, ownerName, rrType, resource string) string {
	return fmt.Sprintf(`
	%[1]s
	data "%[2]s" "data_%[3]s" {
		zone_name = "%[4]s"
		owner_name = "${resource.%[2]s.%[3]s.owner_name}"
		record_type = "%[5]s"
	}
	`, resource, dataSourceType, datasourceName, zoneName, rrType)
}

func TestAccDataSourceProbe(dataSourceType, datasourceName, zoneName, ownerName, poolType, resource string) string {
	return fmt.Sprintf(`
	%[1]s
	data "%[2]s" "data_%[3]s" {
		zone_name = "%[4]s"
		owner_name = "${resource.%[2]s.%[3]s.owner_name}"
		pool_type = "%[5]s"
	}
	`, resource, dataSourceType, datasourceName, zoneName, poolType)
}

func TestAccDataSourceProbeWithOptions(dataSourceType, datasourceName, zoneName, ownerName, poolType, interval, poolRecord, resource string) string {
	return fmt.Sprintf(`
	%[1]s
	data "%[2]s" "data_%[3]s" {
		zone_name = "%[4]s"
		owner_name = "${resource.%[2]s.%[3]s.owner_name}"
		pool_type = "%[7]s"
		interval = "%[5]s"
		pool_record = "%[6]s"
	}
	`, resource, dataSourceType, datasourceName, zoneName, interval, poolRecord, poolType)
}

func TestAccDataSourceCDN(datasourceName, resourceName, resource string) string {
	return fmt.Sprintf(`
	%[1]s
	data "ultradns_cdn" "data_%[2]s" {
		account_name = "${resource.ultradns_cdn.%[3]s.account_name}"
		fqdn = "${resource.ultradns_cdn.%[3]s.fqdn}"
	}
	`, resource, datasourceName, resourceName)
}

func TestAccDataSourceCDNs(datasourceName, resourceName string, page, size int, resource string) string {
	return fmt.Sprintf(`
	%[1]s
	data "ultradns_cdn_list" "data_%[2]s" {
		account_name = "${resource.ultradns_cdn.%[3]s.account_name}"
		page = %[4]d
		size = %[5]d
	}
	`, resource, datasourceName, resourceName, page, size)
}
