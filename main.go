package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
	"zarbat_test/internal/config"
	"zarbat_test/internal/files"
	"zarbat_test/internal/godog/services"
	"zarbat_test/internal/godog/test"
	l "zarbat_test/internal/logging"
	"zarbat_test/internal/steps"

	"github.com/cucumber/godog"
)

var RegMap map[string]*test.FeatureTest
var logPtr *string
var logLevelPtr *string
var configPtr *string
var triesPtr *int

func main() {
	RegMap = test.InitRegister()
	tests, tempDir := initArgs(RegMap)
	if checkEmpty(tests) != nil {
		fmt.Println("Failed to find ctlang file.")
		os.Exit(2)
	}
	initLoggers()
	l.Info.Println("****************************************")
	l.Info.Println("START OF TEST SUITE")
	logArgs(tests)
	status := 0
	passed := 0
	failed := 0
	for i := 0; i < len(tests); i++ {
		ft := RegMap[tests[i].Name]
		fmt.Println("******")
		fmt.Println(" Test: " + tests[i].Name)
		fmt.Println("******")
		steps.TestHash = ft.Hash
		services.TestHash = ft.Hash
		fmt.Println(ft.Hash)
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
			} else {
				failed++
			}
		} else {
			logResult(tests[i].Name, "OK")
			passed++
		}
		time.Sleep(5 * time.Second)
	}
	l.Info.Println("Passed ", passed)
	l.Info.Println("Failed ", failed)
	l.Info.Println("...END OF TEST SUITE")
	os.RemoveAll(tempDir)
	os.Exit(status)
}

func checkEmpty(tests []test.FeatureTest) error {
	if len(tests) == 0 {
		return fmt.Errorf("File not found exception.")
	}
	return nil
}

func initArgs(regMap map[string]*test.FeatureTest) (fts []test.FeatureTest, tempDir string) {

	configPtr = flag.String("config", "config/config.ini", "A configuration file")
	config.ConfigPath = *configPtr
	services.BaseUrl = config.NewConfig().BaseUrl
	triesPtr = flag.Int("n", 5, "number of tries")
	logPtr = flag.String("l", "log/.log", "log location")
	logLevelPtr = flag.String("level", "summary", "options: info, debug")
	testPtr := flag.String("test", "buy", "ctlang")

	flag.Parse()

	if strings.HasSuffix(*testPtr, ".ctlang") {
		tempFiles, tempDir := files.NewTempFiles(*testPtr)
		return files.NewFeatureTests(tempFiles, regMap), tempDir
	} else {
		var tests []string
		tests = append(tests, *testPtr)
		addons := flag.Args() // tail of the arguments
		for _, a := range addons {
			tests = append(tests, a)
		}
		files.NewSingleFile(tests)
		return files.GetFeatureTestsFromMap(tests, regMap), ""
	}
}

func initLoggers() {
	infoLogFile, err := os.OpenFile(nameLogFile("INFO", *logPtr), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	l.Info = log.New(infoLogFile, "INFO: ", log.LstdFlags|log.Lmsgprefix)
	debugLogFile, err := os.OpenFile(nameLogFile("DEBUG", *logPtr), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	l.Debug = log.New(debugLogFile, "DEBUG: ", log.LstdFlags|log.Lmsgprefix)
}

func nameLogFile(level, rootName string) string {
	re := regexp.MustCompile(`(?P<RootName>(.*))(?P<Extension>\..*)`)
	matches := re.FindStringSubmatch(rootName)
	rootIndex := re.SubexpIndex("RootName")
	namePrefix := strings.TrimSpace(matches[rootIndex])
	extensionIndex := re.SubexpIndex("Extension")
	extension := strings.TrimSpace(matches[extensionIndex])
	return namePrefix + "_" + level + extension
}

func logArgs(tests []test.FeatureTest) {
	l.Info.Println("Config:", *configPtr)
	l.Info.Println("Number of Tries:", *triesPtr)
	l.Info.Println("Log:", *logPtr)
	var tsts = ""
	for _, t := range tests {
		tsts += "[" + t.Name + "]"
	}
	l.Info.Println("Tests:", tsts)
	l.Info.Println(".........................................")
	l.Debug.Println("************************************************")
	l.Debug.Println("*** Config:", *configPtr)
	l.Debug.Println("*** Number of Tries:", *triesPtr)
	l.Debug.Println("*** Log:", *logPtr)
	l.Debug.Println("*** Logging Level:", *logLevelPtr)
	l.Debug.Printf("*** Features: %v\n", tsts)
	l.Debug.Println("************************************************")
}

func logResult(test, result string) {
	l.Info.Printf("* Feature/Scenario: %s - Status: %s\n", strings.ToUpper(test), result)
}
