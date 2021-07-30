package record

import (
	"encoding/xml"
	"fmt"
	"time"
	"zarbat_test/internal/adapters/primary"
	"zarbat_test/internal/adapters/secondary"
	"zarbat_test/internal/config"
	"zarbat_test/internal/godog/services"
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

func ConfiguredToRecordCalls(number string) error {
	r := &domains.Record{
		Background:         services.Background,
		MaxLength:          services.MaxLength,
		FileFormat:         services.FileFormat,
		Transcribe:         true,
		TranscribeCallback: services.BaseUrl + "/TranscribeCallback",
	}
	p := &domains.Pause{
		Length: 3,
	}
	ResponseRecord.Pause = *p
	ResponseRecord.Record = *r
	x, _ := xml.MarshalIndent(ResponseRecord, "", "")
	strXML := domains.Header + string(x)
	println(strXML)
	services.WriteActionXML("record", strXML)
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(number)
	NumberPrimaryPort.UpdateNumber()
	return nil
}

func ConfiguredToSay(numberA, speech string) error {
	s := &domains.ResponseSay{
		Pause: domains.Pause{
			Length: 5,
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
	return nil
}

func IMakeACallFromTo(numberA, numberB string) error {
	Configuration.From, Configuration.FromSid = Configuration.SelectNumber(numberA)
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(numberB)
	Configuration.ActionUrl = services.BaseUrl + "/Say"
	CallPrimaryPort.MakeCall()
	return nil
}
func MyTestSetupRuns() error {
	Configuration = config.NewConfig()
	go services.RunServer(Ch, false)
	//Configuration.Fallback = services.BaseUrl + "/Fallback"
	Configuration.StatusCallback = services.BaseUrl + "/Callback"
	//Configuration.VoiceUrl = services.BaseUrl + "/Gather"
	Configuration.VoiceUrl = services.BaseUrl + "/Record"
	CallSecondaryPort = secondary.NewCallsApi(&Configuration)
	CallPrimaryPort = primary.NewCallsService(CallSecondaryPort)
	NumberSecondaryPort = secondary.NewNumbersApi(&Configuration)
	NumberPrimaryPort = primary.NewNumbersService(NumberSecondaryPort)
	return nil
}

func ShouldGetTranscription(transcription string) error {
	transcriptionResult := ""
	select {
	case transcriptionResult = <-Ch:
		fmt.Printf("Result: %s\n", transcriptionResult)
	case <-time.After(time.Duration(services.TestTimeout) * time.Second):
		fmt.Println("timeout")
		Ch = nil
		return fmt.Errorf("timeout")
	}
	if transcription != transcriptionResult {
		return fmt.Errorf("Error %s", "Not able to listen correct transcription.")
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^"([^"]*)" configured to record calls$`, ConfiguredToRecordCalls)
	ctx.Step(`^"([^"]*)" configured to say "([^"]*)"$`, ConfiguredToSay)
	ctx.Step(`^I make a call from "([^"]*)" to "([^"]*)"$`, IMakeACallFromTo)
	ctx.Step(`^my test setup runs$`, MyTestSetupRuns)
	ctx.Step(`^should get transcription "([^"]*)"$`, ShouldGetTranscription)
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
}
