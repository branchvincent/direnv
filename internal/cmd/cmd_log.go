package cmd

import (
	"errors"
	"fmt"

	"github.com/urfave/cli/v3"
)

// CmdLog is `direnv log [--status | --error] <message>`
var CmdLog = &cli.Command{
	Name:      "log",
	Usage:     "Logs a given message",
	Arguments: []cli.Argument{&cli.StringArg{Name: "message", Max: 1}},
	Flags: []cli.Flag{
		&cli.BoolFlag{Name: "status", Usage: ""},
		&cli.BoolFlag{Name: "error", Usage: ""},
	},
	Action: actionWithConfig(cmdLog),
}

func cmdLog(_ Env, args []string, c *Config) (err error) {
	if len(args) != 3 {
		return errors.New("invalid arguments")
	}
	logType := args[1]
	message := args[2]
	if logType == "--status" || logType == "-status" {
		logStatus(c, message)
	} else if logType == "--error" || logType == "-error" {
		logError(c, message)
	} else {
		return fmt.Errorf("invalid log-type '%s'", logType)
	}
	return nil
}
