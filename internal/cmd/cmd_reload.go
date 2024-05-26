package cmd

import (
	"fmt"

	"github.com/urfave/cli/v3"
)

// CmdReload is `direnv reload`
var CmdReload = &cli.Command{
	Name:  "reload",
	Usage: "Triggers an env reload",
	Action: actionWithConfig(func(_ Env, _ []string, config *Config) error {
		foundRC, err := config.FindRC()
		if err != nil {
			return err
		}
		if foundRC == nil {
			return fmt.Errorf(".envrc not found")
		}

		return foundRC.Touch()
	}),
}
