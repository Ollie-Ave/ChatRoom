package pages

import "github.com/gorilla/mux"

func AddPages(r *mux.Router) {
	r.HandleFunc("/", chatRoomHandler)
	r.HandleFunc("/login", loginHandler)
}
