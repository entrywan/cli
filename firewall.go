package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

type rule struct {
	Protocol string `json:"protocol"`
	Port     string `port:"port"`
	Src      string `json:"src"`
}

type Firewall struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Rules []rule `json:"rules"`
}

func listFirewalls() {
	var firewalls []Firewall
	_, b := get("/firewall")
	json.Unmarshal(b, &firewalls)
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, minwidth, tabwidth, padding, '\t', 0)
	fmt.Fprintf(w, "%s\t%s\t%s\t\n",
		"ID",
		"NAME",
		"RULES")
	for _, firewall := range firewalls {
		var rules []string = []string{}
		for _, r := range firewall.Rules {
			var rparams []string
			if r.Protocol != "" {
				rparams = append(rparams, r.Protocol)
			}
			if r.Port != "" {
				rparams = append(rparams, r.Port)
			}
			if r.Src != "" {
				rparams = append(rparams, r.Src)
			}
			rules = append(rules, strings.Join(rparams, ":"))
		}
		fmt.Fprintf(w, "%s\t%s\t%s\n",
			firewall.Id,
			firewall.Name,
			strings.Join(rules, ", "))
	}
	w.Flush()
}

func deleteFirewall(id string) {
	delete("/firewall/" + id)
}

func usageCreateFirewall() {
	fmt.Println(`usage: entrywan firewall create [--name <name>] --protocol <protocol> [--port <port>] [--src <src address>]`)
}

type firewallParams struct {
	Name  string `json:"name"`
	Rules []rule `json:rules`
}

func createFirewall() {
	opts := parseArgs()
	params := firewallParams{}
	r := rule{}
	if opts["name"] != "" {
		params.Name = opts["name"]
	}
	if opts["protocol"] != "" {
		r.Protocol = opts["protocol"]
	}
	if opts["port"] != "" {
		r.Port = opts["port"]
	}
	if opts["src"] != "" {
		r.Src = opts["src"]
	}
	params.Rules = []rule{r}
	b, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	post("/firewall", b)
}

func firewall() {
	if len(os.Args) < 3 {
		usageFirewall()
		os.Exit(0)
	}
	if os.Args[2] == "list" || os.Args[2] == "ls" {
		listFirewalls()
	} else if os.Args[2] == "delete" {
		if len(os.Args) != 4 {
			usageFirewall()
		} else {
			deleteFirewall(os.Args[3])
		}
	} else if os.Args[2] == "create" {
		if len(os.Args) < 4 {
			usageCreateFirewall()
		} else {
			createFirewall()
		}
	} else {
		usageFirewall()
	}
}
