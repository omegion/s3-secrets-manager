package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"

	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	log "github.com/sirupsen/logrus"

	"github.com/omegion/s3-secret-manager/internal/s3"
	"github.com/omegion/s3-secret-manager/pkg/secret"
)

// SecretController is a struct for arithmetic operations.
type SecretController struct {
	s3API s3.APIInterface
}

// NewSecretController is a factory for SecretController.
func NewSecretController(api s3.APIInterface) *SecretController {
	return &SecretController{api}
}

// Get gets secret in given path from S3.
func (c SecretController) Get(secret *secret.Secret) error {
	object, err := c.s3API.GetObject(&s3.GetObjectOptions{
		Bucket: secret.Bucket,
		Path:   secret.Path,
	})
	if err != nil {
		return err
	}

	defer object.Close()

	existingValue := make(map[string]string)

	objectBodyByte, err := ioutil.ReadAll(object)
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

// Set sets a secret to S3 bucket.
func (c SecretController) Set(secret *secret.Secret) error {
	err := c.Get(secret)
	if err != nil {
		var nsk *types.NoSuchKey
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

	_, err = c.s3API.PutObject(&s3.PutObjectOptions{
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

// Delete deletes secret from S3 bucket.
func (c SecretController) Delete(secret *secret.Secret) error {
	_, err := c.s3API.DeleteObject(&s3.DeleteObjectOptions{
		Bucket: secret.Bucket,
		Path:   secret.Path,
	})
	if err != nil {
		return err
	}

	return nil
}
