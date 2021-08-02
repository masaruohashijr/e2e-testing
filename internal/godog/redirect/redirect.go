package redirect

import (
	"encoding/xml"
	"fmt"
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
var ResponseRedirect domains.ResponseRedirect
var Ch = make(chan string)

func ConfiguredToRedirectToPingURL(numberB string) error {
	r := &domains.Redirect{
		Value: services.BaseUrl + "/Pinging",
	}
	ResponseRedirect.Redirect = *r
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(numberB)
	Configuration.VoiceUrl = services.BaseUrl + "/Redirect"
	NumbersSecondaryPort.UpdateNumber()
	return nil
}

func IMakeACallFromTo(numberA, numberB string) error {
	Configuration.From, Configuration.FromSid = Configuration.SelectNumber(numberA)
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(numberB)
	Configuration.StatusCallback = services.BaseUrl + "/Callback"
	x, _ := xml.MarshalIndent(ResponseRedirect, "", "")
	strXML := domains.Header + string(x)
	println(strXML)
	services.WriteActionXML("redirect", strXML)
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
	return nil
}

func ShouldGetAPingRequestOnTheURL() error {
	select {
	case res := <-Ch:
		fmt.Printf("Result: %s\n", res)
	case <-time.After(time.Duration(services.TestTimeout) * time.Second):
		fmt.Println("timeout")
		Ch = nil
		return fmt.Errorf("timeout")
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^my test setup runs$`, MyTestSetupRuns)
	ctx.Step(`^"([^"]*)" configured to redirect to ping URL$`, ConfiguredToRedirectToPingURL)
	ctx.Step(`^I make a call from "([^"]*)" to "([^"]*)"$`, IMakeACallFromTo)
	ctx.Step(`^Should get a ping request on the URL$`, ShouldGetAPingRequestOnTheURL)
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
}
