package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("header.html", "index.html", "footer.html")
	if err != nil {
		fmt.Println(w, err.Error(), "не удалось открыть главную страничку")
	}
	t.ExecuteTemplate(w, "index", nil)
}
func create(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("header.html", "create.html", "footer.html")
	if err != nil {
		fmt.Println(w, err.Error(), "не удалось открыть страничку создания заказа")
	}
	result := dbreadMasters()
	t.ExecuteTemplate(w, "create", result)
}
func newUser(w http.ResponseWriter, r *http.Request) {

	userLastName := r.FormValue("UserLastName")
	userFirstName := r.FormValue("UserFirstName")
	userMidlName := r.FormValue("UserMidlName")
	phoneNombe := r.FormValue("PhoneNombe")
	typeEquipment := r.FormValue("TypeEquipment")
	brand := r.FormValue("Brand")
	model := r.FormValue("Model")
	sn := r.FormValue("SN")
	defect := r.FormValue("defect")
	master := r.FormValue("Id")
	status := ("1")

	uw := Order{
		Status: Status{
			StatusOrder: status},
		Masters: Masters{
			Id: master},
		User: User{
			FirstName: userFirstName,
			LastName:  userLastName,
			MidlName:  userMidlName,
			Phone:     phoneNombe},
		Device: Device{
			TypeEquipment: typeEquipment,
			Brand:         brand,
			Model:         model,
			Sn:            sn,
			Defect:        defect}}
	err := dbWrite(uw)
	if err != nil {
		fmt.Println(err)
	} else {
		http.Redirect(w, r, "/index", http.StatusSeeOther)
	}
}
func userStatusPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("header.html", "userStatus.html", "footer.html")
	if err != nil {
		fmt.Println(w, err.Error())
	}
	id := r.FormValue("id")
	var result Order
	if id != "" {
		result = dbRead(id)
	}
	t.ExecuteTemplate(w, "userStatus", result)
}
func makeChanges(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("header.html", "makeChanges.html", "footer.html")
	if err != nil {
		fmt.Println(w, err.Error(), "не удалось открыть страничку создания заказа")
	}
	var result Order
	pas := dbreadParts()
	result.OllParts = append(result.OllParts, pas.OllParts...)
	id := r.FormValue("id")
	if id != "" {
		result = dbRead(id)
	}
	var partsVhang []string
	err = r.ParseForm()
	if err != nil {
		fmt.Println("не удалось считать форму")
	}
	for _, i := range r.Form {
		partsVhang = append(partsVhang, i...)
	}
	fmt.Println(partsVhang)
	t.ExecuteTemplate(w, "makeChanges", result)
}
func handleFunc() {
	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	http.HandleFunc("/newUser", newUser)
	http.HandleFunc("/userStatus", userStatusPage)
	http.HandleFunc("/makeChanges", makeChanges)
	http.ListenAndServe(":8080", nil)
}
