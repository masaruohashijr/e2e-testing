package primary

import (
	"zarbat_test/pkg/domains"
	"zarbat_test/pkg/ports/mms"
)

type port_mms struct {
	driven mms.SecondaryPort
}

func NewMmsService(driven mms.SecondaryPort) mms.PrimaryPort {
	return &port_mms{
		driven,
	}
}

func (p *port_mms) ViewMMS(mmsSid string) (domains.Mms, error) {
	mms, err := p.driven.ViewMMS(mmsSid)
	return mms, err
}

func (p *port_mms) ListMMS(from, to string) ([]domains.Mms, error) {
	mmss, err := p.driven.ListMMS(from, to)
	return mmss, err
}

func (p *port_mms) SendMMS(from, to, message string) error {
	err := p.driven.SendMMS(from, to, message)
	return err
}
