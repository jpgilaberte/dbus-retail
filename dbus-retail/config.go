package dbusretail

import "github.com/codegangsta/cli"

var (
	FlAddr = cli.StringFlag{
		Name:   "ip",
		Usage:  "<ip>:<port> to listen on",
		Value:  "127.0.0.1:9050",
		EnvVar: "CERVANTES_IP",
	}
	FlLog = cli.StringFlag{
		Name:   "log",
		Usage:  "[INFO, DEBUG]",
		Value:  "DEBUG",
		EnvVar: "CERVANTES_LOG",
	}
)
