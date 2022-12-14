package main

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
	"unicode"

	"golang.org/x/exp/maps"
)

type Terun struct {
	Configuration *Configuration
	ArgsReader    ArgsReader
}

func (t *Terun) getCommand(command string) (CommandDefinition, error) {
	all_configuration, err := t.Configuration.getConfigurationYMLToCommandDefinition()
	return all_configuration.Commands[command], err
}

func (t *Terun) Init() error {
	err := t.Configuration.createConfigurationFile()
	if err != nil {
		return err
	}

	fmt.Println("terun.yml created inside the folder! Happy generation 😃")

	return nil
}

func (t *Terun) createFuncMap() template.FuncMap {
	return template.FuncMap{
		"lowercase": func(value string) string {
			return strings.ToLower(value)
		},
		"uppercase": func(value string) string {
			return strings.ToUpper(value)
		},
		"underscore": func(value string) string {
			// convert every letter to lower case
			var words []string
			l := 0
			for s := value; s != ""; s = s[l:] {
				l = strings.IndexFunc(s[1:], unicode.IsUpper) + 1
				if l <= 0 {
					l = len(s)
				}
				words = append(words, s[:l])
			}

			return strings.Join(words, "_")
		},
	}
}

func (t *Terun) Make(command string) error {
	fmt.Printf("🧰 Executing command: %s\n", command)
	defaultTemplate := template.New("worker")
	defaultTemplate.Funcs(t.createFuncMap())

	// 1 - Get command
	commandItem, err := t.getCommand(command)
	if err != nil {
		return err
	}

	// 2 - Request global arguments
	argumentsStore := make(map[string]string)
	for _, arg := range commandItem.Args {
		argumentsStore[arg] = t.ArgsReader.ReadGlobalArg(arg)
	}

	// 3 - Go through each transport
	for _, transport := range commandItem.Transports {
		localStore := make(map[string]string)
		maps.Copy(localStore, argumentsStore)

		fmt.Printf("📦 Reading: %s\n", transport.Name)
		// 3.1 - Request the arguments
		for _, arg := range transport.Args {
			localStore[arg] = t.ArgsReader.ReadLocalArg(arg)
		}

		// 3.2 - Read from file
		var outputFromPath bytes.Buffer
		defaultTemplate.Parse(transport.From)
		err = defaultTemplate.Execute(&outputFromPath, localStore)
		if err != nil {
			return err
		}

		fromFilePath := t.Configuration.getTransportFullPath(outputFromPath.String())
		fileContent, err := t.Configuration.readFile(fromFilePath)
		if err != nil {
			return err
		}

		// 3.3 - Build from `from` content file
		var outputFromContent bytes.Buffer
		defaultTemplate.Parse(fileContent)
		err = defaultTemplate.Execute(&outputFromContent, localStore)
		if err != nil {
			return err
		}

		// 3.4 - Transport the result for the `to` property
		var outputToPath bytes.Buffer
		defaultTemplate.Parse(transport.To)
		err = defaultTemplate.Execute(&outputToPath, localStore)
		if err != nil {
			return err
		}

		toFilePath := t.Configuration.getTransportFullPath(outputToPath.String())
		err = t.Configuration.writeFile(toFilePath, outputFromContent.Bytes())
		if err != nil {
			return err
		}

		fmt.Println("	✅ Done!")

	}

	fmt.Println("✅ Done all!")

	return nil
}

func createTerun(basePath string) *Terun {
	return &Terun{
		Configuration: createConfiguration(basePath, "terun.yml", "base-terun.yml"),
		ArgsReader:    createArgsConsole(),
	}
}
