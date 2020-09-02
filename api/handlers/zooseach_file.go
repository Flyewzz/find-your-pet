package handlers

import (
	"database/sql"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/Kotyarich/find-your-pet/errs"
	"github.com/spf13/viper"
)

func (hd *HandlerData) ImageZoosearchHandler(w http.ResponseWriter, r *http.Request) {
	strId := r.URL.Query().Get("id")
	if strId == "" {
		errs.ErrHandler(hd.DebugMode, errors.New("Id is missing"),
			&w, http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(strId)
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusBadRequest)
		return
	}
	if id < 1 {
		errs.ErrHandler(hd.DebugMode, errors.New("Id is incorrect"), &w, http.StatusBadRequest)
		return
	}
	file, err := hd.FileController.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			errs.ErrHandler(hd.DebugMode, err, &w, http.StatusNotFound)
			return
		}
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
		return
	}
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
