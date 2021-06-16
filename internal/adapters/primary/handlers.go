package primary

import (
	"e2e-testing/internal/config"
)

func (a *CLIPrimaryAdapter) HandleCall(c config.Operations) (string, error) {
	message := "Call established"
	switch c {
	case config.MakeCall:
		a.service.MakeCall()
	case config.ListCall:
		a.service.MakeCall() //TODO
	case config.ViewCall:
		a.service.MakeCall() //TODO
	}
	return message, nil
}

func (a *HTTPPrimaryAdapter) HandleCall(c config.Operations) (string, error) {
	message := "Call established"
	switch c {
	case config.MakeCall:
		a.service.MakeCall()
	case config.ListCall:
		a.service.MakeCall() //TODO
	case config.ViewCall:
		a.service.MakeCall() //TODO
	}
	return message, nil
}
