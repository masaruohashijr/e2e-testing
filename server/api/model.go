package api

import "zarbat_test/internal/config"

type TestRun struct {
	Name          string         `json:"name"`
	ScenarioId    string         `json:"scenarioId"`
	ListOfSteps   string         `json:"listOfSteps"`
	Description   string         `json:"description"`
	Result        string         `json:"result"`
	Logs          string         `json:"logs"`
	EnvironmentId string         `json:"environmentId"`
	ContextId     string         `json:"contextId"`
	UserId        string         `json:"userId"`
	RunAt         string         `json:"runAt"`
	Args          Arguments      `json:"args"`
	Context       config.Context `json:"context"`
}

type Arguments struct {
	Config   string `json:"config"`
	Url      string `json:"url"`
	Port     string `json:"port"`
	Log      string `json:"log"`
	LogLevel string `json:"logLevel"`
	NTries   string `json:"numberOfTries"`
	Test     string `json:"test"`
}
