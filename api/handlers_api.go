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

	// // University
	// r.HandleFunc("/universities", hd.UniversitiesHandler).Methods("GET")
	// r.HandleFunc("/university", hd.UniversityByIdGetHandler).Methods("GET")
	// r.HandleFunc("/university", hd.AddUniversityHandler).Methods("POST")
	// r.HandleFunc("/university", hd.UniversityByIdRemoveHandler).Methods("DELETE")
	// r.HandleFunc("/universities", hd.AllUniversitiesRemoveHandler).Methods("DELETE")

	// // Subject
	// r.HandleFunc("/university/{id}/subject", hd.AddSubjectHandler).Methods("POST")
	// r.HandleFunc("/university/{id}/subjects", hd.SubjectsHandler).Methods("GET")
	// r.HandleFunc("/subject", hd.SubjectByIdGetHandler).Methods("GET")
	// r.HandleFunc("/subject", hd.SubjectByIdRemoveHandler).Methods("DELETE")
	// r.HandleFunc("/university/{id}/subjects", hd.AllSubjectsRemoveHandler).Methods("DELETE")

	// // Material
	// r.HandleFunc("/subject/{id}/material", hd.AddMaterialHandler).Methods("POST")
	// r.HandleFunc("/subject/{id}/materials", hd.MaterialsHandler).Methods("GET")
	// r.HandleFunc("/material", hd.MaterialByIdGetHandler).Methods("GET")
	// r.HandleFunc("/material", hd.MaterialByIdRemoveHandler).Methods("DELETE")
	// r.HandleFunc("/subject/{id}/materials", hd.AllMaterialsRemoveHandler).Methods("DELETE")
}
