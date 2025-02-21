package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
)

type Cluster struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Version  string `json:"version"`
	Cni      string `json:"cni"`
	Size     int    `json:"size"`
	State    string `json:"state"`
}

func kubeconfig(id string) {
	_, b := get("/cluster/" + id + "/kubeconfig")
	fmt.Printf(string(b))
}

func scale(id string, size int) {
	body := []byte(fmt.Sprintf(`{"size": %d}`, size))
	res, b := put("/cluster/"+id+"/scale", body)
	if res != 200 {
		fmt.Println(string(b))
	}
}

func listClusters() {
	var clusters []Cluster
	_, b := get("/cluster")
	json.Unmarshal(b, &clusters)
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, minwidth, tabwidth, padding, '\t', 0)
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\n",
		"ID",
		"NAME",
		"STATE",
		"SIZE",
		"LOCATION",
		"VERSION")
	for _, cluster := range clusters {
		fmt.Fprintf(w, "%s\t%s\t%s\t%d\t%s\t%s\n",
			cluster.Id,
			cluster.Name,
			cluster.State,
			cluster.Size,
			cluster.Location,
			cluster.Version)
	}
	w.Flush()
}

func deleteCluster(id string) {
	delete("/cluster/" + id)
}

func usageCreateCluster() {
	fmt.Println(`usage: entrywan cluster create --name <name> --location <location> --version <version> --cni <flannel|calico|cilium> --size <size>`)
}

type clusterParams struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	Version  string `json:"version"`
	Cni      string `json:"cni"`
	Size     int    `json:"size"`
}

func createCluster() {
	opts := parseArgs()
	params := clusterParams{}
	if opts["name"] != "" {
		params.Name = opts["name"]
	}
	if opts["location"] != "" {
		params.Location = opts["location"]
	}
	if opts["version"] != "" {
		params.Version = opts["version"]
	}
	if opts["cni"] != "" {
		params.Cni = opts["cni"]
	}
	if opts["size"] != "" {
		params.Size, _ = strconv.Atoi(opts["size"])
	}
	b, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	post("/cluster", b)
}

func cluster() {
	if len(os.Args) < 3 {
		usageCluster()
		os.Exit(0)
	}
	if os.Args[2] == "list" || os.Args[2] == "ls" {
		listClusters()
	} else if os.Args[2] == "delete" || os.Args[2] == "rm" {
		if len(os.Args) != 4 {
			usageCluster()
		} else {
			deleteCluster(os.Args[3])
		}
	} else if os.Args[2] == "kubeconfig" {
		if len(os.Args) != 4 {
			usageCluster()
		} else {
			kubeconfig(os.Args[3])
		}
	} else if os.Args[2] == "scale" {
		if len(os.Args) != 5 {
			usageCluster()
		} else {
			size, err := strconv.Atoi(os.Args[4])
			if err != nil {
				fmt.Println("unable to convert size to integer")
				os.Exit(1)
			}
			scale(os.Args[3], size)
		}
	} else if os.Args[2] == "create" {
		if len(os.Args) < 4 {
			usageCreateCluster()
		} else {
			createCluster()
		}
	} else {
		usageCluster()
	}
}
