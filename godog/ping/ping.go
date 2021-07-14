package ping

import (
	"encoding/xml"
	"fmt"
	"time"
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
var ResponsePing domains.ResponsePing
var CallsSecondaryPort calls.SecondaryPort
var CallsPrimaryPort calls.PrimaryPort
var NumbersSecondaryPort numbers.SecondaryPort
var NumbersPrimaryPort numbers.PrimaryPort
var Ch = make(chan string)

func ConfiguredToPingURL(numberA string) error {
	p := &domains.Ping{
		Value: services.BaseUrl + "/Pinging",
	}
	ResponsePing.Ping = *p
	x, _ := xml.MarshalIndent(p, "", "")
	println(string(x))
	return nil
}

func IMakeACallFromTo(numberA, numberB string) error {
	Configuration.From, Configuration.FromSid = Configuration.SelectNumber(numberA)
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(numberB)
	Configuration.VoiceUrl = ""
	NumbersSecondaryPort.UpdateNumber()
	x, _ := xml.MarshalIndent(ResponsePing, "", "")
	strXML := domains.Header + string(x)
	//println(strXML)
	services.WriteActionXML("ping", strXML)
	CallsPrimaryPort.MakeCall()
	return nil
}

func MyTestSetupRuns() error {
	Configuration = config.NewConfig()
	go services.RunServer(Ch)
	Configuration.ActionUrl = services.BaseUrl + "/Ping"
	//println(Configuration.AccountSid)
	CallsSecondaryPort = secondary.NewCallsApi(&Configuration)
	CallsPrimaryPort = primary.NewCallsService(CallsSecondaryPort)
	NumbersSecondaryPort = secondary.NewNumbersApi(&Configuration)
	NumbersPrimaryPort = primary.NewNumbersService(NumbersSecondaryPort)
	// instantiate the proper Response
	return nil
}

func ShouldGetAPingRequestOnTheURL() error {
	select {
	case res := <-Ch:
		fmt.Println(res)
	case <-time.After(time.Duration(services.TestTimeout) * time.Second):
		fmt.Println("timeout 60")
		Ch = nil
		return fmt.Errorf("timeout 60")
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^"([^"]*)" configured to ping URL$`, ConfiguredToPingURL)
	ctx.Step(`^I make a call from "([^"]*)" to "([^"]*)"$`, IMakeACallFromTo)
	ctx.Step(`^my test setup runs$`, MyTestSetupRuns)
	ctx.Step(`^should get a ping request on the URL$`, ShouldGetAPingRequestOnTheURL)
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {

}
