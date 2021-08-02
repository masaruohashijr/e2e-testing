package steps

import (
	"zarbat_test/internal/godog/services"
	"zarbat_test/pkg/domains"
)

func ConfiguredToRedirectToPingURL(number string) error {
	services.CloseChannel = true
	r := &domains.Redirect{
		Value: services.BaseUrl + "/Pinging",
	}
	ResponseRedirect.Redirect = *r
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(number)
	Configuration.VoiceUrl = services.BaseUrl + "/Redirect"
	NumberSecondaryPort.UpdateNumber()
	return nil
}
