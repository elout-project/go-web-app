package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"web-app/model"
)

func Index(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Welcome to my website Bro!")
	tmpl, err := template.ParseFiles("./public/index.html")
	if err != nil {
		fmt.Println(err)
	}
	customer := model.Customer{
		Nama: "Mudha",
		Umur: 47,
	}

	tmpl.Execute(w, customer)
}


