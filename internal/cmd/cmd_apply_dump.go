package cmd

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

// CmdApplyDump is `direnv apply_dump FILE`
var CmdApplyDump = &cli.Command{
	Name:      "apply_dump",
	Usage:     "Accepts a filename containing `direnv dump` output and generates a series of bash export statements to apply the given env",
	ArgsUsage: "FILE",
	Hidden:    true,
	Action:    actionSimple(cmdApplyDumpAction),
}

func cmdApplyDumpAction(env Env, args []string) (err error) {
	if len(args) < 2 {
		return fmt.Errorf("not enough arguments")
	}

	if len(args) > 2 {
		return fmt.Errorf("too many arguments")
	}
	filename := args[1]

	dumped, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	dumpedEnv, err := LoadEnv(string(dumped))
	if err != nil {
		return err
	}

	diff := env.Diff(dumpedEnv)

	exports := diff.ToShell(Bash)

	_, err = fmt.Println(exports)
	if err != nil {
		return err
	}

	return
}
