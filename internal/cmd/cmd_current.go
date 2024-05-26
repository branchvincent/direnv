package cmd

import (
	"errors"

	"github.com/urfave/cli/v3"
)

// CmdCurrent is `direnv current`
var CmdCurrent = &cli.Command{
	Name:      "current",
	Usage:     "Reports whether direnv's view of a file is current (or stale)",
	Arguments: []cli.Argument{&cli.StringArg{Name: "PATH", Max: 1}},
	Hidden:    true,
	Action:    actionSimple(cmdCurrentAction),
}

func cmdCurrentAction(env Env, args []string) (err error) {
	if len(args) < 2 {
		err = errors.New("missing PATH argument")
		return
	}

	path := args[1]
	watches := NewFileTimes()
	watchString, ok := env[DIRENV_WATCHES]
	if ok {
		err = watches.Unmarshal(watchString)
		if err != nil {
			return
		}
	}

	err = watches.CheckOne(path)

	return
}
