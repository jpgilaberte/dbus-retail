package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"fmt"
	"os"
	"context"
	"os/signal"
	"syscall"
	"time"
	"github.com/jpgilaberte/dbus-retail/exchange/rabbitmq"
)

const(
	Version = "0.0.1-SNAPSHOT"
	Build = "building"
)

var (
	FlAddr = cli.StringFlag{
		Name:   "ip",
		Usage:  "<ip>:<port> to listen on",
		Value:  "127.0.0.1:9050",
		EnvVar: "DBUS_RETAIL_IP",
	}
	FlLog = cli.StringFlag{
		Name:   "log",
		Usage:  "[INFO, DEBUG]",
		Value:  "DEBUG",
		EnvVar: "DBUS_RETAIL_LOG",
	}
)

var(
	agentCommand = cli.Command{
		Name:      "agent",
		ShortName: "a",
		Usage:     "Dbus Agent",
		Flags:     []cli.Flag{FlAddr, FlLog},
		Action:    action(agentAction),
	}
	helpCommand = cli.Command{
		Name:      "help",
		ShortName: "h",
		Usage:     "Usage",
		Action: func(_ *cli.Context) {
			fmt.Printf("Usage: dbus-retail agent\n")
		},
	}
	cliCommands = []cli.Command {agentCommand, helpCommand}
)

func main() {
	app := cli.NewApp()
	app.Name = "dbus-retail"
	app.Commands = append(cliCommands)
	app.Version = fmt.Sprintf("%s-%s", Version, Build)
	app.Run(os.Args)
}

func action(f func(c *cli.Context) error) func(c *cli.Context) {
	return func(c *cli.Context) {
		setLogLevel(c)
		err := f(c)
		if err != nil {
			log.Error("PANIC: " + err.Error())
			os.Exit(1)
		}
	}
}

func agentAction(c *cli.Context) error {
	log.Info("Start server")
	ctx := context.WithValue(context.Background(), "properties", c)
	ctx, cancel := context.WithCancel(ctx)
	//go backend.Dbus(ctx)
	go rabbitmq.Run(ctx)
	manageSignals(cancel)
	return nil
}

func manageSignals(cancel context.CancelFunc){
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigs
	cancel()
	log.Infof("Waiting for context close")
	time.Sleep(1000000)
	log.Infof("Gracefull shutdown - Signal: %v", sig)
}

func setLogLevel(c *cli.Context) {
	level, err := log.ParseLevel(c.String("log"))
	if err != nil {
		log.Fatalf("Wrong log level: %v. Use INFO or DEBUG options", err.Error())
		os.Exit(1)
	}
	log.Infof("Log is ready in level: %v", level)
	log.SetLevel(level)
}
