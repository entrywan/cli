package main

import (
	"os"
)

var version = "dev" // updated by release tooling

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(0)
	}
	parseConfig()
	runCommand()
}
