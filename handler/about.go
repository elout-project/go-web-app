package handler

import (
	"fmt"
	"net/http"
)

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is About Pages")
}
