package primary

import (
	"e2e-testing/pkg/ports/numbers"
)

type port_number struct {
	driven numbers.SecondaryPort
}

func NewNumbersService(driven numbers.SecondaryPort) numbers.PrimaryPort {
	return &port_number{
		driven,
	}
}

func (p *port_number) UpdateNumber() error {
	err := p.driven.UpdateNumber()
	return err
}
