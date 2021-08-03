package files

import (
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"zarbat_test/internal/godog/test"
	"zarbat_test/internal/steps"
)

func ReadFileAsLines(testFile string) (lines []string, err error) {
	f, err := ioutil.ReadFile(testFile)
	if err != nil {
		return nil, err
	}
	s := string(f)
	lines = strings.Split(s, "\n")
	return lines, nil
}

func ExtractFeature(line string) (feature string) {
	re := regexp.MustCompile(`(?P<Feature>(.*)Feature:)(?P<FeatureDescription>.*)`)
	matches := re.FindStringSubmatch(line)
	featureIndex := re.SubexpIndex("FeatureDescription")
	feature = strings.TrimSpace(matches[featureIndex])
	return feature
}

func NewFeatureTests(tempFiles []*os.File, regMap map[string]*test.FeatureTest) (fts []test.FeatureTest) {
	ft := &test.FeatureTest{}
	for _, t := range tempFiles {
		key := strings.ToLower(extractFeatureKeyFromFileName(t.Name()))
		if _, ok := regMap[key]; !ok {
			ft = &test.FeatureTest{
				Name:                key,
				Path:                t.Name(),
				Hash:                test.Hash(key),
				ScenarioInitializer: steps.InitializeScenario,
			}
			fts = append(fts, *ft)
			regMap[key] = ft
		}
	}
	return
}

func extractFeatureKeyFromFileName(fName string) (feature string) {
	//~test.Dial.450292574.feature
	re := regexp.MustCompile(`(?P<Feature>~test.)(?P<FeatureDescription>.*)(?P<FeatureExtension>\.{1}\d+.feature)`)
	matches := re.FindStringSubmatch(fName)
	featureIndex := re.SubexpIndex("FeatureDescription")
	feature = strings.TrimSpace(matches[featureIndex])
	return
}

func GetFeatureTestsFromMap(tests []string, regMap map[string]*test.FeatureTest) (fts []test.FeatureTest) {
	for _, t := range tests {
		ft := regMap[t]
		fts = append(fts, *ft)
	}
	return
}
