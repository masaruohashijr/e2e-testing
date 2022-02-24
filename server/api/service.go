package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"zarbat_test/internal/godog/services"
	"zarbat_test/internal/godog/test"
	"zarbat_test/internal/logging"
	"zarbat_test/internal/steps"

	"github.com/cucumber/godog"
)

const TempRootDir = "./temp"

var ApiLogger string
var callbackUrlPtr *string
var callbackPortPtr *string
var testPtr *string
var triesPtr *string
var configPtr *string
var logPtr *string
var logLevelPtr *string
var regMap map[string]*test.FeatureTest

func RunSingleTest(testRun TestRun) TestRun {
	tmpTriesPtr := testRun.Args.NTries
	triesPtr = &tmpTriesPtr
	tmpLogPtr := testRun.Args.Log
	logPtr = &tmpLogPtr
	logLevel := testRun.Args.LogLevel
	logLevelPtr = &logLevel

	_beginMark_INFO := beginLog("INFO")
	_beginMark_DEBUG := beginLog("DEBUG")

	// Generate Temporary Test Dir
	tempDir := NewTempDir()
	// Generate Temporary Test File
	tempFile := NewTempFile(tempDir, testRun.FeatureName)
	lines := GetListOfSteps(testRun.ListOfSteps)
	AddLineToTemp(tempFile, "Feature: "+testRun.FeatureName)
	AddLineToTemp(tempFile, "Scenario: "+testRun.Name)
	for _, line := range lines {
		AddLineToTemp(tempFile, line)
	}
	var tests []test.FeatureTest
	ft := &test.FeatureTest{
		Name:                testRun.FeatureName,
		Path:                tempFile.Name(),
		Hash:                test.Hash(testRun.FeatureName),
		ScenarioInitializer: steps.InitializeScenario,
	}
	regMap = test.InitRegister()
	regMap[testRun.FeatureName] = ft
	tests = append(tests, *ft)
	tests = ExecuteTest(tempDir, regMap, tests, &testRun.Args)
	if tests[0].Result {
		testRun.Result = "PASSED"
	} else {
		testRun.Result = "FAILED"
	}
	_endMark_INFO := endLog("INFO")
	_endMark_DEBUG := endLog("DEBUG")
	logsINFO := getLogs("INFO", _beginMark_INFO, _endMark_INFO)
	logsDEBUG := getLogs("DEBUG", _beginMark_DEBUG, _endMark_DEBUG)
	testRun.Logs = logsINFO + "\n" + logsDEBUG + "\n"
	return testRun
}

func ExecuteTest(tempDir string, regMap map[string]*test.FeatureTest, tests []test.FeatureTest, args *Arguments) []test.FeatureTest {
	setArgs(args)
	if *callbackUrlPtr != "" {
		services.BaseUrl = *callbackUrlPtr
		if *callbackPortPtr != "" {
			configPort, _ := strconv.Atoi(*callbackPortPtr)
			services.BasePort = configPort
		}
	}
	if CheckEmpty(tests) != nil {
		logging.Debug.Println("Failed to find feature file.")
		os.Exit(2)
	}
	logging.Info.Println("****************************************")
	logging.Info.Println("START OF TEST SUITE")
	LogArgs(tests, logPtr, logLevelPtr, triesPtr)
	mainStatus := 0
	status := 0
	passed := 0
	failed := 0
	tmpTriesPtr, _ := strconv.Atoi(*triesPtr)
	for i := 0; i < len(tests); i++ {
		ft := regMap[tests[i].Name]
		logging.Debug.Println("******")
		logging.Debug.Println(" Test: " + tests[i].Name)
		logging.Debug.Println("******")
		steps.TestHash = ft.Hash
		services.TestHash = ft.Hash
		logging.Debug.Println(ft.Hash)
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
			mainStatus = status
			tests[i].Result = false
			LogResult(tests[i].Name, "Not OK")
			ft.Tries += 1
			if ft.Tries < tmpTriesPtr {
				i--
			} else {
				failed++
			}
		} else {
			LogResult(tests[i].Name, "OK")
			tests[i].Result = true
			passed++
		}
		time.Sleep(5 * time.Second)
	}
	logging.Info.Println("Passed ", passed)
	logging.Info.Println("Failed ", failed)
	logging.Info.Println("...END OF TEST SUITE")
	os.RemoveAll(tempDir)
	if mainStatus != 0 {
		logging.Debug.Println("Errors were found during last tests.")
	}
	return tests
}

