package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"

	// "log"
	"net/http"

	"github.com/Kotyarich/find-your-pet/features"
	"github.com/Kotyarich/find-your-pet/features/paginator"
	"github.com/Kotyarich/find-your-pet/models"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
)

func (hd *HandlerData) LostHandler(w http.ResponseWriter, r *http.Request) {
	// var page int = 0
	// var err error
	// strPage := r.URL.Query().Get("page")
	// if strPage != "" {
	// 	page, err = strconv.Atoi(strPage)
	// 	if err != nil {
	// 		http.Error(w, "Bad request", http.StatusBadRequest)
	// 		return
	// 	}
	// }
	arguments := r.URL.Query()
	strTypeId := arguments.Get("type_id")
	var typeId int
	var err error
	typeId, err = strconv.Atoi(strTypeId)
	sex := arguments.Get("sex")
	breed := arguments.Get("breed")
	description := arguments.Get("description")
	var latitude, longitude float64
	strLatitude := arguments.Get("latitude")
	if strLatitude != "" {
		latitude, err = strconv.ParseFloat(strLatitude, 64)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
	}
	strLongitude := arguments.Get("longitude")
	if strLongitude != "" {
		longitude, err = strconv.ParseFloat(strLongitude, 64)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
	}
	date := arguments.Get("date")
	lost := &models.Lost{
		TypeId:      typeId,
		Sex:         sex,
		Breed:       breed,
		Description: description,
		Date:        date,
		Latitude:    latitude,
		Longitude:   longitude,
	}
	losts, err := hd.LostController.Search(lost)
	// MOCK
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Not found", http.StatusNotFound)
		} else {
			http.Error(w, "Server Internal Error", http.StatusInternalServerError)
		}
		return
	}
	pagesCount := paginator.CalculatePageCount(len(losts),
		hd.LostController.GetItemsPerPageCount())
	lostsEncoded, err := json.Marshal(losts)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	pagesData := paginator.PaginatorData{
		Pages:   pagesCount,
		Payload: lostsEncoded,
	}
	data, err := json.Marshal(pagesData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func (hd *HandlerData) LostByIdGetHandler(w http.ResponseWriter, r *http.Request) {
	strId := r.URL.Query().Get("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	lost, err := hd.LostController.GetById(id)
	if err != nil {
		http.Error(w, "Server Internal Error", http.StatusInternalServerError)
		return
	}
	data, _ := json.Marshal(lost)
	w.Write(data)
}

func (hd *HandlerData) AddLostHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(0)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	params := r.FormValue
	strTypeId := params("type_id")
	typeId, err := strconv.Atoi(strTypeId)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	// author_id is a temprorary parameter
	strAuthorId := params("author_id")
	authorId, err := strconv.Atoi(strAuthorId)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	sex := params("sex")
	breed := params("breed")
	description := params("description")
	strLatitude := params("latitude")
	latitude, err := strconv.ParseFloat(strLatitude, 64)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	strLongitude := params("longitude")
	longitude, err := strconv.ParseFloat(strLongitude, 64)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	// It's a real file. The user sent it
	file, header, err := r.FormFile("picture")
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	extension := features.GetExtension(header.Filename)
	if !features.IsExtensionPicture(extension) {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	defer file.Close()
	// in MB
	fileMaxSize := viper.GetInt64("lost.files.max_size") * 1024 * 1024
	if header.Size > fileMaxSize {
		w.WriteHeader(http.StatusBadRequest)
		type ErrStruct struct {
			err string `json:"error"`
		}
		errMsg := ErrStruct{
			err: fmt.Sprintf("File cannot be larger than %d MB",
				fileMaxSize),
		}
		data, err := json.Marshal(errMsg)
		if err != nil {
			http.Error(w, "Server Internal Error", http.StatusInternalServerError)
			return
		}
		w.Write(data)
		return
	}

	lostParams := &models.Lost{
		TypeId:      typeId,
		AuthorId:    authorId,
		Sex:         sex,
		Breed:       breed,
		Description: description,
		Latitude:    latitude,
		Longitude:   longitude,
	}
	ctx, cancel := context.WithCancel(context.Background())
	lostIdCh := make(chan int)
	errCh := make(chan error, 1)
	// mFile - model of file. It's not a real file. It's only a record
	mFileCh := make(chan *models.File)
	// file is the real file a user sent
	go hd.LostAddingManager.Add(ctx, lostParams, lostIdCh,
		mFileCh, errCh)

	var lostId int

addLostId:
	for {
		select {
		case lostId = <-lostIdCh:
			break addLostId
		case err = <-errCh:
			http.Error(w, "Server Internal Error", http.StatusInternalServerError)
			return
		}
	}

	baseLostDirectoryPath := viper.GetString("lost.files.directory")
	lostDirectoryPath := strconv.Itoa(lostId)
	fullDirectoryPath := filepath.Join(baseLostDirectoryPath,
		lostDirectoryPath)
	err = os.MkdirAll(fullDirectoryPath,
		os.ModePerm)
	if err != nil {
		cancel()
		http.Error(w, "Server Internal Error", http.StatusInternalServerError)
		return
	}
	// Generate UUID key as a filename to store it into the temporary folder
	// uuid will also contain a file extension
	uuid := uuid.NewV4().String() + "." + extension
	fileName := header.Filename
	//Create a name with an extension for the file
	dst, err := os.Create(filepath.Join(
		fullDirectoryPath,
		uuid))
	if err != nil {
		cancel()
		http.Error(w, "Server Internal Error", http.StatusInternalServerError)
		return
	}
	mFile := &models.File{
		Name: fileName,
		Path: filepath.Join(lostDirectoryPath, uuid),
	}
	_, err = io.Copy(dst, file)
	if err != nil {
		cancel()
		http.Error(w, "Server Internal Error", http.StatusInternalServerError)
		return
	}
	select {
	case err = <-errCh:
		http.Error(w, "Server Internal Error", http.StatusInternalServerError)
		return
	default:
		mFileCh <- mFile
	}

	if err = <-errCh; err != nil {
		http.Error(w, "Server Internal Error", http.StatusInternalServerError)
		return
	}
	// Send id to the client
	w.Write([]byte(strconv.Itoa(lostId)))
}
