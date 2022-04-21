package steps

import (
	"encoding/xml"
	"fmt"
	"time"
	"zarbat_test/internal/godog/services"
	"zarbat_test/internal/logging"
	"zarbat_test/pkg/domains"
)

func ConfiguredToRejectCall(number string) error {
	p := &domains.Reject{}
	ResponseReject.Reject = *p
	x, _ := xml.MarshalIndent(p, "", "")
	Configuration.StatusCallback = services.BaseUrl + "/RejectCallBack"
	Configuration.VoiceUrl = services.BaseUrl + "/Reject"
	strXML := domains.Header + string(x)
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(number)
	Configuration.VoiceUrl = ""
	NumberSecondaryPort.UpdateNumber()
	logging.Debug.Println(strXML)
	services.WriteActionXML("reject", strXML)
	return nil
}

func ShouldGetCallCancelStatus(number string) error {
	logging.Debug.Println("Timer has started.")
	select {
	case res := <-Ch:
		logging.Debug.Println(res)
	case <-time.After(time.Duration(services.TestTimeout) * time.Second):
		logging.Debug.Println("timeout")
		return fmt.Errorf("timeout")
	}
	// Reset
	ShouldBeReset(number)
	return nil
}
