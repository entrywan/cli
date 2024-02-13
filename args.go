package main

import (
	"os"
	"strings"
)

func parseArgs() map[string]string {
	m := make(map[string]string)
	if len(os.Args) < 4 {
		return m
	}
	for i := 3; i < len(os.Args); i += 2 {
		m[strings.Replace(os.Args[i], "--", "", 1)] = os.Args[i+1]
	}
	return m
}
