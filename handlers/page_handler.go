package handlers

import (
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/"  {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}
	t, err := template.ParseFiles("templates/home.html")
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, nil); err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
}
