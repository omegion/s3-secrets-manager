package secret

import (
	"github.com/omegion/s3-secret-manager/internal/prompt"
	"github.com/omegion/s3-secret-manager/internal/s3"
	"github.com/omegion/s3-secret-manager/pkg/secret"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/omegion/s3-secret-manager/internal/client"
)

// Delete deletes the secret from S3.
func Delete() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete",
		Short:   "Adds two numbers",
		Long:    "Adds two numbers",
		Example: "  add --num1 1 --num2 2",
		RunE:    client.With(deleteSecretE),
	}

	cmd.Flags().String("path", "", "Path of the secret")
	if err := cmd.MarkFlagRequired("path"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}

	cmd.Flags().String("bucket", "", "S3 bucket name")
	if err := cmd.MarkFlagRequired("bucket"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}

	return cmd
}

func deleteSecretE(c client.Interface, cmd *cobra.Command, args []string) error {
	path, _ := cmd.Flags().GetString("path")
	bucket, _ := cmd.Flags().GetString("bucket")
	interactive, _ := cmd.Flags().GetBool("interactive")

	scrt := &secret.Secret{
		Bucket: bucket,
		Path:   path,
	}

	api, err := s3.NewAPI()
	if err != nil {
		return err
	}

	c.SetS3API(api)

	if interactive {
		_, err = prompt.NewPrompt(prompt.Options{}).DeletionConfirm()
		if err != nil {
			return nil
		}
	}

	err = c.DeleteSecret(scrt)
	if err != nil {
		return err
	}

	return nil
}
