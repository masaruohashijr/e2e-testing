package reject

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
var NumberSecondaryPort numbers.SecondaryPort
var NumberPrimaryPort numbers.PrimaryPort
var ResponseReject domains.ResponseReject

var Ch = make(chan string)

func ConfiguredToRejectCall(numberA string) error {
	p := &domains.Reject{}
	ResponseReject.Reject = *p
	x, _ := xml.MarshalIndent(p, "", "")
	println(string(x))
	return nil
}

func IMakeACallFromTo(numberA, numberB string) error {
	Configuration.From, Configuration.FromSid = Configuration.SelectNumber(numberA)
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(numberB)
	Configuration.VoiceUrl = ""
	NumberSecondaryPort.UpdateNumber()
	x, _ := xml.MarshalIndent(ResponseReject, "", "")
	strXML := domains.Header + string(x)
	println(strXML)
	services.WriteActionXML("reject", strXML)
	CallPrimaryPort.MakeCall()
	return nil
}

func MyTestSetupRuns() error {
	ConfigurationSetup()
	println(Configuration.AccountSid)
	CallSecondaryPort = secondary.NewCallsApi(&Configuration)
	CallPrimaryPort = primary.NewCallsService(CallSecondaryPort)
	NumberSecondaryPort = secondary.NewNumbersApi(&Configuration)
	NumberPrimaryPort = primary.NewNumbersService(NumberSecondaryPort)
	return nil
}

func ConfigurationSetup() {
	Configuration = config.NewConfig()
	go services.RunServer(Ch, false)
	Configuration.StatusCallback = services.BaseUrl + "/RejectCallBack"
	Configuration.VoiceUrl = services.BaseUrl + "/Reject"
}

func ShouldGetCallCancelStatus() error {
	println("Timer has started.")
	select {
	case res := <-Ch:
		fmt.Println(res)
	case <-time.After(time.Duration(services.TestTimeout) * time.Second):
		fmt.Println("timeout")
		return fmt.Errorf("timeout")
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^"([^"]*)" configured to reject call$`, ConfiguredToRejectCall)
	ctx.Step(`^I make a call from "([^"]*)" to "([^"]*)"$`, IMakeACallFromTo)
	ctx.Step(`^my test setup runs$`, MyTestSetupRuns)
	ctx.Step(`^"([^"]*)" should get call cancel status$`, ShouldGetCallCancelStatus)
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
}
