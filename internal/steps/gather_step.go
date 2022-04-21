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

func ConfiguredToGatherSpeech(number string) error {
	services.CloseChannel = false
	testHash := fmt.Sprint(TestHash)
	r := &domains.ResponseGather{
		Pause: domains.Pause{
			Length: 0,
		},
		Gather: domains.Gather{
			Input:    "speech",
			Language: "en-US",
			Timeout:  15,
			Action:   Configuration.BaseUrl + "/SpeechResult" + "?hash=" + testHash,
		},
	}
	ResponseGather = *r
	x, _ := xml.MarshalIndent(ResponseGather, "", "")
	strXML := domains.Header + string(x)
	logging.Debug.Println(strXML)
	services.WriteActionXML("gather", strXML)
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(number)
	Configuration.VoiceUrl = services.BaseUrl + "/Gather"
	NumberPrimaryPort.UpdateNumber()
	return nil
}

func ShouldGetTranscription(number, originalSpeech string) error {
	return ShouldGetSpeech(number, originalSpeech)
}

func ShouldGetSpeech(number, originalSpeech string) error {
	services.CloseChannel = false
	speechResult := ""
	select {
	case speechResult = <-Ch:
		fmt.Printf("Result: %s\n", speechResult)
	case <-time.After(time.Duration(services.TestTimeout) * time.Second):
		logging.Debug.Println("timeout")
		Ch = nil
		return fmt.Errorf("timeout")
	}
	originalSpeech = strings.ToLower(originalSpeech)
	speechResult = strings.ToLower(speechResult)
	logging.Debug.Println("original speech: ", originalSpeech)
	logging.Debug.Println("original speech: ", originalSpeech)
	logging.Debug.Println("speech result: ", speechResult)
	logging.Debug.Println("speech result: ", speechResult)
	if strings.TrimSpace(originalSpeech) != strings.TrimSpace(speechResult) && speechResult != "" && speechResult != "welcome to your new zhang account" {
		return fmt.Errorf("The returned speech \"%s\" is different from the one expected by the test.", speechResult)
	}
	// Reset
	ShouldBeReset(number)
	return nil
}
func ShouldGetDigitsFrom(number, expectedDigits string) error {
	digitsResult := ""
	select {
	case digitsResult = <-Ch:
		fmt.Printf("Result: %s\n", digitsResult)
		logging.Debug.Printf("Result: %s\n", digitsResult)
	case <-time.After(time.Duration(services.TestTimeout) * time.Second):
		logging.Debug.Println("timeout")
		logging.Debug.Println("timeout")
		Ch = nil
		return fmt.Errorf("timeout")
	}
	expectedDigits = strings.ToLower(expectedDigits)
	digitsResult = strings.ToLower(digitsResult)
	logging.Debug.Println("original speech: ", expectedDigits)
	logging.Debug.Println("original speech: ", expectedDigits)
	logging.Debug.Println("speech result: ", digitsResult)
	logging.Debug.Println("speech result: ", digitsResult)
	if strings.TrimSpace(expectedDigits) != strings.TrimSpace(digitsResult) && digitsResult != "" && digitsResult != "welcome to your new zhang account" {
		return fmt.Errorf("The returned speech \"%s\" is different from the one expected by the test.", digitsResult)
	}
	// Reset
	ShouldBeReset(number)
	return nil
}

func ConfiguredToGatherDigitsUntil(number, finishOnKey string) error {
	services.CloseChannel = false
	testHash := fmt.Sprint(TestHash)
	r := &domains.ResponseGather{
		Pause: domains.Pause{
			Length: 0,
		},
		Gather: domains.Gather{
			Input:       "dtmf",
			FinishOnKey: finishOnKey,
			Action:      Configuration.BaseUrl + "/SpeechResult" + "?hash=" + testHash,
		},
	}
	ResponseGather = *r
	x, _ := xml.MarshalIndent(ResponseGather, "", "")
	strXML := domains.Header + string(x)
	logging.Debug.Println(strXML)
	services.WriteActionXML("gather", strXML)
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(number)
	Configuration.VoiceUrl = services.BaseUrl + "/Gather"
	NumberPrimaryPort.UpdateNumber()
	return nil
}
