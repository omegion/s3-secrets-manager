package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"

	awsTypes "github.com/aws/aws-sdk-go-v2/service/s3/types"
	log "github.com/sirupsen/logrus"

	"github.com/omegion/s3-secrets-manager/internal/api"
	"github.com/omegion/s3-secrets-manager/pkg/types"
)

// SecretController is a struct for arithmetic operations.
type SecretController struct {
	s3API api.Interface
}

// ListOptions is option for list secrets.
type ListOptions struct {
	Bucket,
	Path string
}

// NewSecretController is a factory for SecretController.
func NewSecretController(api api.Interface) *SecretController {
	return &SecretController{api}
}

// Get gets secret in given path from S3.
func (c SecretController) Get(secret *types.Secret) error {
	resp, err := c.s3API.GetObject(&api.GetObjectOptions{
		Bucket:    secret.Bucket,
		Path:      secret.Path,
		VersionID: secret.VersionID,
	})
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	existingValue := make(map[string]string)

	objectBodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(objectBodyByte, &existingValue)
	if err != nil {
		return err
	}

	if secret.Value != nil {
		for k, v := range existingValue {
			if changedValue, ok := secret.Value[k]; ok {
				v = changedValue
			}

			secret.Value[k] = v
		}
	} else {
		secret.Value = existingValue
	}

	return nil
}

// ListVersions lists secret versions in given path from S3.
func (c SecretController) ListVersions(scrt *types.Secret) error {
	resp, err := c.s3API.ListObjectVersions(&api.ListObjectVersionsOptions{
		Bucket: scrt.Bucket,
		Path:   scrt.Path,
	})
	if err != nil {
		return err
	}

	for _, version := range resp.Versions {
		scrt.Versions = append(scrt.Versions, &types.Version{
			ID:           *version.VersionId,
			LastModified: version.LastModified,
		})
	}

	return nil
}

// List lists secrets in given path from S3.
func (c SecretController) List(options *ListOptions) (*types.Secrets, error) {
	resp, err := c.s3API.ListObjects(&api.ListObjectOptions{
		Bucket: options.Bucket,
		Path:   options.Path,
	})
	if err != nil {
		return nil, err
	}

	secrets := &types.Secrets{}

	for _, item := range resp.Contents {
		secrets.Items = append(secrets.Items, &types.Secret{
			Bucket:       options.Bucket,
			Path:         *item.Key,
			LastModified: item.LastModified,
		})
	}

	return secrets, nil
}

// Set sets a secret to S3 bucket.
func (c SecretController) Set(secret *types.Secret) error {
	err := c.Get(secret)
	if err != nil {
		var nsk *awsTypes.NoSuchKey
		if errors.As(err, &nsk) {
			log.Debugln("A Secret object is not found, it will be created. Error:", nsk)
		} else {
			return err
		}
	}

	encodedValue, err := secret.EncodedValue()
	if err != nil {
		return err
	}

	_, err = c.s3API.PutObject(&api.PutObjectOptions{
		Bucket:      secret.Bucket,
		Path:        secret.Path,
		Value:       bytes.NewReader(encodedValue),
		EncodedTags: secret.EncodeTags(),
	})
	if err != nil {
		return err
	}

	return nil
}

// DeleteVersion deletes secret from S3 bucket.
func (c SecretController) DeleteVersion(secret *types.Secret) error {
	_, err := c.s3API.DeleteObject(&api.DeleteObjectOptions{
		Bucket:    secret.Bucket,
		Path:      secret.Path,
		VersionID: secret.VersionID,
	})
	if err != nil {
		return err
	}

	return nil
}

// Delete deletes secret from S3 bucket.
func (c SecretController) Delete(secret *types.Secret) error {
	err := c.ListVersions(secret)
	if err != nil {
		return err
	}

	_, err = c.s3API.DeleteObjects(&api.DeleteObjectsOptions{
		Bucket:  secret.Bucket,
		Objects: secret.GetVersionObjects(),
	})
	if err != nil {
		return err
	}

	return nil
}
