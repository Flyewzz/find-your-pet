package errs

import "github.com/pkg/errors"

var (
	TheFoundNotFound = errors.New("The found pet is not found")
)
