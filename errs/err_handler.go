package errs

import (
	"net/http"
)

func ErrHandler(debug bool, err error,
	w *http.ResponseWriter, statusCode int) {
	if debug {
		http.Error(*w, err.Error(), statusCode)
	} else {
		(*w).WriteHeader(statusCode)
	}
}
