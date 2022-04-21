package test

import (
	"hash/fnv"

	"github.com/cucumber/godog"
)

type FeatureTest struct {
	Name                 string
	Path                 string
	Hash                 uint32
	ScenarioInitializer  func(ctx *godog.ScenarioContext)
	TestSuiteInitializer func(ctx *godog.TestSuiteContext)
	Tries                int
	Result               bool
}

func Hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func InitRegister() (RegMap map[string]*FeatureTest) {
	RegMap = make(map[string]*FeatureTest)
	return
}
