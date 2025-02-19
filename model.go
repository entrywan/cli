package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"text/tabwriter"

	"github.com/google/uuid"
)

type Model struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
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

func queryModel(modelName string) {
	var model Model
	uid, err := uuid.Parse(modelName)
	if err != nil {
		var models []Model
		_, b := get("/model")
		json.Unmarshal(b, &models)
		for _, m := range models {
			if m.Name == modelName {
				model = m
			}
		}
	} else {
		var models []Model
		_, b := get("/model")
		json.Unmarshal(b, &models)
		for _, m := range models {
			if m.Name == uid.String() {
				model = m
			}
		}
	}

	if model.Endpoint == "" {
		fmt.Println("Can't find the model.  Please choose another one.")
		os.Exit(1)
	}

	client := &http.Client{}

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("The model is ready.")
	fmt.Println()
	fmt.Printf("> ")

	for scanner.Scan() {
		txt := scanner.Text()

		bs := bytes.NewBuffer([]byte(txt))
		req, err := http.NewRequest("POST", "https://"+model.Endpoint+".entrywan.app/text", bs)
		if err != nil {
			fmt.Fprintln(os.Stderr,
				"error: please try again later or contact support")
			os.Exit(1)
		}

		req.Header.Set("Authorization", "Bearer "+model.Token)
		req.Header.Set("User-Agent", "cli "+runtime.GOOS+"/"+runtime.GOARCH+" version "+version)

		res, err := client.Do(req)
		if err != nil {
			fmt.Fprintln(os.Stderr,
				"error: unable to connect, check your internet connection or try again later")
			os.Exit(1)
		}
		b, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Fprintln(os.Stderr,
				"error: please try again later or contact support")
			os.Exit(1)
		}
		if res.StatusCode != 200 {
			fmt.Fprintf(os.Stderr,
				"error: %s\n", string(b))
			os.Exit(1)
		}

		fmt.Println(string(b))

		fmt.Println()
		fmt.Printf("> ")
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

func usageCreateModel() {
	fmt.Println(`usage: entrywan model create --name <name> --location <location> --type <type>`)
}

type modelParams struct {
	Name     string `json:"name"`
	Location string `json:"location"`
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
	} else if os.Args[2] == "query" {
		queryModel(os.Args[3])
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
