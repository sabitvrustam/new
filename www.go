package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
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
		http.Redirect(w, r, "/create", http.StatusSeeOther)
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
	vars := mux.Vars(r)
	id := vars["id"]
	result = dbRead(id)

	id = r.FormValue("id")
	if id != "" {
		result = dbRead(id)

	}
	t.ExecuteTemplate(w, "makeChanges", result)
}
func makeChangesOrder(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("header.html", "makeChangesOrder.html", "footer.html")
	if err != nil {
		fmt.Println(w, err.Error(), "не удалось открыть страничку создания заказа")

	}
	id := r.FormValue("id")
	if id != "" {
		url := fmt.Sprintf("/makeChanges/%s", id)
		http.Redirect(w, r, url, http.StatusSeeOther)

	}
	t.ExecuteTemplate(w, "makeChangesOrder", nil)
}
func makeChangesParts(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("header.html", "makechangesparts.html", "footer.html")
	if err != nil {
		fmt.Println(w, err.Error(), "не удалось открыть страничку создания заказа")
	}
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println(id)
	result := dbreadParts()
	result.IdOrder = id

	t.ExecuteTemplate(w, "makechangesparts", result)

}
func savePartsOrder(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	var saveParts []string
	for _, i := range r.Form {
		saveParts = append(saveParts, i...)

	}
	url := fmt.Sprintf("/makeChanges/%s", id)
	fmt.Println(saveParts, id)
	http.Redirect(w, r, url, http.StatusSeeOther)

}
func parts(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("header.html", "parts.html", "footer.html")
	if err != nil {
		fmt.Println(w, err.Error())
	}
	result := dbreadParts()
	t.ExecuteTemplate(w, "parts", result)
}
func newParts(w http.ResponseWriter, r *http.Request) {
	partsName := r.FormValue("partsName")
	partsPrice := r.FormValue("partsPrice")
	newParts := Part{
		PartsName:  partsName,
		PartsPrice: partsPrice,
	}
	dbWriteParts(newParts)
	http.Redirect(w, r, "/parts", http.StatusSeeOther)
}
func handleFunc() {
	r := mux.NewRouter()
	r.HandleFunc("/", index)
	r.HandleFunc("/create", create)
	r.HandleFunc("/newUser", newUser)
	r.HandleFunc("/userStatus", userStatusPage)
	r.HandleFunc("/makeChangesOrder", makeChangesOrder)
	r.HandleFunc("/makeChanges/{id:[0-9]+}", makeChanges)
	r.HandleFunc("/makeChangesParts/{id:[0-9]+}", makeChangesParts)
	r.HandleFunc("/parts", parts)
	r.HandleFunc("/newParts", newParts)
	r.HandleFunc("/makeChangesParts/savePartsOrder", savePartsOrder)
	http.Handle("/", r)
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
