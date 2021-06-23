package main

import (
	"e2e-testing/godog/services"
	"e2e-testing/internal/adapters/primary"
	"e2e-testing/internal/adapters/secondary"
	"e2e-testing/internal/config"
	"e2e-testing/pkg/domains"
	"encoding/xml"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/cucumber/godog"
)

func configuredToPlayTone(number, tone string) error {
	Configuration.From, _ = Configuration.SelectNumber(number)
	p := &domains.Play{
		Value: "tone_stream://%(" + tone + ")",
		Loop:  services.PlayLoop,
	}
	ResponsePlay.Play = *p
	x, _ := xml.MarshalIndent(p, "", "")
	println(string(x))
	return nil
}

func configuredToRecordCalls(number string) error {
	//Configuration.To = "+5561984385415"
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(number) //"+5561984385415"
	r := &domains.Record{
		Background: services.Background,
		MaxLength:  services.MaxLength,
		FileFormat: services.FileFormat,
		Action:     services.BaseUrl + "/RecordAction",
	}
	ResponseRecord.Record = *r
	x, _ := xml.MarshalIndent(ResponseRecord, "", "")
	strXML := domains.Header + string(x)
	println(strXML)
	services.WriteActionXML("record", strXML)
	NumberPrimaryPort.UpdateNumber()
	return nil
}

func iMakeACallFromTo(arg1, arg2 string) error {
	Configuration.Timeout = services.Timeout
	x, _ := xml.MarshalIndent(ResponsePlay, "", "")
	strXML := domains.Header + string(x)
	println(strXML)
	services.WriteActionXML("play", strXML)
	CallPrimaryPort.MakeCall()
	return nil
}

func myTestSetupRuns() error {
	configurationSetup()
	println(Configuration.AccountSid)
	CallSecondaryPort = secondary.NewCallsApi(&Configuration)
	CallPrimaryPort = primary.NewCallsService(CallSecondaryPort)
	NumberSecondaryPort = secondary.NewNumbersApi(&Configuration)
	NumberPrimaryPort = primary.NewNumbersService(NumberSecondaryPort)
	// instantiate the proper Response
	return nil
}

func configurationSetup() {
	Configuration = config.NewConfig()
	go services.RunServer(Ch)
	Configuration.ActionUrl = services.BaseUrl + "/Play"
	//Configuration.Fallback = services.BaseUrl + "/Fallback"
	//Configuration.StatusCallback = services.BaseUrl + "/Callback"
	//Configuration.VoiceUrl = services.BaseUrl + "/Gather"
	Configuration.VoiceUrl = services.BaseUrl + "/Record"
}

func shouldBeAbleToListenToFrequencies(frequencies string) error {
	recordUrl := ""
	select {
	case recordUrl = <-Ch:
		fmt.Printf("Result: %s\n", recordUrl)
	case <-time.After(time.Duration(services.TestTimeout) * time.Second):
		Ch = nil
		return fmt.Errorf("timeout")
	}
	err := services.DownloadFile("../../media/record.wav", recordUrl)
	if err != nil {
		return fmt.Errorf("Error %s", "Not able to download the record.")
	}
	frequenciesFromFile := services.GetFrequencies("../../media/record.wav")

	if frequencies != frequenciesFromFile {
		return fmt.Errorf("Error %s", "Not able to listen correct frequencies.")
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^"([^"]*)" configured to play tone "([^"]*)"$`, configuredToPlayTone)
	ctx.Step(`^"([^"]*)" configured to record calls$`, configuredToRecordCalls)
	ctx.Step(`^I make a call from "([^"]*)" to "([^"]*)"$`, iMakeACallFromTo)
	ctx.Step(`^my test setup runs$`, myTestSetupRuns)
	ctx.Step(`^"([^"]*)" should be able to listen to frequencies$`, shouldBeAbleToListenToFrequencies)
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
}

func TestMain(m *testing.M) {
	opts := godog.Options{
		Format:    "progress",
		Paths:     []string{"../../features/" + services.FeatureFolder},
		Randomize: time.Now().UTC().UnixNano(),
	}

	status := godog.TestSuite{
		Name:                 "e2e-testing",
		TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  InitializeScenario,
		Options:              &opts,
	}.Run()

	// Optional: Run `testing` package's logic besides godog.
	if st := m.Run(); st > status {
		status = st
	}

	os.Exit(status)
}
