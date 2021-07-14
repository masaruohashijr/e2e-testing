package primary

import "zarbat_test/pkg/ports/calls"

type port struct {
	driven calls.SecondaryPort
}

func (p *port) MakeCall() error {
	err := p.driven.MakeCall()
	return err
}
