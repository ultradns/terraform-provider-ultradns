package ultradns

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/ultradns/ultradns-go-sdk/ultradns"
)

func mapZoneSchema(zr *ultradns.ZoneResponse, rd *schema.ResourceData) error {
	if err := rd.Set("name", zr.Properties.Name); err != nil {
		return err
	}

	if err := rd.Set("account_name", zr.Properties.AccountName); err != nil {
		return err
	}

	if err := rd.Set("type", zr.Properties.Type); err != nil {
		return err
	}

	return nil
}

func mapPrimaryZoneSchema(zr *ultradns.ZoneResponse, rd *schema.ResourceData) error {
	err := mapZoneSchema(zr, rd)

	if err != nil {
		return err
	}

	set := &schema.Set{
		F: zeroIndexHash,
	}

	if val, ok := rd.GetOk("primary_create_info"); ok && val.(*schema.Set).Len() > 0 {
		primaryCreateInfo := val.(*schema.Set).List()[0].(map[string]interface{})

		if zr.NotifyAddresses != nil && len(*zr.NotifyAddresses) > 0 {
			s := &schema.Set{
				F: schema.HashSchema(&schema.Schema{Type: schema.TypeMap}),
			}

			for _, notifyAddressData := range *zr.NotifyAddresses {
				notifyAddress := make(map[string]interface{})

				notifyAddress["notify_address"] = notifyAddressData.NotifyAddress
				notifyAddress["description"] = notifyAddressData.Description

				s.Add(notifyAddress)
			}

			primaryCreateInfo["notify_addresses"] = s

		}

		if zr.RestrictIPList != nil && len(*zr.RestrictIPList) > 0 {
			s := &schema.Set{
				F: schema.HashSchema(&schema.Schema{Type: schema.TypeMap}),
			}

			for _, restrictIpData := range *zr.RestrictIPList {
				restrictIp := make(map[string]interface{})

				restrictIp["start_ip"] = restrictIpData.StartIp
				restrictIp["end_ip"] = restrictIpData.EndIp
				restrictIp["cidr"] = restrictIpData.Cidr
				restrictIp["single_ip"] = restrictIpData.SingleIp
				restrictIp["comment"] = restrictIpData.Comment

				s.Add(restrictIp)
			}

			primaryCreateInfo["restrict_ip"] = s

		}

		if zr.Tsig != nil {
			tsig := make(map[string]interface{})

			tsig["tsig_key_name"] = zr.Tsig.TsigKeyName
			tsig["tsig_key_value"] = zr.Tsig.TsigKeyValue
			tsig["tsig_algorithm"] = zr.Tsig.TsigAlgorithm
			tsig["description"] = zr.Tsig.Description

			s := &schema.Set{
				F: schema.HashSchema(&schema.Schema{Type: schema.TypeMap}),
			}

			s.Add(tsig)

			primaryCreateInfo["tsig"] = s

		}

		set.Add(primaryCreateInfo)
	}

	if err := rd.Set("primary_create_info", set); err != nil {
		return err
	}

	return nil
}

func mapSecondaryZoneSchema(zr *ultradns.ZoneResponse, rd *schema.ResourceData) error {
	err := mapZoneSchema(zr, rd)

	if err != nil {
		return err
	}

	return nil
}

func mapAliasZoneSchema(zr *ultradns.ZoneResponse, rd *schema.ResourceData) error {
	err := mapZoneSchema(zr, rd)

	if err != nil {
		return err
	}

	aliasCreateInfo := make(map[string]interface{})
	aliasCreateInfo["original_zone_name"] = zr.OriginalZoneName

	set := &schema.Set{
		F: zeroIndexHash,
	}

	set.Add(aliasCreateInfo)

	if err := rd.Set("alias_create_info", set); err != nil {
		return err
	}

	return nil
}
