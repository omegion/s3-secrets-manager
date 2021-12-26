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

func TestSet(t *testing.T) {
	ctrl := gomock.NewController(t)
	clientMock := mocks.NewMockInterface(ctrl)
	api := mocks2.NewMockInterface(ctrl)

	expectedSecretName := "test"
	expectedSecretValue := "TESTSECRET"
	expectedBucket := "test-bucket"
	expectedPath := "test/foo/boo"
	expectedSecret := &secret.Secret{
		Bucket: expectedBucket,
		Path:   expectedPath,
		Value:  map[string]string{"test": "TESTSECRET"},
	}

	clientMock.EXPECT().GetS3API().Return(api, nil).Times(1)
	clientMock.EXPECT().SetSecret(api, expectedSecret).Return(nil).Times(1)

	cmd := &cobra.Command{}
	cmd.Flags().String("name", expectedSecretName, "")
	cmd.Flags().String("value", expectedSecretValue, "")
	cmd.Flags().String("bucket", expectedBucket, "")
	cmd.Flags().String("path", expectedPath, "")

	err := setSecretE(clientMock, cmd, []string{})

	assert.NoError(t, err)
}
