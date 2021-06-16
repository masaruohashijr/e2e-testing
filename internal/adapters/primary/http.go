package primary

import "e2e-testing/pkg/ports/calls"

type HTTPPrimaryAdapter struct {
	service calls.PrimaryPort
}

func NewHTTPPrimaryAdapter(s calls.PrimaryPort) *HTTPPrimaryAdapter {
	return &HTTPPrimaryAdapter{
		s,
	}
}
