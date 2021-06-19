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
	"strconv"
	"testing"
	"time"

	"github.com/cucumber/godog"
)

func configuredToHangupAfterSeconds(NumberA string, timeInSeconds int) error {
	p := &domains.Pause{
		Length: timeInSeconds,
	}
	ResponseHangup.Pause = *p
	h := &domains.Hangup{}
	ResponseHangup.Hangup = *h
	x, _ := xml.MarshalIndent(p, "", "")
	y, _ := xml.MarshalIndent(h, "", "")
	println(string(x), string(y))
	return nil

}

func iMakeACallFromTo(NumberA, NumberB string) error {
	x, _ := xml.MarshalIndent(ResponseHangup, "", "")
	strXML := domains.Header + string(x)
	println(strXML)
	general.WriteActionXML(strXML)
	PrimaryPort.MakeCall()
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^"([^"]*)" configured to hangup after (\d+) seconds$`, configuredToHangupAfterSeconds)
	ctx.Step(`^I make a call from "([^"]*)" to "([^"]*)"$`, iMakeACallFromTo)
	ctx.Step(`^my test setup runs$`, myTestSetupRuns)
	ctx.Step(`^"([^"]*)" should get last call duration equals or more than (\d+)$`, shouldGetLastCallDurationEqualsOrMoreThan)
}

func myTestSetupRuns() error {
	Configuration = config.NewConfig()
	go general.RunServer(Ch)
	Configuration.From = "+12267781734"
	Configuration.To = "+13432022744"
	Configuration.StatusCallback = general.BaseUrl + "/Callback"
	Configuration.ActionUrl = general.BaseUrl + "/InboundXml"
	SecondaryPort = secondary.NewCallsApi(&Configuration)
	PrimaryPort = primary.NewService(SecondaryPort)
	return nil
}

func shouldGetLastCallDurationEqualsOrMoreThan(number string, timeInSeconds int) error {
	bodyContent := <-Ch
	url_parameters, e := url.ParseQuery(bodyContent)
	if e != nil {
		panic(e)
	}
	durationInSeconds, _ := strconv.Atoi(url_parameters["CallDuration"][0])
	fmt.Printf("Duration %d\n ", durationInSeconds)
	if durationInSeconds < timeInSeconds {
		return fmt.Errorf("The call duration should be more than the pause interval."+
			" Expected at least %d seconds but got %d seconds.", timeInSeconds, durationInSeconds)
	}
	return nil
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
}

func TestMain(m *testing.M) {
	opts := godog.Options{
		Format:    "progress",
		Paths:     []string{"../../features/hangup"},
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
