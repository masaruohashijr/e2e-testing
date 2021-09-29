package recordings

import "zarbat_test/pkg/domains"

type PrimaryPort interface {
	RecordCall(callSid string, timeInSeconds int) error
	DeleteRecording(recordingSid string) (domains.Recording, error)
	ViewRecording(recordingSid string) (domains.Recording, error)
	ListRecordings(callSid string) ([]domains.Recording, error)
}

type SecondaryPort interface {
	RecordCall(callSid string, timeInSeconds int) error
	DeleteRecording(recordingSid string) (domains.Recording, error)
	ViewRecording(recordingSid string) (domains.Recording, error)
	ListRecordings(callSid string) ([]domains.Recording, error)
}
