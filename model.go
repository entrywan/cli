package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
)

type Model struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Size     int    `json:"size"`
	State    string `json:"state"`
	Type     string `json:"type"`
	Token    string `json:"token"`
	Endpoint string `json:"endpoint"`
}

func listModels() {
	var models []Model
	_, b := get("/model")
	json.Unmarshal(b, &models)
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, minwidth, tabwidth, padding, '\t', 0)
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\n",
		"ID",
		"NAME",
		"ENDPOINT",
		"TYPE",
		"TOKEN",
		"STATE")
	for _, model := range models {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\n",
			model.Id,
			model.Name,
			model.Endpoint,
			model.Type,
			model.Token,
			model.State)
	}
	w.Flush()
}
func deleteModel(id string) {
	delete("/model/" + id)
}

func usageCreateModel() {
	fmt.Println(`usage: entrywan model create --name <name> --location <location> --type <type> --size <size>`)
}

type modelParams struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Size     int    `json:"size"`
	Type     string `json:"type"`
}

func createModel() {
	opts := parseArgs()
	params := modelParams{}
	if opts["name"] != "" {
		params.Name = opts["name"]
	}
	if opts["location"] != "" {
		params.Location = opts["location"]
	}
	if opts["type"] != "" {
		params.Type = opts["type"]
	}
	if opts["size"] != "" {
		params.Size, _ = strconv.Atoi(opts["size"])
	}
	b, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	post("/model", b)
}

func model() {
	if len(os.Args) < 3 {
		usageModel()
		os.Exit(0)
	}
	if os.Args[2] == "list" || os.Args[2] == "ls" {
		listModels()
	} else if os.Args[2] == "delete" || os.Args[2] == "rm" {
		if len(os.Args) != 4 {
			usageModel()
		} else {
			deleteModel(os.Args[3])
		}
	} else if os.Args[2] == "create" {
		if len(os.Args) < 4 {
			usageCreateModel()
		} else {
			createModel()
		}
	} else {
		usageModel()
	}
}
