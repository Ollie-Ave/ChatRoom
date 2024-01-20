package api

import (
	"net/http"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	if email == "oliver@citruslime.com" && password == "millie88" {
		w.Header().Set("HX-Redirect", "/")
	} else {
		w.Write([]byte("Invalid email or password"))
	}
}
