package pause

import (
	"e2e-testing/godog/services"
	"e2e-testing/internal/adapters/primary"
	"e2e-testing/internal/adapters/secondary"
	"e2e-testing/internal/config"
	"e2e-testing/pkg/domains"
	d "e2e-testing/pkg/domains"
	"e2e-testing/pkg/ports/calls"
	"e2e-testing/pkg/ports/numbers"
	"encoding/xml"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/cucumber/godog"
)

var Configuration config.ConfigType
var CallsSecondaryPort calls.SecondaryPort
var CallsPrimaryPort calls.PrimaryPort
var NumbersSecondaryPort numbers.SecondaryPort
var NumbersPrimaryPort numbers.PrimaryPort
var ResponsePause d.ResponsePause
var Ch = make(chan string)

func AppendToConfigHangup(numberA string) error {
	h := &domains.Hangup{}
	ResponsePause.Hangup = *h
	//x, _ := xml.MarshalIndent(h, "", "")
	//println(string(x))
	return nil
}

func ConfiguredToPauseSeconds(numberB string, timeInSeconds int) error {
	p := &domains.Pause{
		Length: timeInSeconds,
	}
	ResponsePause.Pause = *p
	//x, _ := xml.MarshalIndent(p, "", "")
	//println(string(x))
	return nil
}

func IMakeACallFromTo(numberA, numberB string) error {
	Configuration.From, Configuration.FromSid = Configuration.SelectNumber(numberA)
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(numberB)
	Configuration.VoiceUrl = ""
	Configuration.ActionUrl = services.BaseUrl + "/Pause"
	NumbersPrimaryPort.UpdateNumber()
	x, _ := xml.MarshalIndent(ResponsePause, "", "")
	strXML := domains.Header + string(x)
	//println(strXML)
	services.WriteActionXML("pause", strXML)
	CallsPrimaryPort.MakeCall()
	return nil
}

func MyTestSetupRuns() error {
	Configuration = config.NewConfig()
	go services.RunServer(Ch)
	Configuration.StatusCallback = services.BaseUrl + "/Callback"
	Configuration.ActionUrl = services.BaseUrl + "/Pause"
	CallsSecondaryPort = secondary.NewCallsApi(&Configuration)
	CallsPrimaryPort = primary.NewCallsService(CallsSecondaryPort)
	NumbersSecondaryPort = secondary.NewNumbersApi(&Configuration)
	NumbersPrimaryPort = primary.NewNumbersService(NumbersSecondaryPort)
	// instantiate the proper Response
	return nil
}

func ShouldGetLastCallDurationMoreThanOrEqualsTo(number string, timeInSeconds int) error {
	bodyContent := ""
	select {
	case bodyContent = <-Ch:
		fmt.Printf("Result: %s\n", bodyContent)
	case <-time.After(time.Duration(services.TestTimeout) * time.Second):
		Ch = nil
		return fmt.Errorf("timeout")
	}
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
	ctx.Step(`^my test setup runs$`, MyTestSetupRuns)
	ctx.Step(`^"([^"]*)" configured to pause (\d+) seconds$`, ConfiguredToPauseSeconds)
	ctx.Step(`^append To "([^"]*)" config hangup$`, AppendToConfigHangup)
	ctx.Step(`^I make a call from "([^"]*)" to "([^"]*)"$`, IMakeACallFromTo)
	ctx.Step(`^"([^"]*)" should get last call duration more than or equals to (\d+)$`, ShouldGetLastCallDurationMoreThanOrEqualsTo)
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
}
