package secret

import (
	"github.com/omegion/s3-secret-manager/pkg/secret"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/omegion/s3-secret-manager/internal/client"
)

var tags map[string]string

// Set sets a secret to S3.
func Set() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "set",
		Short:   "Adds two numbers",
		Long:    "Adds two numbers",
		Example: "  add --num1 1 --num2 2",
		RunE:    client.With(setSecretE),
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

	cmd.Flags().String("bucket", "", "S3 bucket name")
	if err := cmd.MarkFlagRequired("bucket"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}

	cmd.Flags().StringToStringVarP(&tags, "tags", "t", nil, "S3 bucket name")

	return cmd
}

func setSecretE(c client.Interface, cmd *cobra.Command, args []string) error {
	name, _ := cmd.Flags().GetString("name")
	value, _ := cmd.Flags().GetString("value")
	path, _ := cmd.Flags().GetString("path")
	bucket, _ := cmd.Flags().GetString("bucket")

	v := make(map[string]string)
	v[name] = value

	scrt := &secret.Secret{
		Bucket: bucket,
		Path:   path,
		Value:  v,
		Tags:   tags,
	}

	api, err := c.GetS3API()
	if err != nil {
		return err
	}

	err = c.SetSecret(api, scrt)
	if err != nil {
		return err
	}

	log.Infoln(scrt)

	return nil
}
