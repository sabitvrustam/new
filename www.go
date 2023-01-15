package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func handler() {

	t := NewTemplates()
	r := mux.NewRouter()

	r.HandleFunc("/", t.indexPage)
	r.HandleFunc("/order/new", t.newOrderPage)
	r.HandleFunc("/order/status", t.statusOrderPage)
	r.HandleFunc("/order/change", t.makeChangesOrder)
	r.HandleFunc("/parts", t.parts)
	r.HandleFunc("/works", t.works)

	r.HandleFunc("/api/order", postApiOrder).Methods("POST")                     //json новый заказ
	r.HandleFunc("/api/order/{id:[0-9]+}", getApiOrder).Methods("GET")           //json статус заказа
	r.HandleFunc("/api/order/{id:[0-9]+}", putApiOrder).Methods("PUT")           //json изменить заказ
	r.HandleFunc("/api/masters", getApiMasters).Methods("GET")                   //json список мастеров
	r.HandleFunc("/api/masters", postApiMasters).Methods("POST")                 //json новый мастер
	r.HandleFunc("/api/masters/{id:[0-9]+}", putApiMasters).Methods("PUT")       //json изменить мастера
	r.HandleFunc("/api/masters/{id:[0-9]+}", deleteApiMasters).Methods("DELETE") //json удалить мастера
	r.HandleFunc("/api/parts", getApiParts).Methods("GET")                       //json список запчастей
	r.HandleFunc("/api/parts", postApiParts).Methods("POST")                     //json новая запчасть
	r.HandleFunc("/api/parts/{id:[0-9]+}", putApiPart).Methods("PUT")            //json изменить запчасть
	r.HandleFunc("/api/parts/{id:[0-9]+}", deleteApiPart).Methods("DELETE")      //json удалить запчасть
	r.HandleFunc("/api/works", getApiWorks).Methods("GET")                       //json список работ
	r.HandleFunc("/api/works", postApiWork).Methods("POST")                      //json новая работа
	r.HandleFunc("/api/works/{id:[0-9]+}", putApiWork).Methods("PUT")            //json изменить работу
	r.HandleFunc("/api/works/{id:[0-9]+}", deleteApiWork).Methods("DELETE")      // json удалить работу
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static"))))
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", r)
}

func getApiOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	result, err := readOrder(id)
	if err != nil {
		fmt.Println(err, "не удалось считать данные ордера из базы данных по ид")
		w.WriteHeader(500)
		return
	}
	if result.IdOrder == 0 {
		fmt.Println("В базе данных не существует такой записи заказа")
		w.WriteHeader(404)
		return
	}
	marshalResult, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err, "не удалось преобразовать данные Ордера в json")
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
	w.Write(marshalResult)

}

func postApiOrder(w http.ResponseWriter, r *http.Request) {
	var result Order
	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err, "не удалось принять данные нового ордера от пользователя")
		w.WriteHeader(408)
		return
	}
	err = json.Unmarshal(b, &result)
	if err != nil {
		fmt.Println(err, "ошибка unmarshal нового заказа")
		w.WriteHeader(500)
		return
	}
	id, err := newOrder(result)
	if err != nil || id == 0 {
		fmt.Println(err, "ошибка базы данных не удалось записать новый заказ")
		w.WriteHeader(500)
		return
	}
	result.IdOrder = id
	m, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err, "не удалось преобразовать данные в json")
		w.WriteHeader(500)
		return
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
	result := readMasters()
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
func deleteApiMasters(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	err = deleteMaster(id)
	if err != nil {
		fmt.Println(err)
	}

	m, err := json.Marshal(id)
	if err != nil {
		fmt.Println(err, "")
	}
	w.WriteHeader(200)
	w.Write(m)
}

func getApiParts(w http.ResponseWriter, r *http.Request) {
	result := readParts()
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
	id, err := newPart(res)
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
func putApiPart(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, _ := io.ReadAll(r.Body)
	var res Part
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
	err = changePart(res)
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
func deleteApiPart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	err = deletePart(id)
	if err != nil {
		fmt.Println(err)
	}

	m, err := json.Marshal(id)
	if err != nil {
		fmt.Println(err, "")
	}
	w.WriteHeader(200)
	w.Write(m)
}

func getApiWorks(w http.ResponseWriter, r *http.Request) {
	result := readWoeks()
	m, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err, "")
		w.WriteHeader(404)
	}
	w.WriteHeader(200)
	w.Write(m)
}

func postApiWork(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, _ := io.ReadAll(r.Body)
	var res Work
	err := json.Unmarshal(b, &res)
	if err != nil {
		fmt.Println(err)
	}
	id, err := writeWork(res)
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
func putApiWork(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, _ := io.ReadAll(r.Body)
	var res Work
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
	err = changeWork(res)
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

func deleteApiWork(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	err = deleteWork(id)
	if err != nil {
		fmt.Println(err)
	}

	m, err := json.Marshal(id)
	if err != nil {
		fmt.Println(err, "")
	}
	w.WriteHeader(200)
	w.Write(m)
}
