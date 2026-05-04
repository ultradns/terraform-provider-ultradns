package cdn

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/errors"
	"github.com/ultradns/terraform-provider-ultradns/internal/service"
	sdkhelper "github.com/ultradns/ultradns-go-sdk/pkg/helper"
)

func ResourceCDN() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCDNCreate,
		ReadContext:   resourceCDNRead,
		UpdateContext: resourceCDNUpdate,
		DeleteContext: resourceCDNDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resourceCDNSchema(),
	}
}

func resourceCDNCreate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Trace(ctx, "CDN resource create context invoked")

	services := meta.(*service.Service)
	accountName := rd.Get("account_name").(string)
	fqdn := sdkhelper.GetZoneFQDN(rd.Get("fqdn").(string))

	payload, err := expandCDNResource(rd, fqdn)
	if err != nil {
		return diag.FromErr(err)
	}

	_, err = services.CDNResourceService.Create(accountName, fqdn, payload)
	if err != nil {
		return diag.FromErr(err)
	}

	rd.SetId(cdnResourceID(accountName, fqdn))

	return resourceCDNRead(ctx, rd, meta)
}

func resourceCDNRead(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Trace(ctx, "CDN resource read context invoked")
	var diags diag.Diagnostics

	accountName, fqdn, err := parseCDNResourceID(rd.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	services := meta.(*service.Service)
	res, payload, err := services.CDNResourceService.Read(accountName, fqdn)
	if err != nil && res != nil && res.StatusCode == http.StatusNotFound {
		tflog.Warn(ctx, errors.ResourceNotFoundError(rd.Id()).Error())
		rd.SetId("")
		return nil
	}

	if err != nil {
		return diag.FromErr(err)
	}

	if err := rd.Set("account_name", accountName); err != nil {
		return diag.FromErr(err)
	}

	if err := flattenCDNResource(payload, rd); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceCDNUpdate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Trace(ctx, "CDN resource update context invoked")

	accountName, fqdn, err := parseCDNResourceID(rd.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	services := meta.(*service.Service)
	payload, err := expandCDNResource(rd, fqdn)
	if err != nil {
		return diag.FromErr(err)
	}

	_, err = services.CDNResourceService.Update(accountName, fqdn, payload)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceCDNRead(ctx, rd, meta)
}

func resourceCDNDelete(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Trace(ctx, "CDN resource delete context invoked")
	var diags diag.Diagnostics

	accountName, fqdn, err := parseCDNResourceID(rd.Id())
	if err != nil {
		rd.SetId("")
		return nil
	}

	services := meta.(*service.Service)
	_, err = services.CDNResourceService.Delete(accountName, fqdn)
	if err != nil {
		rd.SetId("")
		return diag.FromErr(err)
	}

	rd.SetId("")
	return diags
}

func parseCDNResourceID(id string) (string, string, error) {
	parts := strings.SplitN(id, ":", 2)
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", "", fmt.Errorf("invalid CDN resource id %q; expected account_name:fqdn", id)
	}

	return parts[0], sdkhelper.GetZoneFQDN(parts[1]), nil
}

func cdnResourceID(accountName, fqdn string) string {
	return fmt.Sprintf("%s:%s", accountName, fqdn)
}
