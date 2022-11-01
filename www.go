package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("header.html", "index.html", "footer.html")
	if err != nil {
		fmt.Println(w, err.Error())

	}
	t.ExecuteTemplate(w, "index", nil)
}

func create(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("create.html", "header.html", "footer.html")
	if err != nil {
		fmt.Println(w, err.Error())

	}
	t.ExecuteTemplate(w, "create", nil)
}

func saveUser(w http.ResponseWriter, r *http.Request) {

	userLastName := r.FormValue("UserLastName")
	userFirstName := r.FormValue("UserFirstName")
	phoneNombe := r.FormValue("PhoneNombe")
	typeEquipment := r.FormValue("TypeEquipment")
	brand := r.FormValue("Brand")
	model := r.FormValue("Model")
	sn := r.FormValue("SN")

	db, err := sql.Open("mysql", pass)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	dw := DataWrite{db: db}
	uw := UserWrite{firstName: userFirstName, lastName: userLastName, phone: phoneNombe}
	eq := Equipment{typeEquipment: typeEquipment, brand: brand, model: model, sn: sn}
	err = dw.dbWrite(uw, eq)
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func handleFunc() {
	http.Handle("/CSS/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	http.HandleFunc("/saveUser", saveUser)
	http.ListenAndServe(":8080", nil)
}
