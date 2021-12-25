package client

import (
	"github.com/omegion/s3-secret-manager/internal/controller"
	"github.com/omegion/s3-secret-manager/internal/s3"
	"github.com/omegion/s3-secret-manager/pkg/secret"
)

// S3Interface is an interface for S3 Client.
type S3Interface interface {
	GetS3API() (s3.APIInterface, error)
	GetSecret(api s3.APIInterface, secret *secret.Secret) error
	ListSecret(api s3.APIInterface, options *controller.ListOptions) (*secret.Secrets, error)
	SetSecret(api s3.APIInterface, secret *secret.Secret) error
	DeleteSecret(api s3.APIInterface, secret *secret.Secret) error
}

// GetS3API returns S3API.
func (c *Client) GetS3API() (s3.APIInterface, error) {
	return s3.NewAPI()
}

// GetSecret gets secret.
func (c Client) GetSecret(api s3.APIInterface, secret *secret.Secret) error {
	return controller.NewSecretController(api).Get(secret)
}

// ListSecret gets secret.
func (c Client) ListSecret(api s3.APIInterface, options *controller.ListOptions) (*secret.Secrets, error) {
	return controller.NewSecretController(api).List(options)
}

// SetSecret adds secret.
func (c Client) SetSecret(api s3.APIInterface, secret *secret.Secret) error {
	return controller.NewSecretController(api).Set(secret)
}

// DeleteSecret deletes secret.
func (c Client) DeleteSecret(api s3.APIInterface, secret *secret.Secret) error {
	return controller.NewSecretController(api).Delete(secret)
}
