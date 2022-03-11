package probeping

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
	"github.com/ultradns/ultradns-go-sdk/pkg/probe/ping"
)

func ResourceProbePING() *schema.Resource {
	return &schema.Resource{

		CreateContext: resourceProbePINGCreate,
		ReadContext:   resourceProbePINGRead,
		UpdateContext: resourceProbePINGUpdate,
		DeleteContext: resourceProbePINGDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resourceProbePINGSchema(),
	}
}

func resourceProbePINGCreate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	services := meta.(*service.Service)
	probeData := getNewProbePING(rd)
	rrSetKeyData := rrset.NewRRSetKey(rd)

	rrSetKeyData.RecordType = probe.RecordTypeA

	res, err := services.ProbeService.Create(rrSetKeyData, probeData)

	if err != nil {
		return diag.FromErr(err)
	}

	uri := res.Header.Get("Location")

	rrSetKeyData.ID = helper.GetProbeIDFromURI(uri)

	rd.SetId(rrSetKeyData.PID())

	return resourceProbePINGRead(ctx, rd, meta)
}

func resourceProbePINGRead(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	services := meta.(*service.Service)
	rrSetKey := probe.GetRRSetKeyFromID(rd.Id())
	rrSetKey.PType = sdkprobe.PING
	_, probeData, err := services.ProbeService.Read(rrSetKey)

	if err != nil {
		rd.SetId("")

		return nil
	}

	if err = probe.FlattenRRSetKey(rrSetKey, rd); err != nil {
		return diag.FromErr(err)
	}

	if err = flattenProbePING(probeData, rd); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceProbePINGUpdate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	services := meta.(*service.Service)
	probeData := getNewProbePING(rd)
	rrSetKeyData := probe.GetRRSetKeyFromID(rd.Id())

	_, err := services.ProbeService.Update(rrSetKeyData, probeData)

	if err != nil {
		return diag.FromErr(err)
	}

	return resourceProbePINGRead(ctx, rd, meta)
}

func resourceProbePINGDelete(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func getNewProbePING(rd *schema.ResourceData) *sdkprobe.Probe {
	probeData := probe.NewProbe(rd, sdkprobe.PING)
	details := &ping.Details{}
	limits := &sdkprobehelper.LimitsInfo{}

	details.Limits = limits
	probeData.Details = details

	if val, ok := rd.GetOk("packets"); ok {
		details.Packets = val.(int)
	}

	if val, ok := rd.GetOk("packet_size"); ok {
		details.PacketSize = val.(int)
	}

	if val, ok := rd.GetOk("loss_percent_limit"); ok && len(val.([]interface{})) > 0 {
		limitData := val.([]interface{})[0].(map[string]interface{})
		limits.LossPercent = probe.GetLimit(limitData)
	}

	if val, ok := rd.GetOk("total_limit"); ok && len(val.([]interface{})) > 0 {
		limitData := val.([]interface{})[0].(map[string]interface{})
		limits.Total = probe.GetLimit(limitData)
	}

	if val, ok := rd.GetOk("average_limit"); ok && len(val.([]interface{})) > 0 {
		limitData := val.([]interface{})[0].(map[string]interface{})
		limits.Average = probe.GetLimit(limitData)
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
