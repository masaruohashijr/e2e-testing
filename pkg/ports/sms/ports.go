package sms

type PrimaryPort interface {
	SendSMS(to, from, message string) error
}

type SecondaryPort interface {
	SendSMS(to, from, message string) error
}
