package secret

import (
	"encoding/json"
	"net/url"
)

// Secret is a struct for secret management.
type Secret struct {
	Bucket      string
	Description string
	Path        string
	Value       map[string]string
	Tags        map[string]string
}

// EncodeTags encodes tags for S3 API.
func (s Secret) EncodeTags() string {
	encoded := url.Values{}
	for k, v := range s.GetTags() {
		encoded.Set(k, v)
	}

	return encoded.Encode()
}

// GetTags gets all tags both user defined and default.
func (s Secret) GetTags() map[string]string {
	defaultTags := s.defaultTags()
	for k, v := range s.Tags {
		defaultTags[k] = v
	}

	return defaultTags
}

func (s Secret) defaultTags() map[string]string {
	return map[string]string{"SecretPath": s.Path}
}

// EncodedValue encodes value.
func (s Secret) EncodedValue() ([]byte, error) {
	return json.Marshal(s.Value)
}

// GetValue gets value.
func (s Secret) GetValue(key string) (string, error) {
	if v, ok := s.Value[key]; ok {
		return v, nil
	}

	return "", NotFoundError{
		Key:    key,
		Secret: &s,
	}
}
