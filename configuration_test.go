package main

import "testing"

func TestShouldReadTerunFile(t *testing.T) {
	conf := createConfiguration("./assets", "base-terun.yml")
	currConf, _ := conf.getConfigurationYMLToCommandDefinition()

	if len(currConf.Commands) == 0 {
		t.Error("Do not have enough commands")
	}
}

func TestShouldThrowErrorIfDoNotExistFile(t *testing.T) {
	conf := createConfiguration("./assets-error", "base-terun.yml")
	_, err := conf.getConfigurationYMLToCommandDefinition()

	if err == nil {
		t.Error("Must have a error")
	}
}
