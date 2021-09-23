package primary

import (
	"zarbat_test/pkg/domains"
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

func (p *port_sms) ViewSMS(smsSid string) (domains.Sms, error) {
	sms, err := p.driven.ViewSMS(smsSid)
	return sms, err
}

func (p *port_sms) ListSMS(from, to string) ([]domains.Sms, error) {
	smss, err := p.driven.ListSMS(from, to)
	return smss, err
}

func (p *port_sms) SendSMS(from, to, message string) error {
	err := p.driven.SendSMS(from, to, message)
	return err
}
