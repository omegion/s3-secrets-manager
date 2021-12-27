package secret

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	mocks2 "github.com/omegion/s3-secrets-manager/internal/api/mocks"
	"github.com/omegion/s3-secrets-manager/internal/client/mocks"
	"github.com/omegion/s3-secrets-manager/pkg/secret"
)

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	clientMock := mocks.NewMockInterface(ctrl)
	api := mocks2.NewMockInterface(ctrl)

	expectedSecret := &secret.Secret{
		Bucket: expectedBucket,
		Path:   expectedPath,
	}

	clientMock.EXPECT().GetS3API().Return(api, nil).Times(1)
	clientMock.EXPECT().GetSecret(api, expectedSecret).Return(nil).Times(1)

	cmd := &cobra.Command{}
	cmd.Flags().String("bucket", expectedBucket, "")
	cmd.Flags().String("path", expectedPath, "")

	err := getSecretE(clientMock, cmd, []string{})
	assert.NoError(t, err)
}
