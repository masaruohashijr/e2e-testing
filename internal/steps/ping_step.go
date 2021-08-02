package steps

import (
	"encoding/xml"
	"fmt"
	"time"
	"zarbat_test/internal/godog/services"
	"zarbat_test/pkg/domains"
)

func ConfiguredToPingURL(number string) error {
	testHash := fmt.Sprint(TestHash)
	p := &domains.Ping{
		Value: services.BaseUrl + "/Pinging" + "?hash=" + testHash,
	}
	ResponsePing.Ping = *p
	x, _ := xml.MarshalIndent(ResponsePing, "", "")
	strXML := domains.Header + string(x)
	services.WriteActionXML("ping", strXML)
	println(string(x))
	Configuration.ActionUrl = "http://zang.io/ivr/welcome/call"
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(number)
	Configuration.VoiceUrl = services.BaseUrl + "/Ping"
	NumberPrimaryPort.UpdateNumber()
	return nil
}

func ShouldGetAPingRequestOnTheURL(number string) error {
	select {
	case res := <-Ch:
		fmt.Println(res)
	case <-time.After(time.Duration(services.TestTimeout) * time.Second):
		fmt.Println("timeout")
		Ch = nil
		return fmt.Errorf("timeout")
	}
	return nil
}
