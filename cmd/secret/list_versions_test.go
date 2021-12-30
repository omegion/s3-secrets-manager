package secret

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	mocks2 "github.com/omegion/s3-secrets-manager/internal/api/mocks"
	"github.com/omegion/s3-secrets-manager/internal/client/mocks"
)

func TestListVersions(t *testing.T) {
	ctrl := gomock.NewController(t)
	clientMock := mocks.NewMockInterface(ctrl)
	api := mocks2.NewMockInterface(ctrl)

	expectedBucket := "test-bucket"
	expectedPath := "test/foo/boo"
	expectedSecret := &secret.Secret{
		Bucket: expectedBucket,
		Path:   expectedPath,
	}

	clientMock.EXPECT().GetS3API().Return(api, nil).Times(1)
	clientMock.EXPECT().ListVersions(api, expectedSecret).Return(nil).Times(1)

	cmd := &cobra.Command{}
	cmd.Flags().String("bucket", expectedBucket, "")
	cmd.Flags().String("path", expectedPath, "")

	err := listVersionsE(clientMock, cmd, []string{})
	assert.NoError(t, err)
}
