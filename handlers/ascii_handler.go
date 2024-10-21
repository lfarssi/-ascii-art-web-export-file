package handlers

import (
	"html/template"
	"net/http"
	"strings"

	ascii "ascii-art/functions"
)

type Data struct {
	Str    string
	Banner string
	Res    string
	A      template.HTML
}

var data Data

func AsciiProcessHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	}
	temp, err := template.ParseFiles("templates/home.html")
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}


	data.Str = r.FormValue("data")
	if data.Str == "" {
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}
	if len(data.Str) > 500 {
		ErrorHandler(w, r, http.StatusBadRequest)
		return
	}

	data.Banner = r.FormValue("banner")
	if !ascii.BannerExists(data.Banner) {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}

	data.Str = strings.ReplaceAll(data.Str, "\r\n", "\n")

	data.Res = ascii.TraitmentData(w, data.Banner, data.Str)
	if data.Res == "" { // If TraitmentData failed to generate the result
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
	data.A = template.HTML(strings.ReplaceAll(data.Res, "\n", "<br>"))

	if err := temp.Execute(w, data); err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError)
		return
	}
}
