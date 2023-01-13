package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Templates struct {
	Main              *template.Template
	CteateOrder       *template.Template
	OrderCreateStatus *template.Template
	OrderStatus       *template.Template
	OrderChange       *template.Template
	MakeOrderChange   *template.Template
	OrderPartsChange  *template.Template
	OrderWorkChange   *template.Template
	Parts             *template.Template
	Works             *template.Template
}

func handler() {

	t := NewTemplates()
	r := mux.NewRouter()

	r.HandleFunc("/", t.indexPage)
	r.HandleFunc("/order/new", t.newOrderPage)
	r.HandleFunc("/order/status", t.statusOrderPage)
	r.HandleFunc("/order/change", t.makeChangesOrder)

	r.HandleFunc("/api/order", postApiOrder).Methods("POST")                     //json новый заказ
	r.HandleFunc("/api/order/{id:[0-9]+}", getApiOrder).Methods("GET")           //json статус заказа
	r.HandleFunc("/api/order/{id:[0-9]+}", putApiOrder).Methods("PUT")           //json изменить заказ
	r.HandleFunc("/api/masters", getApiMasters).Methods("GET")                   //json список мастеров
	r.HandleFunc("/api/masters", postApiMasters).Methods("POST")                 //json новый мастер
	r.HandleFunc("/api/masters/{id:[0-9]+}", putApiMasters).Methods("PUT")       //json изменить мастера
	r.HandleFunc("/api/masters/{id:[0-9]+}", deleteApiMasters).Methods("DELETE") //json удалить мастера
	r.HandleFunc("/api/parts", getApiParts).Methods("GET")                       //json список запчастей
	r.HandleFunc("/api/parts", postApiParts).Methods("POST")                     //json новая запчасть
	r.HandleFunc("/api/parts/{id:[0-9]+}", putApiParts).Methods("PUT")           //json изменить запчасть
	r.HandleFunc("/api/works", getApiWorks).Methods("GET")                       //json список работ
	r.HandleFunc("/api/works", postApiWork).Methods("POST")                      //json новая работа

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
	tpl, err = template.ParseFiles("web/html/header.html", "web/html/createPage.html", "web/html/footer.html")
	t.CteateOrder = tpl
	if err != nil {
		fmt.Println(err, "не удалось открыть страничку создания заказа")
	}
	tpl, err = template.ParseFiles("web/html/header.html", "web/html/orderCreateStatus.html", "web/html/footer.html")
	t.OrderCreateStatus = tpl
	if err != nil {
		fmt.Println(err, "не удалось загрузить шаблон страницы состояния созданного заказа")
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

func (t *Templates) indexPage(w http.ResponseWriter, r *http.Request) {
	t.Main.ExecuteTemplate(w, "index", nil)
}
func (t *Templates) newOrderPage(w http.ResponseWriter, r *http.Request) {
	t.CteateOrder.ExecuteTemplate(w, "createPage", nil)
}
func (t *Templates) statusOrderPage(w http.ResponseWriter, r *http.Request) {
	t.OrderStatus.ExecuteTemplate(w, "userStatus", nil)
}
func (t *Templates) makeChangesOrder(w http.ResponseWriter, r *http.Request) {
	t.MakeOrderChange.ExecuteTemplate(w, "makeChangesOrder", nil)
}

func getApiOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	result := readOrder(id)
	m, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err, "")
	}
	w.WriteHeader(200)
	w.Write(m)
}

func postApiOrder(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, _ := io.ReadAll(r.Body)
	var res Order
	err := json.Unmarshal(b, &res)
	if err != nil {
		fmt.Println(err)
	}
	idOrder, err := newOrder(res)
	if err != nil {
		fmt.Println(err)
	}
	res.IdOrder = idOrder
	m, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err, "")
	}
	w.WriteHeader(200)
	w.Write(m)
}
func putApiOrder(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, _ := io.ReadAll(r.Body)
	var res Order
	err := json.Unmarshal(b, &res)
	if err != nil {
		fmt.Println(err)
	}
	idOrder, err := newOrder(res)
	if err != nil {
		fmt.Println(err)
	}
	res.IdOrder = idOrder
	m, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err, "")
	}
	w.WriteHeader(200)
	w.Write(m)
}

func getApiMasters(w http.ResponseWriter, r *http.Request) {
	result := dbreadMasters()
	m, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err, "")
		w.WriteHeader(404)
	}
	w.WriteHeader(200)
	w.Write(m)
}

func postApiMasters(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, _ := io.ReadAll(r.Body)
	var res Masters
	err := json.Unmarshal(b, &res)
	if err != nil {
		fmt.Println(err)
	}
	id, err := newMaster(res)
	if err != nil {
		fmt.Println(err)
	}
	res.Id = id
	m, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err, "")
	}
	w.WriteHeader(200)
	w.Write(m)
}
func putApiMasters(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, _ := io.ReadAll(r.Body)
	var res Masters
	err := json.Unmarshal(b, &res)
	if err != nil {
		fmt.Println(err)
	}
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	res.Id = id
	err = changMaster(res)
	if err != nil {
		fmt.Println(err)
	}

	m, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err, "")
	}
	w.WriteHeader(200)
	w.Write(m)
}
func getApiParts(w http.ResponseWriter, r *http.Request) {
	result := dbreadParts()
	m, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err, "")
		w.WriteHeader(404)
	}
	w.WriteHeader(200)
	w.Write(m)
}

