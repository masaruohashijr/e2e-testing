package main

import (
	"os"
	"time"
	"zarbat_test/godog/dial"
	"zarbat_test/godog/hangup"
	"zarbat_test/godog/pause"
	"zarbat_test/godog/ping"
	"zarbat_test/godog/play"
	"zarbat_test/godog/record"
	"zarbat_test/godog/redirect"
	"zarbat_test/godog/reject"
	"zarbat_test/godog/say"

	"github.com/cucumber/godog"
)

var RegMap map[string]FeatureTest

func initRegister() {
	RegMap = make(map[string]FeatureTest)
	RegMap["play"] = FeatureTest{
		Path:                "features/play",
		ScenarioInitializer: play.InitializeScenario,
	}
	RegMap["ping"] = FeatureTest{
		Path:                "features/ping",
		ScenarioInitializer: ping.InitializeScenario,
	}
	RegMap["pause"] = FeatureTest{
		Path:                "features/pause",
		ScenarioInitializer: pause.InitializeScenario,
	}
	RegMap["dial"] = FeatureTest{
		Path:                "features/dial",
		ScenarioInitializer: dial.InitializeScenario,
	}
	RegMap["redirect"] = FeatureTest{
		Path:                "features/redirect",
		ScenarioInitializer: redirect.InitializeScenario,
	}
	RegMap["reject"] = FeatureTest{
		Path:                "features/reject",
		ScenarioInitializer: reject.InitializeScenario,
	}
	RegMap["hangup"] = FeatureTest{
		Path:                "features/hangup",
		ScenarioInitializer: hangup.InitializeScenario,
	}
	RegMap["say"] = FeatureTest{
		Path:                "features/say",
		ScenarioInitializer: say.InitializeScenario,
	}
	RegMap["record"] = FeatureTest{
		Path:                "features/record",
		ScenarioInitializer: record.InitializeScenario,
	}
}

func main() {
	initRegister()
	args := os.Args[1:]
	status := 0
	for _, a := range args {
		ft := RegMap[a]
		opts := godog.Options{
			Format:    "progress",
			Paths:     []string{ft.Path},
			Randomize: time.Now().UTC().UnixNano(),
		}
		println(opts.Paths[0], a)
		status = godog.TestSuite{
			Name:                "zarbat_test",
			ScenarioInitializer: ft.ScenarioInitializer,
			Options:             &opts,
		}.Run()
	}
	os.Exit(status)
}

type FeatureTest struct {
	Path                 string
	ScenarioInitializer  func(ctx *godog.ScenarioContext)
	TestSuiteInitializer func(ctx *godog.TestSuiteContext)
}
