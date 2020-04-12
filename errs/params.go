package errs

import "github.com/pkg/errors"

var (
	IncorrectGender = errors.New("The gender is incorrect")
)
