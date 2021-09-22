package steps

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"
	"zarbat_test/internal/godog/services"
	"zarbat_test/internal/logging"
	"zarbat_test/pkg/domains"
)

func IMakeACallFromTo(numberA, numberB string) error {
	Configuration.From, Configuration.FromSid = Configuration.SelectNumber(numberA)
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(numberB)
	Configuration.Timeout = services.Timeout
	checkSayGather()
	checkUrlEmpty()
	CallPrimaryPort.MakeCall()
	return nil
}

func checkUrlEmpty() {
	if Configuration.ActionUrl == "" {
		Configuration.ActionUrl = "http://zang.io/ivr/welcome/call"
	}
}

func checkSayGather() {
	if strings.HasSuffix(Configuration.VoiceUrl, "/Gather") && strings.HasSuffix(Configuration.ActionUrl, "/Say") {
		ResponseSay.Say.Loop = 2
		x, _ := xml.MarshalIndent(ResponseSay, "", "")
		strXML := domains.Header + string(x)
		logging.Debug.Println(strXML)
		services.WriteActionXML("say", strXML)
		Configuration.ActionUrl = services.BaseUrl + "/Say"
	}
}

func ListCallsAfterSeconds(timeInSeconds int) error {
	time.Sleep(time.Duration(timeInSeconds) * time.Second)
	CallPrimaryPort.ListCalls()
	return nil
}

func AfterWaitingForSeconds(timeInSeconds int) error {
	time.Sleep(time.Duration(timeInSeconds) * time.Second)
	return nil
}

func IShouldGetLastCallDurationGreaterThanOrEqualToSeconds(timeInSeconds int) error {
	calls, err := CallPrimaryPort.ListCalls()
	if err != nil {
		return fmt.Errorf("Error: An error has occured within list calls.")
	}
	call, err := CallPrimaryPort.ViewCall(calls[0].Sid)
	if err != nil {
		return fmt.Errorf("Error: An error has occured within view call.")
	}
	if call.Duration < timeInSeconds {
		return fmt.Errorf("Error: Expected call duration >= %d and got %d.", timeInSeconds, calls[0].Duration)
	}
	return nil
}

func IShouldListAtLeastCall(numberOfCalls int) error {
	calls, err := CallPrimaryPort.ListCalls()
	println("Number of calls: ", len(calls))
	if err != nil {
		return fmt.Errorf("Error: An error has occured within list calls.")
	}
	if len(calls) <= numberOfCalls {
		return fmt.Errorf("The returned number of calls \"%d\" is different from the one expected by the test.", len(calls))
	}
	return nil
}
