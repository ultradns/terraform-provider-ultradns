package pool

import (
	"github.com/ultradns/ultradns-go-sdk/pkg/record/pool"
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

func GetMonitorList(monitorData *pool.Monitor) []interface{} {
	var list []interface{}

	if monitorData != nil {
		list = make([]interface{}, 1)
		monitor := make(map[string]interface{})
		monitor["url"] = monitorData.URL
		monitor["method"] = monitorData.Method
		monitor["transmitted_data"] = monitorData.TransmittedData
		monitor["search_string"] = monitorData.SearchString
		list[0] = monitor
	}

	return list
}

func GetBackupRecord(backupRecordData map[string]interface{}) *pool.BackupRecord {
	backupRecord := &pool.BackupRecord{}

	if val, ok := backupRecordData["rdata"]; ok {
		backupRecord.RData = val.(string)
	}

	if val, ok := backupRecordData["failover_delay"]; ok {
		backupRecord.FailOverDelay = val.(int)
	}

	return backupRecord
}
