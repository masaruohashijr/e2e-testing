package steps

import (
	"encoding/xml"
	"zarbat_test/internal/godog/services"
	"zarbat_test/pkg/domains"
)

func ConfiguredToRedirectToPingURL(number string) error {
	services.CloseChannel = true
	r := &domains.Redirect{
		Value: services.BaseUrl + "/Pinging",
	}
	ResponseRedirect.Redirect = *r
	x, _ := xml.MarshalIndent(ResponseRedirect, "", "")
	strXML := domains.Header + string(x)
	println(strXML)
	services.WriteActionXML("redirect", strXML)
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(number)
	Configuration.VoiceUrl = services.BaseUrl + "/Redirect"
	NumberSecondaryPort.UpdateNumber()
	return nil
}
