package playlastrecording

import (
	"encoding/xml"
	"fmt"
	"strings"
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
var ResponsePlayLastRecording domains.ResponsePlayLastRecording
var ResponseGather domains.ResponseGather
var ResponseSay domains.ResponseSay
var ResponseRecord domains.ResponseRecord
var Ch = make(chan string)

func IMakeACallFromTo(numberA, numberB string) error {
	Configuration.From, Configuration.FromSid = Configuration.SelectNumber(numberA)
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(numberB) // "NumberBR1"
	Configuration.Timeout = services.Timeout
	CallPrimaryPort.MakeCall()
	return nil
}

func ConfiguredToPlayLastRecording(number string) error {
	strXML := domains.Header + string("<Response><PlayLastRecording/></Response>")
	logging.Debug.Println(strXML)
	services.WriteActionXML("playlastrecording", strXML)
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
	logging.Debug.Println(strXML)
	services.WriteActionXML("say", strXML)
	Configuration.ActionUrl = services.BaseUrl + "/Say"
	return nil
}

func ShouldGetSpeech(originalSpeech string) error {
	speechResult := ""
	select {
	case speechResult = <-Ch:
		fmt.Printf("Result: %s\n", speechResult)
	case <-time.After(time.Duration(services.TestTimeout) * time.Second):
		logging.Debug.Println("timeout")
		Ch = nil
		return fmt.Errorf("timeout")
	}
	originalSpeech = strings.ToLower(originalSpeech)
	speechResult = strings.ToLower(speechResult)
	logging.Debug.Println("original speech: ", originalSpeech)
	logging.Debug.Println("original speech: ", originalSpeech)
	logging.Debug.Println("speech result: ", speechResult)
	logging.Debug.Println("speech result: ", speechResult)
	if strings.TrimSpace(originalSpeech) != strings.TrimSpace(speechResult) && !strings.HasPrefix(strings.TrimSpace(originalSpeech), strings.TrimSpace(speechResult)) {
		return fmt.Errorf("Error: The returned speech \"%s\" is different from the one expected by the test.", speechResult)
	}
	return nil
}

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
	logging.Debug.Println(strXML)
	services.WriteActionXML("record", strXML)
	Configuration.VoiceUrl = services.BaseUrl + "/Record"
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(number)
	NumberPrimaryPort.UpdateNumber()
	return nil
}

func ConfiguredToGatherSpeech(numberA string) error {
	r := &domains.ResponseGather{
		Pause: domains.Pause{
			Length: 5,
		},
		Gather: domains.Gather{
			Input:    "speech",
			Language: "en-US",
			Timeout:  20,
			Action:   Configuration.BaseUrl + "/SpeechResult",
		},
	}
	ResponseGather = *r
	x, _ := xml.MarshalIndent(ResponseGather, "", "")
	strXML := domains.Header + string(x)
	logging.Debug.Println(strXML)
	services.WriteActionXML("gather", strXML)
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(numberA)
	Configuration.VoiceUrl = services.BaseUrl + "/Gather"
	NumberPrimaryPort.UpdateNumber()
	return nil
}

func MyTestSetupRuns() error {
	Configuration = config.NewConfig()
	go services.RunServer(Ch, false)
	CallSecondaryPort = secondary.NewCallsApi(&Configuration)
	CallPrimaryPort = primary.NewCallsService(CallSecondaryPort)
	NumberSecondaryPort = secondary.NewNumbersApi(&Configuration)
	NumberPrimaryPort = primary.NewNumbersService(NumberSecondaryPort)
	Configuration.StatusCallback = services.BaseUrl + "/Callback"
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^my test setup runs$`, MyTestSetupRuns)
	ctx.Step(`^"([^"]*)" configured to say "([^"]*)"$`, ConfiguredToSay)
	ctx.Step(`^"([^"]*)" configured to record calls$`, ConfiguredToRecordCalls)
	ctx.Step(`^I make a call from "([^"]*)" to "([^"]*)"$`, IMakeACallFromTo)
	ctx.Step(`^"([^"]*)" configured to gather speech$`, ConfiguredToGatherSpeech)
	ctx.Step(`^"([^"]*)" configured to play last recording$`, ConfiguredToPlayLastRecording)
	ctx.Step(`^should get speech "([^"]*)"$`, ShouldGetSpeech)
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {

}
