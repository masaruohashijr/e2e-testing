package primary

import "e2e-testing/pkg/ports/calls"

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
