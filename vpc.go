package main

import (
	"encoding/json"
	"fmt"
	"os"
	"text/tabwriter"
)

type Member struct {
	Ippublic  string `json:"ippublic"`
	Ipprivate string `json:"ipprivate"`
}

type Vpc struct {
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Prefix  string   `json:"prefix"`
	Members []Member `json:"members"`
}

func listVpcs() {
	var vpcs []Vpc
	_, b := get("/vpc")
	json.Unmarshal(b, &vpcs)
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, minwidth, tabwidth, padding, '\t', 0)
	fmt.Fprintf(w, "%s\t%s\t%s\t%s\n",
		"ID",
		"NAME",
		"PREFIX",
		"MEMBERS")
	for _, vpc := range vpcs {
		var members string
		if vpc.Members == nil {
			members = "none"
		} else {
			members = ""
			for _, member := range vpc.Members {
				members += member.Ippublic + "/" + member.Ipprivate + " "
			}
		}
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n",
			vpc.Id,
			vpc.Name,
			vpc.Prefix,
			members)
	}
	w.Flush()
}

func deleteVpc(id string) {
	delete("/vpc/" + id)
}

func usageCreateVpc() {
	fmt.Println(`usage: entrywan vpc create --name <name> --prefix <prefix>`)
}

func usageAddVpcMember() {
	fmt.Println(`usage: entrywan vpc add --vpc <name> --ippublic <ip> [--ipprivate <ip>]`)
}

func usageRemoveVpcMember() {
	fmt.Println(`usage: entrywan vpc remove --vpc <name> --ipprivate <ip>`)
}

type vpcParams struct {
	Name   string `json:"name"`
	Prefix string `json:"prefix"`
}

func createVpc() {
	opts := parseArgs()
	params := vpcParams{}
	if opts["name"] != "" {
		params.Name = opts["name"]
	}
	if opts["prefix"] != "" {
		params.Prefix = opts["prefix"]
	}
	b, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	post("/vpc", b)
}

type vpcAddParams struct {
	Vpc       string `json:"vpc"`
	IpPublic  string `json:"ip4public"`
	IpPrivate string `json:"ip4private"`
}

func addVpcMember() {
	opts := parseArgs()
	params := vpcAddParams{}
	if opts["vpc"] != "" {
		params.Vpc = opts["vpc"]
	}
	if opts["ippublic"] != "" {
		params.IpPublic = opts["ippublic"]
	}
	if opts["ipprivate"] != "" {
		params.IpPrivate = opts["ipprivate"]
	}
	b, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	put("/vpc/"+params.Vpc, b)
}

type vpcRemoveParams struct {
	Vpc       string `json:"vpc"`
	IpPrivate string `json:"ip4private"`
}

func removeVpcMember() {
	opts := parseArgs()
	params := vpcRemoveParams{}
	if opts["vpc"] != "" {
		params.Vpc = opts["vpc"]
	}
	if opts["ipprivate"] != "" {
		params.IpPrivate = opts["ipprivate"]
	}
	b, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	patch("/vpc/"+params.Vpc, b)
}

func vpc() {
	if len(os.Args) < 3 {
		usageVpc()
		os.Exit(0)
	}
	if os.Args[2] == "list" || os.Args[2] == "ls" {
		listVpcs()
	} else if os.Args[2] == "delete" {
		if len(os.Args) != 4 {
			usageVpc()
		} else {
			deleteVpc(os.Args[3])
		}
	} else if os.Args[2] == "create" {
		if len(os.Args) < 4 {
			usageCreateVpc()
		} else {
			createVpc()
		}
	} else if os.Args[2] == "add" {
		if len(os.Args) < 5 {
			usageAddVpcMember()
		} else {
			addVpcMember()
		}
	} else if os.Args[2] == "remove" {
		if len(os.Args) < 5 {
			usageRemoveVpcMember()
		} else {
			removeVpcMember()
		}
	} else {
		usageVpc()
	}
}
