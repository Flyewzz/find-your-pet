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
	r.HandleFunc("/lost", hd.RemoveLostHandler).Methods("DELETE")
	r.HandleFunc("/lost/notification", hd.LostNotifyHandler).Methods("GET")

	// LostFile
	r.HandleFunc("/lost/img", hd.LostImageHandler).Methods("GET")

	// Found
	r.HandleFunc("/founds", hd.FoundHandler).Methods("GET")
	r.HandleFunc("/found", hd.FoundByIdGetHandler).Methods("GET")
	r.HandleFunc("/found", hd.AddFoundHandler).Methods("POST")
	r.HandleFunc("/found", hd.RemoveFoundHandler).Methods("DELETE")
	r.HandleFunc("/found/notification", hd.FoundNotifyHandler).Methods("GET")

	// FoundFile
	r.HandleFunc("/found/img", hd.FoundImageHandler).Methods("GET")

	// Profile
	r.HandleFunc("/profile/lost", hd.ProfileLostHandler).Methods("GET")
	r.HandleFunc("/lost", hd.ProfileLostOpeningHandler).Methods("PUT")
	r.HandleFunc("/profile/found", hd.ProfileFoundHandler).Methods("GET")
	r.HandleFunc("/found", hd.ProfileFoundOpeningHandler).Methods("PUT")

	// Breed classifier (Python)
	r.HandleFunc("/breed", hd.BreedClassifierHandler).Methods("POST")

}
