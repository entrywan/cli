Entrywan's command line interface.  Read the short [blog post](https://www.entrywan.com/blog/2024-06-26-cli-released).

### Install

Download [official binaries](https://github.com/entrywan/cli/releases/latest) for Linux,
macOS, NetBSD, FreeBSD, OpenBSD, Windows and Solaris.  Alternatively,
if you have a working [go](https://go.dev/) installation, you can
clone this repository and run `go install .` from the root of the
tree.

### Configure

Create a `~/.config/entrywan/config.toml` with the following:

```
token = "<IAM token>"
```

using an [IAM token](https://entrywan.com/docs#iam).

### Examples

List clusters and fetch kubeconfig for a cluster:

```
$ entrywan cluster ls
ID					NAME	STATE		SIZE	LOCATION	VERSION
7cdc1e38-23f4-450e-b4ce-40d8f3b5456f	foo	running		4	us1		1.29

$ entrywan cluster kubeconfig 7cdc1e38-23f4-450e-b4ce-40d8f3b5456f
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: LS0tLS1CRUdJTiBDRVJ
[...]
```

"ls" is an alias for "list" and "rm" is an alias for "delete".

Create an app and fetch its logs:


```
$ entrywan app create --name nginx3 --location us1 --source oci --image nginx --size 256 --port 80

$ entrywan app list
ID					NAME	STATE		SIZE (MB)	LOCATION
7ca38ff7-3efd-4deb-9084-13540c388167	nginx3	running		256		us1

$ entrywan app logs 7ca38ff7-3efd-4deb-9084-13540c388167
/docker-entrypoint.sh: /docker-entrypoint.d/ is not empty, will attempt to perform configuration
/docker-entrypoint.sh: Looking for shell scripts in /docker-entrypoint.d/
/docker-entrypoint.sh: Launching /docker-entrypoint.d/10-listen-on-ipv6-by-default.sh
10-listen-on-ipv6-by-default.sh: info: Getting the checksum of /etc/nginx/conf.d/default.conf
10-listen-on-ipv6-by-default.sh: info: Enabled listen on IPv6 in /etc/nginx/conf.d/default.conf
/docker-entrypoint.sh: Sourcing /docker-entrypoint.d/15-local-resolvers.envsh
/docker-entrypoint.sh: Launching /docker-entrypoint.d/20-envsubst-on-templates.sh
/docker-entrypoint.sh: Launching /docker-entrypoint.d/30-tune-worker-processes.sh
/docker-entrypoint.sh: Configuration complete; ready for start up
2024/06/26 13:28:39 [notice] 1#1: using the "epoll" event method
2024/06/26 13:28:39 [notice] 1#1: nginx/1.27.0
2024/06/26 13:28:39 [notice] 1#1: built by gcc 12.2.0 (Debian 12.2.0-14) 
2024/06/26 13:28:39 [notice] 1#1: OS: Linux 5.14.0-362.13.1.el9_3.x86_64
2024/06/26 13:28:39 [notice] 1#1: getrlimit(RLIMIT_NOFILE): 524288:524288
2024/06/26 13:28:39 [notice] 1#1: start worker processes
2024/06/26 13:28:39 [notice] 1#1: start worker process 30
2024/06/26 13:28:39 [notice] 1#1: start worker process 31
```

Usage information is built into the program:

```
$ entrywan
usage: entrywan <command> [<args>]

The commands are:

   instance      Manage instances
   cluster       Manage kubernetes clusters
   sshkey        Manage ssh keys
   firewall      Manage firewalls
   loadbalancer  Manage load balancers
   vpc           Manage VPC networks
   app           Manage apps
   version       Show cli version

'entrywan <subcommand>' to show usage for each subcommand.
```

### Contributing

Issues and PRs are welcome.

### License

GPL-3.0
