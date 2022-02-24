package files

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"zarbat_test/internal/logging"
)

var finalFileContent bytes.Buffer
var featuresDir = "./features"
var TempRootDir = "./temp"
var testDir = "./test"

func NewTempFiles(testFile string) (tmps []*os.File, tempDir string) {
	tempDir = NewTempDir()
	lines, _ := ReadFileAsLines(testFile)
	var tmp *os.File
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Feature:") {
			// TODO Tratar espaço em branco no meio da feature
			featureName := ExtractFeature(line)
			tmp = NewTempFile(tempDir, featureName)
			tmps = append(tmps, tmp)
		}
		AddLineToTemp(tmp, line)
	}
	return tmps, tempDir
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
