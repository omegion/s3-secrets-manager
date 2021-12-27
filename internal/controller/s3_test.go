package controller

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	s32 "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/omegion/s3-secrets-manager/internal/api"
	"github.com/omegion/s3-secrets-manager/internal/api/mocks"
	"github.com/omegion/s3-secrets-manager/pkg/secret"
)

const (
	expectedBucket = "test-bucket"
	expectedPath   = "test-bucket"
)

func TestNewController(t *testing.T) {
	expectedAPI := api.API{}
	ctrl := NewSecretController(expectedAPI)

	assert.Equal(t, expectedAPI, ctrl.s3API)
}

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	apiMock := mocks.NewMockInterface(ctrl)

	expectedValueKey := "password"
	expectedValueValue := "MYSECRET"
	expectedSecret := secret.Secret{Bucket: expectedBucket, Path: expectedPath}
	expectedValue := map[string]string{expectedValueKey: expectedValueValue}

	options := &api.GetObjectOptions{
		Bucket: expectedBucket,
		Path:   expectedPath,
	}

	stringReader := strings.NewReader(fmt.Sprintf("{\"%s\":\"%s\"}", expectedValueKey, expectedValueValue))
	stringReadCloser := io.NopCloser(stringReader)

	output := &s32.GetObjectOutput{
		Body: stringReadCloser,
	}

	apiMock.EXPECT().GetObject(options).Return(output, nil).Times(1)

	controller := NewSecretController(apiMock)
	err := controller.Get(&expectedSecret)

	assert.NoError(t, err)
	assert.Equal(t, expectedSecret.Value, expectedValue)
}

func TestListVersions(t *testing.T) {
	ctrl := gomock.NewController(t)
	apiMock := mocks.NewMockInterface(ctrl)

	expectedVersionID := "VERSION-ID"
	expectedLastModified := time.Now()
	expectedSecret := secret.Secret{Bucket: expectedBucket, Path: expectedPath}

	options := &api.ListObjectVersionsOptions{
		Bucket: expectedBucket,
		Path:   expectedPath,
	}

	output := &s32.ListObjectVersionsOutput{
		Versions: []types.ObjectVersion{
			{
				VersionId:    &expectedVersionID,
				LastModified: &expectedLastModified,
			},
		},
	}

	apiMock.EXPECT().ListObjectVersions(options).Return(output, nil).Times(1)

	controller := NewSecretController(apiMock)
	err := controller.ListVersions(&expectedSecret)

	assert.NoError(t, err)

	for k, value := range expectedSecret.Versions {
		assert.Equal(t, value.ID, *output.Versions[k].VersionId)
		assert.Equal(t, value.LastModified, output.Versions[k].LastModified)
	}
}

func TestList(t *testing.T) {
	ctrl := gomock.NewController(t)
	apiMock := mocks.NewMockInterface(ctrl)

	expectedSecrets := []secret.Secret{
		{
			Bucket: expectedBucket,
			Path:   expectedPath,
		},
	}

	options := &api.ListObjectOptions{
		Path: expectedPath,
	}

	output := &s32.ListObjectsV2Output{
		Contents: []types.Object{
			{
				Key: aws.String(expectedPath),
			},
		},
	}

	apiMock.EXPECT().ListObjects(options).Return(output, nil).Times(1)

	controller := NewSecretController(apiMock)
	secrets, err := controller.List(&ListOptions{
		Path: expectedPath,
	})

	for k, v := range secrets.Items {
		assert.Equal(t, expectedSecrets[k].Path, v.Path)
	}

	assert.NoError(t, err)
}

func TestSet(t *testing.T) {
	ctrl := gomock.NewController(t)
	apiMock := mocks.NewMockInterface(ctrl)

	expectedValueKey := "password"
	expectedValueValue := "MYSECRET"
	expectedValue := map[string]string{expectedValueKey: expectedValueValue}
	expectedSecret := secret.Secret{Bucket: expectedBucket, Path: expectedPath, Value: expectedValue}

	getObjectOptions := &api.GetObjectOptions{
		Bucket: expectedBucket,
		Path:   expectedPath,
	}

	encodedValue, err := expectedSecret.EncodedValue()
	assert.NoError(t, err)

	putObjectOptions := &api.PutObjectOptions{
		Bucket:      expectedBucket,
		Path:        expectedPath,
		Value:       bytes.NewReader(encodedValue),
		EncodedTags: expectedSecret.EncodeTags(),
	}

	stringReader := strings.NewReader(fmt.Sprintf("{\"%s\":\"%s\"}", expectedValueKey, expectedValueValue))
	stringReadCloser := io.NopCloser(stringReader)

	output := &s32.GetObjectOutput{
		Body: stringReadCloser,
	}

	apiMock.EXPECT().GetObject(getObjectOptions).Return(output, nil).Times(1)
	apiMock.EXPECT().PutObject(putObjectOptions).Return(nil, nil).Times(1)

	controller := NewSecretController(apiMock)
	err = controller.Set(&expectedSecret)

	assert.NoError(t, err)
	assert.Equal(t, expectedSecret.Value, expectedValue)
}

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	apiMock := mocks.NewMockInterface(ctrl)

	expectedSecret := secret.Secret{Bucket: expectedBucket, Path: expectedPath}

	options := &api.DeleteObjectOptions{
		Bucket: expectedBucket,
		Path:   expectedPath,
	}

	apiMock.EXPECT().DeleteObject(options).Return(nil, nil).Times(1)

	controller := NewSecretController(apiMock)
	err := controller.Delete(&expectedSecret)

	assert.NoError(t, err)
}
