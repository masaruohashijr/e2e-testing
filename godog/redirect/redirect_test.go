package main

import (
	"encoding/xml"
	"fmt"
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

func configuredToRedirectToPingURL(numberB string) error {
	r := &domains.Redirect{
		Value: services.BaseUrl + "/Pinging",
	}
	ResponseRedirect.Redirect = *r
	x, _ := xml.MarshalIndent(r, "", "")
	println(string(x))
	return nil
}

func iMakeACallFromTo(numberA, numberB string) error {
	Configuration.From, _ = Configuration.SelectNumber(numberA)
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(numberB)
	Configuration.VoiceUrl = ""
	x, _ := xml.MarshalIndent(ResponseRedirect, "", "")
	strXML := domains.Header + string(x)
	println(strXML)
	services.WriteActionXML("redirect", strXML)
	CallsPrimaryPort.MakeCall()
	return nil
}

func myTestSetupRuns() error {
	configurationSetup()
	println(Configuration.AccountSid)
	Configuration.ActionUrl = services.BaseUrl + "/Redirect"
	CallsSecondaryPort = secondary.NewCallsApi(&Configuration)
	CallsPrimaryPort = primary.NewCallsService(CallsSecondaryPort)
	NumbersSecondaryPort = secondary.NewNumbersApi(&Configuration)
	NumbersPrimaryPort = primary.NewNumbersService(NumbersSecondaryPort)
	return nil

}

func configurationSetup() {
	Configuration = config.NewConfig()
	go services.RunServer(Ch)
}

func shouldGetAPingRequestOnTheURL() error {
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
	ctx.Step(`^my test setup runs$`, myTestSetupRuns)
	ctx.Step(`^"([^"]*)" configured to redirect to ping URL$`, configuredToRedirectToPingURL)
	ctx.Step(`^I make a call from "([^"]*)" to "([^"]*)"$`, iMakeACallFromTo)
	ctx.Step(`^should get a ping request on the URL$`, shouldGetAPingRequestOnTheURL)
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
}

func TestMain(m *testing.M) {
	opts := godog.Options{
		Format:    "progress",
		Paths:     []string{"../../features/redirect"},
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
