package account

import "zarbat_test/pkg/domains"

type PrimaryPort interface {
	UpdateAccount(string) error
	ViewAccount() (*domains.Account, error)
}

type SecondaryPort interface {
	UpdateAccount(string) error
	ViewAccount() (*domains.Account, error)
}
