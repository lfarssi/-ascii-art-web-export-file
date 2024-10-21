package handlers

import (
	"net/http"
	"os"
)

func AssetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		files, err := os.ReadDir("assets")
		if err != nil {
			ErrorHandler(w, r, http.StatusInternalServerError)
			return
		}
		for _, file := range files {
			if r.URL.Path == "/assets/"+file.Name() {
				fs := http.Dir("assets")
				http.StripPrefix("/assets/", http.FileServer(fs)).ServeHTTP(w, r)
				return
			}
		}
		ErrorHandler(w, r, http.StatusNotFound)
		return
	}
	ErrorHandler(w, r, http.StatusMethodNotAllowed)
}
