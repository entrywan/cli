package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

type Loadbalancer struct {
	Id        string     `json:"id"`
	Name      string     `json:"name"`
	Ip        string     `json:"ip"`
	Protocol  string     `json:"protocol"`
	Location  string     `json:"location"`
	Listeners []listener `json:"listeners"`
}

func listLoadbalancers() {
	var loadbalancers []Loadbalancer
	_, b := get("/loadbalancer")
	json.Unmarshal(b, &loadbalancers)
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, minwidth, tabwidth, padding, '\t', 0)
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\n",
		"ID",
		"NAME",
		"PROTOCOL",
		"IP",
		"LOCATION",
		"LISTENERS")
	for _, loadbalancer := range loadbalancers {
		var ls string
		for _, l := range loadbalancer.Listeners {
			ls += strconv.Itoa(l.Port) + "->"
			for i, tr := range l.Targets {
				ls += tr.Ip + ":" + strconv.Itoa(tr.Port)
				if i != len(l.Targets)-1 {
					ls += ","
				}
			}
		}
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\n",
			loadbalancer.Id,
			loadbalancer.Name,
			loadbalancer.Protocol,
			loadbalancer.Ip,
			loadbalancer.Location,
			ls)
	}
	w.Flush()
}

func deleteLoadbalancer(id string) {
	delete("/loadbalancer/" + id)
}

func usageCreateLoadbalancer() {
	fmt.Println(`usage: entrywan loadbalancer create --name <name> --location <location> --protocol <protocol> --port <port> --targets <host:port,host:port,host:port`)
}

type target struct {
	Ip   string `json:"ip"`
	Port int    `json:"port"`
}

type listener struct {
	Port    int      `json:"port"`
	Targets []target `json:"targets"`
}

type loadbalancerParams struct {
	Name      string     `json:"name"`
	Location  string     `json:"location"`
	Protocol  string     `json:"protocol"`
	Listeners []listener `json:"listeners"`
}

func createLoadbalancer() {
	opts := parseArgs()
	params := loadbalancerParams{}
	if opts["name"] != "" {
		params.Name = opts["name"]
	}
	if opts["location"] != "" {
		params.Location = opts["location"]
	}
	if opts["protocol"] != "" {
		params.Protocol = opts["protocol"]
	}
	l := listener{}
	if opts["port"] != "" {
		port, _ := strconv.Atoi(opts["port"])
		l.Port = port
	}
	if opts["targets"] != "" {
		tr := target{}
		ts := strings.Split(opts["targets"], ",")
		for _, t := range ts {
			ta := strings.Split(t, ":")
			tr.Ip = ta[0]
			port, err := strconv.Atoi(ta[1])
			if err != nil {
				fmt.Println("enable to parse load balancer target port")
				os.Exit(1)
			}
			tr.Port = port
			l.Targets = append(l.Targets, tr)
		}
	}
	params.Listeners = []listener{l}
	b, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	post("/loadbalancer", b)
}

func loadbalancer() {
	if len(os.Args) < 3 {
		usageLoadbalancer()
		os.Exit(0)
	}
	if os.Args[2] == "list" || os.Args[2] == "ls" {
		listLoadbalancers()
	} else if os.Args[2] == "delete" || os.Args[2] == "rm" {
		if len(os.Args) != 4 {
			usageLoadbalancer()
		} else {
			deleteLoadbalancer(os.Args[3])
		}
	} else if os.Args[2] == "create" {
		if len(os.Args) < 4 {
			usageCreateLoadbalancer()
		} else {
			createLoadbalancer()
		}
	} else {
		usageLoadbalancer()
	}
}
