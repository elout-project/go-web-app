package main

// https://gowebexamples.com/basic-middleware/

import (
	"fmt"
	"log"
	"net/http"
)

func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// your middleware code  start here
		log.Println(r.URL.Path)
		// your code end
		f(w, r)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "from foo")
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "from bar")
}

func main() {
	http.HandleFunc("/foo", logging(foo))
	http.HandleFunc("/bar", logging(bar))

	// Start server
	http.ListenAndServe(":80", nil)
}
