package errors

import (
	"errors"
	"fmt"
)

var (
	errNotFound = errors.New("not found")
)

func ResourceNotFoundError(key string) error {
	return fmt.Errorf("resource - %s %w", key, errNotFound)
}
