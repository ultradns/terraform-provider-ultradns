package webforward

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/errors"
	"github.com/ultradns/terraform-provider-ultradns/internal/service"
	sdkerrors "github.com/ultradns/ultradns-go-sdk/pkg/errors"
	sdkhelper "github.com/ultradns/ultradns-go-sdk/pkg/helper"
	sdkwebforward "github.com/ultradns/ultradns-go-sdk/pkg/webforward"
)

func DataSourceWebForward() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceWebForwardRead,
		Schema:      dataSourceWebForwardSchema(),
	}
}

func dataSourceWebForwardRead(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	zoneName := sdkhelper.GetZoneFQDN(rd.Get("zone_name").(string))

	services := meta.(*service.Service)
	webForward, err := findWebForward(services, zoneName, rd)
	if err != nil {
		return diag.FromErr(err)
	}

	rd.SetId(webForwardID(zoneName, webForward.GUID))
	if err := flattenWebForward(webForward, rd); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

// findWebForward looks a web forward up by guid when given, and otherwise
// lists the zone and matches on request_to, the forward's natural key.
func findWebForward(services *service.Service, zoneName string, rd *schema.ResourceData) (*sdkwebforward.WebForward, error) {
	if guid, ok := rd.GetOk("guid"); ok {
		_, webForward, err := services.WebForwardService.Read(zoneName, guid.(string))
		return webForward, err
	}

	requestTo := rd.Get("request_to").(string)
	_, listResponse, err := services.WebForwardService.List(zoneName, nil)

	if err != nil {
		return nil, err
	}

	var matches []*sdkwebforward.WebForward

	for _, webForward := range listResponse.WebForwards {
		if strings.EqualFold(webForward.RequestTo, requestTo) {
			matches = append(matches, webForward)
		}
	}

	switch len(matches) {
	case 0:
		return nil, errors.ResourceNotFoundError(fmt.Sprintf("%s:%s", zoneName, requestTo))
	case 1:
		return matches[0], nil
	default:
		return nil, sdkerrors.MultipleResourceFoundError("WebForward", requestTo)
	}
}
