package probe

import (
	"strings"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/helper"
	sdkhelper "github.com/ultradns/ultradns-go-sdk/pkg/helper"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe"
	sdkprobehelper "github.com/ultradns/ultradns-go-sdk/pkg/probe/helper"
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

	if err := rd.Set("agents", helper.GetSchemaSetFromList(probeData.Agents)); err != nil {
		return err
	}

	return nil
}

func FlattenRRSetKey(rrSetKeyData *rrset.RRSetKey, rd *schema.ResourceData) error {
	if err := rd.Set("zone_name", sdkhelper.GetZoneFQDN(rrSetKeyData.Zone)); err != nil {
		return err
	}

	if err := rd.Set("owner_name", sdkhelper.GetOwnerFQDN(rrSetKeyData.Owner, rrSetKeyData.Zone)); err != nil {
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
		rrSetKeyData.RecordType = sdkhelper.GetRecordTypeString(splitStringData[2])
		rrSetKeyData.ID = splitStringData[3]
	}

	return rrSetKeyData
}

func GetLimit(limitData map[string]interface{}) *sdkprobehelper.Limit {
	limit := &sdkprobehelper.Limit{}

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

func GetSearchString(searchStringData map[string]interface{}) *sdkprobehelper.SearchString {
	searchString := &sdkprobehelper.SearchString{}

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

func GetLimitList(limitData *sdkprobehelper.Limit) []interface{} {
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

func GetSearchStringList(searchStringData *sdkprobehelper.SearchString) []interface{} {
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

func ValidateProbeFilterOptions(probeType string, probeData *probe.Probe, rd *schema.ResourceData) bool {
	var agents *schema.Set

	threshold := 0
	interval := ""

	if val, ok := rd.GetOk("threshold"); ok {
		threshold = val.(int)
	}

	if val, ok := rd.GetOk("interval"); ok {
		interval = val.(string)
	}

	if val, ok := rd.GetOk("agents"); ok {
		agents = val.(*schema.Set)
	}

	if probeData.Type != probeType {
		return false
	}

	if threshold != 0 && probeData.Threshold != threshold {
		return false
	}

	if interval != "" && probeData.Interval != interval {
		return false
	}

	agentSet := helper.GetSchemaSetFromList(probeData.Agents)

	if agents != nil && !agentSet.Equal(agents) {
		return false
	}

	return true
}

func poolTypeValidation(i interface{}, p cty.Path) diag.Diagnostics {
	var diags diag.Diagnostics

	supportedRRType := map[string]bool{
		"A": true, "1": true,
		"AAAA": true, "28": true,
	}

	recordType := i.(string)
	_, ok := supportedRRType[recordType]

	if !ok {
		return diag.Errorf("invalid or unsupported record type")
	}

	return diags
}
