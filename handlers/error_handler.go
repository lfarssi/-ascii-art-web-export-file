package handlers

import (
	"html/template"
	"net/http"
)

type MyErr struct {
	StatusCode int
	Error      string
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, statusCode int) {
	t, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, "500 | Internal Server Error !", http.StatusInternalServerError)
		return

	}
	typeError := MyErr{
		StatusCode: statusCode,
		Error: http.StatusText(statusCode),
	}
	w.WriteHeader(statusCode)
	if err := t.Execute(w, typeError); err != nil {
		http.Error(w, "500 | Internal Server Error !", http.StatusInternalServerError)
		return
		}
}
