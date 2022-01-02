package secret

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Secret is command collection of Secret operations.
func Secret() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "secret",
		Short: "Secret operations.",
		Long:  "Secret operations in S3.",
	}

	cmd.AddCommand(
		Get(),
		Versions(),
		List(),
		Set(),
		Delete(),
		Inject(),
	)

	cmd.PersistentFlags().String("bucket", "", "S3 bucket name")

	if err := cmd.MarkPersistentFlagRequired("bucket"); err != nil {
		log.Fatalf("Lethal damage: %s\n\n", err)
	}

	return cmd
}
