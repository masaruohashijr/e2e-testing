package test

import (
	"zarbat_test/internal/godog/dial"
	"zarbat_test/internal/godog/hangup"
	"zarbat_test/internal/godog/number"
	"zarbat_test/internal/godog/pause"
	"zarbat_test/internal/godog/ping"
	"zarbat_test/internal/godog/play"
	"zarbat_test/internal/godog/record"
	"zarbat_test/internal/godog/redirect"
	"zarbat_test/internal/godog/reject"
	"zarbat_test/internal/godog/say"
	"zarbat_test/internal/godog/sms"

	"github.com/cucumber/godog"
)

type FeatureTest struct {
	Name                 string
	Path                 string
	ScenarioInitializer  func(ctx *godog.ScenarioContext)
	TestSuiteInitializer func(ctx *godog.TestSuiteContext)
	Tries                int
}

func InitRegister() (RegMap map[string]*FeatureTest) {
	RegMap = make(map[string]*FeatureTest)
	RegMap["play"] = &FeatureTest{
		Name:                "play",
		Path:                "features/play",
		ScenarioInitializer: play.InitializeScenario,
	}
	RegMap["ping"] = &FeatureTest{
		Name:                "ping",
		Path:                "features/ping",
		ScenarioInitializer: ping.InitializeScenario,
	}
	RegMap["pause"] = &FeatureTest{
		Name:                "pause",
		Path:                "features/pause",
		ScenarioInitializer: pause.InitializeScenario,
	}
	RegMap["dial"] = &FeatureTest{
		Name:                "dial",
		Path:                "features/dial",
		ScenarioInitializer: dial.InitializeScenario,
	}
	RegMap["redirect"] = &FeatureTest{
		Name:                "redirect",
		Path:                "features/redirect",
		ScenarioInitializer: redirect.InitializeScenario,
	}
	RegMap["reject"] = &FeatureTest{
		Name:                "reject",
		Path:                "features/reject",
		ScenarioInitializer: reject.InitializeScenario,
	}
	RegMap["hangup"] = &FeatureTest{
		Name:                "hangup",
		Path:                "features/hangup",
		ScenarioInitializer: hangup.InitializeScenario,
	}
	RegMap["say"] = &FeatureTest{
		Name:                "say",
		Path:                "features/say",
		ScenarioInitializer: say.InitializeScenario,
	}
	RegMap["gather"] = &FeatureTest{
		Name:                "gather",
		Path:                "features/gather",
		ScenarioInitializer: say.InitializeScenario,
	}
	RegMap["record"] = &FeatureTest{
		Name:                "record",
		Path:                "features/record",
		ScenarioInitializer: record.InitializeScenario,
	}
	RegMap["buy"] = &FeatureTest{
		Name:                "buy",
		Path:                "features/buy",
		ScenarioInitializer: number.InitializeScenario,
	}
	RegMap["sms"] = &FeatureTest{
		Name:                "sms",
		Path:                "features/sms",
		ScenarioInitializer: sms.InitializeScenario,
	}
	return
}
