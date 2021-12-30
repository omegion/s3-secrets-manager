package secret

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	apiMock "github.com/omegion/s3-secrets-manager/internal/api/mocks"
	"github.com/omegion/s3-secrets-manager/internal/client/mocks"
)

const (
	expectedPath   = "test/foo/boo"
	expectedBucket = "test-bucket"
)

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	clientMock := mocks.NewMockInterface(ctrl)
	api := apiMock.NewMockInterface(ctrl)

	expectedSecret := &secret.Secret{
		Bucket: expectedBucket,
		Path:   expectedPath,
	}

	clientMock.EXPECT().GetS3API().Return(api, nil).Times(1)
	clientMock.EXPECT().DeleteSecret(api, expectedSecret).Return(nil).Times(1)

	cmd := &cobra.Command{}
	cmd.Flags().String("bucket", expectedBucket, "")
	cmd.Flags().String("path", expectedPath, "")

	err := deleteSecretE(clientMock, cmd, []string{})

	assert.NoError(t, err)
}
