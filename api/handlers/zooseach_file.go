package handlers

import (
	"io"
	"net/http"

	"github.com/Kotyarich/find-your-pet/errs"
	"github.com/Kotyarich/find-your-pet/models"
	"github.com/spf13/viper"
)

func (hd *HandlerData) ImageZoosearchHandler(w http.ResponseWriter, r *http.Request) {
	file := r.Context().Value("file").(*models.File)
	url := viper.GetString("zoosearch.files.path")
	fullPath := url + file.Path
	resp, err := http.Get(fullPath)
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	_, err = io.Copy(w, resp.Body) //'Copy' the file to the client
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
		return
	}
}
