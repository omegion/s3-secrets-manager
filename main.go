package main

import (
	"os"

	cmd2 "github.com/omegion/cobra-commander"
	"github.com/spf13/cobra"

	"github.com/omegion/s3-secret-manager/cmd"
	"github.com/omegion/s3-secret-manager/cmd/secret"
)

func main() {
	root := &cobra.Command{
		Use:          "s3sm",
		Short:        "S3 Secrets Management.",
		Long:         "S3 Secrets Management for AWS S3.",
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
