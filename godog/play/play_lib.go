package main

import (
	"e2e-testing/godog/services"
	"e2e-testing/pkg/domains"
	"encoding/xml"
	"fmt"
	"time"
)

func ShouldBeAbleToListen(arg1 string) error {
	speechResult := ""
	select {
	case speechResult = <-Ch:
		fmt.Printf("Result: %s\n", speechResult)
	case <-time.After(time.Duration(services.TestTimeout) * time.Second):
		Ch = nil
		return fmt.Errorf("timeout")
	}
	b := services.CompareTwoSentences(speechResult, "how are you this is a test the test has ended okay", 80)
	println(b)
	if !b {
		return fmt.Errorf("Error %s", "Not able to listen.")
	}
	return nil
}

func ConfiguredToGatherSpeech(number string) error {
	Configuration.To = "+5561984385415"
	// Configuration.To, Configuration.ToSid = Configuration.SelectNumber(number) //""
	p := &domains.Pause{
		Length: services.GatherPause,
	}
	g := &domains.Gather{
		Input:    "speech",
		Language: "en-US",
		Timeout:  services.GatherTimeOut,
		Action:   services.BaseUrl + "/SpeechResult",
	}
	ResponseGather.Pause = *p
	ResponseGather.Gather = *g
	x, _ := xml.MarshalIndent(ResponseGather, "", "")
	strXML := domains.Header + string(x)
	println(strXML)
	services.WriteActionXML("gather", strXML)
	// NumberPrimaryPort.UpdateNumber()
	return nil
}

func ConfiguredToPlay(number, mp3File string) error {
	Configuration.From, _ = Configuration.SelectNumber(number)
	a := &domains.Pause{
		Length: services.PlayPause,
	}
	p := &domains.Play{
		Value: services.BaseUrl + "/mp3/" + mp3File + ".mp3",
		Loop:  services.PlayLoop,
	}
	ResponsePlay.Pause = *a
	ResponsePlay.Play = *p
	x, _ := xml.MarshalIndent(p, "", "")
	println(string(x))
	return nil
}
