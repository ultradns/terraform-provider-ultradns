package dirgroupip

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/service"
	"github.com/ultradns/ultradns-go-sdk/pkg/dirgroup/ip"
	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
)

func ResourceIPGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIPGroupCreate,
		ReadContext:   resourceIPGroupRead,
		UpdateContext: resourceIPGroupUpdate,
		DeleteContext: resourceIPGroupDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resourceIPGroupSchema(),
	}
}

func resourceIPGroupCreate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	services := meta.(*service.Service)
	ipGroupData := newIPGroup(rd)

	_, err := services.DirGroupIPService.Create(ipGroupData)

	if err != nil {
		return diag.FromErr(err)
	}

	rd.SetId(ipGroupData.DirGroupIPID())

	return resourceIPGroupRead(ctx, rd, meta)
}

func resourceIPGroupRead(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	services := meta.(*service.Service)
	ipID := rd.Id()

	_, ipGroup, _, err := services.DirGroupIPService.Read(ipID)
	//_, _, err := services.DirGroupIPService.Read(ipID)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := rd.Set("name", ipGroup.Name); err != nil {
		return diag.FromErr(err)
	}

	if err := rd.Set("account_name", helper.GetAccountName(ipID)); err != nil {
		return diag.FromErr(err)
	}

	if err := rd.Set("description", ipGroup.Description); err != nil {
		return diag.FromErr(err)
	}

	if err := rd.Set("ip", getSourceIPInfoSet(ipGroup.IPs)); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceIPGroupUpdate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	services := meta.(*service.Service)
	ipGroupData := newIPGroup(rd)

	_, err := services.DirGroupIPService.Update(ipGroupData)

	if err != nil {
		return diag.FromErr(err)
	}

	return resourceIPGroupRead(ctx, rd, meta)
}

func resourceIPGroupDelete(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	services := meta.(*service.Service)
	ipID := rd.Id()
	_, err := services.DirGroupIPService.Delete(ipID)

	if err != nil {
		rd.SetId("")

		return diag.FromErr(err)
	}

	rd.SetId("")

	return diags
}

func newIPGroup(rd *schema.ResourceData) *ip.DirGroupIP {
	ipData := &ip.DirGroupIP{}

	if val, ok := rd.GetOk("name"); ok {
		ipData.Name = val.(string)
	}
	if val, ok := rd.GetOk("account_name"); ok {
		ipData.AccountName = val.(string)
	}
	if val, ok := rd.GetOk("description"); ok {
		ipData.Description = val.(string)
	}
	if val, ok := rd.GetOk("ip"); ok {
		sourceIPInfoDataList := val.(*schema.Set).List()
		ipData.IPs = getSourceIPAddressList(sourceIPInfoDataList)
	}

	return ipData
}

func getSourceIPAddressList(sourceIPAddressDataList []interface{}) []*ip.IPAddress {
	sourceIPAddressList := make([]*ip.IPAddress, len(sourceIPAddressDataList))

	for i, d := range sourceIPAddressDataList {
		sourceIPAddressData := d.(map[string]interface{})
		sourceIPAddressList[i] = getSourceIPAddress(sourceIPAddressData)
	}

	return sourceIPAddressList
}

func getSourceIPAddress(sourceIPAddressData map[string]interface{}) *ip.IPAddress {
	sourceIPAddress := &ip.IPAddress{}

	if val, ok := sourceIPAddressData["start"]; ok {
		sourceIPAddress.Start = val.(string)
	}

	if val, ok := sourceIPAddressData["end"]; ok {
		sourceIPAddress.End = val.(string)
	}

	if val, ok := sourceIPAddressData["cidr"]; ok {
		sourceIPAddress.Cidr = val.(string)
	}

	if val, ok := sourceIPAddressData["address"]; ok {
		sourceIPAddress.Address = val.(string)
	}

	return sourceIPAddress
}

// temp here
func getSourceIPInfoSet(sourceIPDataList []*ip.IPAddress) *schema.Set {
	set := &schema.Set{F: schema.HashResource(sourceIPResource())}

	for _, sourceIPData := range sourceIPDataList {
		sourceIP := make(map[string]interface{})
		sourceIP["start"] = strings.ToLower(sourceIPData.Start)
		sourceIP["end"] = strings.ToLower(sourceIPData.End)
		sourceIP["cidr"] = strings.ToLower(sourceIPData.Cidr)
		sourceIP["address"] = strings.ToLower(sourceIPData.Address)
		set.Add(sourceIP)
	}

	return set
}
