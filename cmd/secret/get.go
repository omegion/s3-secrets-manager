package secret

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/omegion/s3-secret-manager/internal/client"
	"github.com/omegion/s3-secret-manager/pkg/secret"
)

// Get gets secret from S3.
func Get() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get secret",
		Long:  "Get secret from S3 with path and name",
		RunE:  client.With(getSecretE),
	}

	cmd.Flags().String("field", "", "Name of the field in the secret")
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

func getSecretE(client client.Interface, cmd *cobra.Command, args []string) error {
	field, _ := cmd.Flags().GetString("field")
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

	err = client.GetSecret(api, scrt)
	if err != nil {
		var nsk *types.NoSuchKey
		if errors.As(err, &nsk) {
			return secret.FieldNotFoundError{
				Field:  field,
				Secret: scrt,
			}
		}

		return err
	}

	if field != "" {
		val, err := scrt.GetValue(field)
		if err != nil {
			return err
		}

		//nolint:forbidigo // fmt is okay.
		fmt.Println(val)
	} else {
		err = scrt.Print()
		if err != nil {
			return err
		}
	}

	return nil
}
