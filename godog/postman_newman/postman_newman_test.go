package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/cucumber/godog"
	"github.com/tidwall/gjson"
)

func allTestCasesMustPass() error {

	file, e := ioutil.ReadFile(postman_data + collection_folder_path + "/newman_results.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	myJson := string(file)
	failed := gjson.Get(myJson, "run.stats.tests.failed").String()
	total_test := gjson.Get(myJson, "run.stats.tests.total").String()

	if failed != "0" {
		return fmt.Errorf("Failed %s tests out of %s tests.", failed, total_test)
	}

	return nil
}

func runningNewmanCollection(collection_url string) error {
	cli_output = "newman run https://www.getpostman.com/collections/" + collection_url + " --folder " + collection_folder_path + " -e " + postman_data + env_file_path + " -d " + postman_data + data_file_path + " --reporters json,htmlextra,cli --reporter-json-export " + postman_data + collection_folder_path + "/newman_results.json  --reporter-htmlextra-export " + postman_data + collection_folder_path + "/report.html --verbose"
	out, err := exec.Command("/bin/sh", "-c", cli_output).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", out)

	return nil
}

func theDataFileAndCollectionFolder(data_file string, collection_folder string) error {

	postman_data = "../../postman_data/"
	collection_folder_path = collection_folder
	data_file_path = collection_folder_path + "/" + data_file
	return nil
}

func theEnvFile(env_file string) error {
	env_file_path = env_file
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^All test cases must pass$`, allTestCasesMustPass)
	ctx.Step(`^Running Newman Collection "([^"]*)"$`, runningNewmanCollection)
	ctx.Step(`^The data file "([^"]*)" and collection folder "([^"]*)"$`, theDataFileAndCollectionFolder)
	ctx.Step(`^The env File "([^"]*)"$`, theEnvFile)
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
}

func TestMain(m *testing.M) {
	opts := godog.Options{
		Format:    "progress",
		Paths:     []string{"../../features/postman_newman"},
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
