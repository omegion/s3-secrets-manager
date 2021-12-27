package secret

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/omegion/s3-secrets-manager/internal/client"
	"github.com/omegion/s3-secrets-manager/pkg/secret"
)

// Versions lists secret versions from S3.
func Versions() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "versions",
		Short: "List secret versions",
		Long:  "List secret versions from S3 with path and name",
		RunE:  client.With(listVersionsE),
	}

	cmd.Flags().String("path", "", "Path of the secret")

	if err := cmd.MarkFlagRequired("path"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}

	return cmd
}

func listVersionsE(client client.Interface, cmd *cobra.Command, args []string) error {
	path, _ := cmd.Flags().GetString("path")
	bucket, _ := cmd.Flags().GetString("bucket")

	scrt := &secret.Secret{
		Bucket: bucket,
		Path:   path,
	}

	api, err := client.GetS3API()
	if err != nil {
		return err
	}

	err = client.ListVersions(api, scrt)
	if err != nil {
		return err
	}

	err = scrt.PrintVersions()
	if err != nil {
		return err
	}

	return nil
}
