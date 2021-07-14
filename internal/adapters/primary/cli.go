package primary

import "zarbat_test/pkg/ports/calls"

type CLIPrimaryAdapter struct {
	service calls.PrimaryPort
}

func NewCLIPrimaryAdapter(s calls.PrimaryPort) *CLIPrimaryAdapter {
	return &CLIPrimaryAdapter{
		s,
	}
}
