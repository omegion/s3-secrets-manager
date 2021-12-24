package controller

import (
	"github.com/omegion/s3-secret-manager/internal/s3"
	"github.com/omegion/s3-secret-manager/pkg/secret"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	api, err := s3.NewAPI()
	assert.NoError(t, err)

	scrt := &secret.Secret{
		Bucket: "omegion-secret-test",
		Path:   "secret/aws/rds-1",
		Value:  map[string]string{"password": "MYSECRETDATA2"},
	}

	ctrl := NewSecretController(api)

	err = ctrl.Set(scrt)
	assert.NoError(t, err)
}

func TestGet(t *testing.T) {
	api, err := s3.NewAPI()
	assert.NoError(t, err)

	scrt := &secret.Secret{
		Bucket: "omegion-secret-test",
		Path:   "secret/aws/rds-1",
	}

	ctrl := NewSecretController(api)

	err = ctrl.Get(scrt)
	assert.NoError(t, err)
}
