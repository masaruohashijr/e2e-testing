package primary

import (
	"zarbat_test/pkg/domains"
	"zarbat_test/pkg/ports/application"
)

type port_application struct {
	driven application.SecondaryPort
}

func NewApplicationsService(driven application.SecondaryPort) application.PrimaryPort {
	return &port_application{
		driven,
	}
}
func (p *port_application) ListApplications() ([]domains.Application, error) {
	va, err := p.driven.ListApplications()
	return va, err
}

func (p *port_application) ViewApplication(applicationSid string) (domains.Application, error) {
	va, err := p.driven.ViewApplication(applicationSid)
	return va, err
}

func (p *port_application) CreateApplication(friendlyName string) (domains.Application, error) {
	va, err := p.driven.CreateApplication(friendlyName)
	return va, err
}

func (p *port_application) UpdateApplication(applicationSid, friendlyName string) (domains.Application, error) {
	va, err := p.driven.UpdateApplication(applicationSid, friendlyName)
	return va, err
}

func (p *port_application) DeleteApplication(applicationSid string) (domains.Application, error) {
	va, err := p.driven.DeleteApplication(applicationSid)
	return va, err
}
