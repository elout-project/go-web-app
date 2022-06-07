package main

import (
	"net/http"
	"web-app/handler"
	"web-app/salah"
)

func main() {
	// Routing
	http.HandleFunc("/", handler.Index)
	http.HandleFunc("/about", handler.About)
	http.HandleFunc("/salah", salah.Salah)

	// Start Server
	http.ListenAndServe(":80", nil)
}
