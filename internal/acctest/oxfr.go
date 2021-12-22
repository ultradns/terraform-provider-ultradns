package acctest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
)

const (
	oxfrBasePath     = "oxfr/v1/zones"
	oxfrVersion      = 1
	oxfrZoneType     = "PRIMARY"
	oxfrFileFormat   = "BIND"
	oxfrZoneAxfrKey  = "zoneAxfr"
	oxfrAxfrFileKey  = "axfrFile"
	jsonContentType  = "application/json"
	octetContentType = "application/octet-stream"
)

var (
	oxfrServerHostURI = os.Getenv("ULTRADNS_UNIT_TEST_OXFR_HOST_URL")
)

func DeleteOxfrZone(zoneName string) {
	req, err := http.NewRequest(http.MethodDelete, getOxfrDeleteURI(zoneName), nil)

	if err != nil {
		return
	}

	do(req)
}

func CreateOxfrZone(zoneName string) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	err := writeZoneAxfrMetaData(writer, zoneName)

	if err != nil {
		return
	}

	err = writeAxfrFileData(writer, zoneName)

	if err != nil {
		return
	}

	writer.Close()

	req, err := http.NewRequest(http.MethodPost, getOxfrURI(), body)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	if err != nil {
		return
	}

	do(req)
}

func do(req *http.Request) {
	_, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Printf("%v", err)
	}
}

func getOxfrDeleteURI(zoneName string) string {
	return fmt.Sprintf("%s/%s", getOxfrURI(), zoneName)
}

func getOxfrURI() string {
	return fmt.Sprintf("%s/%s", oxfrServerHostURI, oxfrBasePath)
}

func getOxfrZoneAxfr(zoneName string) *ZoneAxfr {
	ipRange := &IPRange{
		StartIP: "1.1.1.1",
		EndIP:   "255.255.255.255",
	}
	transferACL := &TransferACL{
		IPRanges: []*IPRange{ipRange},
	}

	return &ZoneAxfr{
		ZoneName:    zoneName,
		Version:     oxfrVersion,
		ZoneType:    oxfrZoneType,
		FileFormat:  oxfrFileFormat,
		TransferACL: transferACL,
	}
}

func writeZoneAxfrMetaData(writer *multipart.Writer, zoneName string) error {
	header := make(textproto.MIMEHeader)
	header.Set("Content-Type", jsonContentType)
	header.Set("Content-Disposition", getFormDataString(oxfrZoneAxfrKey))
	partWriter, err := writer.CreatePart(header)

	if err != nil {
		return err
	}

	jsonData := new(bytes.Buffer)
	oxfrZoneAxfr := getOxfrZoneAxfr(zoneName)
	err = json.NewEncoder(jsonData).Encode(oxfrZoneAxfr)

	if err != nil {
		return err
	}

	_, err = io.Copy(partWriter, jsonData)

	if err != nil {
		return err
	}

	return nil
}

func writeAxfrFileData(writer *multipart.Writer, zoneName string) error {
	header := make(textproto.MIMEHeader)
	header.Set("Content-Type", octetContentType)
	header.Set("Content-Disposition", getFormDataString(oxfrAxfrFileKey))
	fileWriter, err := writer.CreatePart(header)

	if err != nil {
		return err
	}

	_, err = io.Copy(fileWriter, bytes.NewBufferString(getBindFormatData(zoneName)))

	if err != nil {
		return err
	}

	return nil
}

func getFormDataString(key string) string {
	return fmt.Sprintf(`form-data; name="%s";`, key)
}

func getBindFormatData(zoneName string) string {
	return fmt.Sprintf(
		`%[1]s		86400	IN	SOA	abc.com. domain.tech.%[1]s 2015060717 300 3600 1728000 86400
%[1]s	86400	IN	NS	nameserver.net.
%[1]s	86400	IN	SOA	abc.com. domain.tech.%[1]s 2015060717 300 3600 1728000 86400`, zoneName)
}
