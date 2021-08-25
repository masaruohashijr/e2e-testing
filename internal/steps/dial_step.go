package steps

import (
	"encoding/xml"
	"fmt"
	"net/url"
	"strconv"
	"time"
	"zarbat_test/internal/godog/services"
	"zarbat_test/internal/logging"
	"zarbat_test/pkg/domains"
)

func ConfiguredToDial(dialerNumber, dialedNumber string) error {
	services.CloseChannel = true
	dialed, _ := Configuration.SelectNumber(dialedNumber)
	d := &domains.Dial{
		Value:       dialed,
		CallBackURL: services.BaseUrl + "/DialCallback",
	}
	ResponseDial.Dial = *d
	p := &domains.Hangup{}
	ResponseDial.Hangup = *p
	x, _ := xml.MarshalIndent(ResponseDial, "", "")
	strXML := domains.Header + string(x)
	services.WriteActionXML("dial", strXML)
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(dialerNumber)
	Configuration.VoiceUrl = services.BaseUrl + "/Dial"
	NumberPrimaryPort.UpdateNumber()
	logging.Debug.Println(string(x))
	return nil
}

func ShouldGetTheIncomingCallFrom(dialedNumber, dialerNumber string) error {
	bodyContent := ""
	select {
	case bodyContent = <-Ch:
		fmt.Println(bodyContent)
	case <-time.After(time.Duration(services.TestTimeout) * time.Second):
		tt := strconv.FormatInt(services.TestTimeout, 10)
		return fmt.Errorf("Timeout %s.", tt)
	}
	url_parameters, e := url.ParseQuery(bodyContent)
	if e != nil {
		panic(e)
	}
	dialed, _ := Configuration.SelectNumber(dialedNumber)
	orig_dialed := dialed
	orig_dialer, _ := Configuration.SelectNumber(dialerNumber)
	if len(url_parameters["From"]) == 0 || len(url_parameters["To"]) == 0 {
		return fmt.Errorf("Expected dialed: %s and found %s.", orig_dialed, "none")
	}
	dialer_number := string(url_parameters["From"][0])
	dialed_number := string(url_parameters["To"][0])

	if dialed_number != orig_dialed {
		return fmt.Errorf("Expected dialed: %s and found %s.", orig_dialed, dialed_number)
	}
	if dialer_number != orig_dialer {
		return fmt.Errorf("Expected From: %s and found %s.", orig_dialer, dialer_number)
	}
	// Reset
	ShouldBeReset(dialerNumber)
	return nil
}
