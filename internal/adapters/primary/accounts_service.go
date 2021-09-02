package primary

import (
	"zarbat_test/pkg/domains"
	"zarbat_test/pkg/ports/account"
)

type port_account struct {
	driven account.SecondaryPort
}

func NewAccountsService(driven account.SecondaryPort) account.PrimaryPort {
	return &port_account{
		driven,
	}
}

func (p *port_account) ViewAccount() (*domains.Account, error) {
	va, err := p.driven.ViewAccount()
	return va, err
}

func (p *port_account) UpdateAccount(friendlyName string) error {
	err := p.driven.UpdateAccount(friendlyName)
	return err
}
