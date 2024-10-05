package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
)

type Instance struct {
	Id       string `json:"id"`
	Hostname string `json:"hostname"`
	Ip4      string `json:"ip4"`
	State    string `json:"state"`
	Location string `json:"location"`
}

func listInstances() {
	var instances []Instance
	_, b := get("/instance")
	json.Unmarshal(b, &instances)
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, minwidth, tabwidth, padding, '\t', 0)
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n",
		"ID",
		"NAME",
		"STATE",
		"IPV4 ADDRESS",
		"LOCATION")
	for _, instance := range instances {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n",
			instance.Id,
			instance.Hostname,
			instance.State,
			instance.Ip4,
			instance.Location)
	}
	w.Flush()
}

func deleteInstance(id string) {
	delete("/instance/" + id)
}

func usageCreateInstance() {
	fmt.Println(`usage: entrywan instance create --hostname <hostname> --location <location> --os <os> --sshkey <sshkeyname> --cpus <cpus> --ram <ram> --disk <disk>`)
}

type instanceParams struct {
	Hostname string `json:"hostname"`
	Location string `json:"location"`
	Os       string `json:"os"`
	Sshkey   string `json:"sshkeyname"`
	Cpus     int    `json:"cpus"`
	Ram      int    `json:"ram"`
	Disk     int    `json:"disk"`
}

func createInstance() {
	opts := parseArgs()
	params := instanceParams{}
	if opts["hostname"] != "" {
		params.Hostname = opts["hostname"]
	}
	if opts["location"] != "" {
		params.Location = opts["location"]
	}
	if opts["os"] != "" {
		params.Os = opts["os"]
	}
	if opts["sshkey"] != "" {
		params.Sshkey = opts["sshkey"]
	}
	if opts["cpus"] != "" {
		params.Cpus, _ = strconv.Atoi(opts["cpus"])
	}
	if opts["ram"] != "" {
		params.Ram, _ = strconv.Atoi(opts["ram"])
	}
	if opts["disk"] != "" {
		params.Disk, _ = strconv.Atoi(opts["disk"])
	}
	b, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	post("/instance", b)
}

func instance() {
	if len(os.Args) < 3 {
		usageInstance()
		os.Exit(0)
	}
	if os.Args[2] == "list" || os.Args[2] == "ls" {
		listInstances()
	} else if os.Args[2] == "delete" || os.Args[2] == "rm" {
		if len(os.Args) != 4 {
			usageInstance()
		} else {
			deleteInstance(os.Args[3])
		}
	} else if os.Args[2] == "create" {
		if len(os.Args) < 4 {
			usageCreateInstance()
		} else {
			createInstance()
		}
	} else {
		usageInstance()
	}
}
