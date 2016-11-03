package cli_test

import (
	"fmt"
	"github.com/appcelerator/amp/api/server"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"time"
	"text/template"
	"math/rand"
	"strconv"
	"bytes"
)

type TestSpec struct {
	Name     string
	Commands []CommandSpec
}

type CommandSpec struct {
	Cmd                string   `yaml:"cmd"`
	Args               []string `yaml:"args"`
	Options            []string `yaml:"options"`
	Expectation        string   `yaml:"expectation"`
	ExpectErrorStatus  bool     `yaml:"expectErrorStatus"`
	APICall            string   `yaml:"apiCall"`
	APICallExpectation string   `yaml:"apiCallExpectation"`
}

var (
	testDir = "./test_samples"
	nameMap = map[string]string{}
	fmap = template.FuncMap {
		"mapName" : mapName,
	}
)

func TestMain(m *testing.M) {
	server.StartTestServer()
	os.Exit(m.Run())
}

func TestCmds(t *testing.T) {
	tests, err := loadTestSpecs()
	if err != nil {
		t.Errorf("unable to load test specs, reason: %v", err)
		return
	}

	for _, test := range tests {
		t.Log("-----------------------------------------------------------------------------------------")
		t.Logf("Running spec: %s", test.Name)
		if err := runTestSpec(t, test); err != nil {
			t.Error(err)
			return
		}
	}
}

func loadTestSpecs() ([]*TestSpec, error) {
	files, err := ioutil.ReadDir(testDir)
	if err != nil {
		return nil, err
	}

	tests := []*TestSpec{}
	for _, file := range files {
		test, err := loadTestSpec(path.Join(testDir, file.Name()))
		if err != nil {
			return nil, err
		}
		if test != nil {
			tests = append(tests, test)
		}
	}
	return tests, nil
}

func loadTestSpec(fileName string) (*TestSpec, error) {
	if filepath.Ext(fileName) != ".yml" {
		return nil, nil
	}
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("unable to load test spec: %s. Error: %v", fileName, err)
	}
	testSpec := &TestSpec{
		Name: fileName,
	}

	var commandMap []CommandSpec
	if err := yaml.Unmarshal(content, &commandMap); err != nil {
		return nil, fmt.Errorf("unable to parse test spec: %s. Error: %v", fileName, err)
	}

	// Keep values only
	for _, command := range commandMap {
		testSpec.Commands = append(testSpec.Commands, command)
	}

	return testSpec, nil
}

func runTestSpec(t *testing.T, test *TestSpec) error {
	for _, cmdSpec := range test.Commands {
		cmdString := generateCmdString(&cmdSpec)
		t.Logf("Running: %s", strings.Join(cmdString, " "))
		cmdActualOutput, err := exec.Command(cmdString[0], cmdString[1:]...).CombinedOutput()
		cmdExpectedOutput := regexp.MustCompile(cmdSpec.Expectation)
		if !cmdExpectedOutput.MatchString(string(cmdActualOutput)) {
			return fmt.Errorf("miss matched expected output: %s", cmdActualOutput)
		}
		if err != nil && !cmdSpec.ExpectErrorStatus {
			return fmt.Errorf("Command was expected to exit with zero status but got: %v", err)
		}
		if err == nil && cmdSpec.ExpectErrorStatus {
			return fmt.Errorf("Command was expected to exit with error status but exited with zero")
		}
		if len(cmdSpec.APICall) != 0 {
			apiString := strings.Fields(cmdSpec.APICall)
			apiOutput, err := exec.Command(apiString[0], apiString[1:]...).CombinedOutput()
			apiExpectedOutput := regexp.MustCompile(cmdSpec.APICallExpectation)
			if !apiExpectedOutput.MatchString(string(apiOutput)) {
				return fmt.Errorf("API check not passed: %s", err)
			}
		}
	}
	return nil
}

func mapName(name string) string {
	if _, ok := nameMap[name]; !ok {
		fmt.Println(name)
		i := rand.Int()
		newName := name + strconv.Itoa(i)
		nameMap[name] = newName
		fmt.Println(newName)
		return newName
	}
	return nameMap[name]
}

func generateCmdString(cmdSpec *CommandSpec) (cmdString []string) {
	cmdSplit := strings.Fields(cmdSpec.Cmd)
	optionsSplit := []string{}
	for _, val := range cmdSpec.Options {
		optionsSplit = append(optionsSplit, strings.Fields(val)...)
	}
	cmdString = append(cmdSplit, cmdSpec.Args...)
	cmdString = append(cmdString, optionsSplit...)
	return
}
