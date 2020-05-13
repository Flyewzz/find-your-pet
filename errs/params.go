package errs

import "github.com/pkg/errors"

var (
	IncorrectGender = errors.New("A gender is incorrect")
)
