package secret

import "fmt"

// NotFound occurs when no secret value found in a path.
type NotFound struct {
	Key    string
	Secret *Secret
}

func (e NotFound) Error() string {
	return fmt.Sprintf("no secret found for %s in path %s", e.Key, e.Secret.Path)
}
