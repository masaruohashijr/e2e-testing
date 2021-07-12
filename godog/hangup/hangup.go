package hangup

import (
	"e2e-testing/godog/services"
	"e2e-testing/internal/adapters/primary"
	"e2e-testing/internal/adapters/secondary"
	"e2e-testing/internal/config"
	"e2e-testing/pkg/domains"
	"e2e-testing/pkg/ports/calls"
	"e2e-testing/pkg/ports/numbers"
	"encoding/xml"
	"fmt"
	"net/url"
	"strconv"

	"github.com/cucumber/godog"
)

var Configuration config.ConfigType
var CallsSecondaryPort calls.SecondaryPort
var CallsPrimaryPort calls.PrimaryPort
var ResponseHangup domains.ResponseHangup
var NumbersSecondaryPort numbers.SecondaryPort
var NumbersPrimaryPort numbers.PrimaryPort
var Ch = make(chan string)

func ConfiguredToHangupAfterSeconds(NumberA string, timeInSeconds int) error {
	Configuration.From, _ = Configuration.SelectNumber(NumberA)
	p := &domains.Pause{
		Length: timeInSeconds,
	}
	ResponseHangup.Pause = *p
	h := &domains.Hangup{}
	ResponseHangup.Hangup = *h
	return nil

}

func IMakeACallFromTo(NumberA, NumberB string) error {
	x, _ := xml.MarshalIndent(ResponseHangup, "", "")
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(NumberB)
	strXML := domains.Header + string(x)
	println(strXML)
	services.WriteActionXML("hangup", strXML)
	CallsPrimaryPort.MakeCall()
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^"([^"]*)" configured to hangup after "([^"]*)" seconds$`, ConfiguredToHangupAfterSeconds)
	ctx.Step(`^"([^"]*)" should get last call duration greater than or equals to "([^"]*)"$`, ShouldGetLastCallDurationGreaterThanOrEqualsTo)
	ctx.Step(`^I make a call from "([^"]*)" to "([^"]*)"$`, IMakeACallFromTo)
	ctx.Step(`^my test setup runs$`, MyTestSetupRuns)
}

func MyTestSetupRuns() error {
	Configuration = config.NewConfig()
	go services.RunServer(Ch)
	Configuration.StatusCallback = services.BaseUrl + "/Callback"
	Configuration.ActionUrl = services.BaseUrl + "/Hangup"
	CallsSecondaryPort = secondary.NewCallsApi(&Configuration)
	CallsPrimaryPort = primary.NewCallsService(CallsSecondaryPort)
	NumbersSecondaryPort = secondary.NewNumbersApi(&Configuration)
	NumbersPrimaryPort = primary.NewNumbersService(NumbersSecondaryPort)
	return nil
}

func ShouldGetLastCallDurationGreaterThanOrEqualsTo(number string, timeInSeconds int) error {
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
	Configuration.VoiceUrl = ""
	NumbersPrimaryPort.UpdateNumber()
	return nil
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
}
