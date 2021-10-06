package carrier

import "zarbat_test/pkg/domains"

type PrimaryPort interface {
	CarrierLookup(phoneNumber string) (domains.CarrierLookup, error)
	CarrierLookupList() ([]domains.CarrierLookup, error)
	CNAMLookup(phoneNumber string) (domains.CNAM, error)
	CNAMLookupList() ([]domains.CNAM, error)
	BNALookup(phoneNumber string) (domains.BNA, error)
	BNALookupList() ([]domains.BNA, error)
}

type SecondaryPort interface {
	CarrierLookup(phoneNumber string) (domains.CarrierLookup, error)
	CarrierLookupList() ([]domains.CarrierLookup, error)
	CNAMLookup(phoneNumber string) (domains.CNAM, error)
	CNAMLookupList() ([]domains.CNAM, error)
	BNALookup(phoneNumber string) (domains.BNA, error)
	BNALookupList() ([]domains.BNA, error)
}
