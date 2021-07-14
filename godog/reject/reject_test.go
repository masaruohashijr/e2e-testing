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

func configuredToRejectCall(numberA string) error {
	p := &domains.Reject{}
	ResponseReject.Reject = *p
	x, _ := xml.MarshalIndent(p, "", "")
	println(string(x))
	return nil
}

func iMakeACallFromTo(numberA, numberB string) error {
	Configuration.From, _ = Configuration.SelectNumber(numberA)
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(numberB)
	x, _ := xml.MarshalIndent(ResponseReject, "", "")
	strXML := domains.Header + string(x)
	println(strXML)
	services.WriteActionXML("reject", strXML)
	CallPrimaryPort.MakeCall()
	return nil
}

func myTestSetupRuns() error {
	configurationSetup()
	println(Configuration.AccountSid)
	CallSecondaryPort = secondary.NewCallsApi(&Configuration)
	CallPrimaryPort = primary.NewCallsService(CallSecondaryPort)
	NumberSecondaryPort = secondary.NewNumbersApi(&Configuration)
	NumberPrimaryPort = primary.NewNumbersService(NumberSecondaryPort)
	return nil

}

func configurationSetup() {
	Configuration = config.NewConfig()
	go services.RunServer(Ch)
	Configuration.StatusCallback = services.BaseUrl + "/RejectCallBack"
	Configuration.VoiceUrl = services.BaseUrl + "/Reject"
}

func shouldGetCallCancelStatus() error {
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
	ctx.Step(`^"([^"]*)" configured to reject call$`, configuredToRejectCall)
	ctx.Step(`^I make a call from "([^"]*)" to "([^"]*)"$`, iMakeACallFromTo)
	ctx.Step(`^my test setup runs$`, myTestSetupRuns)
	ctx.Step(`^"([^"]*)" should get call cancel status$`, shouldGetCallCancelStatus)
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
}

func TestMain(m *testing.M) {
	opts := godog.Options{
		Format:    "progress",
		Paths:     []string{"../../features/reject"},
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
