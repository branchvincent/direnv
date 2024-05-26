package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/urfave/cli/v3"
)

// HookContext are the variables available during hook template evaluation
type HookContext struct {
	// SelfPath is the unescaped absolute path to direnv
	SelfPath string
}

// CmdHook is `direnv hook $0`
var CmdHook = &cli.Command{
	Name:      "hook",
	Usage:     "Used to setup the shell hook",
	Arguments: []cli.Argument{&cli.StringArg{Name: "SHELL", Min: 1, Max: 1}},
	Action:    cmdHookAction,
}

func cmdHookAction(_ context.Context, cmd *cli.Command) (err error) {
	target := cmd.Args().First()

	selfPath, err := os.Executable()
	if err != nil {
		return err
	}

	// Convert Windows path if needed
	selfPath = strings.Replace(selfPath, "\\", "/", -1)
	ctx := HookContext{selfPath}

	shell := DetectShell(target)
	if shell == nil {
		return fmt.Errorf("unknown target shell '%s'", target)
	}

	hookStr, err := shell.Hook()
	if err != nil {
		return err
	}

	hookTemplate, err := template.New("hook").Parse(hookStr)
	if err != nil {
		return err
	}

	err = hookTemplate.Execute(os.Stdout, ctx)
	if err != nil {
		return err
	}

	return
}
