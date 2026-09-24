package webforward

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	sdkerrors "github.com/ultradns/ultradns-go-sdk/pkg/errors"
	sdkwebforward "github.com/ultradns/ultradns-go-sdk/pkg/webforward"
)

func TestResourceWebForwardSchema(t *testing.T) {
	resource := ResourceWebForward()
	if err := resource.InternalValidate(resource.Schema, true); err != nil {
		t.Fatal(err)
	}

	if !resource.Schema["zone_name"].ForceNew {
		t.Error("zone_name must force a new resource")
	}
	if !resource.Schema["request_to"].ForceNew {
		t.Error("request_to must force a new resource")
	}
}

func TestDataSourceWebForwardSchema(t *testing.T) {
	resource := DataSourceWebForward()
	if err := resource.InternalValidate(resource.Schema, false); err != nil {
		t.Fatal(err)
	}

	if !resource.Schema["guid"].Optional || !resource.Schema["guid"].Computed {
		t.Error("guid must be optional and computed")
	}
	if !resource.Schema["request_to"].Optional || !resource.Schema["request_to"].Computed {
		t.Error("request_to must be optional and computed")
	}
	if len(resource.Schema["guid"].ExactlyOneOf) != 2 || len(resource.Schema["request_to"].ExactlyOneOf) != 2 {
		t.Error("guid and request_to must be mutually exclusive lookup keys")
	}
	if !resource.Schema["default_redirect_to"].Computed {
		t.Error("default_redirect_to must be computed")
	}
}

func TestExpandAndFlattenWebForward(t *testing.T) {
	resource := ResourceWebForward()
	rd := schema.TestResourceDataRaw(t, resource.Schema, map[string]interface{}{
		"zone_name":             "example.com.",
		"request_to":            "www.example.com",
		"default_redirect_to":   "https://example.com",
		"default_forward_type":  sdkwebforward.HTTP301Redirect,
		"relative_forward_type": sdkwebforward.ParameterAndPath,
		"record": []interface{}{map[string]interface{}{
			"redirect_to":  "https://store.example.com",
			"forward_type": sdkwebforward.HTTP302Redirect,
			"priority":     10,
			"rule": []interface{}{map[string]interface{}{
				"header":           "User-Agent",
				"match_criteria":   "CONTAINS",
				"value":            "mobile",
				"case_insensitive": true,
			}},
		}},
	})

	expanded := expandWebForward(rd)
	if expanded.DefaultRedirectTo != "https://example.com" {
		t.Fatalf("unexpected redirect URL: %q", expanded.DefaultRedirectTo)
	}
	if len(expanded.Records) != 1 || len(expanded.Records[0].Rules) != 1 {
		t.Fatalf("unexpected expanded records: %#v", expanded.Records)
	}

	expanded.GUID = "ABC123"
	expanded.State = "ACTIVE"
	if err := flattenWebForward(expanded, rd); err != nil {
		t.Fatal(err)
	}
	if got := rd.Get("guid").(string); got != "ABC123" {
		t.Errorf("guid = %q, want ABC123", got)
	}
	if got := rd.Get("record.#").(int); got != 1 {
		t.Errorf("record count = %d, want 1", got)
	}
}

func TestIsWebForwardNotFound(t *testing.T) {
	testCases := []struct {
		name string
		res  *http.Response
		err  error
		want bool
	}{
		{"no error", nil, nil, false},
		{"zone with no forwards returns 404", &http.Response{StatusCode: http.StatusNotFound}, fmt.Errorf("list failed"), true},
		{"guid absent from successful list", &http.Response{StatusCode: http.StatusOK}, sdkerrors.ResourceTypeNotFoundError("WebForward", "webForward", "ABC123"), true},
		{"already deleted", &http.Response{StatusCode: http.StatusBadRequest}, fmt.Errorf("{ code: '59002', message: 'Web forward with given guid does not exist under the zone.' }"), true},
		{"other api error", &http.Response{StatusCode: http.StatusInternalServerError}, fmt.Errorf("server error"), false},
		{"no response", nil, fmt.Errorf("connection refused"), false},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			if got := isWebForwardNotFound(testCase.res, testCase.err); got != testCase.want {
				t.Errorf("isWebForwardNotFound = %v, want %v", got, testCase.want)
			}
		})
	}
}

func TestWebForwardID(t *testing.T) {
	id := webForwardID("EXAMPLE.COM", "ABC123")
	if id != "example.com.:ABC123" {
		t.Errorf("id = %q, want example.com.:ABC123", id)
	}

	zoneName, guid, err := parseWebForwardID(id)
	if err != nil {
		t.Fatal(err)
	}
	if zoneName != "example.com." || guid != "ABC123" {
		t.Errorf("parsed id = %q, %q", zoneName, guid)
	}

	for _, invalid := range []string{"", "no-colon", ":missing-guid", "zone:"} {
		if _, _, err := parseWebForwardID(invalid); err == nil {
			t.Errorf("expected error parsing id %q", invalid)
		}
	}
}
