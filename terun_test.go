package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestShouldInitFileYML(t *testing.T) {
	tempDir := t.TempDir()
	terun := createTerun(tempDir)

	err := terun.Init()

	if err != nil {
		t.Error("Error on init the file: " + err.Error())
	}

	file, _ := os.ReadFile(filepath.Join(tempDir, "terun.yml"))
	currentDir, _ := os.Getwd()
	expectedFileContent, _ := os.ReadFile(filepath.Join(currentDir, "assets", "base-terun.yml"))

	if string(file) != string(expectedFileContent) {
		t.Error("The content are different")
	}
}

type ArgsMock struct{}

func (a *ArgsMock) ReadGlobalArg(argKey string) string {
	return "FastPerson"
}
func (a *ArgsMock) ReadLocalArg(argKey string) string {
	return ""
}

func TestShouldTransportFile(t *testing.T) {
	tempDir := t.TempDir()
	terun := Terun{
		Configuration: createConfiguration(tempDir, "terun.yml"),
		ArgsReader:    &ArgsMock{},
	}

	terun.Init()

	// Create controller.template inside the temp dir
	currentDir, _ := os.Getwd()
	sourceFile, _ := os.ReadFile(filepath.Join(currentDir, "assets", "test_controller.template"))
	os.WriteFile(filepath.Join(tempDir, "controller.template"), sourceFile, 0644)

	err := terun.Make("example")
	if err != nil {
		t.Error("Error on transport the files: " + err.Error())
	}

	transportedFile, _ := os.ReadFile(filepath.Join(tempDir, "fast_person_controller.py"))

	if string(transportedFile) != "class FastPersonEntity{\n    constructor(){}\n}" {
		t.Error("Content transported are different")
	}
}
