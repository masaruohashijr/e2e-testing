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
	"strconv"
	"testing"
	"time"

	"github.com/cucumber/godog"
)

func appendToConfigHangup(numberA string) error {
	h := &domains.Hangup{}
	x, _ := xml.MarshalIndent(h, "", "")
	ResponsePause.Hangup = *h
	println(string(x))
	return nil
}

func configuredToPauseSeconds(numberB string, timeInSeconds int) error {
	p := &domains.Pause{
		Length: timeInSeconds,
	}
	ResponsePause.Pause = *p
	x, _ := xml.MarshalIndent(p, "", "")
	println(string(x))
	return nil
}

func iMakeACallFromTo(numberA, numberB string) error {
	Configuration.From, Configuration.FromSid = Configuration.SelectNumber(numberA)
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(numberB)
	Configuration.VoiceUrl = ""
	NumbersPrimaryPort.UpdateNumber()
	Configuration.ActionUrl = services.BaseUrl + "/Pause"
	x, _ := xml.MarshalIndent(ResponsePause, "", "")
	strXML := domains.Header + string(x)
	println(strXML)
	services.WriteActionXML("pause", strXML)
	CallsPrimaryPort.MakeCall()
	return nil
}

func myTestSetupRuns() error {
	Configuration = config.NewConfig()
	go services.RunServer(Ch)
	Configuration.StatusCallback = services.BaseUrl + "/Callback"
	Configuration.ActionUrl = services.BaseUrl + "/Pause"
	println(Configuration.AccountSid)
	CallsSecondaryPort = secondary.NewCallsApi(&Configuration)
	CallsPrimaryPort = primary.NewCallsService(CallsSecondaryPort)
	NumbersSecondaryPort = secondary.NewNumbersApi(&Configuration)
	NumbersPrimaryPort = primary.NewNumbersService(NumbersSecondaryPort)
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

	if st := m.Run(); st > status {
		status = st
	}

	os.Exit(status)
}
