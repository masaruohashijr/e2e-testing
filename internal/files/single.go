package files

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func NewSingleFile(tests []string) {
	tempDir, err := ioutil.TempDir(featuresDir, "~ctlang*")
	if err != nil {
		log.Fatal(err)
	}
	tempFile, err := ioutil.TempFile(tempDir, "~test.*.ctlang")
	if err != nil {
		log.Fatal(err)
	}
	mergeFeatureFiles(tempFile, featuresDir, tests)
	tempFile.Write(finalFileContent.Bytes())
	println(tempFile.Name())
	tempFile.Close()
}

func mergeFeatureFiles(tempFile *os.File, path string, tests []string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if f.IsDir() {
			if !containsDirWithTest(tests, f.Name()) {
				continue
			}
			spath := path + "/" + f.Name()
			mergeFeatureFiles(tempFile, spath, tests)
		} else if strings.HasSuffix(f.Name(), ".feature") {
			mergeFeatureFileToTempFile(tempFile, path, f)
		}
	}
}

func containsDirWithTest(tests []string, name string) bool {
	for _, v := range tests {
		if name == v {
			println(v)
			return true
		}
	}
	return false
}

func mergeFeatureFileToTempFile(tempFile *os.File, path string, f fs.FileInfo) {
	subFileContent, _ := ioutil.ReadFile(path + "/" + f.Name())
	fmt.Println(string(subFileContent))
	finalFileContent.Write(subFileContent)
	finalFileContent.Write([]byte("\n\n\n"))
	fmt.Println("finalFileContent:", len(finalFileContent.Bytes()))
}
