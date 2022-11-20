package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type CommandDefinition struct {
	Args       []string `yml:"args"`
	Transports []struct {
		Name string   `yml:"name"`
		From string   `yml:"from"`
		To   string   `yml:"to"`
		Args []string `yml:"args"`
	} `yml:"transports"`
}

type Configuration struct {
	BasePATH  string
	TerunFile string
	Commands  map[string]CommandDefinition `yml:"commands"`
}

func (c *Configuration) getConfigurationYMLToCommandDefinition() (*Configuration, error) {
	yamlFile, err := ioutil.ReadFile(filepath.Join(c.BasePATH, c.TerunFile))
	if err != nil {
		return c, err
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return c, err
	}

	return c, nil
}

func (c *Configuration) getTransportFullPath(rest string) string {
	return filepath.Join(c.BasePATH, rest)
}

func (t *Configuration) readFile(path string) (string, error) {
	dat, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(dat), nil
}

func createConfiguration(basePath string, terunFile string) *Configuration {
	return &Configuration{
		BasePATH:  basePath,
		TerunFile: terunFile,
	}
}
