package secret

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/omegion/s3-secrets-manager/internal/client"
	"github.com/omegion/s3-secrets-manager/internal/prompt"
	"github.com/omegion/s3-secrets-manager/pkg/secret"
)

// Delete deletes the secret from S3.
func Delete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "DeleteVersion secret.",
		Long:  "DeleteVersion secret from S3 with path.",
		RunE:  client.With(deleteSecretE),
	}

	cmd.Flags().String("version", "", "Version ID to get specific version of the secret")
	cmd.Flags().String("path", "", "Path of the secret")

	if err := cmd.MarkFlagRequired("path"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}

	return cmd
}

func deleteSecretE(client client.Interface, cmd *cobra.Command, args []string) error {
	path, _ := cmd.Flags().GetString("path")
	bucket, _ := cmd.Flags().GetString("bucket")
	version, _ := cmd.Flags().GetString("version")
	interactive, _ := cmd.Flags().GetBool("interactive")

	scrt := &secret.Secret{
		Bucket: bucket,
		Path:   path,
	}

	if version != "" {
		scrt.VersionID = &version
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

	if scrt.VersionID != nil {
		err = client.DeleteSecretVersion(api, scrt)
		if err != nil {
			return err
		}
	} else {
		err = client.DeleteSecret(api, scrt)
		if err != nil {
			return err
		}
	}

	return nil
}
