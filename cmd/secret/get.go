package secret

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/omegion/s3-secret-manager/pkg/secret"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/omegion/s3-secret-manager/internal/client"
)

// Get gets secret from S3.
func Get() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "get",
		Short:   "Adds two numbers",
		Long:    "Adds two numbers",
		Example: "  add --num1 1 --num2 2",
		RunE:    client.With(getSecretE),
	}

	cmd.Flags().String("name", "", "Name of the secret")
	if err := cmd.MarkFlagRequired("name"); err != nil {
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

	return cmd
}

func getSecretE(c client.Interface, cmd *cobra.Command, args []string) error {
	name, _ := cmd.Flags().GetString("name")
	path, _ := cmd.Flags().GetString("path")
	bucket, _ := cmd.Flags().GetString("bucket")

	scrt := &secret.Secret{
		Bucket: bucket,
		Path:   path,
	}

	api, err := c.GetS3API()
	if err != nil {
		return err
	}

	err = c.GetSecret(api, scrt)
	if err != nil {
		var nsk *types.NoSuchKey
		if errors.As(err, &nsk) {
			return secret.NotFound{
				Key:    name,
				Secret: scrt,
			}
		} else {
			return err
		}
	}

	val, err := scrt.GetValue(name)
	if err != nil {
		return err
	}

	fmt.Println(val)

	return nil
}
