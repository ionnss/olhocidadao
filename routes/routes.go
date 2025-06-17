package routes

import (
	"net/http"

	"olhocidadao/handlers"

	"github.com/gorilla/mux"
)

func CreateRoutes() *mux.Router {
	r := mux.NewRouter()

	// Server static files
	fileServer := http.FileServer(http.Dir("static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fileServer))

	// Pages Routes
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	}).Methods("GET")

	r.HandleFunc("/report", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/report.html")
	}).Methods("GET")

	// Report Submission
	r.HandleFunc("/submit-report", handlers.SubmitReport).Methods("POST")

	return r
}
