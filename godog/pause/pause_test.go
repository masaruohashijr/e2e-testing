package main

import (
	"e2e-testing/internal/config"
	"encoding/xml"
	"os"
	"testing"
	"time"

	"github.com/cucumber/godog"
)

var ResponsePause config.ResponsePause

func appendToConfigHangup(arg1 string) error {
	return nil
}

func configuredToPauseSeconds(numberB string, timeInSeconds int) error {
	p := &config.Pause{
		Length: timeInSeconds,
	}
	ResponsePause.Pause = *p
	x, _ := xml.MarshalIndent(p, "", "")
	println(string(x))
	return nil
}

func iMakeACallFromTo(arg1, arg2 string) error {
	return nil
}

func myTestSetupRuns() error {
	//Configuration = config.NewConfig()
	/*secondaryPort := secondary.NewCallsApi(Configuration) // The Secondary Adapter
	primaryPort := primary.NewService(secondaryPort)
	primaryAdapter := primary.NewCLIPrimaryAdapter(primaryPort)*/
	// instantiate the proper Response
	return nil
}

func shouldGetLastCallDurationEqualsTo(arg1 string, arg2 int) error {
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^my test setup runs$`, myTestSetupRuns)
	ctx.Step(`^"([^"]*)" configured to pause (\d+) seconds$`, configuredToPauseSeconds)
	ctx.Step(`^append To "([^"]*)" config hangup$`, appendToConfigHangup)
	ctx.Step(`^I make a call from "([^"]*)" to "([^"]*)"$`, iMakeACallFromTo)
	ctx.Step(`^"([^"]*)" should get last call duration equals to (\d+)$`, shouldGetLastCallDurationEqualsTo)
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