func postApiParts(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, _ := io.ReadAll(r.Body)
	var res Part
	err := json.Unmarshal(b, &res)
	if err != nil {
		fmt.Println(err)
	}
	id, err := dbWriteParts(res)
	if err != nil {
		fmt.Println(err)
	}
	res.Id = id
	m, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err, "")
	}
	w.WriteHeader(200)
	w.Write(m)
}
func putApiParts(w http.ResponseWriter, r *http.Request) {

}

func getApiWorks(w http.ResponseWriter, r *http.Request) {
	result := dbreadWorks()
	m, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err, "")
		w.WriteHeader(404)
	}
	w.WriteHeader(200)
	w.Write(m)
}

// func saveWorksOrder(w http.ResponseWriter, r *http.Request) {

// 	err := r.ParseForm()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	var workId Order
// 	id, err := strconv.ParseInt(r.FormValue("id"), 10, 64)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	workId.IdOrder = id
// 	for n, i := range r.Form {
// 		if n == "id" {
// 		} else {
// 			for _, m := range i {
// 				workId.Work.Id = m
// 				dbWriteWorksOrder(workId)
// 			}
// 		}
// 	}
// 	url := fmt.Sprintf("/makeChanges/%d", id)

// 	http.Redirect(w, r, url, http.StatusSeeOther)

// }
func postApiWork(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, _ := io.ReadAll(r.Body)
	var res Work
	err := json.Unmarshal(b, &res)
	if err != nil {
		fmt.Println(err)
	}
	id, err := dbWriteWork(res)
	if err != nil {
		fmt.Println(err)
	}
	res.Id = id
	m, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err, "")
	}
	w.WriteHeader(200)
	w.Write(m)
}

// func (t *Templates) makeChangesParts(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	id, err := strconv.ParseInt(vars["id"], 10, 64)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(id)
// 	result := dbreadParts()
// 	result = id
// 	t.OrderPartsChange.ExecuteTemplate(w, "makechangesparts", result)
// }

//	func (t *Templates) makeChangesWork(w http.ResponseWriter, r *http.Request) {
//		vars := mux.Vars(r)
//		id, err := strconv.ParseInt(vars["id"], 10, 64)
//		if err != nil {
//			fmt.Println(err)
//		}
//		fmt.Println(id)
//		result := dbreadWorks()
//		result.IdOrder = id
//		t.OrderWorkChange.ExecuteTemplate(w, "makeChangesWork", result)
//	}
// func (t *Templates) parts(w http.ResponseWriter, r *http.Request) {
// 	result := dbreadParts()
// 	t.Parts.ExecuteTemplate(w, "parts", result)
// }
// func (t *Templates) works(w http.ResponseWriter, r *http.Request) {
// 	result := dbreadWorks()
// 	t.Works.ExecuteTemplate(w, "works", result)
// }

// func savePartsOrder(w http.ResponseWriter, r *http.Request) {

// 	err := r.ParseForm()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	var partId Order
// 	id, err := strconv.ParseInt(r.FormValue("id"), 10, 64)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	partId.IdOrder = id
// 	for n, i := range r.Form {
// 		if n == "id" {
// 		} else {
// 			for _, m := range i {
// 				partId.Part.Id = m
// 				dbWritePartsOrder(partId)
// 			}
// 		}
// 	}
// 	url := fmt.Sprintf("/makeChanges/%d", id)

// 	http.Redirect(w, r, url, http.StatusSeeOther)

// }
// func makeChangesDleleteParts(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	var part Order

// 	idOrder, err := strconv.ParseInt(vars["idOrder"], 10, 64)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	part.IdOrder = idOrder
// 	part.Part.Id = vars["idPart"]
// 	fmt.Println(part.IdOrder, part.Part.Id)

// 	dbDeletePartsOrder(part)
// 	url := fmt.Sprintf("/makeChanges/%d", part.IdOrder)

// 	http.Redirect(w, r, url, http.StatusSeeOther)

// }

// func makeChangesDleleteWorks(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	var work Order
// 	idOrder, err := strconv.ParseInt(vars["idOrder"], 10, 64)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	work.IdOrder = idOrder
// 	work.Work.Id = vars["idWork"]
// 	fmt.Println(work.IdOrder, work.Work.Id)

// 	dbDeleteWorksOrder(work)
// 	url := fmt.Sprintf("/makeChanges/%d", work.IdOrder)

// 	http.Redirect(w, r, url, http.StatusSeeOther)

// }

// func newWork(w http.ResponseWriter, r *http.Request) {
// 	workName := r.FormValue("workName")
// 	workPrice := r.FormValue("workPrice")
// 	newParts := Work{
// 		WorkName:  workName,
// 		WorkPrice: workPrice,
// 	}
// 	dbWriteWork(newParts)
// 	http.Redirect(w, r, "/works", http.StatusSeeOther)
// }
