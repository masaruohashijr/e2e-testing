package steps

import (
	"encoding/xml"
	"fmt"
	"net/url"
	"strconv"
	"zarbat_test/internal/godog/services"
	"zarbat_test/internal/logging"
	"zarbat_test/pkg/domains"
)

func ConfiguredToHangUpAfterSeconds(Number string, timeInSeconds int) error {
	services.CloseChannel = true
	p := &domains.Pause{
		Length: timeInSeconds,
	}
	ResponseHangup.Pause = *p
	h := &domains.Hangup{}
	ResponseHangup.Hangup = *h
	x, _ := xml.MarshalIndent(ResponseHangup, "", "")
	strXML := domains.Header + string(x)
	logging.Debug.Println(strXML)
	services.WriteActionXML("hangup", strXML)
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(Number)
	Configuration.VoiceUrl = services.BaseUrl + "/Hangup"
	NumberSecondaryPort.UpdateNumber()
	return nil
}

func ShouldGetLastCallDurationMoreThanOrEqualsTo(number string, timeInSeconds int) error {
	bodyContent := <-Ch
	url_parameters, e := url.ParseQuery(bodyContent)
	if len(url_parameters["CallDuration"]) == 0 {
		return fmt.Errorf("The call duration should be more than the pause interval."+
			" Expected at least %d seconds but got %d seconds.", timeInSeconds, 0)
	}
	if e != nil {
		panic(e)
	}
	durationInSeconds, _ := strconv.Atoi(url_parameters["CallDuration"][0])
	fmt.Printf("Duration %d\n ", durationInSeconds)
	if durationInSeconds < timeInSeconds {
		return fmt.Errorf("The call duration should be more than the pause interval."+
			" Expected at least %d seconds but got %d seconds.", timeInSeconds, durationInSeconds)
	}
	return nil
}
