package main

import (
	"e2e-testing/godog/services"
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

func configuredToDial(dialerNumber, dialedNumber string) error {
	number, _ := Configuration.SelectNumber(dialedNumber)
	d := &domains.Dial{
		Value:       number,
		CallBackURL: services.BaseUrl + "/Callback",
	}
	ResponseDial.Dial = *d
	p := &domains.Hangup{}
	ResponseDial.Hangup = *p
	x, _ := xml.MarshalIndent(ResponseDial, "", "")
	strXML := domains.Header + string(x)
	services.WriteActionXML("dial", strXML)
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(dialerNumber)
	NumbersPrimaryPort.UpdateNumber()
	println(string(x))
	return nil
}

func iMakeACallFromTo(numberA, numberB string) error {
	Configuration.From, Configuration.FromSid = Configuration.SelectNumber(numberA)
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(numberB)
	Configuration.VoiceUrl = ""
	x, _ := xml.MarshalIndent(ResponseDial, "", "")
	strXML := domains.Header + string(x)
	println(strXML)
	CallsPrimaryPort.MakeCall()
	return nil
}

func myTestSetupRuns() error {
	Configuration = config.NewConfig()
	go services.RunServer(Ch)
	Configuration.VoiceUrl = services.BaseUrl + "/Dial"
	CallsSecondaryPort = secondary.NewCallsApi(&Configuration)
	CallsPrimaryPort = primary.NewCallsService(CallsSecondaryPort)
	NumbersSecondaryPort = secondary.NewNumbersApi(&Configuration)
	NumbersPrimaryPort = primary.NewNumbersService(NumbersSecondaryPort)
	return nil
}

func shouldGetTheIncomingCallFrom(dialedNumber, dialerNumber string) error {
	bodyContent := <-Ch
	url_parameters, e := url.ParseQuery(bodyContent)
	if e != nil {
		panic(e)
	}
	dialed, _ := Configuration.SelectNumber(dialedNumber)
	orig_dialed := dialed
	orig_dialer := Configuration.To
	dialed_number := string(url_parameters["To"][0])
	dialer_number := string(url_parameters["From"][0])

	if dialed_number != orig_dialed {
		return fmt.Errorf("Expected dialed: %s and found %s.", orig_dialed, dialed_number)
	}
	if dialer_number != orig_dialer {
		return fmt.Errorf("Expected From: %s and found %s.", orig_dialer, dialer_number)
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
