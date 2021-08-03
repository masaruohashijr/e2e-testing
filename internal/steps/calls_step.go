package steps

import (
	"encoding/xml"
	"strings"
	"zarbat_test/internal/godog/services"
	"zarbat_test/pkg/domains"
)

func IMakeACallFromTo(numberA, numberB string) error {
	Configuration.From, Configuration.FromSid = Configuration.SelectNumber(numberA)
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(numberB)
	Configuration.Timeout = services.Timeout
	checkSayGather()
	checkUrlEmpty()
	CallPrimaryPort.MakeCall()
	return nil
}

func checkUrlEmpty() {
	if Configuration.ActionUrl == "" {
		Configuration.ActionUrl = "http://zang.io/ivr/welcome/call"
	}
}

func checkSayGather() {
	if strings.HasSuffix(Configuration.VoiceUrl, "/Gather") && strings.HasSuffix(Configuration.ActionUrl, "/Say") {
		ResponseSay.Say.Loop = 2
		x, _ := xml.MarshalIndent(ResponseSay, "", "")
		strXML := domains.Header + string(x)
		println(strXML)
		services.WriteActionXML("say", strXML)
		Configuration.ActionUrl = services.BaseUrl + "/Say"
	}
}
