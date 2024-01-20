package pages

import (
	"app/shared"
	"net/http"
	"text/template"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles(
		shared.GetTemplatePath("pages/login.html"),
		shared.GetTemplatePath("components/spinner.html"),
	)

	if err != nil {
		panic(err)
	}

	template.Execute(w, nil)
}
