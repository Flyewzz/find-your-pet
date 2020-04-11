package errs

import "github.com/pkg/errors"

var (
	FileError                = errors.New("File error")
	FileOperationInterrupted = errors.New("File operation was interrupted")
	MissedTransaction        = errors.New("A transaction is missed")
)
