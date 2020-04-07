package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"

	// "log"
	"net/http"

	"github.com/Kotyarich/find-your-pet/features"
	"github.com/Kotyarich/find-your-pet/models"
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
	// name := r.URL.Query().Get("name")
	losts, err := hd.LostController.Search(nil)
	// MOCK
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Not found", http.StatusNotFound)
		} else {
			http.Error(w, "Server Internal Error", http.StatusInternalServerError)
		}
		return
	}
	pagesCount := features.CalculatePageCount(len(losts),
		hd.LostController.GetItemsPerPageCount())
	lostsEncoded, err := json.Marshal(losts)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	pagesData := features.PaginatorData{
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
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	params := r.PostFormValue
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
	place := params("place")
	lost := &models.Lost{
		TypeId:      typeId,
		AuthorId:    authorId,
		Sex:         sex,
		Breed:       breed,
		Description: description,
		Place:       place,
	}
	addedId, err := hd.LostController.Add(lost)
	if err != nil {
		http.Error(w, "Server Internal Error", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(fmt.Sprintf("Added with id %d\n", addedId)))
}
