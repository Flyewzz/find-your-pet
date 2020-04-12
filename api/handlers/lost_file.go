package handlers

import (
	"database/sql"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Kotyarich/find-your-pet/errs"
	"github.com/spf13/viper"
)

func (hd *HandlerData) LostImageHandler(w http.ResponseWriter, r *http.Request) {
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
	file, err := hd.LostFileController.GetById(id)
	if err != nil {
		if err == sql.ErrNoRows {
			errs.ErrHandler(hd.DebugMode, err, &w, http.StatusNotFound)
			return
		}
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
		return
	}
	baseLostFileDirectory := viper.GetString("lost.files.directory")
	// A database stores only the smallest part of the file path
	openedFile, err := os.Open(filepath.Join(
		baseLostFileDirectory,
		file.Path),
	)
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
		return
	}
	defer openedFile.Close()
	//Get the Content-Type of the file
	//Create a buffer to store the header of the file in
	FileHeader := make([]byte, 512)
	//Copy the headers into the FileHeader buffer
	openedFile.Read(FileHeader)
	//Get content type of file
	FileContentType := http.DetectContentType(FileHeader)

	//Get the file size
	FileStat, _ := openedFile.Stat()                   //Get info from the file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string
	//Send the headers
	// w.Header().Set("Content-Disposition", "attachment; filename="+file.Name)
	w.Header().Set("Content-Type", FileContentType)
	w.Header().Set("Content-Length", FileSize)

	//Send the file
	//We read 512 bytes from the file already, so we reset the offset back to 0
	openedFile.Seek(0, 0)
	io.Copy(w, openedFile) //'Copy' the file to the client
	data, err := ioutil.ReadAll(openedFile)
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
		return
	}
	w.Write(data)
}
