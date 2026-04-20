package cdn

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/service"
)

func DataSourceCDN() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCDNRead,
		Schema:      dataSourceCDNSchema(),
	}
}

func dataSourceCDNRead(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	services := meta.(*service.Service)
	accountName := rd.Get("account_name").(string)
	fqdn := strings.ToLower(rd.Get("fqdn").(string))

	_, payload, err := services.CDNResourceService.Read(accountName, fqdn)
	if err != nil {
		return diag.FromErr(err)
	}

	rd.SetId(fmt.Sprintf("%s:%s", accountName, fqdn))

	if err := flattenCDNResource(payload, rd); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
