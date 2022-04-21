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

func IShouldListAtLeastSMSFromTo(number int, numberFrom, numberTo string) error {
	from, _ := Configuration.SelectNumber(numberFrom)
	to, _ := Configuration.SelectNumber(numberTo)
	smss, err1 := SmsPrimaryPort.ListSMS(from, to)
	if err1 != nil {
		return fmt.Errorf("Error found in list SMSs.")
	}
	if len(smss) < number {
		return fmt.Errorf("Error. Minimum number of sms expected is %d and found %d.", number, len(smss))
	}
	return nil
}
func IShouldViewTheSMSFromTo(message, numberFrom, numberTo string) error {
	from, _ := Configuration.SelectNumber(numberFrom)
	to, _ := Configuration.SelectNumber(numberTo)
	sms, err2 := SmsPrimaryPort.ViewSMS(SmsSid)
	if err2 != nil {
		return fmt.Errorf("Error found in view SMS.")
	}
	if sms.Status != "sent" {
		return fmt.Errorf("Status expected is %s and found %s.", "sent", sms.Status)
	} else if sms.From != from {
		return fmt.Errorf("From Number expected is %s and found %s.", from, sms.From)
	} else if sms.To != to {
		return fmt.Errorf("To Number expected is %s and found %s.", to, sms.To)
	} else if message != sms.Body {
		return fmt.Errorf("Message expected is %s and found %s.", message, sms.Body)
	}
	return nil
}

func ISendSMSFromTo(message, numberFrom, numberTo string) error {
	Configuration.From, Configuration.FromSid = Configuration.SelectNumber(numberFrom)
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(numberTo)
	sms, _ := SmsPrimaryPort.SendSMS(Configuration.From, Configuration.To, message)
	SmsSid = sms.Sid
	return nil
}

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
	logging.Debug.Println(strXML)
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
		logging.Debug.Println("timeout")
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
