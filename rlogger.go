package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"log/syslog"
	"os"
	"strconv"
)

func main() {
	var host string
	var proto string
	var tag string
	var level string
	var port int

	app := cli.NewApp()
	app.Name = "rlogger"
	app.Usage = "Easily send syslog mesasge to remote host!"
	app.Version = "0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "host",
			Usage:       "peer host",
			Value:       "127.0.0.1",
			Destination: &host,
		},
		cli.IntFlag{
			Name:        "port",
			Usage:       "peer port",
			Value:       514,
			Destination: &port,
		},
		cli.StringFlag{
			Name:        "protocol",
			Usage:       "tcp or udp",
			Value:       "udp",
			Destination: &proto,
		},
		cli.StringFlag{
			Name:        "tag",
			Usage:       "syslog message tag",
			Value:       "rlogger",
			Destination: &tag,
		},
		cli.StringFlag{
			Name:        "level",
			Usage:       "DEBUG,INFO,NOTICE,WARNING,ERR,CRIT,ALERT,EMERG",
			Value:       "INFO",
			Destination: &level,
		},
	}

	app.Action = func(c *cli.Context) error {
		// check proto
		if proto != "udp" && proto != "tcp" {
			fmt.Println("proto incorrect.")
			return nil
		}

		var log_level syslog.Priority
		switch level {
		case "DEBUG":
			log_level = syslog.LOG_DEBUG
		case "INFO":
			log_level = syslog.LOG_INFO
		case "NOTICE":
			log_level = syslog.LOG_NOTICE
		case "WARNING":
			log_level = syslog.LOG_WARNING
		case "ERR":
			log_level = syslog.LOG_ERR
		case "CRIT":
			log_level = syslog.LOG_CRIT
		case "ALERT":
			log_level = syslog.LOG_ALERT
		case "EMERG":
			log_level = syslog.LOG_EMERG
		default:
			fmt.Println("level incorrect.")
			return nil
		}

		// if args was passed, replace msg
		msg := "test message"
		if c.NArg() > 0 {
			msg = c.Args().Get(0)
		}

		logger, err := syslog.Dial(
			proto,
			host+":"+strconv.Itoa(port),
			log_level,
			tag)
		defer logger.Close()
		if err != nil {
			log.Fatal("error")
		}

		// send message
		logger.Write([]byte(msg))

		return nil
	}

	app.Run(os.Args)
}
