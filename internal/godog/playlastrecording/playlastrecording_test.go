package main

import (
	"os"
	"testing"
	"time"

	"github.com/cucumber/godog"
)

func configuredToPlayLastRecording(arg1 string) error {
	return godog.ErrPending
}

func configuredToRecordCalls(arg1 string) error {
	return godog.ErrPending
}

func configuredToSay(arg1, arg2 string) error {
	return godog.ErrPending
}

func shouldGetTranscription(arg1, arg2 string) error {
	return godog.ErrPending
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^"([^"]*)" configured to play last recording$`, configuredToPlayLastRecording)
	ctx.Step(`^"([^"]*)" configured to record calls$`, configuredToRecordCalls)
	ctx.Step(`^"([^"]*)" configured to say "([^"]*)"$`, configuredToSay)
	ctx.Step(`^"([^"]*)" should get transcription "([^"]*)"$`, shouldGetTranscription)
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {

}

func TestMain(m *testing.M) {
	opts := godog.Options{
		Format:    "progress",
		Paths:     []string{"../../features/playlastrecording"},
		Randomize: time.Now().UTC().UnixNano(),
	}

	status := godog.TestSuite{
		Name:                 "zarbat_test",
		TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  InitializeScenario,
		Options:              &opts,
	}.Run()

	// Optional: Run `testing` package's logic besides godog.
	if st := m.Run(); st > status {
		status = st
	}

	os.Exit(status)
}
