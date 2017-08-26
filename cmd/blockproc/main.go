package main

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/tendermint/tmlibs/cli"

	"github.com/tendermint/basecoin/cmd/basecoin/commands"
	"github.com/torresashjian/blockproc/plugin"
	"github.com/tendermint/basecoin/types"
)

func main() {
	var RootCmd = &cobra.Command{
		Use:   "blockproc",
		Short: "Block Process plugin for basecoin",
	}

	RootCmd.AddCommand(
		commands.InitCmd,
		commands.StartCmd,
		commands.UnsafeResetAllCmd,
		commands.VersionCmd,
	)

	commands.RegisterStartPlugin("blockproc", func() types.Plugin { return plugin.New() })
	cmd := cli.PrepareMainCmd(RootCmd, "BP", os.ExpandEnv("$HOME/.blockproc"))
	cmd.Execute()
}
