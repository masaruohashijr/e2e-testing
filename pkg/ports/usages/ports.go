package usage

import "zarbat_test/pkg/domains"

type PrimaryPort interface {
	ViewUsage(usageSid string) (domains.Usage, error)
	ListUsage() ([]domains.Usage, error)
}

type SecondaryPort interface {
	ViewUsage(usageSid string) (domains.Usage, error)
	ListUsage() ([]domains.Usage, error)
}
