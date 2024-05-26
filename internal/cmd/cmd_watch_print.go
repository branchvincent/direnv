package cmd

import (
	"fmt"

	"github.com/urfave/cli/v3"
)

// CmdWatchPrint is `direnv watch-print`
var CmdWatchPrint = &cli.Command{
	Name:  "watch-print",
	Usage: "prints the watched paths",
	Flags: []cli.Flag{
		&cli.BoolFlag{Name: "null", Usage: "Print null-terminated paths"},
	},
	Hidden: true,
	Action: actionSimple(cmdWatchPrintAction),
}

func cmdWatchPrintAction(env Env, args []string) (err error) {
	watches := NewFileTimes()
	watchString, ok := env[DIRENV_WATCHES]
	separator := '\n'
	if len(args) > 1 && args[1] == "--null" {
		separator = 0
	}

	if ok {
		err = watches.Unmarshal(watchString)
		if err != nil {
			return
		}
	}

	for _, watch := range *watches.list {
		fmt.Printf("%s%c", watch.Path, separator)
	}

	return
}
