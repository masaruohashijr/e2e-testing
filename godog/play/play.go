package play

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"time"
	"zarbat_test/godog/services"
	"zarbat_test/internal/adapters/primary"
	"zarbat_test/internal/adapters/secondary"
	"zarbat_test/internal/config"
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
var ResponsePlay domains.ResponsePlay
var ResponseGather domains.ResponseGather
var ResponseRecord domains.ResponseRecord
var Ch = make(chan string)

func ConfiguredToPlayTone(number, tone string) error {
	Configuration.From, Configuration.FromSid = Configuration.SelectNumber(number) //"+5561984385415"
	//Configuration.From, _ = Configuration.SelectNumber(number)
	p := &domains.Play{
		Value: "tone_stream://%(" + tone + ")",
		Loop:  services.PlayLoop,
	}
	ResponsePlay.Play = *p
	x, _ := xml.MarshalIndent(p, "", "")
	println(string(x))
	return nil
}

func ConfiguredToRecordCalls(number string) error {
	// Configuration.To = "+5561984385415"
	// Configuration.To, Configuration.ToSid = Configuration.SelectNumber(number) //"+5561984385415"
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(number) //"+5561984385415"
	r := &domains.Record{
		Background: services.Background,
		MaxLength:  services.MaxLength,
		FileFormat: services.FileFormat,
		Action:     services.BaseUrl + "/RecordAction",
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
	NumberPrimaryPort.UpdateNumber()
	return nil
}

func IMakeACallFromTo(numberA, numberB string) error {
	Configuration.Timeout = services.Timeout
	x, _ := xml.MarshalIndent(ResponsePlay, "", "")
	strXML := domains.Header + string(x)
	println(strXML)
	services.WriteActionXML("play", strXML)
	CallPrimaryPort.MakeCall()
	return nil
}

func MyTestSetupRuns() error {
	ConfigurationSetup()
	// println(Configuration.AccountSid)
	CallSecondaryPort = secondary.NewCallsApi(&Configuration)
	CallPrimaryPort = primary.NewCallsService(CallSecondaryPort)
	NumberSecondaryPort = secondary.NewNumbersApi(&Configuration)
	NumberPrimaryPort = primary.NewNumbersService(NumberSecondaryPort)
	// instantiate the proper Response
	return nil
}

func ConfigurationSetup() {
	Configuration = config.NewConfig()
	go services.RunServer(Ch, false)
	Configuration.ActionUrl = services.BaseUrl + "/Play"
	//Configuration.Fallback = services.BaseUrl + "/Fallback"
	Configuration.StatusCallback = services.BaseUrl + "/Callback"
	//Configuration.VoiceUrl = services.BaseUrl + "/Gather"
	Configuration.VoiceUrl = services.BaseUrl + "/Record"
}

func ShouldBeAbleToListenToFrequencies(number, frequencies string) error {
	recordUrl := ""
	select {
	case recordUrl = <-Ch:
		fmt.Printf("Result: %s\n", recordUrl)
	case <-time.After(time.Duration(services.TestTimeout) * time.Second):
		fmt.Println("timeout")
		Ch = nil
		return fmt.Errorf("timeout")
	}
	err := services.DownloadFile("media/record.wav", recordUrl)
	if err != nil {
		return fmt.Errorf("Error %s", "Not able to download the record.")
	}
	iFrequencies, _ := strconv.Atoi(frequencies)
	err = services.GetFrequencies("media/record.wav", iFrequencies, 90)
	if err != nil {
		return fmt.Errorf("Error %s", "Not able to listen correct frequencies.")
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^"([^"]*)" configured to play tone "([^"]*)"$`, ConfiguredToPlayTone)
	ctx.Step(`^"([^"]*)" configured to record calls$`, ConfiguredToRecordCalls)
	ctx.Step(`^I make a call from "([^"]*)" to "([^"]*)"$`, IMakeACallFromTo)
	ctx.Step(`^my test setup runs$`, MyTestSetupRuns)
	ctx.Step(`^"([^"]*)" should be able to listen to frequencies "([^"]*)"$`, ShouldBeAbleToListenToFrequencies)
}
