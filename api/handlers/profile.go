package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Kotyarich/find-your-pet/errs"
	"github.com/Kotyarich/find-your-pet/features"
	"github.com/Kotyarich/find-your-pet/features/paginator"
	"github.com/spf13/viper"
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
	closeId := viper.GetInt("lost.close_id")
	ctx := context.WithValue(
		context.Background(),
		"close_id",
		closeId,
	)
	lost, err := hd.ProfileController.GetLost(ctx, userId)
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
		return
	}
	pagesCount := paginator.CalculatePageCount(len(lost),
		hd.ProfileController.GetItemsPerPageCount())
	lostsEncoded, err := json.Marshal(lost)
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
		return
	}
	pagesData := paginator.PaginatorData{
		Pages:   pagesCount,
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
	opened, err := strconv.ParseBool(params("opened"))
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusBadRequest)
		return
	}
	ctx := context.WithValue(
		context.Background(),
		"params",
		features.StatusIdParams{
			OpenId:  viper.GetInt("lost.open_id"),
			CloseId: viper.GetInt("lost.close_id"),
		},
	)
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
		return
	}
	viper.GetInt("lost.open_id")
	err = hd.ProfileController.SetLostOpening(ctx, lostId, opened)
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
	closeId := viper.GetInt("found.close_id")
	ctx := context.WithValue(
		context.Background(),
		"close_id",
		closeId,
	)
	found, err := hd.ProfileController.GetFound(ctx, userId)
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
		return
	}
	pagesCount := paginator.CalculatePageCount(len(found),
		hd.ProfileController.GetItemsPerPageCount())
	foundEncoded, err := json.Marshal(found)
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
		return
	}
	pagesData := paginator.PaginatorData{
		Pages:   pagesCount,
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
	opened, err := strconv.ParseBool(params("opened"))
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusBadRequest)
		return
	}
	ctx := context.WithValue(
		context.Background(),
		"params",
		features.StatusIdParams{
			OpenId:  viper.GetInt("found.open_id"),
			CloseId: viper.GetInt("found.close_id"),
		},
	)
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
		return
	}
	viper.GetInt("found.open_id")
	err = hd.ProfileController.SetFoundOpening(ctx, foundId, opened)
	if err != nil {
		if err == errs.TheFoundNotFound {
			errs.ErrHandler(hd.DebugMode, err, &w, http.StatusNotFound)
		} else {
			errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
		}
		return
	}
}
