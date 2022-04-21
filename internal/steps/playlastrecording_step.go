package steps

import (
	"zarbat_test/internal/godog/services"
	"zarbat_test/internal/logging"
	"zarbat_test/pkg/domains"
)

func ConfiguredToPlayLastRecording(number string) error {
	strXML := domains.Header + string("<Response><PlayLastRecording/></Response>")
	logging.Debug.Println(strXML)
	services.WriteActionXML("playlastrecording", strXML)
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(number)
	NumberPrimaryPort.UpdateNumber()
	return nil
}
