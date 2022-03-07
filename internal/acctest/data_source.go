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

func TestAccDataSourceProbe(dataSourceType, datasourceName, zoneName, ownerName, resource string) string {
	return fmt.Sprintf(`
	%[1]s
	data "%[2]s" "data_%[3]s" {
		zone_name = "%[4]s"
		owner_name = "${resource.%[2]s.%[3]s.owner_name}"
	}
	`, resource, dataSourceType, datasourceName, zoneName)
}

func TestAccDataSourceProbeWithOptions(dataSourceType, datasourceName, zoneName, ownerName, interval, poolRecord, resource string) string {
	return fmt.Sprintf(`
	%[1]s
	data "%[2]s" "data_%[3]s" {
		zone_name = "%[4]s"
		owner_name = "${resource.%[2]s.%[3]s.owner_name}"
		interval = "%[5]s"
		pool_record = "%[6]s"
	}
	`, resource, dataSourceType, datasourceName, zoneName, interval, poolRecord)
}
