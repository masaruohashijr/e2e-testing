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

func appendToConfigHangup(numberA string) error {
	h := &domains.Hangup{}
	x, _ := xml.MarshalIndent(h, "", "")
	println(string(x))
	return nil
}

func configuredToPauseSeconds(numberA string, timeInSeconds int) error {
	p := &domains.Pause{
		Length: timeInSeconds,
	}
	ResponsePause.Pause = *p
	x, _ := xml.MarshalIndent(p, "", "")
	println(string(x))
	return nil
}

func iMakeACallFromTo(numberA, numberB string) error {
	x, _ := xml.MarshalIndent(ResponsePause, "", "")
	strXML := domains.Header + string(x)
	println(strXML)
	general.WriteActionXML(strXML)
	PrimaryPort.MakeCall()
	return nil
}

func myTestSetupRuns() error {
	Configuration = config.NewConfig()
	go general.RunServer(Ch)
	Configuration.From = "+558140423562"
	Configuration.To = "+5561984385415"
	Configuration.StatusCallback = general.BaseUrl + "/Callback"
	Configuration.ActionUrl = general.BaseUrl + "/InboundXml"
	println(Configuration.AccountSid)
	SecondaryPort = secondary.NewCallsApi(&Configuration)
	PrimaryPort = primary.NewCallsService(SecondaryPort)
	// instantiate the proper Response
	return nil
}

func shouldGetLastCallDurationMoreThanOrEqualsTo(number string, timeInSeconds int) error {
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

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^my test setup runs$`, myTestSetupRuns)
	ctx.Step(`^"([^"]*)" configured to pause (\d+) seconds$`, configuredToPauseSeconds)
	ctx.Step(`^append To "([^"]*)" config hangup$`, appendToConfigHangup)
	ctx.Step(`^I make a call from "([^"]*)" to "([^"]*)"$`, iMakeACallFromTo)
	ctx.Step(`^"([^"]*)" should get last call duration more than or equals to (\d+)$`, shouldGetLastCallDurationMoreThanOrEqualsTo)
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
}

func TestMain(m *testing.M) {
	opts := godog.Options{
		Format:    "progress",
		Paths:     []string{"../../features/pause"},
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

func SelectNumber(option string) string {
	switch option {
	case "NumberA":
		return Configuration.NumberA
	case "NumberB":
		return Configuration.NumberB
	}
	return ""
}
