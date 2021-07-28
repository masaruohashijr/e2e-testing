package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	"zarbat_test/godog/dial"
	"zarbat_test/godog/hangup"
	"zarbat_test/godog/number"
	"zarbat_test/godog/pause"
	"zarbat_test/godog/ping"
	"zarbat_test/godog/play"
	"zarbat_test/godog/record"
	"zarbat_test/godog/redirect"
	"zarbat_test/godog/reject"
	"zarbat_test/godog/say"
	"zarbat_test/godog/sms"
	"zarbat_test/internal/config"

	"github.com/cucumber/godog"
)

var RegMap map[string]*FeatureTest
var logPtr *string
var triesPtr *int

func initRegister() {
	RegMap = make(map[string]*FeatureTest)
	RegMap["play"] = &FeatureTest{
		Path:                "features/play",
		ScenarioInitializer: play.InitializeScenario,
	}
	RegMap["ping"] = &FeatureTest{
		Path:                "features/ping",
		ScenarioInitializer: ping.InitializeScenario,
	}
	RegMap["pause"] = &FeatureTest{
		Path:                "features/pause",
		ScenarioInitializer: pause.InitializeScenario,
	}
	RegMap["dial"] = &FeatureTest{
		Path:                "features/dial",
		ScenarioInitializer: dial.InitializeScenario,
	}
	RegMap["redirect"] = &FeatureTest{
		Path:                "features/redirect",
		ScenarioInitializer: redirect.InitializeScenario,
	}
	RegMap["reject"] = &FeatureTest{
		Path:                "features/reject",
		ScenarioInitializer: reject.InitializeScenario,
	}
	RegMap["hangup"] = &FeatureTest{
		Path:                "features/hangup",
		ScenarioInitializer: hangup.InitializeScenario,
	}
	RegMap["say"] = &FeatureTest{
		Path:                "features/say",
		ScenarioInitializer: say.InitializeScenario,
	}
	RegMap["record"] = &FeatureTest{
		Path:                "features/record",
		ScenarioInitializer: record.InitializeScenario,
	}
	RegMap["buy"] = &FeatureTest{
		Path:                "features/number",
		ScenarioInitializer: number.InitializeScenario,
	}
	RegMap["sms"] = &FeatureTest{
		Path:                "features/sms",
		ScenarioInitializer: sms.InitializeScenario,
	}
}

func main() {
	initRegister()
	tests := initArgs()
	initLogger()
	log.Println("**********************************")
	log.Println("START OF TEST SUITE")
	status := 0
	result := "OK"
	for i := 0; i < len(tests); i++ {
		ft := RegMap[tests[i]]
		opts := godog.Options{
			Format:    "progress",
			Paths:     []string{ft.Path},
			Randomize: time.Now().UTC().UnixNano(),
		}
		status = godog.TestSuite{
			Name:                "zarbat_test",
			ScenarioInitializer: ft.ScenarioInitializer,
			Options:             &opts,
		}.Run()
		if status != 0 {
			result = "Not OK"
			log.Println("* Feature/Scenario: ", strings.ToUpper(tests[i]), result)
			ft.tries += 1
			if ft.tries <= *triesPtr {
				i--
			}
			time.Sleep(5 * time.Second)
		} else {
			result = "OK"
			log.Println("* Feature/Scenario: ", strings.ToUpper(tests[i]), result)
		}
	}
	log.Println("...END OF TEST SUITE")
	os.Exit(status)
}

type FeatureTest struct {
	Path                 string
	ScenarioInitializer  func(ctx *godog.ScenarioContext)
	TestSuiteInitializer func(ctx *godog.TestSuiteContext)
	tries                int
}

func initArgs() []string {
	var tests []string
	configPtr := flag.String("config", "config/config.ini", "a configuration file")
	config.ConfigPath = *configPtr
	triesPtr = flag.Int("n", 2, "number of tries")
	logPtr = flag.String("l", "log/.log", "log location")
	logLevelPtr := flag.String("level", "summary", "options: info, summary, debug, error")
	testPtr := flag.String("test", "buy", "ctlang")
	flag.Parse()
	addons := flag.Args()
	tests = append(tests, *testPtr)
	for _, a := range addons {
		tests = append(tests, a)
	}
	fmt.Println("************************************************")
	fmt.Println("*** Config:", *configPtr)
	fmt.Println("*** Number of Tries:", *triesPtr)
	fmt.Println("*** Log:", *logPtr)
	fmt.Println("*** Logging Level:", *logLevelPtr)
	fmt.Printf("*** Tests: %+q\n", tests)
	fmt.Println("************************************************")
	return tests
}

func initLogger() {
	file, err := os.OpenFile(*logPtr, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
}
