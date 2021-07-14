package primary

import "zarbat_test/pkg/ports/calls"

type HTTPPrimaryAdapter struct {
	service calls.PrimaryPort
}

func NewHTTPPrimaryAdapter(s calls.PrimaryPort) *HTTPPrimaryAdapter {
	return &HTTPPrimaryAdapter{
		s,
	}
}
