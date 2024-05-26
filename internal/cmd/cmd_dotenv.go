package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/direnv/direnv/v2/pkg/dotenv"
	"github.com/urfave/cli/v3"
)

// CmdDotEnv is `direnv dotenv [SHELL [PATH_TO_DOTENV]]`
// Transforms a .env file to evaluatable `export KEY=PAIR` statements.
//
// See: https://github.com/bkeepers/dotenv and https://github.com/ddollar/foreman
var CmdDotEnv = &cli.Command{
	Name:      "dotenv",
	Usage:     "Transforms a .env file to evaluatable `export KEY=PAIR` statements",
	ArgsUsage: "[SHELL] [PATH_TO_DOTENV]",
	Hidden:    true,
	Action:    actionSimple(cmdDotEnvAction),
}

func cmdDotEnvAction(_ Env, args []string) (err error) {
	var shell Shell
	var newenv Env
	var target string

	if len(args) > 1 {
		shell = DetectShell(args[1])
	} else {
		shell = Bash
	}

	if len(args) > 2 {
		target = args[2]
	}

	if target == "" {
		target = ".env"
	}

	var data []byte
	if data, err = os.ReadFile(target); err != nil {
		return
	}

	// Set PWD env var to the directory the .env file resides in. This results
	// in the least amount of surprise, as a dotenv file is most often defined
	// in the same directory it's loaded from, so referring to PWD should match
	// the directory of the .env file.
	path, err := filepath.Abs(target)
	if err != nil {
		return err
	}
	os.Setenv("PWD", filepath.Dir(path))

	newenv, err = dotenv.Parse(string(data))
	if err != nil {
		return err
	}

	str := newenv.ToShell(shell)
	fmt.Println(str)

	return
}
