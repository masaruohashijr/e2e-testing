package say

import (
	"encoding/xml"
	"fmt"
	"time"
	"zarbat_test/internal/adapters/primary"
	"zarbat_test/internal/adapters/secondary"
	"zarbat_test/internal/config"
	"zarbat_test/internal/godog/services"
	"zarbat_test/internal/logging"
	"zarbat_test/pkg/domains"
	"zarbat_test/pkg/ports/calls"
	"zarbat_test/pkg/ports/numbers"

	"github.com/cucumber/godog"
)

var Configuration config.ConfigType
var CallSecondaryPort calls.SecondaryPort
var CallPrimaryPort calls.PrimaryPort
var NumberSecondaryPort numbers.SecondaryPort
var NumberPrimaryPort numbers.PrimaryPort
var ResponseSay domains.ResponseSay
var ResponseGather domains.ResponseGather
var ResponseRecord domains.ResponseRecord
var Ch = make(chan string)

func ConfiguredToGatherSpeech(number string) error {
	r := &domains.ResponseGather{
		Pause: domains.Pause{
			Length: 5,
		},
		Gather: domains.Gather{
			Input:    "speech",
			Language: "en-US",
			Timeout:  15,
			Action:   Configuration.BaseUrl + "/SpeechResult",
		},
	}
	ResponseGather = *r
	x, _ := xml.MarshalIndent(ResponseGather, "", "")
	strXML := domains.Header + string(x)
	logging.Debug.Println(strXML)
	services.WriteActionXML("gather", strXML)
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(number)
	Configuration.VoiceUrl = services.BaseUrl + "/Gather"
	NumberPrimaryPort.UpdateNumber()
	return nil
}

func ConfiguredToSay(number, speech string) error {
	s := &domains.ResponseSay{
		Pause: domains.Pause{
			Length: 5,
		},
		Say: domains.Say{
			Value:    speech,
			Voice:    "man",
			Language: "en-US",
			Loop:     3,
		},
	}
	ResponseSay = *s
	x, _ := xml.MarshalIndent(ResponseSay, "", "")
	strXML := domains.Header + string(x)
	logging.Debug.Println(strXML)
	services.WriteActionXML("say", strXML)
	return nil
}

func IMakeACallFromTo(numberA, numberB string) error {
	Configuration.From, Configuration.FromSid = Configuration.SelectNumber(numberA)
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(numberB)
	Configuration.Timeout = services.Timeout
	CallPrimaryPort.MakeCall()
	return nil
}

func MyTestSetupRuns() error {
	Configuration = config.NewConfig()
	go services.RunServer(Ch, false)
	CallSecondaryPort = secondary.NewCallsApi(&Configuration)
	CallPrimaryPort = primary.NewCallsService(CallSecondaryPort)
	NumberSecondaryPort = secondary.NewNumbersApi(&Configuration)
	NumberPrimaryPort = primary.NewNumbersService(NumberSecondaryPort)
	Configuration.ActionUrl = services.BaseUrl + "/Say"
	Configuration.StatusCallback = services.BaseUrl + "/Callback"
	return nil
}

func ShouldGetSpeech(speechOriginal string) error {
	speechResult := ""
	select {
	case speechResult = <-Ch:
		fmt.Printf("Result: %s\n", speechResult)
	case <-time.After(time.Duration(services.TestTimeout) * time.Second):
		logging.Debug.Println("timeout")
		Ch = nil
		return fmt.Errorf("timeout")
	}
	if speechResult != speechOriginal {
		return fmt.Errorf("Error %s", "The returned speech is different from the one expected by the test.")
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^"([^"]*)" configured to gather speech$`, ConfiguredToGatherSpeech)
	ctx.Step(`^"([^"]*)" configured to say "([^"]*)"$`, ConfiguredToSay)
	ctx.Step(`^I make a call from "([^"]*)" to "([^"]*)"$`, IMakeACallFromTo)
	ctx.Step(`^my test setup runs$`, MyTestSetupRuns)
	ctx.Step(`^Should get speech "([^"]*)"$`, ShouldGetSpeech)
}