func setArgs(args *Arguments) {
	configPtr = &args.Config
	logPtr = &args.Log
	logLevelPtr = &args.LogLevel
	triesPtr = &args.NTries
	testPtr = &args.Test
	callbackUrlPtr = &args.Url
	callbackPortPtr = &args.Port
}

func GetListOfSteps(s string) []string {
	lines := strings.Split(s, "\n")
	return lines
}

func NewTempDir() string {
	test := ""
	logging.Debug.Println(test)

	tempDir, err := ioutil.TempDir(TempRootDir, "~ctlang")
	if err != nil {
		log.Fatal(err)
	}
	return tempDir
}

func NewTempFile(tempDir, featureName string) *os.File {
	tempFile, err := ioutil.TempFile(tempDir, "~test."+featureName+".*.feature")
	if err != nil {
		logging.Debug.Println(err.Error())
	}
	return tempFile
}

func AddLineToTemp(tmp *os.File, line string) {
	if _, err := tmp.Write([]byte(line + "\n")); err != nil {
		logging.Debug.Println("Failed to write to temporary file", err)
	}
}

func GetFeatureName(id string) string {
	return "TEST"
}

func InitLoggers(logPtr *string) {
	infoLogFile, err := os.OpenFile(NameLogFile("INFO", *logPtr), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	logging.Info = log.New(infoLogFile, "INFO: ", log.LstdFlags|log.Lmsgprefix)
	debugLogFile, err := os.OpenFile(NameLogFile("DEBUG", *logPtr), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	logging.Debug = log.New(debugLogFile, "DEBUG: ", log.LstdFlags|log.Lmsgprefix)
}

func NameLogFile(level, rootName string) string {
	re := regexp.MustCompile(`(?P<RootName>(.*))(?P<Extension>\..*)`)
	matches := re.FindStringSubmatch(rootName)
	rootIndex := re.SubexpIndex("RootName")
	namePrefix := strings.TrimSpace(matches[rootIndex])
	extensionIndex := re.SubexpIndex("Extension")
	extension := strings.TrimSpace(matches[extensionIndex])
	return namePrefix + "_" + level + extension
}

func CheckEmpty(tests []test.FeatureTest) error {
	if len(tests) == 0 {
		return fmt.Errorf("File not found exception.")
	}
	return nil
}

func LogResult(test, result string) {
	logging.Info.Printf("* Feature/Scenario: %s - Status: %s\n", strings.ToUpper(test), result)
}

func LogArgs(tests []test.FeatureTest, logPtr *string, logLevelPtr *string, triesPtr *string) {
	logging.Info.Println("Config:", *configPtr)
	logging.Info.Println("Number of Tries:", *triesPtr)
	logging.Info.Println("Log:", *logPtr)
	var tsts = ""
	for _, t := range tests {
		tsts += "[" + t.Name + "]"
	}
	logging.Info.Println("Tests:", tsts)
	logging.Info.Println(".........................................")
	logging.Debug.Println("************************************************")
	logging.Debug.Println("*** Config:", *configPtr)
	logging.Debug.Println("*** Number of Tries:", *triesPtr)
	logging.Debug.Println("*** Log:", *logPtr)
	logging.Debug.Println("*** Logging Level:", *logLevelPtr)
	logging.Debug.Printf("*** Features: %v\n", tsts)
	logging.Debug.Println("************************************************")
}
