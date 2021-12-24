package secret

import (
	"encoding/json"
	"net/url"
)

type Secret struct {
	Bucket      string
	Description string
	Path        string
	Value       map[string]string
	Tags        map[string]string
}

func (s Secret) EncodeTags() string {
	encoded := url.Values{}
	for k, v := range s.GetTags() {
		encoded.Set(k, v)
	}

	return encoded.Encode()
}

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

func (s Secret) EncodedValue() ([]byte, error) {
	return json.Marshal(s.Value)
}

func (s Secret) GetValue(key string) (string, error) {
	if v, ok := s.Value[key]; ok {
		return v, nil
	}

	return "", NotFound{
		Key:    key,
		Secret: &s,
	}
}
