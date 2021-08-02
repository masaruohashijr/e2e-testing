package steps

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"time"
	"zarbat_test/internal/godog/services"
	"zarbat_test/pkg/domains"
)

func ConfiguredToPlayTone(number, tone string) error {
	Configuration.From, Configuration.FromSid = Configuration.SelectNumber(number)
	p := &domains.Play{
		Value: "tone_stream://%(" + tone + ")",
		Loop:  services.PlayLoop,
	}
	ResponsePlay.Play = *p
	x, _ := xml.MarshalIndent(ResponsePlay, "", "")
	strXML := domains.Header + string(x)
	println(strXML)
	services.WriteActionXML("play", strXML)
	Configuration.ActionUrl = services.BaseUrl + "/Play"
	return nil
}

func ShouldBeAlaybleToListenToFrequencies(number, frequencies string) error {
	recordUrl := ""
	select {
	case recordUrl = <-Ch:
		fmt.Printf("Result: %s\n", recordUrl)
	case <-time.After(time.Duration(services.TestTimeout) * time.Second):
		fmt.Println("timeout")
		Ch = nil
		return fmt.Errorf("timeout")
	}
	time.Sleep(1 * time.Second)
	err := services.DownloadFile("media/record.wav", recordUrl)
	if err != nil {
		return fmt.Errorf("Error %s", "Not able to download the record.")
	}
	iFrequencies, _ := strconv.Atoi(frequencies)
	err = services.GetFrequencies("media/record.wav", iFrequencies, 90)
	if err != nil {
		return fmt.Errorf("Error %s", "Not able to listen correct frequencies.")
	}
	// Reset
	ShouldBeReset(number)
	return nil
}
