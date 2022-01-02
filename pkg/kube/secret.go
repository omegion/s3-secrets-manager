package kube

import (
	"bytes"
	"encoding/base64"
	"errors"
	"io"
	"reflect"
	"regexp"
	"strings"

	types2 "github.com/aws/aws-sdk-go-v2/service/s3/types"
	log "github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	k8yaml "k8s.io/apimachinery/pkg/util/yaml"
	"sigs.k8s.io/yaml"

	"github.com/omegion/s3-secrets-manager/internal/api"
	"github.com/omegion/s3-secrets-manager/internal/client"
	"github.com/omegion/s3-secrets-manager/pkg/types"
)

// Manifest is used for injection of secrets to Kubernetes Secret resource.
type Manifest struct {
	Resources []*Resource
	S3Bucket  string
	S3Client  client.S3Interface
	S3API     api.Interface
}

// Resource is the basis for all Templates.
type Resource struct {
	Kind         string
	TemplateData map[string]interface{}
	Data         map[string]interface{}
	Annotations  map[string]string
}

// LoadResources loads Kubernetes Unstructured resources from given file data.
func (m *Manifest) LoadResources(fileData []byte) error {
	for _, resourceYAML := range strings.Split(string(fileData), "---") {
		//nolint:gomnd // if file has empty line in the end.
		if len(resourceYAML) < 2 {
			continue
		}

		decoder := k8yaml.NewYAMLOrJSONDecoder(bytes.NewReader([]byte(resourceYAML)), 1)
		manifest := unstructured.Unstructured{}

		err := decoder.Decode(&manifest)
		if err != nil {
			//nolint:errorlint // this is okay.
			if err == io.EOF {
				break
			}

			return err
		}

		m.Resources = append(m.Resources, &Resource{
			Kind:         manifest.GetKind(),
			TemplateData: manifest.Object,
			Annotations:  manifest.GetAnnotations(),
		})
	}

	return nil
}

// Inject injects secrets to Resources.
func (m *Manifest) Inject() error {
	for _, resource := range m.Resources {
		if resource.Kind == "Secret" {
			err := resource.Replace(m)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// ToYAML serializes the completed template into YAML to be consumed by Kubernetes.
func (m *Manifest) ToYAML() (string, error) {
	var result []string

	for _, r := range m.Resources {
		if r != nil {
			resource, err := yaml.Marshal(r.TemplateData)
			if err != nil {
				return "", err
			}

			result = append(result, string(resource))
		}
	}

	return strings.Join(result, "---\n"), nil
}

//nolint:gocognit,cyclop // complexity of this function is fine.
// Replace replaces secret values with the manifest resource.
func (r *Resource) Replace(manifest *Manifest) error {
	//nolint:nestif // nested ifs are okay.
	if path, ok := r.Annotations[types.PathAnnotation]; ok {
		secret := &types.Secret{
			Bucket: manifest.S3Bucket,
			Path:   path,
		}

		err := manifest.S3Client.GetSecret(manifest.S3API, secret)
		if err != nil {
			var nsk *types2.NoSuchKey
			if errors.As(err, &nsk) {
				log.Debugf("secret path %s is not found in S3.", path)

				return nil
			}

			return err
		}

		if data, ok := r.TemplateData["data"]; ok {
			dataType := reflect.ValueOf(data).Kind()

			if dataType == reflect.Map {
				inner, ok := data.(map[string]interface{})
				if !ok {
					return nil
				}

				for key, placeholder := range inner {
					// Only if data field's value is string. e.g. <password>.
					if reflect.ValueOf(placeholder).Kind() == reflect.String {
						// Match placeholder data fields.
						r := regexp.MustCompile("<(.*?)>")
						if r.MatchString(placeholder.(string)) {
							// Extract placeholder name for Secret field.
							secretField := r.ReplaceAllString(placeholder.(string), "$1")

							val, err := secret.GetValue(secretField)
							if err != nil {
								log.Debugf("secret field %s is not found in path %s.", secretField, path)

								continue
							}

							inner[key] = base64.StdEncoding.EncodeToString([]byte(val))
						}
					}
				}
			}
		}
	}

	return nil
}
