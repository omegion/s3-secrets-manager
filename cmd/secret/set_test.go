package secret

import (
	"github.com/omegion/s3-secret-manager/pkg/secret"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	"github.com/omegion/s3-secret-manager/internal/client/mocks"
)

func TestSet(t *testing.T) {
	ctrl := gomock.NewController(t)
	c := mocks.NewMockInterface(ctrl)

	expectedSecretName := "test"
	expectedSecretValue := "TESTSECRET"
	expectedBucket := "test-bucket"
	expectedPath := "test/foo/boo"
	expectedSecret := &secret.Secret{
		Bucket: expectedBucket,
		Path:   expectedPath,
		Value:  map[string]string{"test": "TESTSECRET"},
	}

	api := APIMock{}

	c.EXPECT().GetS3API().Return(api, nil).Times(1)
	c.EXPECT().SetSecret(api, expectedSecret).Return(nil).Times(1)

	cmd := &cobra.Command{}
	cmd.Flags().String("name", expectedSecretName, "")
	cmd.Flags().String("value", expectedSecretValue, "")
	cmd.Flags().String("bucket", expectedBucket, "")
	cmd.Flags().String("path", expectedPath, "")

	err := setSecretE(c, cmd, []string{})

	assert.NoError(t, err)
}
