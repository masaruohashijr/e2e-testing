package steps

import (
	"encoding/xml"
	"fmt"
	"net/url"
	"time"
	"zarbat_test/internal/godog/services"
	"zarbat_test/internal/logging"
	"zarbat_test/pkg/domains"
)

func ConfiguredToSendMMSAndMediaTo(numberB, message, media, numberC string) error {
	Configuration.From, _ = Configuration.SelectNumber(numberB)
	Configuration.To, _ = Configuration.SelectNumber(numberC)
	m := &domains.Mms{
		Value:          message,
		From:           Configuration.From,
		To:             Configuration.To,
		MediaUrl:       services.BaseUrl + "/Media",
		StatusCallback: services.BaseUrl + "/MmsStatus",
	}
	ResponseMMS.Mms = *m
	x, _ := xml.MarshalIndent(ResponseMMS, "", "")
	strXML := domains.Header + string(x)
	println(strXML)
	services.WriteActionXML("mms", strXML)
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(numberB)
	Configuration.VoiceUrl = services.BaseUrl + "/Mms"
	NumberPrimaryPort.UpdateNumber()
	return nil
}

func ShouldBeAbleToViewTheMMSWithMedia(number, message string, mediaName string) error {
	bodyContent := ""
	select {
	case bodyContent = <-Ch:
		fmt.Printf("Result: %s\n", bodyContent)
		logging.Debug.Printf("Result: %s\n", bodyContent)
	case <-time.After(time.Duration(services.TestTimeout) * time.Second):
		fmt.Println("timeout")
		logging.Debug.Println("timeout")
		Ch = nil
		return fmt.Errorf("timeout")
	}

	url_parameters, _ := url.ParseQuery(bodyContent)
	status := url_parameters["DlrStatus"][0]
	if status != "sent" {
		return fmt.Errorf("MMS not sent")
	}
	body := url_parameters["Body"][0]
	if body != message {
		return fmt.Errorf("Expected message %s different from %s.", message, body)
	}
	Configuration.VoiceUrl = ""
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(number)
	NumberPrimaryPort.UpdateNumber()
	return nil
}
