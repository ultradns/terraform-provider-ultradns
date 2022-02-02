package pool

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/ultradns-go-sdk/pkg/pool"
)

func GetMonitor(monitorData map[string]interface{}) *pool.Monitor {
	monitor := &pool.Monitor{}

	if val, ok := monitorData["url"]; ok {
		monitor.URL = val.(string)
	}

	if val, ok := monitorData["method"]; ok {
		monitor.Method = val.(string)
	}

	if val, ok := monitorData["transmitted_data"]; ok {
		monitor.TransmittedData = val.(string)
	}

	if val, ok := monitorData["search_string"]; ok {
		monitor.SearchString = val.(string)
	}

	return monitor
}

func GetMonitorSet(monitorData *pool.Monitor) *schema.Set {
	set := &schema.Set{F: schema.HashResource(MonitorResource())}
	monitor := make(map[string]interface{})
	monitor["url"] = monitorData.URL
	monitor["method"] = monitorData.Method
	monitor["transmitted_data"] = monitorData.TransmittedData
	monitor["search_string"] = monitorData.SearchString
	set.Add(monitor)

	return set
}
