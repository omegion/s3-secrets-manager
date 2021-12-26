package main

import (
	"os"

	cmd2 "github.com/omegion/cobra-commander"
	"github.com/spf13/cobra"

	"github.com/omegion/s3-secret-manager/cmd"
	"github.com/omegion/s3-secret-manager/cmd/secret"
)

const (
	// Config file name where a config file will be created.
	// For example, $HOME/.s3sm/config.yaml.
	configFileName = "s3sm"

	// The environment variable prefix of all environment variables bound to our command line flags.
	// For example, --bucket is bound to S3SM_BUCKET.
	configEnvPrefix = "S3SM"
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
		SetConfig(getConfig()).
		Init()

	if err := commander.Execute(); err != nil {
		os.Exit(1)
	}
}

func getConfig() *cmd2.Config {
	configName := configFileName
	environmentPrefix := configEnvPrefix

	return &cmd2.Config{Name: &configName, EnvironmentPrefix: &environmentPrefix}
}
