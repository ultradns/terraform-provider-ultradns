package cdn

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	cdnresource "github.com/ultradns/ultradns-go-sdk/pkg/cdn/resource"
)

func TestFlattenCDNResource(t *testing.T) {
	// Use resourceCDNSchema so that all fields set by flattenCDNResource are
	// present in the ResourceData (using dataSourceCDNSchema would miss the
	// Required writable fields added to the resource schema).
	rd := schema.TestResourceDataRaw(t, resourceCDNSchema(), map[string]interface{}{
		"account_name": "acct",
		"fqdn":         "cdn.example.com.",
		// Required fields that must be present for the schema to accept Set calls.
		"cdn_providers": []interface{}{},
		"config_properties": map[string]interface{}{},
		"preference_properties": map[string]interface{}{},
	})

	payload := &cdnresource.Resource{
		FQDN:        "cdn.example.com.",
		Type:        cdnresource.TypeSynthetic,
		ResourceID:  101,
		Name:        "cdn-resource",
		Description: "synthetic profile",
		TTL:         300,
		ContentType: "static",
		Version:     "v1",
		LastUpdated: "2026-01-01T00:00:00Z",
		OwnerName:   "acct",
	}

	if err := flattenCDNResource(payload, rd); err != nil {
		t.Fatalf("flattenCDNResource returned error: %v", err)
	}

	if got := rd.Get("fqdn").(string); got != payload.FQDN {
		t.Fatalf("unexpected fqdn: got %q, want %q", got, payload.FQDN)
	}
	if got := rd.Get("type").(string); got != payload.Type {
		t.Fatalf("unexpected type: got %q, want %q", got, payload.Type)
	}
	if got := rd.Get("resource_id").(int); got != payload.ResourceID {
		t.Fatalf("unexpected resource_id: got %d, want %d", got, payload.ResourceID)
	}
	if got := rd.Get("name").(string); got != payload.Name {
		t.Fatalf("unexpected name: got %q, want %q", got, payload.Name)
	}
	if got := rd.Get("description").(string); got != payload.Description {
		t.Fatalf("unexpected description: got %q, want %q", got, payload.Description)
	}
	if got := rd.Get("ttl").(int); got != payload.TTL {
		t.Fatalf("unexpected ttl: got %d, want %d", got, payload.TTL)
	}
	if got := rd.Get("content_type").(string); got != payload.ContentType {
		t.Fatalf("unexpected content_type: got %q, want %q", got, payload.ContentType)
	}
	if got := rd.Get("version").(string); got != payload.Version {
		t.Fatalf("unexpected version: got %q, want %q", got, payload.Version)
	}
	if got := rd.Get("last_updated").(string); got != payload.LastUpdated {
		t.Fatalf("unexpected last_updated: got %q, want %q", got, payload.LastUpdated)
	}
	if got := rd.Get("owner_name").(string); got != payload.OwnerName {
		t.Fatalf("unexpected owner_name: got %q, want %q", got, payload.OwnerName)
	}
}

func TestFlattenCDNList(t *testing.T) {
	rd := schema.TestResourceDataRaw(t, dataSourceCDNsSchema(), map[string]interface{}{
		"account_name": "acct",
		"page":         1,
		"size":         100,
	})

	payload := &cdnresource.ResponseList{
		Content: []*cdnresource.Item{
			{
				FQDN:       "a.example.com.",
				Type:       cdnresource.TypeBYOD,
				ResourceID: 201,
				Name:       "a",
			},
			{
				FQDN:       "b.example.com.",
				Type:       cdnresource.TypeSynthetic,
				ResourceID: 202,
				Name:       "b",
			},
		},
		TotalPages:    3,
		TotalElements: 2,
	}

	if err := flattenCDNList(payload, rd); err != nil {
		t.Fatalf("flattenCDNList returned error: %v", err)
	}

	if got := rd.Get("total_pages").(int); got != payload.TotalPages {
		t.Fatalf("unexpected total_pages: got %d, want %d", got, payload.TotalPages)
	}
	if got := rd.Get("total_elements").(int); got != payload.TotalElements {
		t.Fatalf("unexpected total_elements: got %d, want %d", got, payload.TotalElements)
	}

	cdns := rd.Get("cdns").([]interface{})
	if len(cdns) != 2 {
		t.Fatalf("unexpected cdns count: got %d, want 2", len(cdns))
	}

	first := cdns[0].(map[string]interface{})
	if first["type"].(string) != cdnresource.TypeBYOD {
		t.Fatalf("unexpected first type: got %q", first["type"].(string))
	}

	second := cdns[1].(map[string]interface{})
	if second["type"].(string) != cdnresource.TypeSynthetic {
		t.Fatalf("unexpected second type: got %q", second["type"].(string))
	}
}
