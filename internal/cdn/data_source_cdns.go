package cdn

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/service"
	cdnresource "github.com/ultradns/ultradns-go-sdk/pkg/cdn/resource"
)

func DataSourceCDNs() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCDNsRead,
		Schema:      dataSourceCDNsSchema(),
	}
}

func dataSourceCDNsRead(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	services := meta.(*service.Service)
	accountName := rd.Get("account_name").(string)
	page := rd.Get("page").(int)
	size := rd.Get("size").(int)

	_, payload, err := services.CDNResourceService.List(accountName, &cdnresource.ListOptions{
		Page: page,
		Size: size,
	})
	if err != nil {
		return diag.FromErr(err)
	}

	rd.SetId(fmt.Sprintf("%s:%d:%d", accountName, page, size))

	if err := flattenCDNList(payload, rd); err != nil {
		return diag.FromErr(err)
	}

	return diags
}
