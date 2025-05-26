package handler

import (
	"html/template"
	"net/http"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
    tpl := template.Must(template.ParseFiles(
    "templates/layout.html",
    "templates/index.html",
		))
	tpl.ExecuteTemplate(w, "layout", nil)
}
