package fraudcontrol

import "zarbat_test/pkg/domains"

type PrimaryPort interface {
	BlockDestination(countryCode string) (domains.Blocked, error)
	AuthorizeDestination(countryCode string) (domains.Authorized, error)
	ExtendDestinationAuthorization(countryCode string) (domains.Authorized, error)
	WhiteListDestination(countryCode string) (domains.WhiteListed, error)
	ListFraudControl() ([]domains.Fraud, error)
}

type SecondaryPort interface {
	BlockDestination(countryCode string) (domains.Blocked, error)
	AuthorizeDestination(countryCode string) (domains.Authorized, error)
	ExtendDestinationAuthorization(countryCode string) (domains.Authorized, error)
	WhiteListDestination(countryCode string) (domains.WhiteListed, error)
	ListFraudControl() ([]domains.Fraud, error)
}
