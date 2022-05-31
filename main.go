package main

import (
	"net/http"
	"web-app/handler"
)

func main() {
	// Routing
	http.HandleFunc("/", handler.Index)
	http.HandleFunc("/about", handler.About)

	// Start Server
	http.ListenAndServe(":80", nil)
}
