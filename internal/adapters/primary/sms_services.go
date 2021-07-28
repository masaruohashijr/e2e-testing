package primary

import (
	"zarbat_test/pkg/ports/sms"
)

type port_sms struct {
	driven sms.SecondaryPort
}

func NewSmsService(driven sms.SecondaryPort) sms.PrimaryPort {
	return &port_sms{
		driven,
	}
}

func (p *port_sms) SendSMS(to, from, message string) error {
	err := p.driven.SendSMS(to, from, message)
	return err
}
