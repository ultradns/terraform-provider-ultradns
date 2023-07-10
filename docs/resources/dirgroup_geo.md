---
subcategory: "Geo"
layout: "ultradns"
page_title: "ULTRADNS: ultradns_dirgroup_geo"
description: |-
  Manages the account-level geoIP groups in UltraDNS.
---

# Resource: ultradns_dirgroup_geo

Use this resource to manage zones in UltraDNS

## Example Usage

### Create Account-level geoIP group

```terraform
resource "ultradns_dirgroup_geo" "v4" {
    name         = "VisegradGroup"
    account_name = "my_account"
    codes        = [ "CZ", "HU", "PL", "SK",  ]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) (String) Name of the geoIP group.
* `account_name` - (Required) (String) 	Name of the account. It must be provided, but it can also be sourced from the `ULTRADNS_ACCOUNT` environment variable.
* `codes` - (Required) (String List) The codes for the geographical territories. [Valid GEO codes](#valid-geo-codes).
* `description` - (Optional) (String) 
## Import

GeoIP group can be imported by combining their `name` and `account_name`.<br/>
Example: `VisegradGroup:my_account`

Example:
```
$ terraform import ultradns_dirgroup_geo.v4 "VisegradGroup:my_account"
```


## Valid GEO Codes:

| Code 	| Meaning | Equivalent ISO codes |
| :--- 	| :----: | :--- |
|_________________________|__________________________________________________|__________________________________________________|
| `A1`  | Anonymous Proxy | None |
|_________________________|__________________________________________________|__________________________________________________|
| `A2`	| Satellite Provider | None |
|_________________________|__________________________________________________|__________________________________________________|
| `A3`	| Unknown / Uncategorized IPs | None |
|_________________________|__________________________________________________|__________________________________________________|
| `NAM`	| North America (including Central America and the Caribbean) | `AG`,`AI`,`AN`,`AW`,`BB`,`BL`,`BM`,</br>`BQ`,`BS`,`BZ`,`CA`,`CR`,`CU`,`CW`,</br>`DM`,`DO`,`GD`,`GL`,`GP`,`GT`,`HN`,</br>`HT`,`JM`,`KN`,`KY`,`LC`,`MF`,`MQ`,</br>`MS`,`MX`,`NI`,`PA`,`PM`,`PR`,`SV`,</br>`SX`,`TC`,`TT`,`U3`,`US`,`VC`,`VG`,</br>`VI` |
|_________________________|__________________________________________________|__________________________________________________|
| `SAM`	| South America | `AR`,`BO`,`BR`,`CL`,`CO`,`EC`,`FK`,</br>`GF`,`GS`,`GY`,`PE`,`PY`,`SR`,`U4`,</br>`UY`,`VE` |
|_________________________|__________________________________________________|__________________________________________________|
| `EUR`	| Europe | `AD`,`AL`,`AM`,`AT`,`AX`,`AZ`,`BA`,</br>`BE`,`BG`,`BY`,`CH`,`CZ`,`DE`,`DK`,</br>`EE`,`ES`,`FI`,`FO`,`FR`,`GB`,`GE`,</br>`GG`,`GI`,`GR`,`HR`,`HU`,`IE`,`IM`,</br>`IS`,`IT`,`JE`,`LI`,`LT`,`LU`,`LV`,</br>`MC`,`MD`,`ME`,`MK`,`MT`,`NL`,`NO`,</br>`PL`,`PT`,`RO`,`RS`,`SE`,`SI`,`SJ`,</br>`SK`,`SM`,`U5`,`UA`,`VA` |
|_________________________|__________________________________________________|__________________________________________________|
| `AFR`	| Africa | `AO`,`BF`,`BI`,`BJ`,`BW`,`CD`,`CF`,</br>`CG`,`CI`,`CM`,`CV`,`DJ`,`DZ`,`EG`,</br>`EH`,`ER`,`ET`,`GA`,`GH`,`GM`,`GN`,</br>`GQ`,`GW`,`KE`,`KM`,`LR`,`LS`,`LY`,</br>`MA`,`MG`,`ML`,`MR`,`MU`,`MW`,`MZ`,</br>`NA`,`NE`,`NG`,`RE`,`RW`,`SC`,`SD`,</br>`SH`,`SL`,`SN`,`SO`,`SS`,`ST`,`SZ`,</br>`TD`,`TG`,`TN`,`TZ`,`U7`,`UG`,`YT`,</br>`ZA`,`ZM`,`ZW` |
|_________________________|__________________________________________________|__________________________________________________|
| `ASI`	| Asia (including Middle East and the Russian Federation) | `AE`,`AF`,`BD`,`BH`,`BN`,`BT`,`CN`,</br>`CY`,`HK`,`ID`,`IL`,`IN`,`IO`,`IQ`,</br>`IR`,`JO`,`JP`,`KG`,`KH`,`KP`,`KR`,</br>`KW`,`KZ`,`LA`,`LB`,`LK`,`MM`,`MN`,</br>`MO`,`MV`,`MY`,`NP`,`OM`,`PH`,`PK`,</br>`PS`,`QA`,`RU`,`SA`,`SG`,`SY`,`TH`,</br>`TJ`,`TL`,`TM`,`TR`,`TW`,`U6`,`U8`,</br>`UZ`,`VN`,`YE` |
|_________________________|__________________________________________________|__________________________________________________|
| `OCN`	| Australia / Oceania | `AS`,`AU`,`CC`,`CK`,`CX`,`FJ`,`FM`,</br>`GU`,`HM`,`KI`,`MH`,`MP`,`NC`,`NF`,</br>`NR`,`NU`,`NZ`,`PF`,`PG`,`PN`,`PW`,</br>`SB`,`TK`,`TO`,`TV`,`U9`,`UM`,`VU`,</br>`WF`,`WS` |
|_________________________|__________________________________________________|__________________________________________________|
| `ANT`	| Antarctica | `AQ`, `TF`, `BV` |
|_________________________|__________________________________________________|__________________________________________________|
