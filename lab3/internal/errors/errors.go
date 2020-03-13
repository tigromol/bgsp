package errors

import "github.com/pkg/errors"

var ErrNotFound = errors.New("not found")

func New(message string) error {
	return errors.New(message)
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}
