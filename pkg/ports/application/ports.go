package application

import "zarbat_test/pkg/domains"

type PrimaryPort interface {
	ListApplications() ([]*domains.Application, error)
	ViewApplication() (*domains.Application, error)
	CreateApplication() (*domains.Application, error)
	UpdateApplication() (*domains.Application, error)
	DeleteApplication() (*domains.Application, error)
}

type SecondaryPort interface {
	ListApplications() ([]*domains.Application, error)
	ViewApplication() (*domains.Application, error)
	CreateApplication() (*domains.Application, error)
	UpdateApplication() (*domains.Application, error)
	DeleteApplication() (*domains.Application, error)
}
