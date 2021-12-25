package secret

import (
	"github.com/spf13/cobra"
)

// Secret is command collection of Secret operations.
func Secret() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "secret",
		Short: "Adds two numbers",
		Long:  "Adds two numbers",
	}

	cmd.AddCommand(
		Get(),
		List(),
		Set(),
		Delete(),
	)

	return cmd
}
