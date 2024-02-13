package main

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"
)

type Sshkey struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

func listSshkeys() {
	var sshkeys []Sshkey
	_, b := get("/sshkey")
	json.Unmarshal(b, &sshkeys)
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, minwidth, tabwidth, padding, '\t', 0)
	fmt.Fprintf(w, "%s\t%s\t%s\n",
		"ID",
		"NAME",
		"DESCRIPTION")
	for _, sshkey := range sshkeys {
		fmt.Fprintf(w, "%s\t%s\t%s\n",
			sshkey.Id,
			sshkey.Name,
			sshkey.Comment)
	}
	w.Flush()
}

func deleteSshkey(id string) {
	delete("/sshkey/" + id)
}

func usageCreateSshkey() {
	fmt.Println(`usage: entrywan sshkey create --name <name> --pub <pub>`)
}

type sshkeyParams struct {
	Name string `json:"name"`
	Pub  string `json:"pub"`
}

func createSshkey() {
	opts := parseArgs()
	params := sshkeyParams{}
	if opts["name"] != "" {
		params.Name = opts["name"]
	}
	if opts["pub"] != "" {
		params.Pub = opts["pub"]
	}
	b, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	post("/sshkey", b)
}

func sshkey() {
	if len(os.Args) < 3 {
		usageSshkey()
		os.Exit(0)
	}
	if os.Args[2] == "list" || os.Args[2] == "ls" {
		listSshkeys()
	} else if os.Args[2] == "delete" {
		deleteSshkey(os.Args[3])
	} else if os.Args[2] == "create" {
		if len(os.Args) < 4 {
			usageCreateSshkey()
		} else {
			createSshkey()
		}
	} else {
		usageSshkey()
	}
}
