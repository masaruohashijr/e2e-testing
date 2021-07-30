package dial

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
var ResponseDial domains.ResponseDial
var Ch = make(chan string)

func ConfiguredToDial(dialerNumber, dialedNumber string) error {
	number, _ := Configuration.SelectNumber(dialedNumber)
	d := &domains.Dial{
		Value:       number,
		CallBackURL: services.BaseUrl + "/Callback",
	}
	ResponseDial.Dial = *d
	p := &domains.Hangup{}
	ResponseDial.Hangup = *p
	x, _ := xml.MarshalIndent(ResponseDial, "", "")
	strXML := domains.Header + string(x)
	services.WriteActionXML("dial", strXML)
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(dialerNumber)
	Configuration.VoiceUrl = services.BaseUrl + "/Dial"
	NumbersPrimaryPort.UpdateNumber()
	println(string(x))
	return nil
}

func IMakeACallFromTo(numberA, numberB string) error {
	Configuration.From, Configuration.FromSid = Configuration.SelectNumber(numberA)
	CallsPrimaryPort.MakeCall()
	return nil
}

func MyTestSetupRuns() error {
	Configuration = config.NewConfig()
	go services.RunServer(Ch, true)
	CallsSecondaryPort = secondary.NewCallsApi(&Configuration)
	CallsPrimaryPort = primary.NewCallsService(CallsSecondaryPort)
	NumbersSecondaryPort = secondary.NewNumbersApi(&Configuration)
	NumbersPrimaryPort = primary.NewNumbersService(NumbersSecondaryPort)
	Configuration.ActionUrl = "http://zang.io/ivr/welcome/call"
	return nil
}

func ShouldGetTheIncomingCallFrom(dialedNumber, dialerNumber string) error {
	bodyContent := ""
	select {
	case bodyContent = <-Ch:
		fmt.Println(bodyContent)
	case <-time.After(time.Duration(services.TestTimeout) * time.Second):
		tt := strconv.FormatInt(services.TestTimeout, 10)
		return fmt.Errorf("Timeout %s.", tt)
	}
	url_parameters, e := url.ParseQuery(bodyContent)
	if e != nil {
		panic(e)
	}
	dialed, _ := Configuration.SelectNumber(dialedNumber)
	orig_dialed := dialed
	orig_dialer := Configuration.To
	if len(url_parameters["From"]) == 0 || len(url_parameters["To"]) == 0 {
		return fmt.Errorf("Expected dialed: %s and found %s.", orig_dialed, "none")
	}
	dialer_number := string(url_parameters["From"][0])
	dialed_number := string(url_parameters["To"][0])

	if dialed_number != orig_dialed {
		return fmt.Errorf("Expected dialed: %s and found %s.", orig_dialed, dialed_number)
	}
	if dialer_number != orig_dialer {
		return fmt.Errorf("Expected From: %s and found %s.", orig_dialer, dialer_number)
	}
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(dialerNumber)
	Configuration.VoiceUrl = ""
	NumbersPrimaryPort.UpdateNumber()
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^"([^"]*)" configured to dial "([^"]*)"$`, ConfiguredToDial)
	ctx.Step(`^I make a call from "([^"]*)" to "([^"]*)"$`, IMakeACallFromTo)
	ctx.Step(`^my test setup runs$`, MyTestSetupRuns)
	ctx.Step(`^"([^"]*)" should get the incoming call from "([^"]*)"$`, ShouldGetTheIncomingCallFrom)
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
}
