package files

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var finalFileContent bytes.Buffer
var featuresDir = "./features"
var testDir = "./test"

func NewTempFiles(testFile string) []*os.File {
	tempDir := NewTempDir()
	lines, _ := ReadFileAsLines(testFile)
	var tmp *os.File
	var tmps []*os.File
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
	return tmps
}

func NewTempDir() string {
	test := ""
	println(test)

	tempDir, err := ioutil.TempDir(testDir, "~ctlang")
	if err != nil {
		log.Fatal(err)
	}
	return tempDir
}

func NewTempFile(tempDir, featureName string) *os.File {
	tempFile, err := ioutil.TempFile(tempDir, "~test."+featureName+".*.feature")
	if err != nil {
		println(err.Error())
	}
	return tempFile
}

func AddLineToTemp(tmp *os.File, line string) {
	if _, err := tmp.Write([]byte(line + "\n")); err != nil {
		fmt.Println("Failed to write to temporary file", err)
	}
}
