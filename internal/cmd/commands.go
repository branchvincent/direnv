package cmd

import (
	"context"
	"time"

	"github.com/urfave/cli/v3"
)

type actionSimpleFn func(env Env, args []string) error

func actionSimple(fn actionSimpleFn) cli.ActionFunc {
	return func(_ context.Context, cmd *cli.Command) error {
		env := cmd.Root().Metadata["env"].(Env)
		return fn(env, cmd.Args().Slice())
	}
}

type actionFn func(env Env, args []string, config *Config) error

func actionWithConfig(fn actionFn) cli.ActionFunc {
	return func(_ context.Context, cmd *cli.Command) error {
		env := cmd.Root().Metadata["env"].(Env)
		config := cmd.Root().Metadata["config"].(*Config)
		return fn(env, cmd.Args().Slice(), config)
	}
}

// CmdList contains the list of all direnv sub-commands
var CmdList = []*cli.Command{
	CmdAllow,
	CmdApplyDump,
	CmdShowDump,
	CmdDeny,
	CmdDotEnv,
	CmdDump,
	CmdEdit,
	CmdExec,
	CmdExport,
	CmdFetchURL,
	CmdHook,
	CmdPrune,
	CmdReload,
	CmdStatus,
	CmdStdlib,
	CmdVersion,
	CmdWatch,
	CmdWatchDir,
	CmdWatchList,
	CmdWatchPrint,
	CmdCurrent,
	CmdLog,
}

func cmdWithWarnTimeout(fn cli.ActionFunc) cli.ActionFunc {
	return func(ctx context.Context, cmd *cli.Command) (err error) {
		config := cmd.Metadata["config"].(*Config)

		// Disable warning if WarnTimeout is <= 0
		if config.WarnTimeout <= 0 {
			return fn(ctx, cmd)
		}

		done := make(chan bool, 1)
		go func() {
			select {
			case <-done:
				return
			case <-time.After(config.WarnTimeout):
				args := cmd.Args().Slice()
				logError(config, "(%v) is taking a while to execute. Use CTRL-C to give up.", args)
			}
		}()

		err = fn(ctx, cmd)
		done <- true
		return err
	}
}

func Run(env Env, args []string) error {
	app := &cli.Command{
		Name:  "direnv",
		Usage: "Load/unload environment variables based on $PWD",
		Before: func(_ context.Context, cmd *cli.Command) (context.Context, error) {
			config, err := LoadConfig(env)
			if err != nil {
				return nil, err
			}
			cmd.Metadata["config"] = config
			cmd.Metadata["env"] = env
			return nil, nil
		},
		Commands:                   CmdList,
		EnableShellCompletion:      true,
		ShellCompletionCommandName: "completion",
		Version:                    version,
	}
	return app.Run(context.Background(), args)
}
