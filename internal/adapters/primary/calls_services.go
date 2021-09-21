package primary

import (
	"zarbat_test/pkg/domains"
	"zarbat_test/pkg/ports/calls"
)

type port_calls struct {
	driven calls.SecondaryPort
}

func NewCallsService(driven calls.SecondaryPort) calls.PrimaryPort {
	return &port_calls{
		driven,
	}
}

func (p *port_calls) MakeCall() error {
	err := p.driven.MakeCall()
	return err
}

func (p *port_calls) ListCalls() ([]domains.Call, error) {
	call, err := p.driven.ListCalls()
	return call, err
}

func (p *port_calls) ViewCall(callSid string) (domains.Call, error) {
	call, err := p.driven.ViewCall(callSid)
	return call, err
}
