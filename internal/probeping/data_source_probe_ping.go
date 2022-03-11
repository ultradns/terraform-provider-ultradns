package probeping

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/errors"
	"github.com/ultradns/terraform-provider-ultradns/internal/probe"
	"github.com/ultradns/terraform-provider-ultradns/internal/rrset"
	"github.com/ultradns/terraform-provider-ultradns/internal/service"
	sdkprobe "github.com/ultradns/ultradns-go-sdk/pkg/probe"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe/ping"
	sdkrrset "github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

func DataSourceprobePING() *schema.Resource {
	return &schema.Resource{

		ReadContext: dataSourceprobePINGRead,

		Schema: dataSourceprobePINGSchema(),
	}
}

func dataSourceprobePINGRead(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	rrSetKey := rrset.NewRRSetKey(rd)
	rrSetKey.PType = sdkprobe.PING
	rrSetKey.RecordType = probe.RecordTypeA

	if val, ok := rd.GetOk("guid"); ok {
		rrSetKey.ID = val.(string)

		return readProbePING(rrSetKey, rd, meta)
	}

	return listProbePING(rrSetKey, rd, meta)
}

func readProbePING(rrSetKey *sdkrrset.RRSetKey, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	services := meta.(*service.Service)

	_, probeData, err := services.ProbeService.Read(rrSetKey)

	if err != nil {
		return diag.FromErr(err)
	}

	return flattenDataSourceProbePING(rrSetKey, probeData, rd)
}

func listProbePING(rrSetKey *sdkrrset.RRSetKey, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	services := meta.(*service.Service)

	query := &sdkprobe.Query{}

	if val, ok := rd.GetOk("pool_record"); ok {
		query.PoolRecord = val.(string)
	}

	_, probeDataList, err := services.ProbeService.List(rrSetKey, query)

	if err != nil {
		return diag.FromErr(err)
	}

	return setMatchedProbePING(rrSetKey, probeDataList.Probes, rd)
}

func flattenDataSourceProbePING(rrSetKey *sdkrrset.RRSetKey, probeData *sdkprobe.Probe, rd *schema.ResourceData) diag.Diagnostics {
	var diags diag.Diagnostics

	rd.SetId(rrSetKey.PID())

	if err := probe.FlattenRRSetKey(rrSetKey, rd); err != nil {
		return diag.FromErr(err)
	}

	if err := flattenProbePING(probeData, rd); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func setMatchedProbePING(rrSetKey *sdkrrset.RRSetKey, probeDataList []*sdkprobe.Probe, rd *schema.ResourceData) diag.Diagnostics {
	var probeData *sdkprobe.Probe

	for _, probeResData := range probeDataList {
		if ok := validateProbeFilterOptions(probeResData, rd); !ok {
			continue
		}

		probeData = probeResData
	}

	if probeData != nil {
		rrSetKey.ID = probeData.ID

		return setProbePINGDetails(rrSetKey, probeData, rd)
	}

	return diag.FromErr(errors.ProbeResourceNotFound(sdkprobe.PING))
}

func validateProbeFilterOptions(probeData *sdkprobe.Probe, rd *schema.ResourceData) bool {
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

	if probeData.Type != sdkprobe.PING {
		return false
	}

	if threshold != 0 && probeData.Threshold != threshold {
		return false
	}

	if interval != "" && probeData.Interval != interval {
		return false
	}

	agentSet := probe.GetAgentSet(probeData.Agents)

	if agents != nil && !agentSet.Equal(agents) {
		return false
	}

	return true
}

func setProbePINGDetails(rrSetKey *sdkrrset.RRSetKey, probeData *sdkprobe.Probe, rd *schema.ResourceData) diag.Diagnostics {
	details := &ping.Details{}
	jsonStr, err := json.Marshal(probeData.Details.(map[string]interface{}))

	if err != nil {
		diag.FromErr(err)
	}

	if err := json.Unmarshal(jsonStr, details); err != nil {
		diag.FromErr(err)
	}

	probeData.Details = details

	return flattenDataSourceProbePING(rrSetKey, probeData, rd)
}
