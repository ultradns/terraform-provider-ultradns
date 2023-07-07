package probetcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/helper"
	"github.com/ultradns/terraform-provider-ultradns/internal/probe"
	"github.com/ultradns/terraform-provider-ultradns/internal/rrset"
	"github.com/ultradns/terraform-provider-ultradns/internal/service"
	sdkprobe "github.com/ultradns/ultradns-go-sdk/pkg/probe"
	sdkprobehelper "github.com/ultradns/ultradns-go-sdk/pkg/probe/helper"
	"github.com/ultradns/ultradns-go-sdk/pkg/probe/tcp"
)

func ResourceProbeTCP() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceProbeTCPCreate,
		ReadContext:   resourceProbeTCPRead,
		UpdateContext: resourceProbeTCPUpdate,
		DeleteContext: resourceProbeTCPDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resourceProbeTCPSchema(),
	}
}

func resourceProbeTCPCreate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	services := meta.(*service.Service)
	probeData := getNewProbeTCP(rd)
	rrSetKeyData := rrset.NewRRSetKey(rd)

	rrSetKeyData.RecordType = probe.RecordTypeA

	res, err := services.ProbeService.Create(rrSetKeyData, probeData)

	if err != nil {
		return diag.FromErr(err)
	}

	uri := res.Header.Get("Location")

	rrSetKeyData.ID = helper.GetProbeIDFromURI(uri)

	rd.SetId(rrSetKeyData.PID())

	return resourceProbeTCPRead(ctx, rd, meta)
}

func resourceProbeTCPRead(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	services := meta.(*service.Service)
	rrSetKey := probe.GetRRSetKeyFromID(rd.Id())
	rrSetKey.PType = sdkprobe.TCP
	_, probeData, err := services.ProbeService.Read(rrSetKey)

	if err != nil {
		rd.SetId("")

		return nil
	}

	if err = probe.FlattenRRSetKey(rrSetKey, rd); err != nil {
		return diag.FromErr(err)
	}

	if err = flattenProbeTCP(probeData, rd); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceProbeTCPUpdate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	services := meta.(*service.Service)
	probeData := getNewProbeTCP(rd)
	rrSetKeyData := probe.GetRRSetKeyFromID(rd.Id())

	_, err := services.ProbeService.Update(rrSetKeyData, probeData)

	if err != nil {
		return diag.FromErr(err)
	}

	return resourceProbeTCPRead(ctx, rd, meta)
}

func resourceProbeTCPDelete(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func getNewProbeTCP(rd *schema.ResourceData) *sdkprobe.Probe {
	probeData := probe.NewProbe(rd, sdkprobe.TCP)
	details := &tcp.Details{}
	limits := &sdkprobehelper.LimitsInfo{}

	details.Limits = limits
	probeData.Details = details

	if val, ok := rd.GetOk("port"); ok {
		details.Port = val.(int)
	}

	if val, ok := rd.GetOk("control_ip"); ok {
		details.ControlIP = val.(string)
	}

	if val, ok := rd.GetOk("connect_limit"); ok && len(val.([]interface{})) > 0 {
		limitData := val.([]interface{})[0].(map[string]interface{})
		limits.Connect = probe.GetLimit(limitData)
	}

	if val, ok := rd.GetOk("avg_connect_limit"); ok && len(val.([]interface{})) > 0 {
		limitData := val.([]interface{})[0].(map[string]interface{})
		limits.AvgConnect = probe.GetLimit(limitData)
	}

	return probeData
}
