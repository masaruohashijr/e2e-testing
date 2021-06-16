package main

import (
	"e2e-testing/godog/general"
	"e2e-testing/internal/adapters/primary"
	"e2e-testing/internal/adapters/secondary"
	"e2e-testing/internal/config"
	d "e2e-testing/pkg/domains"
	"encoding/xml"
	"log"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/cucumber/godog"
)

func appendToConfigHangup(numberB string) error {
	h := &d.Hangup{}
	x, _ := xml.MarshalIndent(h, "", "")
	println(string(x))
	return nil
}

func configuredToPauseSeconds(numberB string, timeInSeconds int) error {
	p := &d.Pause{
		Length: timeInSeconds,
	}
	ResponsePause.Pause = *p
	x, _ := xml.MarshalIndent(p, "", "")
	println(string(x))
	return nil
}

func iMakeACallFromTo(numberA, numberB string) error {
	x, _ := xml.MarshalIndent(ResponsePause, "", "")
	strXML := d.Header + string(x)
	println(strXML)
	writeActionXML(strXML)
	PrimaryPort.MakeCall()
	return nil
}

func myTestSetupRuns() error {
	Configuration = config.NewConfig()
	go general.RunServer(Ch)
	Configuration.From = "+558140423562" //+558140421695
	Configuration.To = "+5561984385415"
	Configuration.StatusCallback = "https://918a4971ed21.ngrok.io/Callback"
	Configuration.ActionUrl = "https://918a4971ed21.ngrok.io/InboundXml"
	println(Configuration.AccountSid)
	SecondaryPort = secondary.NewCallsApi(&Configuration)
	PrimaryPort = primary.NewService(SecondaryPort)
	// instantiate the proper Response

	return nil
}

func shouldGetLastCallDurationMoreThanOrEqualsTo(number string, timeInSeconds int) error {
	<-Ch
	println("GOT IT")
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

func selectNumber(option string) string {
	switch option {
	case "NumberA":
		return Configuration.NumberA
	case "NumberB":
		return Configuration.NumberB
	}
	return ""
}

func writeActionXML(strXML string) {
	f, err := os.Create("../../xml/inbound.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	strXML = strings.Replace(strXML, "<Hangup></Hangup>", "<Hangup/>", 1)
	_, err2 := f.WriteString(strXML)
	if err2 != nil {
		log.Fatal(err2)
	}
}
