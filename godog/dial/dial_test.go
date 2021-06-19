package main

import (
	"e2e-testing/godog/general"
	"e2e-testing/internal/adapters/primary"
	"e2e-testing/internal/adapters/secondary"
	"e2e-testing/internal/config"
	"e2e-testing/pkg/domains"
	"encoding/xml"
	"fmt"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/cucumber/godog"
)

func configuredToDial(arg1, arg2 string) error {
	d := &domains.Dial{
		Value:       "+15146627677",
		CallBackURL: "http://fe6732d93b0e.ngrok.io/Callback",
	}
	ResponseDial.Dial = *d
	p := &domains.Hangup{}
	ResponseDial.Hangup = *p
	x, _ := xml.MarshalIndent(d, "", "")
	y, _ := xml.MarshalIndent(p, "", "")

	println(string(x), string(y))
	return nil
}

func iMakeACallFromTo(arg1, arg2 string) error {
	x, _ := xml.MarshalIndent(ResponseDial, "", "")
	strXML := domains.Header + string(x)
	println(strXML)
	general.WriteActionXML(strXML)
	PrimaryPort.MakeCall()
	return nil
}

func myTestSetupRuns() error {
	Configuration = config.NewConfig()
	go general.RunServer(Ch)
	Configuration.From = "+12267781734" //+558140421695
	Configuration.To = "+13432022744"
	//Configuration.StatusCallback = "http://fe6732d93b0e.ngrok.io/Callback"
	Configuration.ActionUrl = general.BaseUrl + "/InboundXml"
	println(Configuration.AccountSid)
	SecondaryPort = secondary.NewCallsApi(&Configuration)
	PrimaryPort = primary.NewCallsService(SecondaryPort)
	return nil
}

func shouldGetTheIncomingCallFrom(arg1, arg2 string) error {
	bodyContent := <-Ch
	url_parameters, e := url.ParseQuery(bodyContent)
	if e != nil {
		panic(e)
	}
	orig_dialed := "+15146627677"
	orig_parent := "+12267781734"
	dialed_number := string(url_parameters["To"][0])
	parent_number := string(url_parameters["From"][0])

	if dialed_number != orig_dialed {
		return fmt.Errorf("Expected dialed: %s and found %s.", orig_dialed, dialed_number)
	}
	if parent_number != orig_parent {
		return fmt.Errorf("Expected From: %s and found %s.", orig_parent, parent_number)
	}

	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^"([^"]*)" configured to dial "([^"]*)"$`, configuredToDial)
	ctx.Step(`^I make a call from "([^"]*)" to "([^"]*)"$`, iMakeACallFromTo)
	ctx.Step(`^my test setup runs$`, myTestSetupRuns)
	ctx.Step(`^"([^"]*)" should get the incoming call from "([^"]*)"$`, shouldGetTheIncomingCallFrom)
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
}

func TestMain(m *testing.M) {
	opts := godog.Options{
		Format:    "progress",
		Paths:     []string{"../../features/dial"},
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
