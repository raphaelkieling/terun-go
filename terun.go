package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"golang.org/x/exp/maps"
)

type Terun struct {
	Configuration *Configuration
}

func (t *Terun) getCommand(command string) (CommandDefinition, error) {
	all_configuration, err := t.Configuration.getConfigurationYMLToCommandDefinition()
	return all_configuration.Commands[command], err
}

func (t *Terun) Init() error {
	path, err := os.Getwd()
	if err != nil {
		return err
	}

	currentBaseTerunFile := filepath.Join(path, "./assets/base-terun.yml")
	fileContent, err := t.Configuration.readFile(currentBaseTerunFile)
	if err != nil {
		return err
	}

	terunBasePath := filepath.Join(t.Configuration.BasePATH, "terun.yml")
	err = os.WriteFile(terunBasePath, []byte(fileContent), 0644)
	if err != nil {
		return err
	}

	fmt.Println("terun.yml created inside the folder! Happy generation 😃")

	return nil
}

func (t *Terun) Make(command string) error {
	fmt.Printf("🧰 Executing command: %s\n", command)
	defaultTemplate := template.New("worker")

	// 1 - Get command
	commandItem, err := t.getCommand(command)
	if err != nil {
		return err
	}

	// 2 - Request global arguments
	argumentsStore := make(map[string]string)
	for _, arg := range commandItem.Args {
		fmt.Printf("🌍 Enter global arg \"%s\" value: ", arg)
		var contentArg string
		fmt.Scanf("%s", &contentArg)

		argumentsStore[arg] = contentArg
	}

	// 3 - Go through each transport
	for _, transport := range commandItem.Transports {
		localStore := make(map[string]string)
		maps.Copy(localStore, argumentsStore)

		fmt.Printf("📦 Reading: %s\n", transport.Name)
		// 3.1 - Request the arguments
		for _, arg := range transport.Args {
			fmt.Printf("	Enter local arg \"%s\" value: ", arg)
			var contentArg string
			fmt.Scanf("%s", &contentArg)

			localStore[arg] = contentArg
		}

		// 3.2 - Read from file
		var outputFromPath bytes.Buffer
		defaultTemplate.Parse(transport.From)
		_ = defaultTemplate.Execute(&outputFromPath, localStore)

		fromFilePath := t.Configuration.getTransportFullPath(outputFromPath.String())
		fileContent, err := t.Configuration.readFile(fromFilePath)
		if err != nil {
			return err
		}

		// 3.3 - Build from `from` content file
		var outputFromContent bytes.Buffer
		defaultTemplate.Parse(fileContent)
		_ = defaultTemplate.Execute(&outputFromContent, localStore)

		// 3.4 - Transport the result for the `to` property
		var outputToPath bytes.Buffer
		defaultTemplate.Parse(transport.To)
		_ = defaultTemplate.Execute(&outputToPath, localStore)

		toFilePath := t.Configuration.getTransportFullPath(outputToPath.String())
		os.WriteFile(toFilePath, outputFromContent.Bytes(), 0644)
		fmt.Println("	✅ Done!")

	}

	fmt.Println("✅ Done all!")

	return nil
}

func createTerun(path string) *Terun {
	return &Terun{
		Configuration: createConfiguration(path, "terun.yml"),
	}
}
