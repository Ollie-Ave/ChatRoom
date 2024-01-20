package main

import (
	"app/api"
	"app/pages"
	"app/shared"

	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	appSettings, err := shared.GetAppSettings()

	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(appSettings.StaticFileDirectory))))

	api.AddApiEndpoints(r)
	pages.AddPages(r)

	fmt.Printf("Server starting, listening on port %d...", appSettings.ApplicationPort)
	http.ListenAndServe(fmt.Sprintf(":%d", appSettings.ApplicationPort), r)
}
