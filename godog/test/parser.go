package test

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func ParseGherkinDocument(filePath string, ctx ScenarioContext) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	r := bufio.NewReader(f)
	feature := &FeatureContext{}
	scenario := &ScenarioContext{}
	var scenarios []ScenarioContext
	background := &StepDefinition{}
	step := &StepDefinition{}

	var s string
	for err == nil {
		s, err = Readln(r)
		feature = findFeature(s, feature)
		background = findBackGround(s, background)
		scenario = findScenario(s, scenario, background)
		step = findStep(s, step)
		scenarios = append(scenarios, *scenario)
	}
}

func findBackGround(s string, background *StepDefinition) *StepDefinition {
	re := regexp.MustCompile(`(?P<Step>(\s*)Background:)(?P<StepExpression>.*)`)
	matches := re.FindStringSubmatch(s)
	if len(matches) == 0 {
		return background
	}
	backgroundIndex := re.SubexpIndex("StepExpression")
	bg := strings.TrimSpace(matches[backgroundIndex])
	println(bg)
	return background
}

func findFeature(s string, feature *FeatureContext) *FeatureContext {
	re := regexp.MustCompile(`(?P<Feature>(.*)Feature:)(?P<FeatureDescription>.*)`)
	matches := re.FindStringSubmatch(s)
	if len(matches) == 0 {
		return feature
	}
	featureIndex := re.SubexpIndex("FeatureDescription")
	feature.Name = strings.TrimSpace(matches[featureIndex])
	return feature
}

func findScenario(s string, scenario *ScenarioContext, background *StepDefinition) *ScenarioContext {
	re := regexp.MustCompile(`(?P<Scenario>(\s*)Scenario:)(?P<ScenarioDescription>.*)`)
	matches := re.FindStringSubmatch(s)
	if len(matches) == 0 {
		return scenario
	}
	featureIndex := re.SubexpIndex("ScenarioDescription")
	scenario.Name = matches[featureIndex]
	return scenario
}

func findStep(s string, step *StepDefinition) *StepDefinition {
	re := regexp.MustCompile(`(?P<Step>(\s*)And:)(?P<StepExpression>.*)`)
	matches := re.FindStringSubmatch(s)
	if len(matches) == 0 {
		return step
	}
	return step
}

func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}
