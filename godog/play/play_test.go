package main

import (
	"e2e-testing/godog/general"
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

func configuredToPlay(arg1, arg2 string) error {
	p := &domains.Play{
		Value: general.BaseUrl + "/mp3/sample.mp3",
		Loop:  1,
	}
	ResponsePlay.Play = *p
	x, _ := xml.MarshalIndent(p, "", "")
	println(string(x))
	return nil
}

func iMakeACallFromTo(arg1, arg2 string) error {
	x, _ := xml.MarshalIndent(ResponsePlay, "", "")
	strXML := domains.Header + string(x)
	println(strXML)
	general.WriteActionXML(strXML)
	CallPrimaryPort.MakeCall()
	return nil
}

func myTestSetupRuns() error {
	configurationSetup()
	println(Configuration.AccountSid)
	CallSecondaryPort = secondary.NewCallsApi(&Configuration)
	CallPrimaryPort = primary.NewCallsService(CallSecondaryPort)
	// instantiate the proper Response
	return nil
}

func shouldBeAbleToListen(arg1 string) error {
	body := <-Ch
	println(body)
	speechResult := ""
	if speechResult != "" {
		return fmt.Errorf("Error %s", arg1)
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^"([^"]*)" configured to play "([^"]*)"$`, configuredToPlay)
	ctx.Step(`^I make a call from "([^"]*)" to "([^"]*)"$`, iMakeACallFromTo)
	ctx.Step(`^my test setup runs$`, myTestSetupRuns)
	ctx.Step(`^"([^"]*)" should be able to listen$`, shouldBeAbleToListen)
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
}

func TestMain(m *testing.M) {
	opts := godog.Options{
		Format:    "progress",
		Paths:     []string{"../../features/play"},
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
	go general.RunServer(Ch)
	Configuration.From = "+558140423562"
	Configuration.To = "+5561984385415"
	Configuration.StatusCallback = general.BaseUrl + "/Callback"
	Configuration.ActionUrl = general.BaseUrl + "/InboundXml"
	Configuration.VoiceUrl = general.BaseUrl + "/Gather"
}
