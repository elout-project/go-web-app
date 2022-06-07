package salah

import (
	"fmt"
	"net/http"
)

func Salah(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Error Bro")
}
