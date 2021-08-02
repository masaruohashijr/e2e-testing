package steps

import (
	"encoding/xml"
	"zarbat_test/internal/godog/services"
	"zarbat_test/pkg/domains"
)

func ConfiguredToSay(numberA, speech string) error {
	s := &domains.ResponseSay{
		Pause: domains.Pause{
			Length: 1,
		},
		Say: domains.Say{
			Value:    speech,
			Voice:    "man",
			Language: "en-US",
			Loop:     1,
		},
	}
	ResponseSay = *s
	x, _ := xml.MarshalIndent(ResponseSay, "", "")
	strXML := domains.Header + string(x)
	println(strXML)
	services.WriteActionXML("say", strXML)
	Configuration.ActionUrl = services.BaseUrl + "/Say"
	return nil
}
