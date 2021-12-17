package helper

import (
	"errors"
	"fmt"
)

var (
	errMismatchData = errors.New("mismatched")
	errZoneName     = errors.New("zone name must be a FQDN")
	errNotFound     = errors.New("not found")
	errNil          = errors.New("nil")
)

func MismatchError(key, expected, found string) error {
	return fmt.Errorf("%s %w : expected - %s : found - %s", key, errMismatchData, expected, found)
}

func ZoneNameValidationError() error {
	return fmt.Errorf("%w", errZoneName)
}

func ResourceNotFoundError(key string) error {
	return fmt.Errorf("resource - %s %w", key, errNotFound)
}

func DataNilError(key string) error {
	return fmt.Errorf("%s %w", key, errNil)
}
