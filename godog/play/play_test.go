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

func configuredToPlay(number, mp3File string) error {
	Configuration.From, _ = Configuration.SelectNumber(number)
	a := &domains.Pause{
		Length: services.PlayPause,
	}
	p := &domains.Play{
		Value: services.BaseUrl + "/mp3/" + mp3File + ".mp3",
		Loop:  services.PlayLoop,
	}
	ResponsePlay.Pause = *a
	ResponsePlay.Play = *p
	x, _ := xml.MarshalIndent(p, "", "")
	println(string(x))
	return nil
}

func configuredToGatherSpeech(number string) error {
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(number) //"+5561984385415"
	p := &domains.Pause{
		Length: services.GatherPause,
	}
	g := &domains.Gather{
		Input:    "speech",
		Language: "en-US",
		Timeout:  services.GatherTimeOut,
		Action:   services.BaseUrl + "/SpeechResult",
	}
	ResponseGather.Pause = *p
	ResponseGather.Gather = *g
	x, _ := xml.MarshalIndent(ResponseGather, "", "")
	strXML := domains.Header + string(x)
	println(strXML)
	services.WriteActionXML("gather", strXML)
	// NumberPrimaryPort.UpdateNumber()
	return nil
}

func iMakeACallFromTo(arg1, arg2 string) error {
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

func shouldBeAbleToListen(arg1 string) error {
	speechResult := ""
	select {
	case speechResult = <-Ch:
		fmt.Printf("Result: %s\n", speechResult)
	case <-time.After(time.Duration(services.TestTimeout) * time.Second):
		Ch = nil
		return fmt.Errorf("timeout")
	}
	b := services.CompareTwoSentences(speechResult, "how are you this is a test the test has ended okay", 80)
	println(b)
	if !b {
		return fmt.Errorf("Error %s", "Not able to listen.")
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^"([^"]*)" configured to play "([^"]*)"$`, configuredToPlay)
	ctx.Step(`^"([^"]*)" configured to gather speech$`, configuredToGatherSpeech)
	ctx.Step(`^I make a call from "([^"]*)" to "([^"]*)"$`, iMakeACallFromTo)
	ctx.Step(`^my test setup runs$`, myTestSetupRuns)
	ctx.Step(`^"([^"]*)" should be able to listen$`, shouldBeAbleToListen)
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

func configurationSetup() {
	Configuration = config.NewConfig()
	go services.RunServer(Ch)
	Configuration.ActionUrl = services.BaseUrl + "/Play"
	//Configuration.Fallback = services.BaseUrl + "/Fallback"
	//Configuration.StatusCallback = services.BaseUrl + "/Callback"
	Configuration.VoiceUrl = services.BaseUrl + "/Gather"
}
