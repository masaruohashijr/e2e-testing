package number

import (
	"fmt"
	"zarbat_test/internal/adapters/primary"
	"zarbat_test/internal/adapters/secondary"
	"zarbat_test/internal/config"
	"zarbat_test/pkg/domains"
	"zarbat_test/pkg/ports/calls"
	"zarbat_test/pkg/ports/numbers"

	"github.com/cucumber/godog"
)

var Configuration config.ConfigType
var CallSecondaryPort calls.SecondaryPort
var CallPrimaryPort calls.PrimaryPort
var NumberSecondaryPort numbers.SecondaryPort
var NumberPrimaryPort numbers.PrimaryPort
var ResponsePlay domains.ResponsePlay
var ResponseGather domains.ResponseGather
var ResponseRecord domains.ResponseRecord
var Ch = make(chan string)
var AvailableNumbers []string
var IncomingNumbers []string

func IListAllAvailableNumbers() error {
	if len(AvailableNumbers) == 0 {
		return fmt.Errorf("Error %s", "Not able to list available numbers.")
	}
	for _, a := range AvailableNumbers {
		println(a)
	}
	return nil
}

func IShouldGetToBuyFromList(amount int) error {
	ok := false
	for i := 0; i < amount; i++ {
		NumberSecondaryPort.AddNumber(AvailableNumbers[i])
		purchased, _ := NumberSecondaryPort.ListNumbers()
		for _, n := range purchased {
			if AvailableNumbers[i] == n {
				ok = true
				break
			}
		}
		if !ok {
			return fmt.Errorf("Error %s", "Not able to list available numbers.")
		}
	}

	return nil
}

func MyTestSetupRuns() error {
	Configuration = config.NewConfig()
	NumberSecondaryPort = secondary.NewNumbersApi(&Configuration)
	NumberPrimaryPort = primary.NewNumbersService(NumberSecondaryPort)
	anumbers, err := NumberSecondaryPort.ListAvailableNumbers()
	AvailableNumbers = anumbers
	if err != nil {
		return fmt.Errorf("Error %s", "Not able to list available numbers.")
	}
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^my test setup runs$`, MyTestSetupRuns)
	ctx.Step(`^I list all available numbers$`, IListAllAvailableNumbers)
	ctx.Step(`^I should get to buy (\d+) from list$`, IShouldGetToBuyFromList)
}
