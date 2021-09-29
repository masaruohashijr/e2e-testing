package steps

import (
	"encoding/xml"
	"fmt"
	"zarbat_test/internal/godog/services"
	"zarbat_test/pkg/domains"
)

func IShouldListAtLeastNotificationFromTo(number string) error {
	testHash := fmt.Sprint(TestHash)
	r := &domains.Record{
		Background: services.Background,
		MaxLength:  services.MaxLength,
		FileFormat: services.FileFormat,
		Transcribe: false,
		Action:     services.BaseUrl + "/RecordAction" + "?hash=" + testHash,
	}
	p := &domains.Pause{
		Length: 0,
	}
	ResponseRecord.Pause = *p
	ResponseRecord.Record = *r
	x, _ := xml.MarshalIndent(ResponseRecord, "", "")
	strXML := domains.Header + string(x)
	println(strXML)
	services.WriteActionXML("record", strXML)
	Configuration.VoiceUrl = services.BaseUrl + "/Record"
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(number)
	NumberPrimaryPort.UpdateNumber()
	return nil
}

func IShouldViewTheLastNotificationFromTo(number string) error {
	testHash := fmt.Sprint(TestHash)
	r := &domains.Record{
		Background: services.Background,
		MaxLength:  services.MaxLength,
		FileFormat: services.FileFormat,
		Transcribe: false,
		Action:     services.BaseUrl + "/RecordAction" + "?hash=" + testHash,
	}
	p := &domains.Pause{
		Length: 0,
	}
	ResponseRecord.Pause = *p
	ResponseRecord.Record = *r
	x, _ := xml.MarshalIndent(ResponseRecord, "", "")
	strXML := domains.Header + string(x)
	println(strXML)
	services.WriteActionXML("record", strXML)
	Configuration.VoiceUrl = services.BaseUrl + "/Record"
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(number)
	NumberPrimaryPort.UpdateNumber()
	return nil
}
