package pause

import (
	"encoding/xml"
	"fmt"
	"net/url"
	"strconv"
	"time"
	"zarbat_test/internal/adapters/primary"
	"zarbat_test/internal/adapters/secondary"
	"zarbat_test/internal/config"
	"zarbat_test/internal/godog/services"
	"zarbat_test/pkg/domains"
	"zarbat_test/pkg/ports/calls"
	"zarbat_test/pkg/ports/numbers"

	"github.com/cucumber/godog"
)

var Configuration config.ConfigType
var CallsSecondaryPort calls.SecondaryPort
var CallsPrimaryPort calls.PrimaryPort
var NumbersSecondaryPort numbers.SecondaryPort
var NumbersPrimaryPort numbers.PrimaryPort
var ResponsePause domains.ResponsePause
var Ch = make(chan string)
var Finish = true

func AppendToConfigHangup(numberA string) error {
	h := &domains.Hangup{}
	ResponsePause.Hangup = *h
	return nil
}

func ConfiguredToPauseSeconds(numberB string, timeInSeconds int) error {
	p := &domains.Pause{
		Length: timeInSeconds,
	}
	ResponsePause.Pause = *p
	return nil
}

func IMakeACallFromTo(numberA, numberB string) error {
	Configuration.From, Configuration.FromSid = Configuration.SelectNumber(numberA)
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(numberB)
	Configuration.ActionUrl = services.BaseUrl + "/Pause"
	x, _ := xml.MarshalIndent(ResponsePause, "", "")
	strXML := domains.Header + string(x)
	services.WriteActionXML("pause", strXML)
	CallsPrimaryPort.MakeCall()
	return nil
}

func MyTestSetupRuns() error {
	Configuration = config.NewConfig()
	go services.RunServer(Ch, true)
	Configuration.StatusCallback = services.BaseUrl + "/Callback"
	Configuration.ActionUrl = services.BaseUrl + "/Pause"
	CallsSecondaryPort = secondary.NewCallsApi(&Configuration)
	CallsPrimaryPort = primary.NewCallsService(CallsSecondaryPort)
	NumbersSecondaryPort = secondary.NewNumbersApi(&Configuration)
	NumbersPrimaryPort = primary.NewNumbersService(NumbersSecondaryPort)
	Configuration.VoiceUrl = ""
	NumbersPrimaryPort.UpdateNumber()
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
	if url_parameters["CallDuration"] == nil {
		Ch = make(chan string)
		go services.RunServer(Ch, true)
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
