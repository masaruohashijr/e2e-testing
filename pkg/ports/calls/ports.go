package calls

import "zarbat_test/pkg/domains"

type PrimaryPort interface {
	MakeCall() error
	ListCalls() ([]domains.Call, error)
	FilterCalls(from, to, status string) ([]domains.Call, error)
	ViewCall(callSid string) (domains.Call, error)
}

type SecondaryPort interface {
	MakeCall() error
	ListCalls() ([]domains.Call, error)
	FilterCalls(from, to, status string) ([]domains.Call, error)
	ViewCall(callSid string) (domains.Call, error)
}
