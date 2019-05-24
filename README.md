# Caddy Builder

> A simple build tool for caddy

## How to use?

**Warning: Caddy Builder will delete the `$GOPATH/src/github.com/mholt/caddy` directory and re-clone at compile**

``` sh
➜  ~ caddybuilder --help

A simple build tool for caddy

Usage:
  caddybuilder [flags]
  caddybuilder [command]

Available Commands:
  help        Help about any command
  version     Print version

Flags:
      --debug            debug mode
  -j, --extjson string   extended caddy plugins json file
  -h, --help             help for caddybuilder
      --modcmd string    custom go mod command file for handling special dependencies
  -o, --output string    caddy binary output path
  -p, --plugins string   comma separated list of caddy builder (default "all")
  -t, --tag string       caddy tag (default "v1.0.0")

Use "caddybuilder [command] --help" for more information about a command.
```

### Host Compile

- 1、Download the precompiled binary from the release page
- 2、Run `caddybuilder` to compile caddy (requires go and git commands, golang version is 1.12)
- 3、Use the `-p` option to select the list of plugins you want to compile
- 4、Use the `-j` option to add a list of extended plugins (and you should specify it in the `-p` option)
- 5、Add the extended `go mod` repair command file with the `-m` option, for example `go mod edit -replace github.com/h2non/gock=gopkg.in/h2non/gock.v1@v1.0.14`

### Docker Compile

**Warning: The `mritd/caddybuilder` image uses `golang:1.12.5-alpine3.9` as the base image.
Since the C library is different (musl c), the successfully compiled caddy binary cannot be run on debian or centos; 
if you want to run it on debian and centos, please modify the Dockerfile file under this project (replace the golang base image)**

- 1、Use the `mritd/caddybuilder` docker image (`FROM mritd/caddybuilder`)
- 2、Run the compile command (`CMD ["-p", "godaddy,ipfilter", "-o", "/caddy"] `)
- 3、Use the docker multi-stage build to copy caddy binary (`COPY --from=builder /caddy /usr/bin/caddy`)


