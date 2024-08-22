package probedns

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
	"github.com/ultradns/ultradns-go-sdk/pkg/probe/dns"
	sdkrrset "github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

func DataSourceprobeDNS() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceprobeDNSRead,

		Schema: dataSourceprobeDNSSchema(),
	}
}

func dataSourceprobeDNSRead(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	rrSetKey := rrset.NewRRSetKey(rd)
	rrSetKey.PType = sdkprobe.DNS

	if val, ok := rd.GetOk("pool_type"); ok {
		rrSetKey.RecordType = val.(string)
	}

	if val, ok := rd.GetOk("guid"); ok {
		rrSetKey.ID = val.(string)

		return readProbeDNS(rrSetKey, rd, meta)
	}

	return listProbeDNS(rrSetKey, rd, meta)
}

func readProbeDNS(rrSetKey *sdkrrset.RRSetKey, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	services := meta.(*service.Service)

	_, probeData, err := services.ProbeService.Read(rrSetKey)
	if err != nil {
		return diag.FromErr(err)
	}

	return flattenDataSourceProbeDNS(rrSetKey, probeData, rd)
}

func listProbeDNS(rrSetKey *sdkrrset.RRSetKey, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	services := meta.(*service.Service)

	query := &sdkprobe.Query{}

	if val, ok := rd.GetOk("pool_record"); ok {
		query.PoolRecord = val.(string)
	}

	_, probeDataList, err := services.ProbeService.List(rrSetKey, query)
	if err != nil {
		return diag.FromErr(err)
	}

	return setMatchedProbeDNS(rrSetKey, probeDataList.Probes, rd)
}

func flattenDataSourceProbeDNS(rrSetKey *sdkrrset.RRSetKey, probeData *sdkprobe.Probe, rd *schema.ResourceData) diag.Diagnostics {
	var diags diag.Diagnostics

	rd.SetId(rrSetKey.PID())

	if err := probe.FlattenRRSetKey(rrSetKey, rd); err != nil {
		return diag.FromErr(err)
	}

	if err := flattenProbeDNS(probeData, rd); err != nil {
		return diag.FromErr(err)
	}

	if err := rd.Set("pool_type", rrSetKey.RecordType); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func setMatchedProbeDNS(rrSetKey *sdkrrset.RRSetKey, probeDataList []*sdkprobe.Probe, rd *schema.ResourceData) diag.Diagnostics {
	var probeData *sdkprobe.Probe

	for _, probeResData := range probeDataList {
		if ok := probe.ValidateProbeFilterOptions(sdkprobe.DNS, probeResData, rd); !ok {
			continue
		}

		probeData = probeResData
	}

	if probeData != nil {
		rrSetKey.ID = probeData.ID

		return setProbeDNSDetails(rrSetKey, probeData, rd)
	}

	return diag.FromErr(errors.ProbeResourceNotFound(sdkprobe.DNS))
}

func setProbeDNSDetails(rrSetKey *sdkrrset.RRSetKey, probeData *sdkprobe.Probe, rd *schema.ResourceData) diag.Diagnostics {
	details := &dns.Details{}
	jsonStr, err := json.Marshal(probeData.Details.(map[string]interface{}))
	if err != nil {
		diag.FromErr(err)
	}

	if err := json.Unmarshal(jsonStr, details); err != nil {
		diag.FromErr(err)
	}

	probeData.Details = details

	return flattenDataSourceProbeDNS(rrSetKey, probeData, rd)
}
