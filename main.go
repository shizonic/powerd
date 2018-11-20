package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli"
)

const Version string = "1.0.0"

func main() {
	daemon := NewDaemon()

	app := cli.NewApp()
	app.Version = Version
	app.Name = "powerd"
	app.Usage = "Simple daemon to handle battery power levels"
	app.Action = func(c *cli.Context) error {
		fmt.Println("Hello friend!")
		return nil
	}

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "check-interval",
			Value: 60,
			Usage: "Sets the interval to check the battery level in seconds",
		},
		cli.IntFlag{
			Name:  "warn-on-level",
			Value: 10,
			Usage: "Sets the battery level which will trigger a warn message",
		},
		cli.StringFlag{
			Name:  "warn-message",
			Value: "Battery level at 10%",
			Usage: "Sets the warning message",
		},
		cli.IntFlag{
			Name:  "exec-on-level",
			Value: 5,
			Usage: "Sets the battery level which will trigger a command execution",
		},
		cli.StringFlag{
			Name:  "exec-command",
			Value: "poweroff",
			Usage: "Sets the exec command",
		},
	}

	app.Action = func(c *cli.Context) {
		daemon.config.checkInterval = time.Duration(c.Int("check-interval"))
		daemon.config.execCommand = c.String("exec-command")
		daemon.config.execOnLevel = c.Int("exec-on-level")
		daemon.config.warnMessage = c.String("warn-message")
		daemon.config.warnOnLevel = c.Int("warn-on-level")

		daemon.Start()
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
