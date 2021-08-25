package steps

import (
	"fmt"
	"zarbat_test/internal/logging"
)

var contextName string

func IWantToWriteMyName(name string) error {
	fmt.Printf("My name is %s", name)
	logging.Debug.Printf("My name is %s", name)
	contextName = name
	return nil
}

func IShouldSeeOnConsole(expectedName string) error {
	if expectedName != contextName {
		logging.Debug.Printf("Error %s", "Expected name "+expectedName+", but received "+contextName)
		return fmt.Errorf("Error %s", "Expected name "+expectedName+", but received "+contextName)
	}
	return nil
}
