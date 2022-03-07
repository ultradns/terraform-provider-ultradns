package probe

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe"
	probehelper "github.com/ultradns/ultradns-go-sdk/pkg/probe/helper"
	"github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

func NewProbe(rd *schema.ResourceData, probeType string) *probe.Probe {
	probeData := &probe.Probe{}

	if val, ok := rd.GetOk("interval"); ok {
		probeData.Interval = val.(string)
	}

	if val, ok := rd.GetOk("pool_record"); ok {
		probeData.PoolRecord = val.(string)
	}

	if val, ok := rd.GetOk("threshold"); ok {
		probeData.Threshold = val.(int)
	}

	if val, ok := rd.GetOk("agents"); ok {
		agentsList := val.(*schema.Set).List()
		probeData.Agents = make([]string, len(agentsList))

		for i, agent := range agentsList {
			probeData.Agents[i] = agent.(string)
		}
	}

	probeData.Type = probeType

	return probeData
}

func FlattenProbe(probeData *probe.Probe, rd *schema.ResourceData) error {
	if err := rd.Set("interval", probeData.Interval); err != nil {
		return err
	}

	if err := rd.Set("pool_record", probeData.PoolRecord); err != nil {
		return err
	}

	if err := rd.Set("threshold", probeData.Threshold); err != nil {
		return err
	}

	if err := rd.Set("guid", probeData.ID); err != nil {
		return err
	}

	if err := rd.Set("agents", GetAgentSet(probeData.Agents)); err != nil {
		return err
	}

	return nil
}

func GetAgentSet(agentsData []string) *schema.Set {
	set := &schema.Set{F: schema.HashString}

	for _, data := range agentsData {
		set.Add(data)
	}

	return set
}

func FlattenRRSetKey(rrSetKeyData *rrset.RRSetKey, rd *schema.ResourceData) error {
	if err := rd.Set("zone_name", helper.GetZoneFQDN(rrSetKeyData.Zone)); err != nil {
		return err
	}

	if err := rd.Set("owner_name", helper.GetOwnerFQDN(rrSetKeyData.Owner, rrSetKeyData.Zone)); err != nil {
		return err
	}

	return nil
}

func GetRRSetKeyFromID(id string) *rrset.RRSetKey {
	rrSetKeyData := &rrset.RRSetKey{}
	splitStringData := strings.Split(id, ":")

	if len(splitStringData) == 4 {
		rrSetKeyData.Owner = splitStringData[0]
		rrSetKeyData.Zone = splitStringData[1]
		rrSetKeyData.RecordType = helper.GetRecordTypeString(splitStringData[2])
		rrSetKeyData.ID = splitStringData[3]
	}

	return rrSetKeyData
}

func GetLimit(limitData map[string]interface{}) *probehelper.Limit {
	limit := &probehelper.Limit{}

	if val, ok := limitData["warning"]; ok {
		limit.Warning = val.(int)
	}

	if val, ok := limitData["critical"]; ok {
		limit.Critical = val.(int)
	}

	if val, ok := limitData["fail"]; ok {
		limit.Fail = val.(int)
	}

	return limit
}

func GetSearchString(searchStringData map[string]interface{}) *probehelper.SearchString {
	searchString := &probehelper.SearchString{}

	if val, ok := searchStringData["warning"]; ok {
		searchString.Warning = val.(string)
	}

	if val, ok := searchStringData["critical"]; ok {
		searchString.Critical = val.(string)
	}

	if val, ok := searchStringData["fail"]; ok {
		searchString.Fail = val.(string)
	}

	return searchString
}

func GetLimitList(limitData *probehelper.Limit) []interface{} {
	var list []interface{}

	if limitData != nil {
		list = make([]interface{}, 1)
		limit := make(map[string]interface{})
		limit["warning"] = limitData.Warning
		limit["critical"] = limitData.Critical
		limit["fail"] = limitData.Fail
		list[0] = limit
	}

	return list
}

func GetSearchStringList(searchStringData *probehelper.SearchString) []interface{} {
	var list []interface{}

	if searchStringData != nil {
		list = make([]interface{}, 1)
		searchString := make(map[string]interface{})
		searchString["warning"] = searchStringData.Warning
		searchString["critical"] = searchStringData.Critical
		searchString["fail"] = searchStringData.Fail
		list[0] = searchString
	}

	return list
}
