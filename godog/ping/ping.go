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
var CallSecondaryPort calls.SecondaryPort
var CallPrimaryPort calls.PrimaryPort
var NumbersSecondaryPort numbers.SecondaryPort
var NumbersPrimaryPort numbers.PrimaryPort
var ResponseGather domains.ResponseGather
var ResponseRecord domains.ResponseRecord
var ResponsePing domains.ResponsePing
var Ch = make(chan string)

func IMakeACallFromTo(numberA, numberB string) error {
	Configuration.From, Configuration.FromSid = Configuration.SelectNumber(numberA) //"+5561984385415"
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(numberB)     //"+5561984385415"
	x, _ := xml.MarshalIndent(ResponsePing, "", "")
	strXML := domains.Header + string(x)
	services.WriteActionXML("ping", strXML)
	CallPrimaryPort.MakeCall()
	return nil
}

func MyTestSetupRuns() error {
	Configuration = config.NewConfig()
	go services.RunServer(Ch, false)
	Configuration.StatusCallback = services.BaseUrl + "/Callback"
	Configuration.ActionUrl = services.BaseUrl + "/Ping"
	CallSecondaryPort = secondary.NewCallsApi(&Configuration)
	CallPrimaryPort = primary.NewCallsService(CallSecondaryPort)
	NumbersSecondaryPort = secondary.NewNumbersApi(&Configuration)
	NumbersPrimaryPort = primary.NewNumbersService(NumbersSecondaryPort)
	Configuration.VoiceUrl = ""
	NumbersPrimaryPort.UpdateNumber()
	// instantiate the proper Response
	return nil
}

func ShouldBeAbleToListenToFrequencies(number, frequencies string) error {
	recordUrl := ""
	select {
	case recordUrl = <-Ch:
		fmt.Printf("Result: %s\n", recordUrl)
	case <-time.After(time.Duration(services.TestTimeout) * time.Second):
		fmt.Println("timeout")
		Ch = nil
		return fmt.Errorf("timeout")
	}
	return nil
}

func ConfiguredToPingURL(numberA string) error {
	p := &domains.Ping{
		Value: services.BaseUrl + "/Pinging",
	}
	ResponsePing.Ping = *p
	x, _ := xml.MarshalIndent(p, "", "")
	println(string(x))
	return nil
}

func ShouldGetAPingRequestOnTheURL() error {
	select {
	case res := <-Ch:
		fmt.Println(res)
	case <-time.After(time.Duration(services.TestTimeout) * time.Second):
		fmt.Println("timeout")
		Ch = nil
		return fmt.Errorf("timeout")
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
