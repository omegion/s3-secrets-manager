package client

import (
	"github.com/omegion/s3-secret-manager/internal/api"
	"github.com/omegion/s3-secret-manager/internal/controller"
	"github.com/omegion/s3-secret-manager/pkg/secret"
)

// S3Interface is an interface for S3 Client.
type S3Interface interface {
	GetS3API() (api.APIInterface, error)
	GetSecret(api api.APIInterface, secret *secret.Secret) error
	ListVersions(api api.APIInterface, secret *secret.Secret) error
	ListSecret(api api.APIInterface, options *controller.ListOptions) (*secret.Secrets, error)
	SetSecret(api api.APIInterface, secret *secret.Secret) error
	DeleteSecret(api api.APIInterface, secret *secret.Secret) error
}

// GetS3API returns S3API.
func (c *Client) GetS3API() (api.APIInterface, error) {
	return api.NewAPI()
}

// GetSecret gets secret.
func (c Client) GetSecret(api api.APIInterface, secret *secret.Secret) error {
	return controller.NewSecretController(api).Get(secret)
}

// ListVersions lists secret versions.
func (c Client) ListVersions(api api.APIInterface, secret *secret.Secret) error {
	return controller.NewSecretController(api).ListVersions(secret)
}

// ListSecret gets secret.
func (c Client) ListSecret(api api.APIInterface, options *controller.ListOptions) (*secret.Secrets, error) {
	return controller.NewSecretController(api).List(options)
}

// SetSecret adds secret.
func (c Client) SetSecret(api api.APIInterface, secret *secret.Secret) error {
	return controller.NewSecretController(api).Set(secret)
}

// DeleteSecret deletes secret.
func (c Client) DeleteSecret(api api.APIInterface, secret *secret.Secret) error {
	return controller.NewSecretController(api).Delete(secret)
}
