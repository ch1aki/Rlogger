rlogger
====

Sending syslog message tool, written in Golang.

## Usage

execute `go run rlogger.go` from command line.

```bash
NAME:
   rlogger - Easily send syslog mesasge to remote host!

USAGE:
   main [global options] command [command options] [arguments...]

VERSION:
   0.1

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --host value      peer host (default: "127.0.0.1")
   --port value      peer port (default: 514)
   --protocol value  tcp or udp (default: "udp")
   --tag value       syslog message tag (default: "rlogger")
   --level value     DEBUG,INFO,NOTICE,WARNING,ERR,CRIT,ALERT,EMERG (default: "INFO")
   --help, -h        show help
   --version, -v     print the version
```

### send via UDP

```bash
go run rlogger.go \
"send message via UDP"
```

### send via TCP

```bash
go run rlogger.go \
--port 601 \
--protocol tcp \
"send message via TCP"
```

## Licence

[MIT](https://github.com/tcnksm/tool/blob/master/LICENCE)

## Author

[gahara22](https://github.com/gahara22)
