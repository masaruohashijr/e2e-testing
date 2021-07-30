package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	"zarbat_test/internal/config"
	"zarbat_test/internal/files"
	"zarbat_test/internal/godog/services"
	"zarbat_test/internal/godog/test"

	"github.com/cucumber/godog"
)

var RegMap map[string]*test.FeatureTest
var logPtr *string
var logLevelPtr *string
var configPtr *string
var triesPtr *int
var tests []test.FeatureTest

func main() {
	RegMap = test.InitRegister()
	tests = initArgs(RegMap)
	initLogger()
	log.Println("****************************************")
	log.Println("START OF TEST SUITE")
	logArgs()
	status := 0
	for i := 0; i < len(tests); i++ {
		ft := RegMap[tests[i].Name]
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
			logResult(tests[i].Name, "Not OK")
			ft.Tries += 1
			if ft.Tries < *triesPtr {
				i--
			}
		} else {
			logResult(tests[i].Name, "OK")
		}
		time.Sleep(5 * time.Second)
	}
	log.Println("...END OF TEST SUITE")
	os.Exit(status)
}

func initArgs(regMap map[string]*test.FeatureTest) (fts []test.FeatureTest) {

	configPtr = flag.String("config", "config/config.ini", "A configuration file")
	config.ConfigPath = *configPtr
	services.BaseUrl = config.NewConfig().BaseUrl
	triesPtr = flag.Int("n", 5, "number of tries")
	logPtr = flag.String("l", "log/.log", "log location")
	logLevelPtr = flag.String("level", "summary", "options: info, summary, debug, error")
	testPtr := flag.String("test", "buy", "ctlang")

	flag.Parse()

	if strings.HasSuffix(*testPtr, ".ctlang") {
		tempFiles := files.NewTempFiles(*testPtr)
		return files.NewFeatureTests(tempFiles, regMap)
	} else {
		var tests []string
		tests = append(tests, *testPtr)
		addons := flag.Args() // tail of the arguments
		for _, a := range addons {
			tests = append(tests, a)
		}
		files.NewSingleFile(tests)
		return files.GetFeatureTestsFromMap(tests, regMap)
	}
}

func printArgs(tests []string) {
	fmt.Println("************************************************")
	fmt.Println("*** Config:", *configPtr)
	fmt.Println("*** Number of Tries:", *triesPtr)
	fmt.Println("*** Log:", *logPtr)
	fmt.Println("*** Logging Level:", *logLevelPtr)
	fmt.Printf("*** Features: %v\n", tests)
	fmt.Println("************************************************")
}

func initLogger() {
	file, err := os.OpenFile(*logPtr, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
}

func logArgs() {
	log.Println("Config:", *configPtr)
	log.Println("Number of Tries:", *triesPtr)
	log.Println("Log:", *logPtr)
	var tsts = ""
	for _, t := range tests {
		tsts += "[" + t.Name + "]"
	}
	log.Println("Tests:", tsts)
	log.Println(".........................................")
}

func logResult(test, result string) {
	log.Printf("* Feature/Scenario: %s - Status: %s\n", strings.ToUpper(test), result)
}
