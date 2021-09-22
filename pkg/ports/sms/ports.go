package sms

import "zarbat_test/pkg/domains"

type PrimaryPort interface {
	SendSMS(from, to, message string) error
	ViewSMS(smsSid string) (domains.Sms, error)
	ListSMS(from, to string) ([]domains.Sms, error)
}

type SecondaryPort interface {
	SendSMS(from, to, message string) error
	ViewSMS(smsSid string) (domains.Sms, error)
	ListSMS(from, to string) ([]domains.Sms, error)
}
