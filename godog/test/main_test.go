package main

import (
	"os"
	"strings"
	"testing"
	"time"
	"zarbat_test/godog/dial"
	"zarbat_test/godog/hangup"
	"zarbat_test/godog/pause"
	"zarbat_test/godog/ping"
	"zarbat_test/godog/play"

	"github.com/cucumber/godog"
)

func (r Register) init() {
	r.register("dial", dial.InitializeScenario)
	r.register("hangup", hangup.InitializeScenario)
	r.register("pause", pause.InitializeScenario)
	r.register("ping", ping.InitializeScenario)
	r.register("play", play.InitializeScenario)
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
}

func TestMain(m *testing.M) {

	r := &Register{
		ScenarioInitializerMap: make(map[string]func(ctx *godog.ScenarioContext)),
	}
	r.init()
	status := 0
	features := r.extractFeatures(os.Args[1:])

	for i := range features {
		println("%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%")
		println(strings.ToUpper(features[i]))
		println("%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%")
		opts := godog.Options{
			Format:    "progress",
			Paths:     []string{"../../features/" + features[i]},
			Randomize: time.Now().UTC().UnixNano(),
		}
		status = godog.TestSuite{
			Name:                 "zarbat_test",
			TestSuiteInitializer: InitializeTestSuite,
			ScenarioInitializer:  r.GetScenarioInitializer(features[i]),
			Options:              &opts,
		}.Run()

		// Optional: Run `testing` package's logic besides godog.
		if st := m.Run(); st > status {
			status = st
		}
	}
	os.Exit(status)
}

type Register struct {
	ScenarioInitializerMap map[string]func(ctx *godog.ScenarioContext)
}

func (r Register) register(feature string, f func(ctx *godog.ScenarioContext)) {
	r.ScenarioInitializerMap[feature] = f
}

func (r Register) GetScenarioInitializer(feature string) func(ctx *godog.ScenarioContext) {
	return r.ScenarioInitializerMap[feature]
}

func (r Register) extractFeatures(array []string) (result []string) {
	for n := range array {
		if _, ok := r.ScenarioInitializerMap[array[n]]; ok {
			result = append(result, array[n])
		}
	}
	return result
}
