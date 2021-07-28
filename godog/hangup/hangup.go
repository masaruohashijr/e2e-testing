package hangup

import (
	"fmt"
	"net/url"
	"strconv"
	"zarbat_test/godog/services"
	"zarbat_test/internal/adapters/primary"
	"zarbat_test/internal/adapters/secondary"
	"zarbat_test/internal/config"
	"zarbat_test/pkg/domains"
	"zarbat_test/pkg/ports/calls"
	"zarbat_test/pkg/ports/numbers"

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
	p := &domains.Pause{
		Length: timeInSeconds,
	}
	ResponseHangup.Pause = *p
	h := &domains.Hangup{}
	ResponseHangup.Hangup = *h
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(NumberA)
	Configuration.VoiceUrl = services.BaseUrl + "/Hangup"
	NumbersSecondaryPort.UpdateNumber()
	return nil
}

func IMakeACallFromTo(NumberA, NumberB string) error {
	Configuration.From, Configuration.FromSid = Configuration.SelectNumber(NumberA)
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(NumberB)
	CallsPrimaryPort.MakeCall()
	return nil
}

func MyTestSetupRuns() error {
	Configuration = config.NewConfig()
	go services.RunServer(Ch, true)
	Configuration.StatusCallback = services.BaseUrl + "/Callback"
	CallsSecondaryPort = secondary.NewCallsApi(&Configuration)
	CallsPrimaryPort = primary.NewCallsService(CallsSecondaryPort)
	NumbersSecondaryPort = secondary.NewNumbersApi(&Configuration)
	NumbersPrimaryPort = primary.NewNumbersService(NumbersSecondaryPort)
	// instantiate the proper Response
	return nil
}

func ShouldGetLastCallDurationGreaterThanOrEqualsTo(number string, timeInSeconds int) error {
	bodyContent := <-Ch
	url_parameters, e := url.ParseQuery(bodyContent)
	if len(url_parameters["CallDuration"]) == 0 {
		return fmt.Errorf("The call duration should be more than the pause interval."+
			" Expected at least %d seconds but got %d seconds.", timeInSeconds, 0)
	}
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
	ctx.Step(`^"([^"]*)" configured to hangup after "([^"]*)" seconds$`, ConfiguredToHangupAfterSeconds)
	ctx.Step(`^"([^"]*)" should get last call duration greater than or equals to "([^"]*)"$`, ShouldGetLastCallDurationGreaterThanOrEqualsTo)
	ctx.Step(`^I make a call from "([^"]*)" to "([^"]*)"$`, IMakeACallFromTo)
	ctx.Step(`^my test setup runs$`, MyTestSetupRuns)
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
}
