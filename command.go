package main

import (
	"fmt"
	"os"
)

const (
	minwidth = 8
	tabwidth = 8
	padding  = 2
)

func usage() {
	fmt.Println(`usage: entrywan <command> [<args>]

The commands are:

   instance      Manage instances
   cluster       Manage kubernetes clusters
   sshkey        Manage ssh keys
   firewall      Manage firewalls
   loadbalancer  Manage load balancers
   vpc           Manage VPC networks
   app           Manage apps
   version       Show cli version

'entrywan <subcommand>' to show usage for each subcommand.`)
}

func usageInstance() {
	fmt.Println(`usage: entrywan instance <subcommand>

The subcommands are:

    list (or ls)          List instances
    create <arguments>    Create a new instance
    delete <instance_id>  Delete an instance

'entrywan instance create' for create arguments`)
}

func usageCluster() {
	fmt.Println(`usage: entrywan cluster <command> [<args]

The subcommands are:

    list (or ls)               List clusters
    create <arguments>         Create a new cluster
    kubeconfig <cluster_id>    Fetch kubeconfig for cluster
    scale <cluster_id> <size>  Scale cluster up or down
    delete <cluster_id>        Delete an cluster

'entrywan cluster create' for create arguments`)
}

func usageApp() {
	fmt.Println(`usage: entrywan app <command> [<args]

The subcommands are:

    list (or ls)            List apps
    create <arguments>      Create a new app
    logs <app_id>           Fetch app logs
    image <app_id> <image>  Deploy new image
    delete <app_id>         Delete an app

'entrywan app create' for create arguments`)
}

func usageSshkey() {
	fmt.Println(`usage: entrywan sshkey <subcommand>

The subcommands are:

    list (or ls)        List sshkeys
    create <arguments>  Create a new sshkey
    delete <sshkey_id>  Delete an sshkey

'entrywan sshkey create' for create arguments`)
}

func usageFirewall() {
	fmt.Println(`usage: entrywan firewall <subcommand>

The subcommands are:

    list (or ls)          List firewalls
    create <arguments>    Create a new firewall
    delete <firewall_id>  Delete a firewall

'entrywan firewall create' for create arguments`)
}

func usageLoadbalancer() {
	fmt.Println(`usage: entrywan loadbalancer list (or ls)

The subcommands are:

    list (or ls)              List load balancers
    create <arguments>        Create a new load balancer
    delete <loadbalancer_id>  Delete a load balancer

'entrywan loadbalancer create' for create arguments`)
}

func usageVpc() {
	fmt.Println(`usage: entrywan vpc <subcommand>

The subcommands are:

    list (or ls)        List vpcs
    create <arguments>  Create a new vpc
    delete <vpc_id>     Delete a vpc
    add <arguments>     Add instance to vpc
    remove <arguments>  Remove instance from vpc

'entrywan vpc create' for create arguments)
'entrywan vpc add' for add arguments)
'entrywan vpc remove' for remove arguments`)
}

func runCommand() {
	switch os.Args[1] {
	case "instance":
		instance()
	case "cluster":
		cluster()
	case "sshkey":
		sshkey()
	case "firewall":
		firewall()
	case "loadbalancer":
		loadbalancer()
	case "vpc":
		vpc()
	case "app":
		app()
	case "version":
		fmt.Printf("entrywan cli version %s\n", version)
	case "help":
		usage()
	default:
		usage()
	}
}
