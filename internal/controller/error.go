package controller

import "fmt"

// NoObjectVersionsFoundError occurs when no provider found.
type NoObjectVersionsFoundError struct {
	Bucket string
	Path   string
}

func (e NoObjectVersionsFoundError) Error() string {
	return fmt.Sprintf("no object versions found in bucket %s with path %s", e.Bucket, e.Path)
}
