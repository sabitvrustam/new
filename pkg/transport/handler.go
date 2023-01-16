package main

import (
	"fmt"
	"net/http"

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
