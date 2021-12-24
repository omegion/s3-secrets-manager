package client

import (
	"github.com/omegion/s3-secret-manager/internal/s3"
	"github.com/spf13/cobra"
)

//nolint:lll // go generate is ugly.
//go:generate mockgen -destination=mocks/client_mock.go -package=mocks github.com/omegion/s3-secret-manager/internal/client Interface
// Interface is an interface entrypoint for the application.
type Interface interface {
	S3Interface
}

// Client is an entrypoint to controllers.
type Client struct {
	S3API s3.APIInterface
}

// NewClient is a factory for Client.
func NewClient() *Client {
	return &Client{}
}

// With is a wrapper for testing.
func With(
	fn func(c Interface, cmd *cobra.Command, args []string) error,
) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		c := NewClient()

		return fn(c, cmd, args)
	}
}
