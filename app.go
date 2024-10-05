package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
)

type App struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Version  string `json:"version"`
	Size     int    `json:"size"`
	State    string `json:"state"`
}

func logs(id string) {
	_, b := get("/app/" + id + "/logs")
	fmt.Println(string(b))
}

func image(id, image string) {
	body := []byte(fmt.Sprintf(`{"image": "%s"}`, image))
	res, b := put("/app/"+id, body)
	if res != 200 {
		fmt.Println(string(b))
	}
}

func listApps() {
	var apps []App
	_, b := get("/app")
	json.Unmarshal(b, &apps)
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, minwidth, tabwidth, padding, '\t', 0)
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\n",
		"ID",
		"NAME",
		"STATE",
		"SIZE (MB)",
		"LOCATION")
	for _, app := range apps {
		fmt.Fprintf(w, "%s\t%s\t%s\t%d\t%s\n",
			app.Id,
			app.Name,
			app.State,
			app.Size,
			app.Location)
	}
	w.Flush()
}
func deleteApp(id string) {
	delete("/app/" + id)
}

func usageCreateApp() {
	fmt.Println(`usage: entrywan app create --name <name> --location <location> --source <source> --image <image> --size <size> --port <port>`)
}

func usageAppImage() {
	fmt.Println(`usage: entrywan app image <app id> <oci image location>`)
}

type appParams struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Source   string `json:"source"`
	Image    string `json:"image"`
	Size     int    `json:"size"`
	Port     int    `json:"port"`
}

func createApp() {
	opts := parseArgs()
	params := appParams{}
	if opts["name"] != "" {
		params.Name = opts["name"]
	}
	if opts["location"] != "" {
		params.Location = opts["location"]
	}
	if opts["source"] != "" {
		params.Source = opts["source"]
	}
	if opts["image"] != "" {
		params.Image = opts["image"]
	}
	if opts["size"] != "" {
		params.Size, _ = strconv.Atoi(opts["size"])
	}
	if opts["port"] != "" {
		params.Port, _ = strconv.Atoi(opts["port"])
	}
	b, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	post("/app", b)
}

func app() {
	if len(os.Args) < 3 {
		usageApp()
		os.Exit(0)
	}
	if os.Args[2] == "list" || os.Args[2] == "ls" {
		listApps()
	} else if os.Args[2] == "delete" || os.Args[2] == "rm" {
		if len(os.Args) != 4 {
			usageApp()
		} else {
			deleteApp(os.Args[3])
		}
	} else if os.Args[2] == "logs" {
		if len(os.Args) != 4 {
			usageApp()
		} else {
			logs(os.Args[3])
		}
	} else if os.Args[2] == "image" {
		if len(os.Args) != 5 {
			usageAppImage()
		} else {
			image(os.Args[3], os.Args[4])
		}
	} else if os.Args[2] == "create" {
		if len(os.Args) < 4 {
			usageCreateApp()
		} else {
			createApp()
		}
	} else {
		usageApp()
	}
}
