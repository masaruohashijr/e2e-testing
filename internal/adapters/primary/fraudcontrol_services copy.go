package primary

import (
	"zarbat_test/pkg/domains"
	"zarbat_test/pkg/ports/fraudcontrol"
)

type port_fraudcontrol struct {
	driven fraudcontrol.SecondaryPort
}

func NewFraudControlService(driven fraudcontrol.SecondaryPort) fraudcontrol.PrimaryPort {
	return &port_fraudcontrol{
		driven,
	}
}

func (p *port_fraudcontrol) BlockDestination(countryCode string) (domains.Blocked, error) {
	fraudcontrolLookup, err := p.driven.BlockDestination(countryCode)
	return fraudcontrolLookup, err
}

func (p *port_fraudcontrol) AuthorizeDestination(countryCode string) (domains.Authorized, error) {
	authorizeDestination, err := p.driven.AuthorizeDestination(countryCode)
	return authorizeDestination, err
}

func (p *port_fraudcontrol) ExtendDestinationAuthorization(countryCode string) (domains.Authorized, error) {
	cnamLookup, err := p.driven.ExtendDestinationAuthorization(countryCode)
	return cnamLookup, err
}

func (p *port_fraudcontrol) WhiteListDestination(countryCode string) (domains.WhiteListed, error) {
	whitelisted, err := p.driven.WhiteListDestination(countryCode)
	return whitelisted, err
}

func (p *port_fraudcontrol) ListFraudControl() ([]domains.Fraud, error) {
	frauds, err := p.driven.ListFraudControl()
	return frauds, err
}
