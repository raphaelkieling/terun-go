package main

import (
	"embed"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

//go:embed assets/*
var assets embed.FS

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
	BasePATH       string
	TerunFile      string
	TerunAssetFile string
	Commands       map[string]CommandDefinition `yml:"commands"`
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

func (t *Configuration) createConfigurationFile() error {
	fileContent, err := assets.ReadFile("assets/" + t.TerunAssetFile)
	if err != nil {
		return err
	}

	terunBasePath := filepath.Join(t.BasePATH, "terun.yml")
	err = os.WriteFile(terunBasePath, []byte(fileContent), 0644)
	if err != nil {
		return err
	}

	return nil
}

func (t *Configuration) readFile(path string) (string, error) {
	dat, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(dat), nil
}

func (t *Configuration) writeFile(path string, content []byte) error {
	directorypath := filepath.Dir(path)
	err := os.MkdirAll(directorypath, os.ModePerm)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, content, 0644)
	if err != nil {
		return err
	}

	return nil
}

func createConfiguration(basePath string, terunFile string, terunAssetFile string) *Configuration {
	return &Configuration{
		BasePATH:       basePath,
		TerunFile:      terunFile,
		TerunAssetFile: terunAssetFile,
	}
}
