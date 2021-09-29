package primary

import (
	"zarbat_test/pkg/domains"
	usage "zarbat_test/pkg/ports/usages"
)

type port_usage struct {
	driven usage.SecondaryPort
}

func NewUsageService(driven usage.SecondaryPort) usage.PrimaryPort {
	return &port_usage{
		driven,
	}
}

func (p *port_usage) ViewUsage(usageSid string) (domains.Usage, error) {
	usage, err := p.driven.ViewUsage(usageSid)
	return usage, err
}

func (p *port_usage) ListUsage() ([]domains.Usage, error) {
	smss, err := p.driven.ListUsage()
	return smss, err
}
