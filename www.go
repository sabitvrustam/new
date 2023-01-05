package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type Templates struct {
	Main             *template.Template
	Cteate           *template.Template
	OrderStatus      *template.Template
	OrderChange      *template.Template
	MakeOrderChange  *template.Template
	OrderPartsChange *template.Template
	OrderWorkChange  *template.Template
	Parts            *template.Template
	Works            *template.Template
}

func handler() {

	t := NewTemplates()
	r := mux.NewRouter()
	r.HandleFunc("/test", test)
	r.HandleFunc("/", t.index)
	r.HandleFunc("/create", t.create)
	r.HandleFunc("/userStatus", t.userStatusPage)
	r.HandleFunc("/makeChangesOrder", t.makeChangesOrder)
	r.HandleFunc("/makeChanges/{id:[0-9]+}", t.makeChanges)
	r.HandleFunc("/makeChangesParts/{id:[0-9]+}", t.makeChangesParts)
	r.HandleFunc("/makeChangesDeleteParts/{idOrder:[0-9]+}/{idPart:[0-9]+}", makeChangesDleleteParts)
	r.HandleFunc("/makeChangesWork/{id:[0-9]+}", t.makeChangesWork)
	r.HandleFunc("/makeChangesDeleteWorks/{idOrder:[0-9]+}/{idWork:[0-9]+}", makeChangesDleleteWorks)
	r.HandleFunc("/parts", t.parts)
	r.HandleFunc("/works", t.works)
	r.HandleFunc("/newParts", newParts)
	r.HandleFunc("/newWork", newWork)
	r.HandleFunc("/makeChangesParts/savePartsOrder", savePartsOrder)
	r.HandleFunc("/makeChangesWork/saveWorksOrder", saveWorksOrder)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static"))))
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", r)
}

func NewTemplates() Templates {
	var t Templates
	tpl, err := template.ParseFiles("web/html/header.html", "web/html/index.html", "web/html/footer.html")
	t.Main = tpl
	if err != nil {
		fmt.Println(err, "не удалось открыть главную страничку")
	}
	tpl, err = template.ParseFiles("web/html/header.html", "web/html/create.html", "web/html/footer.html")
	t.Cteate = tpl
	if err != nil {
		fmt.Println(err, "не удалось открыть страничку создания заказа")
	}
	tpl, err = template.ParseFiles("web/html/header.html", "web/html/userStatus.html", "web/html/footer.html")
	t.OrderStatus = tpl
	if err != nil {
		fmt.Println(err, "Не удалось открыть страницу состояния заказа")
	}
	tpl, err = template.ParseFiles("web/html/header.html", "web/html/makeChanges.html", "web/html/footer.html")
	t.OrderChange = tpl
	if err != nil {
		fmt.Println(err, "не удалось открыть страничку изменения заказа")
	}
	tpl, err = template.ParseFiles("web/html/header.html", "web/html/makeChangesOrder.html", "web/html/footer.html")
	t.MakeOrderChange = tpl
	if err != nil {
		fmt.Println(err, "не удалось открыть страничку изменения заказа")
	}
	tpl, err = template.ParseFiles("web/html/header.html", "web/html/makechangesparts.html", "web/html/footer.html")
	t.OrderPartsChange = tpl
	if err != nil {
		fmt.Println(err, "не удалось открыть страничку изменения запчастей в заказе")
	}
	tpl, err = template.ParseFiles("web/html/header.html", "web/html/makeChangesWork.html", "web/html/footer.html")
	t.OrderWorkChange = tpl
	if err != nil {
		fmt.Println(err, "не удалось открыть страничку изменения работ в заказе")
	}
	tpl, err = template.ParseFiles("web/html/header.html", "web/html/parts.html", "web/html/footer.html")
	t.Parts = tpl
	if err != nil {
		fmt.Println(err, "не удалось открыть страничку с запчастями")
	}
	tpl, err = template.ParseFiles("web/html/header.html", "web/html/works.html", "web/html/footer.html")
	t.Works = tpl
	if err != nil {
		fmt.Println(err, "не удалось открыть страничку с работами")
	}

	return t
}

func (t *Templates) index(w http.ResponseWriter, r *http.Request) {
	t.Main.ExecuteTemplate(w, "index", nil)
}

