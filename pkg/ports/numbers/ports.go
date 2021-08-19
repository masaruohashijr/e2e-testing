package numbers

import "zarbat_test/pkg/domains"

type PrimaryPort interface {
	AddNumber(string) error
	UpdateNumber() error
	ViewNumber(string) (*domains.IncomingPhoneNumber, error)
	DeleteNumber(string) error
	ListAvailableNumbers() ([]string, error)
	ListNumbers() ([]string, error)
}

type SecondaryPort interface {
	AddNumber(string) error
	UpdateNumber() error
	ViewNumber(string) (*domains.IncomingPhoneNumber, error)
	DeleteNumber(string) error
	ListAvailableNumbers() ([]string, error)
	ListNumbers() ([]string, error)
}
