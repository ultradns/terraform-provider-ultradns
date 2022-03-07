package errors

import (
	"errors"
	"fmt"
)

var (
	errNotFound     = errors.New("not found")
	errNotDestroyed = errors.New("not destroyed")
	errMismatched   = errors.New("mismatched")
)

func ResourceNotFoundError(key string) error {
	return fmt.Errorf("resource - %s %w", key, errNotFound)
}

func ResourceNotDestroyedError(key string) error {
	return fmt.Errorf("resource - %s %w", key, errNotDestroyed)
}

func ResourceTypeMismatched(expected, found string) error {
	return fmt.Errorf("resource schema %w : expected - %s : found - %s", errMismatched, expected, found)
}

func ProbeResourceNotFound(key string) error {
	return fmt.Errorf("probe resource of type %s %w : try with other criteria fields", key, errNotFound)
}
