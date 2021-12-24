package secret

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeTags(t *testing.T) {
	secret := Secret{
		Path: "test/foo/boo",
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
	expectedPath := "test/foo/boo"

	secret := Secret{
		Path:  expectedPath,
		Value: map[string]string{"password": "SECRET"},
	}

	actualValue, err := secret.GetValue("not-set-key")
	assert.EqualError(t, err, fmt.Sprintf("no secret found for %s in path %s", "not-set-key", expectedPath))
	assert.Equal(t, actualValue, "")
}
