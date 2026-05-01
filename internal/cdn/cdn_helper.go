package cdn

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cdnresource "github.com/ultradns/ultradns-go-sdk/pkg/cdn/resource"
)

func flattenCDNResource(payload *cdnresource.Resource, rd *schema.ResourceData) error {
	if payload == nil {
		return nil
	}

	if err := rd.Set("fqdn", payload.FQDN); err != nil {
		return err
	}
	if err := rd.Set("type", payload.Type); err != nil {
		return err
	}
	if err := rd.Set("resource_id", payload.ResourceID); err != nil {
		return err
	}
	if err := rd.Set("name", payload.Name); err != nil {
		return err
	}
	if err := rd.Set("description", payload.Description); err != nil {
		return err
	}
	if err := rd.Set("ttl", payload.TTL); err != nil {
		return err
	}
	if err := rd.Set("content_type", payload.ContentType); err != nil {
		return err
	}
	if err := rd.Set("version", payload.Version); err != nil {
		return err
	}
	if err := rd.Set("last_updated", payload.LastUpdated); err != nil {
		return err
	}
	if err := rd.Set("owner_name", payload.OwnerName); err != nil {
		return err
	}

	if payload.Configs != nil {
		if err := rd.Set("cdn_providers", flattenCdnProviders(payload.Configs.CDNs)); err != nil {
			return err
		}
		if err := rd.Set("config_properties", flattenAdditionalProperties(filterAdditionalProperties(payload.Configs.AdditionalProperties, rd, "config_properties"))); err != nil {
			return err
		}
	}

	if payload.Preferences != nil {
		if err := rd.Set("preference_properties", flattenAdditionalProperties(filterAdditionalProperties(payload.Preferences.AdditionalProperties, rd, "preference_properties"))); err != nil {
			return err
		}
	}

	return nil
}

func flattenCdnProviders(cdns []*cdnresource.CdnConfig) []interface{} {
	result := make([]interface{}, 0, len(cdns))
	for _, c := range cdns {
		if c == nil {
			continue
		}
		result = append(result, map[string]interface{}{
			"client_cdn_id": c.ClientCdnID,
			"cdn_name":      c.CdnName,
			"description":   c.Description,
			"fqdn":          c.FQDN,
		})
	}
	return result
}

// flattenAdditionalProperties encodes each value as a JSON string so it can be
// stored in a TypeMap(TypeString) schema attribute.
func flattenAdditionalProperties(props map[string]interface{}) map[string]string {
	out := make(map[string]string, len(props))
	for k, v := range props {
		b, err := json.Marshal(v)
		if err != nil {
			out[k] = fmt.Sprintf("%v", v)
		} else {
			out[k] = string(b)
		}
	}
	return out
}

// filterAdditionalProperties keeps user-declared keys stable in state and
// drops server-managed metadata keys that cause perpetual diffs.
func filterAdditionalProperties(props map[string]interface{}, rd *schema.ResourceData, attr string) map[string]interface{} {
	if len(props) == 0 {
		return props
	}

	desired := map[string]struct{}{}
	if raw, ok := rd.Get(attr).(map[string]interface{}); ok {
		for k := range raw {
			desired[k] = struct{}{}
		}
	}

	out := make(map[string]interface{})
	for k, v := range props {
		if isServerManagedPropertyKey(k) {
			continue
		}
		// If config declared keys, keep state aligned to only those keys.
		if len(desired) > 0 {
			if _, keep := desired[k]; !keep {
				continue
			}
		}
		out[k] = v
	}

	return out
}

func isServerManagedPropertyKey(key string) bool {
	switch strings.ToLower(strings.TrimSpace(key)) {
	case "id", "accountid", "resourceid", "created", "modified", "version", "ownername":
		return true
	default:
		return false
	}
}

// expandCDNResource builds the SDK Resource from Terraform ResourceData.
func expandCDNResource(rd *schema.ResourceData, fqdn string) (*cdnresource.Resource, error) {
	payload := &cdnresource.Resource{
		FQDN: fqdn,
		Type: rd.Get("type").(string),
		Name: rd.Get("name").(string),
	}

	if v, ok := rd.GetOkExists("description"); ok {
		payload.Description = v.(string)
	}
	if v, ok := rd.GetOkExists("ttl"); ok {
		payload.TTL = v.(int)
	}
	if v, ok := rd.GetOkExists("content_type"); ok {
		payload.ContentType = v.(string)
	}

	// cdn_providers → Configs.CDNs
	if v, ok := rd.GetOk("cdn_providers"); ok {
		rawList := v.([]interface{})
		cdns := make([]*cdnresource.CdnConfig, 0, len(rawList))
		for _, raw := range rawList {
			m := raw.(map[string]interface{})
			cdns = append(cdns, &cdnresource.CdnConfig{
				ClientCdnID: m["client_cdn_id"].(string),
				CdnName:     m["cdn_name"].(string),
				Description: m["description"].(string),
				FQDN:        m["fqdn"].(string),
			})
		}
		if payload.Configs == nil {
			payload.Configs = &cdnresource.Configs{}
		}
		payload.Configs.CDNs = cdns
	}

	// config_properties → Configs.AdditionalProperties
	if v, ok := rd.GetOk("config_properties"); ok {
		rawMap := v.(map[string]interface{})
		props, err := expandAdditionalProperties(rawMap, "config_properties")
		if err != nil {
			return nil, err
		}
		if payload.Configs == nil {
			payload.Configs = &cdnresource.Configs{}
		}
		payload.Configs.AdditionalProperties = props
	}

	// preference_properties → Preferences.AdditionalProperties
	if v, ok := rd.GetOk("preference_properties"); ok {
		rawMap := v.(map[string]interface{})
		props, err := expandAdditionalProperties(rawMap, "preference_properties")
		if err != nil {
			return nil, err
		}
		payload.Preferences = &cdnresource.Preferences{AdditionalProperties: props}
	}

	return payload, nil
}

func expandAdditionalProperties(rawMap map[string]interface{}, field string) (map[string]interface{}, error) {
	props := make(map[string]interface{}, len(rawMap))
	for k, raw := range rawMap {
		strVal, ok := raw.(string)
		if !ok {
			return nil, fmt.Errorf("%s[%s] must be a JSON string", field, k)
		}

		var decoded interface{}
		if err := json.Unmarshal([]byte(strVal), &decoded); err != nil {
			return nil, fmt.Errorf("%s[%s] must contain valid JSON: %w", field, k, err)
		}

		props[k] = decoded
	}

	return props, nil
}

func flattenCDNList(payload *cdnresource.ResponseList, rd *schema.ResourceData) error {
	if payload == nil {
		return nil
	}

	cdnItems := make([]interface{}, 0, len(payload.Content))
	for _, item := range payload.Content {
		if item == nil {
			continue
		}

		cdnItems = append(cdnItems, map[string]interface{}{
			"fqdn":        item.FQDN,
			"type":        item.Type,
			"resource_id": item.ResourceID,
			"name":        item.Name,
		})
	}

	if err := rd.Set("cdns", cdnItems); err != nil {
		return err
	}
	if err := rd.Set("total_pages", payload.TotalPages); err != nil {
		return err
	}
	if err := rd.Set("total_elements", payload.TotalElements); err != nil {
		return err
	}

	return nil
}
