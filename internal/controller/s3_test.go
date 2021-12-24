package controller

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/omegion/s3-secret-manager/internal/s3"
	"github.com/omegion/s3-secret-manager/internal/s3/mocks"
	"github.com/omegion/s3-secret-manager/pkg/secret"
)

const (
	expectedBucket = "test-bucket"
	expectedPath   = "test-bucket"
)

func TestNewController(t *testing.T) {
	expectedAPI := s3.API{}
	ctrl := NewSecretController(expectedAPI)

	assert.Equal(t, expectedAPI, ctrl.s3API)
}

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	api := mocks.NewMockAPIInterface(ctrl)

	expectedValueKey := "password"
	expectedValueValue := "MYSECRET"
	expectedSecret := secret.Secret{Bucket: expectedBucket, Path: expectedPath}
	expectedValue := map[string]string{expectedValueKey: expectedValueValue}

	options := &s3.GetObjectOptions{
		Bucket: expectedBucket,
		Path:   expectedPath,
	}

	stringReader := strings.NewReader(fmt.Sprintf("{\"%s\":\"%s\"}", expectedValueKey, expectedValueValue))
	stringReadCloser := io.NopCloser(stringReader)

	api.EXPECT().GetObject(options).Return(stringReadCloser, nil).Times(1)

	controller := NewSecretController(api)
	err := controller.Get(&expectedSecret)

	assert.NoError(t, err)
	assert.Equal(t, expectedSecret.Value, expectedValue)
}

func TestSet(t *testing.T) {
	ctrl := gomock.NewController(t)
	api := mocks.NewMockAPIInterface(ctrl)

	expectedValueKey := "password"
	expectedValueValue := "MYSECRET"
	expectedValue := map[string]string{expectedValueKey: expectedValueValue}
	expectedSecret := secret.Secret{Bucket: expectedBucket, Path: expectedPath, Value: expectedValue}

	getObjectOptions := &s3.GetObjectOptions{
		Bucket: expectedBucket,
		Path:   expectedPath,
	}

	encodedValue, err := expectedSecret.EncodedValue()
	assert.NoError(t, err)

	putObjectOptions := &s3.PutObjectOptions{
		Bucket:      expectedBucket,
		Path:        expectedPath,
		Value:       bytes.NewReader(encodedValue),
		EncodedTags: expectedSecret.EncodeTags(),
	}

	stringReader := strings.NewReader(fmt.Sprintf("{\"%s\":\"%s\"}", expectedValueKey, expectedValueValue))
	stringReadCloser := io.NopCloser(stringReader)

	api.EXPECT().GetObject(getObjectOptions).Return(stringReadCloser, nil).Times(1)
	api.EXPECT().PutObject(putObjectOptions).Return(nil, nil).Times(1)

	controller := NewSecretController(api)
	err = controller.Set(&expectedSecret)

	assert.NoError(t, err)
	assert.Equal(t, expectedSecret.Value, expectedValue)
}

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	api := mocks.NewMockAPIInterface(ctrl)

	expectedSecret := secret.Secret{Bucket: expectedBucket, Path: expectedPath}

	options := &s3.DeleteObjectOptions{
		Bucket: expectedBucket,
		Path:   expectedPath,
	}

	api.EXPECT().DeleteObject(options).Return(nil, nil).Times(1)

	controller := NewSecretController(api)
	err := controller.Delete(&expectedSecret)

	assert.NoError(t, err)
}
