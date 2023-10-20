package probedns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/helper"
	"github.com/ultradns/terraform-provider-ultradns/internal/probe"
	"github.com/ultradns/terraform-provider-ultradns/internal/rrset"
	"github.com/ultradns/terraform-provider-ultradns/internal/service"
	sdkprobe "github.com/ultradns/ultradns-go-sdk/pkg/probe"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe/dns"
	sdkprobehelper "github.com/ultradns/ultradns-go-sdk/pkg/probe/helper"
)

func ResourceProbeDNS() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceProbeDNSCreate,
		ReadContext:   resourceProbeDNSRead,
		UpdateContext: resourceProbeDNSUpdate,
		DeleteContext: resourceProbeDNSDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resourceProbeDNSSchema(),
	}
}

func resourceProbeDNSCreate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	services := meta.(*service.Service)
	probeData := getNewProbeDNS(rd)
	rrSetKeyData := rrset.NewRRSetKey(rd)

	rrSetKeyData.RecordType = probe.RecordTypeA

	res, err := services.ProbeService.Create(rrSetKeyData, probeData)
	if err != nil {
		return diag.FromErr(err)
	}

	uri := res.Header.Get("Location")

	rrSetKeyData.ID = helper.GetProbeIDFromURI(uri)

	rd.SetId(rrSetKeyData.PID())

	return resourceProbeDNSRead(ctx, rd, meta)
}

func resourceProbeDNSRead(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	services := meta.(*service.Service)
	rrSetKey := probe.GetRRSetKeyFromID(rd.Id())
	rrSetKey.PType = sdkprobe.DNS
	_, probeData, err := services.ProbeService.Read(rrSetKey)
	if err != nil {
		rd.SetId("")
		tflog.Error(ctx, err.Error())
		return nil
	}

	if err = probe.FlattenRRSetKey(rrSetKey, rd); err != nil {
		return diag.FromErr(err)
	}

	if err = flattenProbeDNS(probeData, rd); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceProbeDNSUpdate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	services := meta.(*service.Service)
	probeData := getNewProbeDNS(rd)
	rrSetKeyData := probe.GetRRSetKeyFromID(rd.Id())

	_, err := services.ProbeService.Update(rrSetKeyData, probeData)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceProbeDNSRead(ctx, rd, meta)
}

func resourceProbeDNSDelete(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	services := meta.(*service.Service)
	rrSetKeyData := probe.GetRRSetKeyFromID(rd.Id())

	_, err := services.ProbeService.Delete(rrSetKeyData)
	if err != nil {
		rd.SetId("")

		return diag.FromErr(err)
	}

	rd.SetId("")

	return diags
}

func getNewProbeDNS(rd *schema.ResourceData) *sdkprobe.Probe {
	probeData := probe.NewProbe(rd, sdkprobe.DNS)
	details := &dns.Details{}
	limits := &sdkprobehelper.LimitsInfo{}

	details.Limits = limits
	probeData.Details = details

	if val, ok := rd.GetOk("port"); ok {
		details.Port = val.(int)
	}

	if val, ok := rd.GetOk("tcp_only"); ok {
		details.TCPOnly = val.(bool)
	}

	if val, ok := rd.GetOk("type"); ok {
		details.Type = val.(string)
	}

	if val, ok := rd.GetOk("query_name"); ok {
		details.OwnerName = val.(string)
	}

	if val, ok := rd.GetOk("response"); ok && len(val.([]interface{})) > 0 {
		limitData := val.([]interface{})[0].(map[string]interface{})
		limits.Response = probe.GetSearchString(limitData)
	}

	if val, ok := rd.GetOk("run_limit"); ok && len(val.([]interface{})) > 0 {
		limitData := val.([]interface{})[0].(map[string]interface{})
		limits.Run = probe.GetLimit(limitData)
	}

	if val, ok := rd.GetOk("avg_run_limit"); ok && len(val.([]interface{})) > 0 {
		limitData := val.([]interface{})[0].(map[string]interface{})
		limits.AvgRun = probe.GetLimit(limitData)
	}

	return probeData
}
