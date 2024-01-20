package api

import "github.com/gorilla/mux"

func AddApiEndpoints(r *mux.Router) {
	r.HandleFunc("/api/login", loginHandler).Methods("POST")
}
