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
	"strings"

	"net/http"

	"github.com/Kotyarich/find-your-pet/errs"
	"github.com/Kotyarich/find-your-pet/features"
	"github.com/Kotyarich/find-your-pet/features/normalizer"
	"github.com/Kotyarich/find-your-pet/features/paginator"
	"github.com/Kotyarich/find-your-pet/models"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
)

func (hd *HandlerData) FoundHandler(w http.ResponseWriter, r *http.Request) {
	arguments := r.URL.Query()
	strTypeId := arguments.Get("type_id")
	var typeId int
	var err error
	if strTypeId != "" {
		typeId, err = strconv.Atoi(strTypeId)
		if err != nil {
			errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
			return
		}
	}
	sex := arguments.Get("sex")
	breed := arguments.Get("breed")
	description := arguments.Get("description")
	var latitude, longitude float64
	strLatitude := arguments.Get("latitude")
	if strLatitude != "" {
		latitude, err = strconv.ParseFloat(strLatitude, 64)
		if err != nil {
			errs.ErrHandler(hd.DebugMode, err, &w, http.StatusBadRequest)
			return
		}
	}
	strLongitude := arguments.Get("longitude")
	if strLongitude != "" {
		longitude, err = strconv.ParseFloat(strLongitude, 64)
		if err != nil {
			errs.ErrHandler(hd.DebugMode, err, &w, http.StatusBadRequest)
			return
		}
	}
	found := &models.Found{
		TypeId:      typeId,
		Sex:         sex,
		Breed:       breed,
		Description: description,
		Latitude:    latitude,
		Longitude:   longitude,
	}
	mapCloseId := make(map[string]interface{})
	mapCloseId["close_id"] = viper.GetInt("found.close_id")
	ctx := context.WithValue(context.Background(), "params", mapCloseId)
	founds, err := hd.FoundController.Search(ctx, found)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Not found", http.StatusNotFound)
		} else {
			errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
		}
		return
	}
	pagesCount := paginator.CalculatePageCount(len(founds),
		hd.FoundController.GetItemsPerPageCount())
	foundsEncoded, err := json.Marshal(founds)
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
		return
	}
	pagesData := paginator.PaginatorData{
		Pages:   pagesCount,
		Payload: foundsEncoded,
	}
	data, err := json.Marshal(pagesData)
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
		return
	}
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
	}
	w.Write(data)
}

func (hd *HandlerData) FoundByIdGetHandler(w http.ResponseWriter, r *http.Request) {
	strId := r.URL.Query().Get("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusBadRequest)
		return
	}
	closeId := viper.GetInt("found.close_id")
	ctx := context.WithValue(context.Background(), "close_id", closeId)
	found, err := hd.FoundController.GetById(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			errs.ErrHandler(hd.DebugMode, err, &w, http.StatusNotFound)
		} else {
			errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
		}
		return
	}
	data, err := json.Marshal(found)
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
	}
	w.Write(data)
}

func (hd *HandlerData) AddFoundHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(0)
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusBadRequest)
		return
	}
	params := r.FormValue
	strTypeId := params("type_id")
	typeId, err := strconv.Atoi(strTypeId)
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusBadRequest)
		return
	}
	// author_id is a temprorary parameter
	strAuthorId := params("vk_id")
	authorId, err := strconv.Atoi(strAuthorId)
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusBadRequest)
		return
	}
	sex, err := normalizer.SexNormalize(params("sex"))
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusBadRequest)
		return
	}
	// The first letter must be in uppercase
	breed := strings.Title(
		strings.ToLower(params("breed")),
	)
	description := params("description")
	strLatitude := params("latitude")
	latitude, err := strconv.ParseFloat(strLatitude, 64)
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusBadRequest)
		return
	}
	strLongitude := params("longitude")
	longitude, err := strconv.ParseFloat(strLongitude, 64)
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusBadRequest)
		return
	}
	// It's a real file. The user sent it
	file, header, err := r.FormFile("picture")
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusBadRequest)
		return
	}
	extension := features.GetExtension(header.Filename)
	defer file.Close()
	// in MB
	fileMaxSize := viper.GetInt64("found.files.max_size") * 1024 * 1024
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
			errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
			return
		}
		w.Write(data)
		return
	}

	foundParams := &models.Found{
		TypeId:      typeId,
		AuthorId:    authorId,
		Sex:         sex,
		Breed:       breed,
		Description: description,
		Latitude:    latitude,
		Longitude:   longitude,
	}
	ctx, cancel := context.WithCancel(context.Background())
	foundIdCh := make(chan int)
	errCh := make(chan error, 1)
	// mFile - model of file. It's not a real file. It's only a record
	mFileCh := make(chan *models.File)
	// file is the real file a user sent
	go hd.FoundAddingManager.Add(ctx, foundParams, foundIdCh,
		mFileCh, errCh)

	var foundId int

addFoundId:
	for {
		select {
		case foundId = <-foundIdCh:
			break addFoundId
		case err = <-errCh:
			errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
			return
		}
	}

	baseFoundDirectoryPath := viper.GetString("found.files.directory")
	foundDirectoryPath := strconv.Itoa(foundId)
	fullDirectoryPath := filepath.Join(baseFoundDirectoryPath,
		foundDirectoryPath)
	err = os.MkdirAll(fullDirectoryPath,
		os.ModePerm)
	if err != nil {
		cancel()
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
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
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
		return
	}
	mFile := &models.File{
		Name: fileName,
		Path: filepath.Join(foundDirectoryPath, uuid),
	}
	_, err = io.Copy(dst, file)
	if err != nil {
		cancel()
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
		return
	}
	select {
	case err = <-errCh:
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
		return
	default:
		mFileCh <- mFile
	}

	if err = <-errCh; err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
		return
	}
	// Send id to the client
	w.Write([]byte(strconv.Itoa(foundId)))
}
