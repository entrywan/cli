package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

var (
	config Config
)

type Config struct {
	Token    string // required
	Endpoint string // optional
}

func parseConfig() {
	home := os.Getenv("HOME")
	f := home + "/.config/entrywan/config.toml"
	if _, err := os.Stat(f); err != nil {
		fmt.Fprintf(os.Stderr,
			"error: no config file found at %s\n", f)
		os.Exit(1)
	}
	_, err := toml.DecodeFile(f, &config)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if config.Endpoint == "" {
		config.Endpoint = "https://api.entrywan.com/v1beta"
	}
	if config.Token == "" {
		fmt.Fprintln(os.Stderr,
			"error: unable to find token value in config")
		os.Exit(1)
	}
}
