package steps

import (
	"encoding/xml"
	"zarbat_test/internal/godog/services"
	"zarbat_test/pkg/domains"
)

func AppendToConfigHangup(numberA string) error {
	h := &domains.Hangup{}
	x, _ := xml.MarshalIndent(*h, "", "")
	strXML := string(x)
	services.AppendActionXML(Configuration.VoiceUrl, strXML)
	return nil
}

func ConfiguredToPauseSeconds(number string, timeInSeconds int) error {
	services.CloseChannel = true
	p := &domains.Pause{
		Length: timeInSeconds,
	}
	ResponsePause.Pause = *p
	x, _ := xml.MarshalIndent(ResponsePause, "", "")
	strXML := domains.Header + string(x)
	services.WriteActionXML("pause", strXML)
	Configuration.ActionUrl = "http://zang.io/ivr/welcome/call"
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(number)
	Configuration.VoiceUrl = services.BaseUrl + "/Pause"
	NumberSecondaryPort.UpdateNumber()
	return nil
}