func (t *Templates) create(w http.ResponseWriter, r *http.Request) {
	result := dbreadMasters()
	t.Cteate.ExecuteTemplate(w, "create", result)
}

func (t *Templates) userStatusPage(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	var result Order
	if id != "" {
		result = dbRead(id)
	}
	t.OrderStatus.ExecuteTemplate(w, "userStatus", result)
}

func (t *Templates) makeChanges(w http.ResponseWriter, r *http.Request) {
	var result Order
	vars := mux.Vars(r)
	id := vars["id"]
	result = dbRead(id)

	id = r.FormValue("id")
	if id != "" {
		result = dbRead(id)
	}
	t.OrderChange.ExecuteTemplate(w, "makeChanges", result)
}

func (t *Templates) makeChangesOrder(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id != "" {
		url := fmt.Sprintf("/makeChanges/%s", id)
		http.Redirect(w, r, url, http.StatusSeeOther)
	}
	t.MakeOrderChange.ExecuteTemplate(w, "makeChangesOrder", nil)
}

func (t *Templates) makeChangesParts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println(id)
	result := dbreadParts()
	result.IdOrder = id
	t.OrderPartsChange.ExecuteTemplate(w, "makechangesparts", result)
}

func (t *Templates) makeChangesWork(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println(id)
	result := dbreadWorks()
	result.IdOrder = id
	t.OrderWorkChange.ExecuteTemplate(w, "makeChangesWork", result)
}
func (t *Templates) parts(w http.ResponseWriter, r *http.Request) {
	result := dbreadParts()
	t.Parts.ExecuteTemplate(w, "parts", result)
}
func (t *Templates) works(w http.ResponseWriter, r *http.Request) {
	result := dbreadWorks()
	t.Works.ExecuteTemplate(w, "works", result)
}
func test(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, _ := io.ReadAll(r.Body)

	var res Order

	err := json.Unmarshal(b, &res)
	if err != nil {
		fmt.Println(err)
	}
	err = dbWrite(res)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
	response := `{"name": "John", "age": 30}`
	w.WriteHeader(500)
	w.Write([]byte(response))

}

func savePartsOrder(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}
	var id string
	var partId Order
	id = r.FormValue("id")
	partId.IdOrder = id
	for n, i := range r.Form {
		if n == "id" {
		} else {
			for _, m := range i {
				partId.Part.Id = m
				dbWritePartsOrder(partId)
			}
		}
	}
	url := fmt.Sprintf("/makeChanges/%s", id)

	http.Redirect(w, r, url, http.StatusSeeOther)

}
func makeChangesDleleteParts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var part Order
	part.IdOrder = vars["idOrder"]
	part.Part.Id = vars["idPart"]
	fmt.Println(part.IdOrder, part.Part.Id)

	dbDeletePartsOrder(part)
	url := fmt.Sprintf("/makeChanges/%s", part.IdOrder)

	http.Redirect(w, r, url, http.StatusSeeOther)

}

func saveWorksOrder(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}
	var id string
	var workId Order
	id = r.FormValue("id")
	workId.IdOrder = id
	for n, i := range r.Form {
		if n == "id" {
		} else {
			for _, m := range i {
				workId.Work.Id = m
				dbWriteWorksOrder(workId)
			}
		}
	}
	url := fmt.Sprintf("/makeChanges/%s", id)

	http.Redirect(w, r, url, http.StatusSeeOther)

}
func makeChangesDleleteWorks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var work Order
	work.IdOrder = vars["idOrder"]
	work.Work.Id = vars["idWork"]
	fmt.Println(work.IdOrder, work.Work.Id)

	dbDeleteWorksOrder(work)
	url := fmt.Sprintf("/makeChanges/%s", work.IdOrder)

	http.Redirect(w, r, url, http.StatusSeeOther)

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

func newWork(w http.ResponseWriter, r *http.Request) {
	workName := r.FormValue("workName")
	workPrice := r.FormValue("workPrice")
	newParts := Work{
		WorkName:  workName,
		WorkPrice: workPrice,
	}
	dbWriteWork(newParts)
	http.Redirect(w, r, "/works", http.StatusSeeOther)
}
