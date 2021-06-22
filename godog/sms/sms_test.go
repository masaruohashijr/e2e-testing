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

func configuredToSendSms(numberA string) error {
	s := &domains.Sms{
		Value:          "Test SMS",
		To:             "+13432022744",
		From:           "+12267781734",
		StatusCallback: "https://018d09d8beb2.ngrok.io/StatusCallback",
	}
	ResponseSMS.Sms = *s
	x, _ := xml.MarshalIndent(s, "", "")
	println(string(x))
	return nil
}

func iMakeACallFromTo(numberA, numberB string) error {
	x, _ := xml.MarshalIndent(ResponseSMS, "", "")
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
	Configuration.ActionUrl = "https://018d09d8beb2.ngrok.io/InboundXml"
	println(Configuration.AccountSid)
	SecondaryPort = secondary.NewCallsApi(&Configuration)
	PrimaryPort = primary.NewService(SecondaryPort)
	// instantiate the proper Response
	return nil
}

func sMSStatusShouldBeSentToCallStatusURL() error {
	bodyContent := <-Ch

	url_parameters, _ := url.ParseQuery(bodyContent)
	println(url_parameters)

	status := url_parameters["DlrStatus"][0]
	from_number := url_parameters["From"][0]
	to_number := url_parameters["To"][0]
	orig_from := "+12267781734"
	orig_to := "+13432022744"
	if status != "delivered" {
		return fmt.Errorf("SMS not delivered")
	}
	if orig_from != from_number {
		return fmt.Errorf("Expected From: %s and found %s.", orig_from, from_number)
	}
	if orig_to != to_number {
		return fmt.Errorf("Expected T: %s and found %s.", orig_to, orig_to)
	}

	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^"([^"]*)" configured to send sms$`, configuredToSendSms)
	ctx.Step(`^I make a call from "([^"]*)" to "([^"]*)"$`, iMakeACallFromTo)
	ctx.Step(`^my test setup runs$`, myTestSetupRuns)
	ctx.Step(`^SMS Status should be sent to call Status URL$`, sMSStatusShouldBeSentToCallStatusURL)
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
}

func TestMain(m *testing.M) {
	opts := godog.Options{
		Format:    "progress",
		Paths:     []string{"../../features/sms"},
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
