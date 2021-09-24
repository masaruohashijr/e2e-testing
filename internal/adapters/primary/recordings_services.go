package primary

import (
	"zarbat_test/pkg/domains"
	"zarbat_test/pkg/ports/recordings"
)

type port_recordings struct {
	driven recordings.SecondaryPort
}

func NewRecordingsService(driven recordings.SecondaryPort) recordings.PrimaryPort {
	return &port_recordings{
		driven,
	}
}

func (p *port_recordings) ViewRecording(recordingSid string) (domains.Recording, error) {
	recording, err := p.driven.ViewRecording(recordingSid)
	return recording, err
}

func (p *port_recordings) DeleteRecording(recordingSid string) (domains.Recording, error) {
	recording, err := p.driven.DeleteRecording(recordingSid)
	return recording, err
}

func (p *port_recordings) ListRecordings(callSid string) ([]domains.Recording, error) {
	recordings, err := p.driven.ListRecordings(callSid)
	return recordings, err
}

func (p *port_recordings) RecordCall(callSid string, timeInSeconds int) error {
	err := p.driven.RecordCall(callSid, timeInSeconds)
	return err
}
