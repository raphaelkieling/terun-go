package main

import "fmt"

// ArgsReader interface
type ArgsReader interface {
	ReadGlobalArg(argKey string) string
	ReadLocalArg(argKey string) string
}

// Using the console reader
type ArgsConsole struct{}

func (a *ArgsConsole) ReadGlobalArg(argKey string) string {
	fmt.Printf("🌍 Enter global arg \"%s\" value: ", argKey)
	var contentArg string
	fmt.Scanf("%s", &contentArg)

	return contentArg
}

func (a *ArgsConsole) ReadLocalArg(argKey string) string {
	fmt.Printf("	Enter local arg \"%s\" value: ", argKey)
	var contentArg string
	fmt.Scanf("%s", &contentArg)

	return contentArg
}

func createArgsConsole() *ArgsConsole {
	return &ArgsConsole{}
}
