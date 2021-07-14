package main

import (
	"encoding/xml"
	"fmt"
	"net/url"
	"os"
	"testing"
	"time"
	"zarbat_test/godog/services"
	"zarbat_test/internal/adapters/primary"
	"zarbat_test/internal/adapters/secondary"
	"zarbat_test/internal/config"
	"zarbat_test/pkg/domains"

	"github.com/cucumber/godog"
)

func configuredToSendSms(numberA string) error {
	s := &domains.Sms{
		Value:          "Test SMS",
		To:             "+13432022744",
		From:           "+12267781734",
		StatusCallback: services.BaseUrl + "/SmsStatus",
	}
	ResponseSMS.Sms = *s
	x, _ := xml.MarshalIndent(s, "", "")
	println(string(x))
	return nil
}

func iMakeACallFromTo(numberA, numberB string) error {
	Configuration.From, _ = Configuration.SelectNumber(numberA)
	Configuration.To, _ = Configuration.SelectNumber(numberB)
	Configuration.VoiceUrl = ""
	x, _ := xml.MarshalIndent(ResponseSMS, "", "")
	strXML := domains.Header + string(x)
	println(strXML)
	services.WriteActionXML("sms", strXML)
	PrimaryPort.MakeCall()
	return nil
}

func myTestSetupRuns() error {
	Configuration = config.NewConfig()
	go services.RunServer(Ch)
	Configuration.ActionUrl = services.BaseUrl + "/sms"
	println(Configuration.AccountSid)
	SecondaryPort = secondary.NewCallsApi(&Configuration)
	PrimaryPort = primary.NewCallsService(SecondaryPort)
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
		Name:                 "zarbat_test",
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
