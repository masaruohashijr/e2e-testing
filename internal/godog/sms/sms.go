package sms

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/url"
	"time"
	"zarbat_test/internal/adapters/primary"
	"zarbat_test/internal/adapters/secondary"
	"zarbat_test/internal/config"
	"zarbat_test/internal/godog/services"
	"zarbat_test/pkg/domains"
	"zarbat_test/pkg/ports/calls"
	"zarbat_test/pkg/ports/sms"

	"github.com/cucumber/godog"
)

var Configuration config.ConfigType
var CallSecondaryPort calls.SecondaryPort
var CallPrimaryPort calls.PrimaryPort
var SmsSecondaryPort sms.SecondaryPort
var SmsPrimaryPort sms.PrimaryPort
var ResponseSMS domains.ResponseSMS
var Ch = make(chan string)

func SendsSMSTo(numberA, message, numberB string) error {
	Configuration.From, _ = Configuration.SelectNumber(numberA)
	Configuration.To, _ = Configuration.SelectNumber(numberB)
	SmsSecondaryPort.SendSMS(Configuration.To, Configuration.From, message)
	return nil
}

func ConfiguredToSendSmsTo(numberA, numberB string) error {
	Configuration.From, _ = Configuration.SelectNumber(numberA)
	Configuration.To, _ = Configuration.SelectNumber(numberB)
	s := &domains.Sms{
		Value:          "SMS Test",
		To:             Configuration.To,
		From:           Configuration.From,
		StatusCallback: services.BaseUrl + "/SmsStatus",
	}
	ResponseSMS.Sms = *s
	x, _ := xml.MarshalIndent(s, "", "")
	log.Println(string(x))
	return nil
}

func IMakeACallFromTo(numberA, numberB string) error {
	Configuration.From, _ = Configuration.SelectNumber(numberA)
	Configuration.To, _ = Configuration.SelectNumber(numberB)
	Configuration.VoiceUrl = ""
	x, _ := xml.MarshalIndent(ResponseSMS, "", "")
	strXML := domains.Header + string(x)
	println(strXML)
	services.WriteActionXML("sms", strXML)
	CallPrimaryPort.MakeCall()
	return nil
}

func MyTestSetupRuns() error {
	Configuration = config.NewConfig()
	go services.RunServer(Ch, false)
	Configuration.ActionUrl = services.BaseUrl + "/sms"
	CallSecondaryPort = secondary.NewCallsApi(&Configuration)
	CallPrimaryPort = primary.NewCallsService(CallSecondaryPort)
	SmsSecondaryPort = secondary.NewSmsApi(&Configuration)
	SmsPrimaryPort = primary.NewSmsService(SmsSecondaryPort)
	return nil
}

func SMSStatusShouldBeSentToCallStatusURL() error {
	bodyContent := ""
	select {
	case bodyContent = <-Ch:
		fmt.Printf("Result: %s\n", bodyContent)
	case <-time.After(time.Duration(services.TestTimeout) * time.Second):
		fmt.Println("timeout")
		Ch = nil
		return fmt.Errorf("timeout")
	}
	url_parameters, _ := url.ParseQuery(bodyContent)
	println(url_parameters)
	status := url_parameters["DlrStatus"][0]
	from_number := url_parameters["From"][0]
	to_number := url_parameters["To"][0]
	orig_from := Configuration.From
	orig_to := Configuration.To
	if status != "sent" {
		return fmt.Errorf("SMS not sent")
	}
	if orig_from != from_number {
		return fmt.Errorf("Expected From: %s and found %s.", orig_from, from_number)
	}
	if orig_to != to_number {
		return fmt.Errorf("Expected To: %s and found %s.", orig_to, orig_to)
	}

	return nil
}
func ShouldBeAbleToViewTheSMS(number, message string) error {
	bodyContent := ""
	select {
	case bodyContent = <-Ch:
		fmt.Printf("Result: %s\n", bodyContent)
	case <-time.After(time.Duration(services.TestTimeout) * time.Second):
		fmt.Println("timeout")
		Ch = nil
		return fmt.Errorf("timeout")
	}

	url_parameters, _ := url.ParseQuery(bodyContent)
	println(url_parameters)

	status := url_parameters["DlrStatus"][0]
	from_number := url_parameters["From"][0]
	to_number := url_parameters["To"][0]
	body := url_parameters["Body"][0]
	orig_from := Configuration.From
	orig_to := Configuration.To
	if status != "sent" {
		return fmt.Errorf("SMS not sent")
	}
	if orig_from != from_number {
		return fmt.Errorf("Expected From: %s and found %s.", orig_from, from_number)
	}
	if orig_to != to_number {
		return fmt.Errorf("Expected To: %s and found %s.", orig_to, orig_to)
	}
	if body != message {
		return fmt.Errorf("Expected Message: %s and found %s.", message, body)
	}

	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^my test setup runs$`, MyTestSetupRuns)
	ctx.Step(`^"([^"]*)" sends SMS "([^"]*)" to "([^"]*)"$`, SendsSMSTo)
	ctx.Step(`^"([^"]*)" should be able to view the SMS "([^"]*)"$`, ShouldBeAbleToViewTheSMS)
	// --------------------------
	ctx.Step(`^I make a call from "([^"]*)" to "([^"]*)"$`, IMakeACallFromTo)
	ctx.Step(`^SMS Status should be sent to call Status URL$`, SMSStatusShouldBeSentToCallStatusURL)
	ctx.Step(`^"([^"]*)" configured to send sms to "([^"]*)"$`, ConfiguredToSendSmsTo)
}
