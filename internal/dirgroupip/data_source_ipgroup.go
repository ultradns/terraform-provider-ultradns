package dirgroupip

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/service"
	"github.com/ultradns/ultradns-go-sdk/pkg/helper"
)

func DataSourceIPGroup() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIPGroupRead,

		Schema: dataSourceIPSchema(),
	}
}

func dataSourceIPGroupRead(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	services := meta.(*service.Service)
	ipGroupData := newIPGroup(rd)

	_, ipGroup, uri, err := services.DirGroupIPService.Read(ipGroupData.DirGroupIPID())

	if err != nil {
		return diag.FromErr(err)
	}

	rd.SetId(ipGroupData.DirGroupIPID())
	if err := rd.Set("name", ipGroup.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := rd.Set("account_name", helper.GetAccountNameFromURI(uri)); err != nil {
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
