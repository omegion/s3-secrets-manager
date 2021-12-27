package secret

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/omegion/s3-secrets-manager/internal/client"
	"github.com/omegion/s3-secrets-manager/pkg/secret"
)

//nolint:gochecknoglobals // it's okay to use global here.
var tags map[string]string

// Set sets a secret to S3.
func Set() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set",
		Short: "Set secret",
		Long:  "Set secret to S3.",
		RunE:  client.With(setSecretE),
	}

	cmd.Flags().String("name", "", "Name of the secret")

	if err := cmd.MarkFlagRequired("name"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}

	cmd.Flags().String("value", "", "Value of the secret")

	if err := cmd.MarkFlagRequired("value"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}

	cmd.Flags().String("path", "", "Path of the secret")

	if err := cmd.MarkFlagRequired("path"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}

	cmd.Flags().StringToStringVarP(&tags, "tags", "t", nil, "S3 bucket name")

	return cmd
}

func setSecretE(client client.Interface, cmd *cobra.Command, args []string) error {
	name, _ := cmd.Flags().GetString("name")
	value, _ := cmd.Flags().GetString("value")
	path, _ := cmd.Flags().GetString("path")
	bucket, _ := cmd.Flags().GetString("bucket")

	scrt := &secret.Secret{
		Bucket: bucket,
		Path:   path,
		Value:  map[string]string{name: value},
		Tags:   tags,
	}

	api, err := client.GetS3API()
	if err != nil {
		return err
	}

	err = client.SetSecret(api, scrt)
	if err != nil {
		return err
	}

	err = scrt.Print()
	if err != nil {
		return err
	}

	return nil
}
