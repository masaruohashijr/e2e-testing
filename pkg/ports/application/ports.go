package application

import "zarbat_test/pkg/domains"

type PrimaryPort interface {
	ListApplications() ([]domains.Application, error)
	ViewApplication(aplicationSid string) (domains.Application, error)
	CreateApplication(friendlyName string) (domains.Application, error)
	UpdateApplication(aplicationSid, friendlyName string) (domains.Application, error)
	DeleteApplication(aplicationSid string) (domains.Application, error)
}

type SecondaryPort interface {
	ListApplications() ([]domains.Application, error)
	ViewApplication(aplicationSid string) (domains.Application, error)
	CreateApplication(friendlyName string) (domains.Application, error)
	UpdateApplication(aplicationSid, friendlyName string) (domains.Application, error)
	DeleteApplication(aplicationSid string) (domains.Application, error)
}
