package secret

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	"github.com/omegion/s3-secret-manager/internal/client/mocks"
	mocks2 "github.com/omegion/s3-secret-manager/internal/s3/mocks"
	"github.com/omegion/s3-secret-manager/pkg/secret"
)

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	clientMock := mocks.NewMockInterface(ctrl)
	api := mocks2.NewMockAPIInterface(ctrl)

	expectedSecretName := "test"
	expectedBucket := "test-bucket"
	expectedPath := "test/foo/boo"
	expectedSecret := &secret.Secret{
		Bucket: expectedBucket,
		Path:   expectedPath,
	}

	clientMock.EXPECT().GetS3API().Return(api, nil).Times(1)
	clientMock.EXPECT().GetSecret(api, expectedSecret).Return(nil).Times(1)

	cmd := &cobra.Command{}
	cmd.Flags().String("name", expectedSecretName, "")
	cmd.Flags().String("bucket", expectedBucket, "")
	cmd.Flags().String("path", expectedPath, "")

	err := getSecretE(clientMock, cmd, []string{})

	assert.EqualError(t, err, fmt.Sprintf("no secret found for %s in path %s", expectedSecretName, expectedPath))
}
