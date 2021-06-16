package primary

import "e2e-testing/pkg/ports/calls"

type CLIPrimaryAdapter struct {
	service calls.PrimaryPort
}

func NewCLIPrimaryAdapter(s calls.PrimaryPort) *CLIPrimaryAdapter {
	return &CLIPrimaryAdapter{
		s,
	}
}
