package steps

import (
	"encoding/xml"
	"zarbat_test/internal/godog/services"
	"zarbat_test/pkg/domains"
)

func ConfiguredToDialAndSendDigitsTo(dialerNumber, digits, dialedNumber string) error {
	services.CloseChannel = true
	dialed, _ := Configuration.SelectNumber(dialedNumber)
	n := &domains.Number{
		Value:      dialed,
		SendDigits: digits,
	}
	d := &domains.DialNumber{
		Number: *n,
	}
	ResponseDialNumber.DialNumber = *d
	p := &domains.Hangup{}
	ResponseDialNumber.Hangup = *p
	x, _ := xml.MarshalIndent(ResponseDialNumber, "", "")
	strXML := domains.Header + string(x)
	services.WriteActionXML("number", strXML)
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(dialerNumber)
	Configuration.VoiceUrl = services.BaseUrl + "/Number"
	NumberPrimaryPort.UpdateNumber()
	println(string(x))
	return nil
}

func ShouldBeReset(number string) error {
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(number)
	Configuration.VoiceUrl = ""
	NumberPrimaryPort.UpdateNumber()
	return nil
}
