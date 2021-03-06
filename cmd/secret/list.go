package secret

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/omegion/s3-secrets-manager/internal/client"
	"github.com/omegion/s3-secrets-manager/internal/controller"
)

// List gets secret from S3.
func List() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List secrets in path",
		Long:  "List secrets in path from S3 with path and name",
		RunE:  client.With(listSecretE),
	}

	cmd.Flags().String("path", "", "Path of the secret")

	if err := cmd.MarkFlagRequired("path"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}

	return cmd
}

func listSecretE(client client.Interface, cmd *cobra.Command, args []string) error {
	path, _ := cmd.Flags().GetString("path")
	bucket, _ := cmd.Flags().GetString("bucket")

	options := &controller.ListOptions{
		Bucket: bucket,
		Path:   path,
	}

	api, err := client.GetS3API()
	if err != nil {
		return err
	}

	secrets, err := client.ListSecret(api, options)
	if err != nil {
		return err
	}

	err = secrets.Print()
	if err != nil {
		return err
	}

	return nil
}
