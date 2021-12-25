package secret

import "fmt"

// NotFoundError occurs when no secret value found in a path.
type NotFoundError struct {
	Key    string
	Secret *Secret
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("no secret found for %s in path %s", e.Key, e.Secret.Path)
}

// FieldNotFoundError occurs when no field found in the secret.
type FieldNotFoundError struct {
	Field  string
	Secret *Secret
}

func (e FieldNotFoundError) Error() string {
	return fmt.Sprintf("no field found for %s in path %s", e.Field, e.Secret.Path)
}
