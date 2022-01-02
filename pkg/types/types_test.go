package types

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/stretchr/testify/assert"
)

const (
	expectedPath = "test/foo/boo"
)

func TestEncodeTags(t *testing.T) {
	secret := Secret{
		Path: expectedPath,
	}

	assert.Equal(t, secret.EncodeTags(), "SecretPath=test%2Ffoo%2Fboo")
}

func TestEncodedValue(t *testing.T) {
	secret := Secret{
		Value: map[string]string{"password": "SECRET"},
	}

	actualValue, err := secret.EncodedValue()
	assert.NoError(t, err)
	assert.Equal(t, actualValue, []byte("{\"password\":\"SECRET\"}"))
}

func TestGetValue(t *testing.T) {
	secret := Secret{
		Value: map[string]string{"password": "SECRET"},
	}

	actualValue, err := secret.GetValue("password")
	assert.NoError(t, err)
	assert.Equal(t, actualValue, "SECRET")
}

func TestGetValueNotFound(t *testing.T) {
	expectedPath := expectedPath

	secret := Secret{
		Path:  expectedPath,
		Value: map[string]string{"password": "SECRET"},
	}

	actualValue, err := secret.GetValue("not-set-key")
	assert.EqualError(t, err, fmt.Sprintf("no secret found for %s in path %s", "not-set-key", expectedPath))
	assert.Equal(t, actualValue, "")
}

func TestGetVersionObjects(t *testing.T) {
	expectedPath := expectedPath
	expectedObjects := []types.ObjectIdentifier{
		{
			VersionId: aws.String("1"),
			Key:       &expectedPath,
		},
	}

	secret := Secret{
		Path:  expectedPath,
		Value: map[string]string{"password": "SECRET"},
		Versions: []*Version{
			{ID: "1"},
		},
	}

	assert.Equal(t, expectedObjects, secret.GetVersionObjects())
}

func TestSecretPrint(t *testing.T) {
	expectedPath := expectedPath

	secret := Secret{
		Path:  expectedPath,
		Value: map[string]string{"password": "SECRET"},
	}

	err := secret.Print()
	assert.NoError(t, err)

	err = secret.PrintVersions()
	assert.NoError(t, err)
}

func TestSecretsPrint(t *testing.T) {
	expectedPath := expectedPath

	secrets := Secrets{
		Items: []*Secret{
			{
				Path:  expectedPath,
				Value: map[string]string{"password": "SECRET"},
			},
		},
	}

	err := secrets.Print()
	assert.NoError(t, err)
}
