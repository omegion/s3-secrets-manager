package client

import (
	"github.com/omegion/s3-secrets-manager/internal/api"
	"github.com/omegion/s3-secrets-manager/internal/controller"
	"github.com/omegion/s3-secrets-manager/pkg/types"
)

// S3Interface is an interface for S3 Client.
type S3Interface interface {
	GetS3API() (api.Interface, error)
	GetSecret(api api.Interface, secret *types.Secret) error
	ListVersions(api api.Interface, secret *types.Secret) error
	ListSecret(api api.Interface, options *controller.ListOptions) (*types.Secrets, error)
	SetSecret(api api.Interface, secret *types.Secret) error
	DeleteSecretVersion(api api.Interface, secret *types.Secret) error
	DeleteSecret(api api.Interface, secret *types.Secret) error
}

// GetS3API returns S3API.
func (c *Client) GetS3API() (api.Interface, error) {
	return api.NewAPI()
}

// GetSecret gets secret.
func (c Client) GetSecret(api api.Interface, secret *types.Secret) error {
	return controller.NewSecretController(api).Get(secret)
}

// ListVersions lists secret versions.
func (c Client) ListVersions(api api.Interface, secret *types.Secret) error {
	return controller.NewSecretController(api).ListVersions(secret)
}

// ListSecret gets secret.
func (c Client) ListSecret(api api.Interface, options *controller.ListOptions) (*types.Secrets, error) {
	return controller.NewSecretController(api).List(options)
}

// SetSecret adds secret.
func (c Client) SetSecret(api api.Interface, secret *types.Secret) error {
	return controller.NewSecretController(api).Set(secret)
}

// DeleteSecretVersion deletes secret.
func (c Client) DeleteSecretVersion(api api.Interface, secret *types.Secret) error {
	return controller.NewSecretController(api).DeleteVersion(secret)
}

// DeleteSecret deletes secret.
func (c Client) DeleteSecret(api api.Interface, secret *types.Secret) error {
	return controller.NewSecretController(api).Delete(secret)
}
