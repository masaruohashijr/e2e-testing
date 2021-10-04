package steps

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"zarbat_test/internal/godog/services"
	"zarbat_test/pkg/domains"
)

func IRecordCurrentCallFromToForSeconds(from, to string, timeInSeconds int) error {
	Configuration.From, Configuration.FromSid = Configuration.SelectNumber(from)
	Configuration.To, Configuration.ToSid = Configuration.SelectNumber(to)
	CallPrimaryPort.MakeCall()
	calls, err := CallPrimaryPort.FilterCalls(Configuration.From, Configuration.To, "in-progress")
	if err != nil {
		return fmt.Errorf("Error %s", err.Error())
	}
	if len(calls) > 0 {
		CallSid = calls[0].Sid
		RecordingPrimaryPort.RecordCall(CallSid, timeInSeconds)
	} else {
		return fmt.Errorf("There is no in-progress call.")
	}
	return nil
}
func IShouldListAtLeastRecordingFromTo(from, to string) error {
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

func ConfiguredToRecordCalls(number string) error {
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

func ConfiguredToRecordCallsForDownload(number string) error {
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

func IDeleteAllRecordingsFromTo() error {
	list, err := RecordingPrimaryPort.ListRecordings(CallSid)
	if err != nil {
		return fmt.Errorf("Could not list recordings from a call.")
	}
	for _, rec := range list {
		RecordingPrimaryPort.DeleteRecording(rec.Sid)
	}
	return nil
}

func IShouldListNoRecordingFromTo() error {
	list, err := RecordingPrimaryPort.ListRecordings(CallSid)
	if err != nil {
		return fmt.Errorf("Could not list recordings from a call.")
	}
	if len(list) > 0 {
		return fmt.Errorf("Could not list recordings from a call.")
	}
	return nil
}

func IShouldGetLastRecordingDurationGreaterThanOrEqualToSeconds(timeInSeconds int) error {
	list, err := RecordingPrimaryPort.ListRecordings(CallSid)
	rec, err := RecordingPrimaryPort.ViewRecording(list[0].Sid)
	if err != nil {
		return fmt.Errorf("Could not list recordings from a call.")
	}
	duration, _ := strconv.Atoi(rec.Duration)
	if duration < timeInSeconds {
		return fmt.Errorf("Could not list recordings from a call.")
	}
	return nil
}
