package secret

import (
	"errors"
	"fmt"

	types2 "github.com/aws/aws-sdk-go-v2/service/s3/types"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/omegion/s3-secrets-manager/internal/client"
	"github.com/omegion/s3-secrets-manager/pkg/types"
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
	cmd.Flags().String("version-id", "", "Version ID to get specific version of the secret")
	cmd.Flags().String("path", "", "Path of the secret")

	if err := cmd.MarkFlagRequired("path"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}

	return cmd
}

func getSecretE(client client.Interface, cmd *cobra.Command, args []string) error {
	field, _ := cmd.Flags().GetString("field")
	path, _ := cmd.Flags().GetString("path")
	bucket, _ := cmd.Flags().GetString("bucket")
	versionID, _ := cmd.Flags().GetString("version-id")

	scrt := &types.Secret{
		Bucket: bucket,
		Path:   path,
	}

	if versionID != "" {
		scrt.VersionID = &versionID
	}

	api, err := client.GetS3API()
	if err != nil {
		return err
	}

	err = client.GetSecret(api, scrt)
	if err != nil {
		var nsk *types2.NoSuchKey
		if errors.As(err, &nsk) {
			return types.FieldNotFoundError{
				Field:  field,
				Secret: scrt,
			}
		}

		return err
	}

	if field != "" {
		var val string
		val, err = scrt.GetValue(field)

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
