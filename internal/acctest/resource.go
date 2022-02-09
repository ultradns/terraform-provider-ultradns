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
