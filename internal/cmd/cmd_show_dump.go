package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/direnv/direnv/v2/gzenv"
	"github.com/urfave/cli/v3"
)

// CmdShowDump is `direnv show_dump`
var CmdShowDump = &cli.Command{
	Name:      "show_dump",
	Usage:     "Show the data inside of a dump for debugging purposes",
	ArgsUsage: "DUMP",
	Hidden:    true,
	Action:    actionSimple(cmdShowDumpAction),
}

func cmdShowDumpAction(_ Env, args []string) (err error) {
	if len(args) < 2 {
		return fmt.Errorf("missing DUMP argument")
	}

	var f interface{}
	err = gzenv.Unmarshal(args[1], &f)
	if err != nil {
		return err
	}

	e := json.NewEncoder(os.Stdout)
	e.SetIndent("", "  ")
	return e.Encode(f)
}
