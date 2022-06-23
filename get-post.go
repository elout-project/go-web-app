package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			w.Write([]byte("POST request"))
		case "GET":
			w.Write([]byte("GET request"))
		default:
			http.Error(w, "", http.StatusBadRequest)
		}
	})

	fmt.Println("Serving http at http://localhost")
	http.ListenAndServe(":80", nil)
}
