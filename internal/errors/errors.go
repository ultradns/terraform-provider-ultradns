package errors

import (
	"errors"
	"fmt"
)

var (
	errNotFound     = errors.New("not found")
	errNotDestroyed = errors.New("not destroyed")
)

func ResourceNotFoundError(key string) error {
	return fmt.Errorf("resource - %s %w", key, errNotFound)
}

func ResourceNotDestroyedError(key string) error {
	return fmt.Errorf("resource - %s %w", key, errNotDestroyed)
}
