package primary

import (
	"zarbat_test/pkg/domains"
	"zarbat_test/pkg/ports/carrier"
)

type port_carrier struct {
	driven carrier.SecondaryPort
}

func NewCarrierService(driven carrier.SecondaryPort) carrier.PrimaryPort {
	return &port_carrier{
		driven,
	}
}

func (p *port_carrier) CarrierLookup(phoneNumber string) (domains.CarrierLookup, error) {
	carrierLookup, err := p.driven.CarrierLookup(phoneNumber)
	return carrierLookup, err
}

func (p *port_carrier) CarrierLookupList() ([]domains.CarrierLookup, error) {
	carrierLookups, err := p.driven.CarrierLookupList()
	return carrierLookups, err
}

func (p *port_carrier) CNAMLookup(phoneNumber string) (domains.CNAM, error) {
	cnamLookup, err := p.driven.CNAMLookup(phoneNumber)
	return cnamLookup, err
}

func (p *port_carrier) CNAMLookupList() ([]domains.CNAM, error) {
	cnamLookups, err := p.driven.CNAMLookupList()
	return cnamLookups, err
}

func (p *port_carrier) BNALookup(phoneNumber string) (domains.BNA, error) {
	bnaLookup, err := p.driven.BNALookup(phoneNumber)
	return bnaLookup, err
}

func (p *port_carrier) BNALookupList() ([]domains.BNA, error) {
	bnaLookups, err := p.driven.BNALookupList()
	return bnaLookups, err
}
