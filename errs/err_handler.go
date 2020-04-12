package errs

import (
	"fmt"
	"net/http"
)

func ErrHandler(debug bool, err error,
	w *http.ResponseWriter, statusCode int) {
	if debug {
		http.Error(*w, err.Error(), statusCode)
	} else {
		http.Error(*w,
			fmt.Sprintf("Status code: %d", statusCode), statusCode)
	}
}
