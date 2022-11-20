package main

import (
	"fmt"
	"os"

	"github.com/dimiro1/banner"
)

func printBanner() {
	templ := `{{ .Title "Terun" "" 4 }}`

	banner.InitString(os.Stdout, true, true, templ)
	fmt.Println("")
}
