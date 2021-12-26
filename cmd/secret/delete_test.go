package secret

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	mocks2 "github.com/omegion/s3-secret-manager/internal/api/mocks"
	"github.com/omegion/s3-secret-manager/internal/client/mocks"
	"github.com/omegion/s3-secret-manager/pkg/secret"
)

const (
	expectedPath   = "test/foo/boo"
	expectedBucket = "test-bucket"
)

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	clientMock := mocks.NewMockInterface(ctrl)
	api := mocks2.NewMockAPIInterface(ctrl)

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
