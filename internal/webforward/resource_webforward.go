package webforward

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
	sdkwebforward "github.com/ultradns/ultradns-go-sdk/pkg/webforward"
)

func ResourceWebForward() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceWebForwardCreate,
		ReadContext:   resourceWebForwardRead,
		UpdateContext: resourceWebForwardUpdate,
		DeleteContext: resourceWebForwardDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resourceWebForwardSchema(),
	}
}

func resourceWebForwardCreate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Trace(ctx, "Web forward resource create context invoked")

	zoneName := sdkhelper.GetZoneFQDN(rd.Get("zone_name").(string))
	payload := expandWebForward(rd)
	services := meta.(*service.Service)
	_, created, err := services.WebForwardService.Create(zoneName, payload)
	if err != nil {
		return diag.FromErr(err)
	}

	rd.SetId(webForwardID(zoneName, created.GUID))
	return resourceWebForwardRead(ctx, rd, meta)
}

func resourceWebForwardRead(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Trace(ctx, "Web forward resource read context invoked")

	zoneName, guid, err := parseWebForwardID(rd.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	services := meta.(*service.Service)
	res, payload, err := services.WebForwardService.Read(zoneName, guid)
	if isWebForwardNotFound(res, err) {
		tflog.Warn(ctx, errors.ResourceNotFoundError(rd.Id()).Error())
		rd.SetId("")
		return nil
	}
	if err != nil {
		return diag.FromErr(err)
	}

	if err := rd.Set("zone_name", zoneName); err != nil {
		return diag.FromErr(err)
	}
	if err := flattenWebForward(payload, rd); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceWebForwardUpdate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Trace(ctx, "Web forward resource update context invoked")

	zoneName, guid, err := parseWebForwardID(rd.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	services := meta.(*service.Service)
	if _, err := services.WebForwardService.Update(zoneName, guid, expandWebForward(rd)); err != nil {
		return diag.FromErr(err)
	}

	return resourceWebForwardRead(ctx, rd, meta)
}

func resourceWebForwardDelete(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	tflog.Trace(ctx, "Web forward resource delete context invoked")

	zoneName, guid, err := parseWebForwardID(rd.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	services := meta.(*service.Service)
	res, err := services.WebForwardService.Delete(zoneName, guid)
	if err != nil && !isWebForwardNotFound(res, err) {
		return diag.FromErr(err)
	}

	rd.SetId("")
	return nil
}

func expandWebForward(rd *schema.ResourceData) *sdkwebforward.WebForward {
	webForward := &sdkwebforward.WebForward{
		RequestTo:              rd.Get("request_to").(string),
		DefaultRedirectTo:      rd.Get("default_redirect_to").(string),
		DefaultForwardType:     rd.Get("default_forward_type").(string),
		RelativeForwardType:    rd.Get("relative_forward_type").(string),
		DefaultRedirectType:    rd.Get("default_redirect_type").(string),
		CertificateID:          rd.Get("certificate_id").(string),
		CertificateManagedType: rd.Get("certificate_managed_type").(string),
		CertificateName:        rd.Get("certificate_name").(string),
		ExpirationDays:         rd.Get("expiration_days").(string),
	}

	if records, ok := rd.GetOk("record"); ok {
		webForward.Records = expandWebForwardRecords(records.([]interface{}))
	}

	return webForward
}

func expandWebForwardRecords(records []interface{}) []sdkwebforward.Record {
	result := make([]sdkwebforward.Record, len(records))
	for index, rawRecord := range records {
		record := rawRecord.(map[string]interface{})
		result[index] = sdkwebforward.Record{
			RedirectTo:  record["redirect_to"].(string),
			ForwardType: record["forward_type"].(string),
			Priority:    record["priority"].(int),
		}
		if rules, ok := record["rule"].([]interface{}); ok {
			result[index].Rules = expandWebForwardRules(rules)
		}
	}

	return result
}

func expandWebForwardRules(rules []interface{}) []sdkwebforward.Rule {
	result := make([]sdkwebforward.Rule, len(rules))
	for index, rawRule := range rules {
		rule := rawRule.(map[string]interface{})
		result[index] = sdkwebforward.Rule{
			Header:          rule["header"].(string),
			MatchCriteria:   rule["match_criteria"].(string),
			Value:           rule["value"].(string),
			CaseInsensitive: rule["case_insensitive"].(bool),
		}
	}

	return result
}

func flattenWebForward(webForward *sdkwebforward.WebForward, rd *schema.ResourceData) error {
	fields := map[string]interface{}{
		"guid":                     webForward.GUID,
		"request_to":               webForward.RequestTo,
		"default_redirect_to":      webForward.DefaultRedirectTo,
		"default_forward_type":     webForward.DefaultForwardType,
		"relative_forward_type":    webForward.RelativeForwardType,
		"default_redirect_type":    webForward.DefaultRedirectType,
		"certificate_id":           webForward.CertificateID,
		"certificate_managed_type": webForward.CertificateManagedType,
		"certificate_name":         webForward.CertificateName,
		"expiration_days":          webForward.ExpirationDays,
		"state":                    webForward.State,
		"error_description":        webForward.ErrorDescription,
		"record":                   flattenWebForwardRecords(webForward.Records),
	}
	for name, value := range fields {
		if err := rd.Set(name, value); err != nil {
			return err
		}
	}

	return nil
}

func flattenWebForwardRecords(records []sdkwebforward.Record) []interface{} {
	result := make([]interface{}, len(records))
	for index, record := range records {
		result[index] = map[string]interface{}{
			"redirect_to":  record.RedirectTo,
			"forward_type": record.ForwardType,
			"priority":     record.Priority,
			"rule":         flattenWebForwardRules(record.Rules),
		}
	}

	return result
}

func flattenWebForwardRules(rules []sdkwebforward.Rule) []interface{} {
	result := make([]interface{}, len(rules))
	for index, rule := range rules {
		result[index] = map[string]interface{}{
			"header":           rule.Header,
			"match_criteria":   rule.MatchCriteria,
			"value":            rule.Value,
			"case_insensitive": rule.CaseInsensitive,
		}
	}

	return result
}

func parseWebForwardID(id string) (string, string, error) {
	parts := strings.SplitN(id, ":", 2)
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", "", fmt.Errorf("invalid web forward resource id %q; expected zone_name:guid", id)
	}

	return sdkhelper.GetZoneFQDN(parts[0]), parts[1], nil
}

func webForwardID(zoneName, guid string) string {
	return fmt.Sprintf("%s:%s", sdkhelper.GetZoneFQDN(zoneName), guid)
}

func isWebForwardNotFound(res *http.Response, err error) bool {
	if err == nil {
		return false
	}

	if res != nil && res.StatusCode == http.StatusNotFound {
		return true
	}

	// The SDK reports a typed not-found error when the zone list succeeds but
	// the guid is absent (there is no single-item GET), and the API answers
	// 59002 when deleting an already-deleted forward.
	return strings.Contains(err.Error(), "Resource not found") ||
		strings.Contains(err.Error(), "59002")
}
