package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

func ExportHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		if data.Res != "" {
			length := strconv.Itoa(len(data.Res))
			w.Header().Set("Content-Type","Text/plan")
			w.Header().Set("Content-Length",length)
			w.Header().Set("Content-Disposition", "attachement; filename=ascii-art.txt")
			fmt.Fprint(w, data.Res)
			return
		} else {
			ErrorHandler(w, r, http.StatusBadRequest)
			return
		}
	} 
	ErrorHandler(w, r, http.StatusMethodNotAllowed)
}
