package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
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
)

func (hd *HandlerData) LostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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
	address := arguments.Get("address")
	description := arguments.Get("description")
	strPage := arguments.Get("page")
	page, err := strconv.Atoi(strPage)
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusBadRequest)
		return
	}
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
	query := arguments.Get("query")
	lost := &models.Lost{
		TypeId:      typeId,
		Sex:         sex,
		Breed:       breed,
		Description: description,
		Latitude:    latitude,
		Longitude:   longitude,
		Address:     address,
	}
	mapCloseId := make(map[string]interface{})
	ctx := context.WithValue(context.Background(), "params", mapCloseId)
	losts, hasMore, err := hd.LostController.Search(ctx, lost, query, page)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Not found", http.StatusNotFound)
		} else {
			errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
		}
		return
	}
	lostsEncoded, err := json.Marshal(losts)
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
		return
	}
	pagesData := paginator.PaginatorData{
		HasMore: hasMore,
		Payload: lostsEncoded,
	}
	data, err := json.Marshal(pagesData)
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func (hd *HandlerData) LostByIdGetHandler(w http.ResponseWriter, r *http.Request) {
	strId := r.URL.Query().Get("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusBadRequest)
		return
	}
	lost, err := hd.LostController.GetById(context.Background(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			errs.ErrHandler(hd.DebugMode, err, &w, http.StatusNotFound)
		} else {
			errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
		}
		return
	}
	data, err := json.Marshal(lost)
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func (hd *HandlerData) AddLostHandler(w http.ResponseWriter, r *http.Request) {
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
	// author_id is a temprorarily parameter
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
	address := params("address")
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
	// It's a real file. The user can send it
	file, header, err := r.FormFile("picture")
	// An empty file is not error
	if err != nil && err != http.ErrMissingFile {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusBadRequest)
		return
	}
	var extension string
	if err != http.ErrMissingFile {
		defer file.Close()
		extension = features.GetExtension(header.Filename)
		// in MB
		fileMaxSize := hd.FileMaxSize * 1024 * 1024
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
	}
	lostParams := &models.Lost{
		TypeId:      typeId,
		AuthorId:    authorId,
		Sex:         sex,
		Breed:       breed,
		Description: description,
		Latitude:    latitude,
		Longitude:   longitude,
		Address:     address,
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
			errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
			return
		}
	}
	if err == http.ErrMissingFile {
		select {
		case err = <-errCh:
			errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
			return
		default:
			mFileCh <- nil
		}

		if err = <-errCh; err != nil {
			errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
			return
		}

		w.Write([]byte(strconv.Itoa(lostId)))
		return
	}
	// Generate UUID key as a filename to store it into the temporary folder
	// uuid will also contain a file extension
	uuid := uuid.NewV4().String() + "." + extension
	fileName := header.Filename
	err = hd.FileStoreController.Save(&file, lostId, "lost", uuid)
	if err != nil {
		cancel()
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
		return
	}
	mFile := &models.File{
		Name: fileName,
		Path: filepath.Join(strconv.Itoa(lostId), uuid),
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
	w.Write([]byte(strconv.Itoa(lostId)))
}

func (hd *HandlerData) RemoveLostHandler(w http.ResponseWriter, r *http.Request) {
	strId := r.URL.Query().Get("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusBadRequest)
		return
	}
	err = hd.LostAddingManager.Remove(id)
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
		return
	}
}

func (hd *HandlerData) LostNotifyHandler(w http.ResponseWriter, r *http.Request) {
	losts, err := hd.LostController.GetAll()
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
		return
	}
	type Notify struct {
		AuthorId int              `json:"vk_id"`
		Similars []models.Similar `json:"similars"`
	}
	var notifyCollection []Notify
	for _, lost := range losts {
		similars, err := hd.FoundController.GetSimilars(&lost)
		if err != nil {
			continue
		}
		if similars == nil {
			continue
		}
		notifyCollection = append(
			notifyCollection,
			Notify{
				AuthorId: lost.AuthorId,
				Similars: similars,
			},
		)
	}
	data, err := json.Marshal(notifyCollection)
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
