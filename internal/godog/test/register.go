package test

import (
	"hash/fnv"
	"zarbat_test/internal/godog/number"
	"zarbat_test/internal/godog/sms"
	"zarbat_test/internal/steps"

	"github.com/cucumber/godog"
)

type FeatureTest struct {
	Name                 string
	Path                 string
	Hash                 uint32
	ScenarioInitializer  func(ctx *godog.ScenarioContext)
	TestSuiteInitializer func(ctx *godog.TestSuiteContext)
	Tries                int
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func InitRegister() (RegMap map[string]*FeatureTest) {
	RegMap = make(map[string]*FeatureTest)
	RegMap["play"] = &FeatureTest{
		Name:                "play",
		Path:                "features/play",
		Hash:                hash("play"),
		ScenarioInitializer: steps.InitializeScenario,
	}
	RegMap["playlastrecording"] = &FeatureTest{
		Name:                "playlastrecording",
		Path:                "features/playlastrecording",
		Hash:                hash("playlastrecording"),
		ScenarioInitializer: steps.InitializeScenario,
	}
	RegMap["ping"] = &FeatureTest{
		Name:                "ping",
		Path:                "features/ping",
		Hash:                hash("ping"),
		ScenarioInitializer: steps.InitializeScenario,
	}
	RegMap["pause"] = &FeatureTest{
		Name:                "pause",
		Path:                "features/pause",
		Hash:                hash("pause"),
		ScenarioInitializer: steps.InitializeScenario,
	}
	RegMap["dial"] = &FeatureTest{
		Name:                "dial",
		Path:                "features/dial",
		Hash:                hash("dial"),
		ScenarioInitializer: steps.InitializeScenario,
	}
	RegMap["number"] = &FeatureTest{
		Name:                "number",
		Path:                "features/number",
		Hash:                hash("number"),
		ScenarioInitializer: steps.InitializeScenario,
	}
	RegMap["redirect"] = &FeatureTest{
		Name:                "redirect",
		Path:                "features/redirect",
		Hash:                hash("redirect"),
		ScenarioInitializer: steps.InitializeScenario,
	}
	RegMap["reject"] = &FeatureTest{
		Name:                "reject",
		Path:                "features/reject",
		Hash:                hash("reject"),
		ScenarioInitializer: steps.InitializeScenario,
	}
	RegMap["hangup"] = &FeatureTest{
		Name:                "hangup",
		Path:                "features/hangup",
		Hash:                hash("hangup"),
		ScenarioInitializer: steps.InitializeScenario,
	}
	RegMap["say"] = &FeatureTest{
		Name:                "say",
		Path:                "features/say",
		Hash:                hash("say"),
		ScenarioInitializer: steps.InitializeScenario,
	}
	RegMap["gather"] = &FeatureTest{
		Name:                "gather",
		Path:                "features/gather",
		Hash:                hash("gather"),
		ScenarioInitializer: steps.InitializeScenario,
	}
	RegMap["record"] = &FeatureTest{
		Name:                "record",
		Path:                "features/record",
		Hash:                hash("record"),
		ScenarioInitializer: steps.InitializeScenario,
	}
	RegMap["buy"] = &FeatureTest{
		Name:                "buy",
		Path:                "features/buy",
		Hash:                hash("buy"),
		ScenarioInitializer: number.InitializeScenario,
	}
	RegMap["sms"] = &FeatureTest{
		Name:                "sms",
		Path:                "features/sms",
		Hash:                hash("sms"),
		ScenarioInitializer: sms.InitializeScenario,
	}
	return
}
