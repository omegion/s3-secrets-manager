package main

import (
	cmd2 "github.com/omegion/cobra-commander"
	"github.com/omegion/s3-secret-manager/cmd/secret"
	"github.com/spf13/cobra"
	"os"

	"github.com/omegion/s3-secret-manager/cmd"
)

func main() {
	root := &cobra.Command{
		Use:          "go-cli",
		Short:        "Go CLI application template",
		Long:         "Go CLI application template for Go projects.",
		SilenceUsage: true,
	}

	root.PersistentFlags().Bool("interactive", true, "Set the interactivity")

	commander := cmd2.NewCommander(root).
		SetCommand(
			secret.Secret(),
			cmd.Version(),
		).
		Init()

	if err := commander.Execute(); err != nil {
		os.Exit(1)
	}
}
