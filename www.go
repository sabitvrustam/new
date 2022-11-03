package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
)

//var result = []UserRead{}

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
	userMidlName := r.FormValue("UserMidlName")
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
	uw := User{FirstName: userFirstName, LastName: userLastName, MidlName: userMidlName, Phone: phoneNombe}
	eq := Equipment{TypeEquipment: typeEquipment, Brand: brand, Model: model, Sn: sn}
	err = dw.dbWrite(uw, eq)
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func userStatus(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("header.html", "userStatus.html", "footer.html")
	if err != nil {
		fmt.Println(w, err.Error())

	}

	db, err := sql.Open("mysql", pass)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	dw := DataRead{db: db}
	id := r.FormValue("id")
	result, trig := dw.dbRead(id)
	fmt.Println(result.FirstName, result.LastName, result.Phone)
	fmt.Println(trig.TypeEquipment, trig.Brand, trig.Model, trig.Sn)
	t.ExecuteTemplate(w, "userStatus", result)

}

func handleFunc() {
	http.Handle("/CSS/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	http.HandleFunc("/saveUser", saveUser)
	http.HandleFunc("/userStatus", userStatus)
	http.ListenAndServe(":8080", nil)
}
