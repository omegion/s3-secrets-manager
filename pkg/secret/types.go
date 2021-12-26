package secret

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"text/tabwriter"
	"time"
)

const (
	writerPadding = 2
)

// Secret is a struct for secret management.
type Secret struct {
	Bucket       string
	Description  string
	Path         string
	LastModified *time.Time
	Value        map[string]string
	Tags         map[string]string
	VersionID    *string
	Versions     []*Version
}

// Secrets is collection of Secret.
type Secrets struct {
	Items []*Secret
}

type Version struct {
	ID           string
	LastModified *time.Time
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

// Print prints Secret details.
func (s Secret) Print() error {
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, writerPadding, ' ', 0)
	fmt.Fprintf(writer, "Key\tValue\n")
	fmt.Fprintf(writer, "----\t----\n")

	for key, value := range s.Value {
		fmt.Fprintf(writer, "%s\t%s\n", key, value)
	}

	writer.Flush()

	return nil
}

// PrintVersions prints Secret Versions.
func (s Secret) PrintVersions() error {
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, writerPadding, ' ', 0)
	fmt.Fprintf(writer, "Order\tVersion ID\tLast Modified\n")
	fmt.Fprintf(writer, "----\t----\t----\n")

	for key, scrt := range s.Versions {
		fmt.Fprintf(writer, "%d\t%s\t%s\n", key+1, scrt.ID, scrt.LastModified)
	}

	writer.Flush()

	return nil
}

// Print prints Secrets details.
func (s Secrets) Print() error {
	writer := tabwriter.NewWriter(os.Stdout, 0, 0, writerPadding, ' ', 0)
	fmt.Fprintf(writer, "Secret\tLast Modified\n")
	fmt.Fprintf(writer, "----\t----\n")

	for _, scrt := range s.Items {
		fmt.Fprintf(writer, "%s\t%s\n", scrt.Path, scrt.LastModified)
	}

	writer.Flush()

	return nil
}
