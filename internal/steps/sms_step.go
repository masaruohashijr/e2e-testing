package steps

import (
	"encoding/xml"
	"fmt"
	"net/url"
	"time"
	"zarbat_test/internal/godog/services"
	"zarbat_test/pkg/domains"
)

func ConfiguredToSendSMSTo(numberB, message, numberC string) error {
	Configuration.From, _ = Configuration.SelectNumber(numberB)
	Configuration.To, _ = Configuration.SelectNumber(numberC)
	s := &domains.Sms{
		Value:          message,
		From:           Configuration.From,
		To:             Configuration.To,
		StatusCallback: services.BaseUrl + "/SmsStatus",
	}
	ResponseSMS.Sms = *s
	x, _ := xml.MarshalIndent(ResponseSMS, "", "")
	strXML := domains.Header + string(x)
	println(strXML)
	services.WriteActionXML("sms", strXML)
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(numberB)
	Configuration.VoiceUrl = services.BaseUrl + "/Sms"
	NumberPrimaryPort.UpdateNumber()
	return nil
}

func ShouldBeAbleToViewTheSMS(number, message string) error {
	bodyContent := ""
	select {
	case bodyContent = <-Ch:
		fmt.Printf("Result: %s\n", bodyContent)
	case <-time.After(time.Duration(services.TestTimeout) * time.Second):
		fmt.Println("timeout")
		Ch = nil
		return fmt.Errorf("timeout")
	}

	url_parameters, _ := url.ParseQuery(bodyContent)
	status := url_parameters["DlrStatus"][0]
	if status != "sent" {
		return fmt.Errorf("SMS not sent")
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
