package secret

import (
	"fmt"
	"github.com/omegion/s3-secret-manager/internal/s3"
	"github.com/omegion/s3-secret-manager/pkg/secret"
	"io"
	"testing"

	awss3 "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/golang/mock/gomock"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	"github.com/omegion/s3-secret-manager/internal/client/mocks"
)

type APIMock struct{}

func (a APIMock) GetObject(options *s3.GetObjectOptions) (io.ReadCloser, error) {
	return nil, nil
}
func (a APIMock) PutObject(options *s3.PutObjectOptions) (*awss3.PutObjectOutput, error) {
	return nil, nil
}
func (a APIMock) DeleteObject(options *s3.DeleteObjectOptions) (*awss3.DeleteObjectOutput, error) {
	return nil, nil
}

func TestGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	c := mocks.NewMockInterface(ctrl)

	expectedSecretName := "test"
	expectedBucket := "test-bucket"
	expectedPath := "test/foo/boo"
	expectedSecret := &secret.Secret{
		Bucket: expectedBucket,
		Path:   expectedPath,
	}

	api := APIMock{}

	c.EXPECT().GetS3API().Return(api, nil).Times(1)
	c.EXPECT().GetSecret(api, expectedSecret).Return(nil).Times(1)

	cmd := &cobra.Command{}
	cmd.Flags().String("name", expectedSecretName, "")
	cmd.Flags().String("bucket", expectedBucket, "")
	cmd.Flags().String("path", expectedPath, "")

	err := getSecretE(c, cmd, []string{})

	assert.EqualError(t, err, fmt.Sprintf("no secret found for %s in path %s", expectedSecretName, expectedPath))
}
