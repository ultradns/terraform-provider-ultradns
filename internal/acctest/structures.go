package acctest

type ZoneAxfr struct {
	ZoneName    string       `json:"zoneName"`
	Version     int          `json:"version"`
	ZoneType    string       `json:"zoneType"`
	FileFormat  string       `json:"fileFormat"`
	TransferACL *TransferACL `json:"transferAcl"`
}

type TransferACL struct {
	IPRanges  []*IPRange  `json:"ipRanges"`
	NotifyIPs []string    `json:"notifyIps"`
	Tsig      interface{} `json:"tsig"`
}

type IPRange struct {
	StartIP string `json:"startIp"`
	EndIP   string `json:"endIp"`
}
