package api

import (
	"github.com/Kotyarich/find-your-pet/api/handlers"
	"github.com/gorilla/mux"
)

func ConfigureHandlers(r *mux.Router, hd *handlers.HandlerData) {
	// Lost
	r.HandleFunc("/losts", hd.LostHandler).Methods("GET")
	r.HandleFunc("/lost", hd.LostByIdGetHandler).Methods("GET")
	r.HandleFunc("/lost", hd.AddLostHandler).Methods("POST")

	// LostFile
	r.HandleFunc("/lost/img", hd.LostImageHandler).Methods("GET")

	// Found
	r.HandleFunc("/founds", hd.FoundHandler).Methods("GET")
	r.HandleFunc("/found", hd.FoundByIdGetHandler).Methods("GET")
	r.HandleFunc("/found", hd.AddFoundHandler).Methods("POST")

	// FoundFile
	r.HandleFunc("/found/img", hd.FoundImageHandler).Methods("GET")

	// Profile
	r.HandleFunc("/profile/lost", hd.ProfileLostHandler).Methods("GET")
	r.HandleFunc("/lost", hd.ProfileLostOpeningHandler).Methods("PUT")
	r.HandleFunc("/profile/found", hd.ProfileFoundHandler).Methods("GET")
	r.HandleFunc("/found", hd.ProfileFoundOpeningHandler).Methods("PUT")

}
