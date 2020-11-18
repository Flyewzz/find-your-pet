package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Kotyarich/find-your-pet/errs"
)

func (hd *HandlerData) BreedClassifierHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var reqPicture struct {
		Picture string `json:"picture"`
	}
	err := decoder.Decode(&reqPicture)
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
		return
	}
	picture := reqPicture.Picture
	b64data := picture[strings.IndexByte(picture, ',')+1:]
	breeds, err := hd.breedClassifier.GetBreeds(b64data)
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(breeds)
	if err != nil {
		errs.ErrHandler(hd.DebugMode, err, &w, http.StatusInternalServerError)
		return
	}
	w.Write(data)
}
