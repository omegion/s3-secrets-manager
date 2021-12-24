package client

import (
	"github.com/omegion/s3-secret-manager/internal/controller"
	"github.com/omegion/s3-secret-manager/internal/s3"
	"github.com/omegion/s3-secret-manager/pkg/secret"
)

// S3Interface is an interface for S3 Client.
type S3Interface interface {
	SetS3API(api s3.APIInterface)
	GetS3API() (s3.APIInterface, error)
	SetSecret(api s3.APIInterface, secret *secret.Secret) error
	GetSecret(api s3.APIInterface, secret *secret.Secret) error
	DeleteSecret(*secret.Secret) error
}

// GetS3API returns S3API.
func (c *Client) GetS3API() (s3.APIInterface, error) {
	return s3.NewAPI()
}

// SetS3API sets S3 API.
func (c *Client) SetS3API(api s3.APIInterface) {
	c.S3API = api
}

// SetSecret adds secret.
func (c Client) SetSecret(api s3.APIInterface, secret *secret.Secret) error {
	return controller.NewSecretController(api).Set(secret)
}

// GetSecret gets secret.
func (c Client) GetSecret(api s3.APIInterface, secret *secret.Secret) error {
	return controller.NewSecretController(api).Get(secret)
}

// DeleteSecret deletes secret.
func (c Client) DeleteSecret(secret *secret.Secret) error {
	return controller.NewSecretController(c.S3API).Delete(secret)
}
