package rdpool

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/terraform-provider-ultradns/internal/errors"
	"github.com/ultradns/terraform-provider-ultradns/internal/rrset"
	"github.com/ultradns/ultradns-go-sdk/pkg/rdpool"
	sdkrrset "github.com/ultradns/ultradns-go-sdk/pkg/rrset"
)

const profileType = "*rdpool.Profile"

func flattenRDPool(resList *sdkrrset.ResponseList, rd *schema.ResourceData) error {
	if err := rrset.FlattenRRSetWithRecordData(resList, rd); err != nil {
		return err
	}

	profile, ok := resList.RRSets[0].Profile.(*rdpool.Profile)

	if !ok {
		return errors.ResourceTypeMismatched(profileType, fmt.Sprintf("%T", profile))
	}

	if err := rd.Set("order", profile.Order); err != nil {
		return err
	}

	if err := rd.Set("description", profile.Description); err != nil {
		return err
	}

	return nil
}
