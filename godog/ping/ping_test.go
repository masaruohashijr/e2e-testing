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

func configuredToPingURL(numberA string) error {
	p := &domains.Ping{
		Value: general.BaseUrl + "/Ping",
	}
	ResponsePing.Ping = *p
	x, _ := xml.MarshalIndent(p, "", "")
	println(string(x))
	return nil
}

func iMakeACallFromTo(arg1, arg2 string) error {
	x, _ := xml.MarshalIndent(ResponsePing, "", "")
	strXML := domains.Header + string(x)
	println(strXML)
	general.WriteActionXML(strXML)
	PrimaryPort.MakeCall()
	return nil
}

func myTestSetupRuns() error {
	Configuration = config.NewConfig()
	go general.RunServer(Ch)
	Configuration.From = "+558140421695"
	Configuration.To = "+5561984385415"
	Configuration.ActionUrl = general.BaseUrl + "/InboundXml"
	println(Configuration.AccountSid)
	SecondaryPort = secondary.NewCallsApi(&Configuration)
	PrimaryPort = primary.NewCallsService(SecondaryPort)
	// instantiate the proper Response
	return nil
}

func shouldGetAPingRequestOnTheURL() error {
	println("Timer has started.")
	select {
	case res := <-Ch:
		fmt.Println(res)
	case <-time.After(60 * time.Second):
		fmt.Println("timeout 60")
		return fmt.Errorf("timeout 60")
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^"([^"]*)" configured to ping URL$`, configuredToPingURL)
	ctx.Step(`^I make a call from "([^"]*)" to "([^"]*)"$`, iMakeACallFromTo)
	ctx.Step(`^my test setup runs$`, myTestSetupRuns)
	ctx.Step(`^should get a ping request on the URL$`, shouldGetAPingRequestOnTheURL)
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {

}

func TestMain(m *testing.M) {
	opts := godog.Options{
		Format:    "progress",
		Paths:     []string{"../../features/ping"},
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
