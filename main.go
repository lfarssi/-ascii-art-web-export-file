package main

import (
	"fmt"
	"net/http"

	"ascii-art/handlers"
)

func main() {
	http.HandleFunc("/assets/", handlers.AssetHandler)
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/ascii-art", handlers.AsciiProcessHandler)
	http.HandleFunc("/export", handlers.ExportHandler)

	
	fmt.Println("Server is running at http://localhost:8088")
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
