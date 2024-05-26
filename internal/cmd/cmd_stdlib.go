package cmd

import (
	"fmt"

	"github.com/urfave/cli/v3"
)

// CmdStdlib is `direnv stdlib`
var CmdStdlib = &cli.Command{
	Name:  "stdlib",
	Usage: "Displays the stdlib available in the .envrc execution context",
	Action: actionWithConfig(func(_ Env, _ []string, config *Config) error {
		fmt.Println(getStdlib(config))
		return nil
	}),
}
