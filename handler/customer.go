package handler

import (
	"fmt"
	"html/template"
	"net/http"
	_ "web-app/model" //_ you may un use this library
)

func Index(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Welcome to my website Bro!")

	tmpl, err := template.ParseFiles("./public/index.html")
	if err != nil {
		fmt.Println(err)
	}
	// customer := model.Customer{
	// 	Nama: "Mudha",
	// 	Umur: 47,
	// }

	customer := struct {
		Nama string
		Umur int
	}{Nama: "Mudha", Umur: 47}

	tmpl.Execute(w, customer)
}
