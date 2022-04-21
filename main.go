package main

import (
	"flag"
	"strconv"
	"strings"
	"zarbat_test/internal/config"
	"zarbat_test/internal/files"
	"zarbat_test/internal/godog/services"
	"zarbat_test/internal/godog/test"
	"zarbat_test/server/api"
)

var RegMap map[string]*test.FeatureTest
var logPtr *string
var logLevelPtr *string
var configPtr *string
var triesPtr *string
var callbackPtr *string
var startPtr *bool
var testPtr *string
var callbackUrlPtr *string
var callbackPortPtr *string

func main() {
	RegMap = test.InitRegister()
	initArgs()
	api.InitLoggers(logPtr)
	if *startPtr {
		api.Start()
	}
	tests, tempDir := getFeaturesTests(RegMap)
	args := &api.Arguments{
		Config:   *configPtr,
		Url:      *callbackUrlPtr,
		Port:     *callbackPortPtr,
		Log:      *logPtr,
		LogLevel: *logLevelPtr,
		NTries:   *triesPtr,
		Test:     *testPtr,
	}
	api.ExecuteTest(tempDir, RegMap, tests, args)
}

func initArgs() {
	startPtr = flag.Bool("start", false, "server api command")
	configPtr = flag.String("config", "config/config.ini", "A configuration file")
	config.ConfigPath = *configPtr
	triesPtr = flag.String("n", "5", "number of tries")
	logPtr = flag.String("l", "log/zarbat.log", "log location")
	logLevelPtr = flag.String("level", "info", "options: info, debug")
	testPtr = flag.String("test", "", "ctlang")
	callbackUrlPtr = flag.String("url", "", "Public IP")
	callbackPortPtr = flag.String("port", "", "Public Port")
	flag.Parse()
	if *callbackUrlPtr != "" {
		services.BaseUrl = *callbackUrlPtr
		if *callbackPortPtr != "" {
			configPort, _ := strconv.Atoi(*callbackPortPtr)
			services.BasePort = configPort
		}
	} else {
		services.BaseUrl = config.NewConfig().BaseUrl
	}
}

func getFeaturesTests(regMap map[string]*test.FeatureTest) (fts []test.FeatureTest, tempDir string) {
	if !files.IsParametersValid(testPtr) || !files.IsParametersValid(configPtr) {
		return fts, ""
	}
	if strings.HasSuffix(*testPtr, ".ctlang") || strings.HasSuffix(*testPtr, ".feature") {
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
