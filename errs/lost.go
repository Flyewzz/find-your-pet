package errs

import "github.com/pkg/errors"

var (
	LostNotFound = errors.New("Lost is not found")
)
