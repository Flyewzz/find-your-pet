package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Kotyarich/find-your-pet/errs"
	"github.com/Kotyarich/find-your-pet/features/paginator"
)

func (hd *HandlerData) ProfileLostHandler(w http.ResponseWriter, r *http.Request) {
	strUserId := r.URL.Query().Get("vk_id")
	var userId int
	var err error
	if strUserId != "" {
		userId, err = strconv.Atoi(strUserId)
		if err != nil {
			errs.ErrHandler(hd.DebugMode, err, &w, http.StatusBadRequest)
			return
		}
	}
	lost, err := hd.ProfileController.GetLost(context.Background(), userId)
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
		return
	}
	lostsEncoded, err := json.Marshal(lost)
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
		return
	}
	pagesData := paginator.PaginatorData{
		HasMore: true,
		Payload: lostsEncoded,
	}
	data, err := json.Marshal(pagesData)
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func (hd *HandlerData) ProfileLostOpeningHandler(w http.ResponseWriter, r *http.Request) {
	params := r.FormValue
	strId := params("lost_id")
	lostId, err := strconv.Atoi(strId)
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusBadRequest)
		return
	}
	statusId, err := strconv.Atoi(params("status_id"))
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusBadRequest)
		return
	}
	err = hd.ProfileController.SetLostOpening(context.Background(), lostId, statusId)
	if err != nil {
		if err == errs.LostNotFound {
			errs.ErrHandler(hd.DebugMode, err, &w, http.StatusNotFound)
		} else {
			errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
		}
		return
	}
}

func (hd *HandlerData) ProfileFoundHandler(w http.ResponseWriter, r *http.Request) {
	strUserId := r.URL.Query().Get("vk_id")
	var userId int
	var err error
	if strUserId != "" {
		userId, err = strconv.Atoi(strUserId)
		if err != nil {
			errs.ErrHandler(hd.DebugMode, err, &w, http.StatusBadRequest)
			return
		}
	}
	found, err := hd.ProfileController.GetFound(context.Background(), userId)
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
		return
	}
	foundEncoded, err := json.Marshal(found)
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
		return
	}
	pagesData := paginator.PaginatorData{
		HasMore: true,
		Payload: foundEncoded,
	}
	data, err := json.Marshal(pagesData)
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func (hd *HandlerData) ProfileFoundOpeningHandler(w http.ResponseWriter, r *http.Request) {
	params := r.FormValue
	strId := params("found_id")
	foundId, err := strconv.Atoi(strId)
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusBadRequest)
		return
	}
	statusId, err := strconv.Atoi(params("status_id"))
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusBadRequest)
		return
	}
	err = hd.ProfileController.SetFoundOpening(context.Background(), foundId, statusId)
	if err != nil {
		if err == errs.TheFoundNotFound {
			errs.ErrHandler(hd.DebugMode, err, &w, http.StatusNotFound)
		} else {
			errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
		}
		return
	}
}
