package secret

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/omegion/s3-secret-manager/internal/client"
	"github.com/omegion/s3-secret-manager/internal/prompt"
	"github.com/omegion/s3-secret-manager/pkg/secret"
)

// Delete deletes the secret from S3.
func Delete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete secret.",
		Long:  "Delete secret from S3 with path.",
		RunE:  client.With(deleteSecretE),
	}

	cmd.Flags().String("path", "", "Path of the secret")

	if err := cmd.MarkFlagRequired("path"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}

	return cmd
}

func deleteSecretE(client client.Interface, cmd *cobra.Command, args []string) error {
	path, _ := cmd.Flags().GetString("path")
	bucket, _ := cmd.Flags().GetString("bucket")
	interactive, _ := cmd.Flags().GetBool("interactive")

	scrt := &secret.Secret{
		Bucket: bucket,
		Path:   path,
	}

	api, err := client.GetS3API()
	if err != nil {
		return err
	}

	if interactive {
		_, err = prompt.NewPrompt(prompt.Options{}).DeletionConfirm()
		if err != nil {
			//nolint:nilerr // it's okay to return nil.
			return nil
		}
	}

	err = client.DeleteSecret(api, scrt)
	if err != nil {
		return err
	}

	return nil
}
