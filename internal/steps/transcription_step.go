package steps

import (
	"encoding/xml"
	"fmt"
	"zarbat_test/internal/godog/services"
	"zarbat_test/pkg/domains"
)

func IShouldListAtLeastTranscription(number int) error {
	notifications, err := NotificationPrimaryPort.ListNotifications()
	if err != nil {
		return fmt.Errorf("Could not list notifications.")
	}
	if len(notifications) < number {
		return fmt.Errorf("Expected %d notifications, but got %d.", number, len(notifications))
	}
	return nil
}
func IProvideAnAudioUrl(from, to string) error {
	calls, err := CallPrimaryPort.FilterCalls(from, to, "completed")
	if err != nil {
		return fmt.Errorf("Error %s", err.Error())
	}
	if len(calls) > 0 {
		recordings, err := RecordingPrimaryPort.ListRecordings(calls[0].Sid)
		if err != nil {
			return fmt.Errorf("Error %s", err.Error())
		}
		if len(recordings) == 0 {
			return fmt.Errorf("Error %s", err.Error())
		}
	}
	return nil
}

func IShouldGetTranscriptionTextAs(number string) error {
	testHash := fmt.Sprint(TestHash)
	r := &domains.Record{
		Background:         services.Background,
		MaxLength:          services.MaxLength,
		FileFormat:         services.FileFormat,
		Transcribe:         true,
		TranscribeCallback: services.BaseUrl + "/TranscribeCallback" + "?hash=" + testHash,
	}
	p := &domains.Pause{
		Length: 0,
	}
	ResponseRecord.Pause = *p
	ResponseRecord.Record = *r
	x, _ := xml.MarshalIndent(ResponseRecord, "", "")
	strXML := domains.Header + string(x)
	println(strXML)
	services.WriteActionXML("record", strXML)
	Configuration.VoiceUrl = services.BaseUrl + "/Record"
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(number)
	NumberPrimaryPort.UpdateNumber()
	return nil
}

func ITranscribeAudioUrl(number string) error {
	testHash := fmt.Sprint(TestHash)
	r := &domains.Record{
		Background: services.Background,
		MaxLength:  services.MaxLength,
		FileFormat: services.FileFormat,
		Transcribe: false,
		Action:     services.BaseUrl + "/RecordAction" + "?hash=" + testHash,
	}
	p := &domains.Pause{
		Length: 0,
	}
	ResponseRecord.Pause = *p
	ResponseRecord.Record = *r
	x, _ := xml.MarshalIndent(ResponseRecord, "", "")
	strXML := domains.Header + string(x)
	println(strXML)
	services.WriteActionXML("record", strXML)
	Configuration.VoiceUrl = services.BaseUrl + "/Record"
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(number)
	NumberPrimaryPort.UpdateNumber()
	return nil
}

func IShouldGetLastTranscriptionTextAs(expected string) error {
	transcriptions, err := TranscriptionPrimaryPort.ListTranscriptions()
	if err != nil {
		return fmt.Errorf("Could not list transcriptions.")
	}
	tr := transcriptions[0]
	txt := tr.TranscriptionText
	if expected != txt {
		return fmt.Errorf("Transcription text '%s' is different from the expected '%s'", txt, expected)
	}
	return nil
}

func ITranscribeLastRecording() error {
	recordings, err := RecordingPrimaryPort.ListRecordings(CallSid)
	if err != nil {
		return fmt.Errorf("Could not list the recordings from call: %s", CallSid)
	}
	if len(recordings) > 0 {
		recording, err := RecordingPrimaryPort.ViewRecording(recordings[0].Sid)
		if err != nil {
			return fmt.Errorf("Could not view recording from call: %s", CallSid)
		}
		TranscriptionPrimaryPort.TranscribeRecording(recording.Sid)
	} else {
		return fmt.Errorf("No recordings created for call: %s", CallSid)
	}
	return nil
}
