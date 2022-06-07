package main

import (
	"net/http"
	"web-app/handler"
	_ "web-app/salah"

	"github.com/gorilla/mux"
)

func main() {
	// // Routing
	// http.HandleFunc("/", handler.Index)
	// http.HandleFunc("/about", handler.About)
	// // http.HandleFunc("/salah", salah.Salah)
	// // Start Server
	// http.ListenAndServe(":80", nil)

	r := mux.NewRouter()
	r.HandleFunc("/", handler.Index)
	r.HandleFunc("/about", handler.About)
	http.ListenAndServe(":80", r)
}
