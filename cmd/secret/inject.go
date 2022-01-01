package secret

import (
	"fmt"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/omegion/s3-secrets-manager/internal/client"
	"github.com/omegion/s3-secrets-manager/pkg/kube"
)

// Inject injects secret from S3 to K8s Secret Resource.
func Inject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "inject",
		Short: "Inject secret",
		Long:  "Inject secret from S3 with path and name",
		RunE:  client.With(injectSecretE),
	}

	cmd.Flags().String("manifest", "", "Path of the manifest")

	if err := cmd.MarkFlagRequired("manifest"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}

	return cmd
}

func injectSecretE(client client.Interface, cmd *cobra.Command, args []string) error {
	bucket, _ := cmd.Flags().GetString("bucket")
	manifestFile, _ := cmd.Flags().GetString("manifest")

	s3API, err := client.GetS3API()
	if err != nil {
		return err
	}

	data, err := ioutil.ReadFile(manifestFile)
	if err != nil {
		return err
	}

	manifest := kube.Manifest{
		S3Bucket: bucket,
		S3API:    s3API,
		S3Client: client,
	}

	err = manifest.LoadResources(data)
	if err != nil {
		return err
	}

	err = manifest.Inject()
	if err != nil {
		return err
	}

	output, err := manifest.ToYAML()
	if err != nil {
		return err
	}

	fmt.Fprintf(cmd.OutOrStdout(), "%s---\n", output)

	return nil
}
