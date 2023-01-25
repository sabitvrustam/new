package http

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sabitvrustam/new/pkg/transport/http/api"
)

func StartHandler(db *sql.DB) {

	t := NewTemplates()
	r := mux.NewRouter()
	deviceAPI := api.NewDeviceAPI(db)
	r.HandleFunc("/", t.indexPage)
	r.HandleFunc("/order/new", t.newOrderPage)
	r.HandleFunc("/order/status", t.statusOrderPage)
	r.HandleFunc("/order/change", t.makeChangesOrder)
	r.HandleFunc("/parts", t.parts)
	r.HandleFunc("/works", t.works)

	r.HandleFunc("/api/users", api.GetUsers).Methods("GET")                                   //json список клиентов
	r.HandleFunc("/api/users/search/{l_name:[А-Яа-яЁё]+}", api.GetUsersSearch).Methods("GET") //json поиск клиента
	r.HandleFunc("/api/users/{id:[0-9]+}", api.GetUser).Methods("GET")                        //json статус клиента
	r.HandleFunc("/api/users", api.PostUser).Methods("POST")                                  //json новый клиент
	r.HandleFunc("/api/users/{id:[0-9]+}", api.PutUser).Methods("PUT")                        //json изменить данные клиента
	r.HandleFunc("/api/users/{id:[0-9]+}", api.DeleteUser).Methods("DELETE")                  //json удалить клиента
	r.HandleFunc("/api/device", deviceAPI.GetDevices).Methods("GET")                          //json список устройств
	r.HandleFunc("/api/device/search/{sn}", deviceAPI.GetDevicesSearch).Methods("GET")        //json поиск устройства
	r.HandleFunc("/api/device/{id:[0-9]+}", deviceAPI.GetDevice).Methods("GET")               //json статус устройства
	r.HandleFunc("/api/device", deviceAPI.PostDevice).Methods("POST")                         //json новое устройство
	r.HandleFunc("/api/device/{id:[0-9]+}", deviceAPI.PutDevice).Methods("PUT")               //json изменение устройства
	r.HandleFunc("/api/device/{id:[0-9]+}", deviceAPI.DeleteDevice).Methods("DELETE")         //json удаление устройства
	r.HandleFunc("/api/order/{limit:[0-9]+}/{offset:[0-9]+}", api.GetOrders).Methods("get")   //json список заказов
	r.HandleFunc("/api/order", api.PostOrder).Methods("POST")                                 //json новый заказ
	r.HandleFunc("/api/order/{id:[0-9]+}", api.GetOrder).Methods("GET")                       //json статус заказа
	r.HandleFunc("/api/order/{id:[0-9]+}", api.PutOrder).Methods("PUT")                       //json изменить заказ
	r.HandleFunc("/api/order/{id:[0-9]+}", api.DeleteOrder).Methods("DELETE")                 //json удалить заказ
	r.HandleFunc("/api/masters", api.GetMasters).Methods("GET")                               //json список мастеров
	r.HandleFunc("/api/masters", api.PostMasters).Methods("POST")                             //json новый мастер
	r.HandleFunc("/api/masters/{id:[0-9]+}", api.PutMasters).Methods("PUT")                   //json изменить мастера
	r.HandleFunc("/api/masters/{id:[0-9]+}", api.DeleteMasters).Methods("DELETE")             //json удалить мастера
	r.HandleFunc("/api/parts", api.GetParts).Methods("GET")                                   //json список запчастей
	r.HandleFunc("/api/parts", api.PostParts).Methods("POST")                                 //json новая запчасть
	r.HandleFunc("/api/parts/{id:[0-9]+}", api.PutPart).Methods("PUT")                        //json изменить запчасть
	r.HandleFunc("/api/parts/{id:[0-9]+}", api.DeletePart).Methods("DELETE")                  //json удалить запчасть
	r.HandleFunc("/api/works", api.GetWorks).Methods("GET")                                   //json список работ
	r.HandleFunc("/api/works", api.PostWork).Methods("POST")                                  //json новая работа
	r.HandleFunc("/api/works/c", api.PutWork).Methods("PUT")                                  //json изменить работу
	r.HandleFunc("/api/works/{id:[0-9]+}", api.DeleteWork).Methods("DELETE")                  // json удалить работу
	r.HandleFunc("/api/orderparts/{id:[0-9]+}", api.GetOrderParts).Methods("GET")             //json запчасти в заказе
	r.HandleFunc("/api/orderparts", api.PostOrderParts).Methods("POST")                       //json добавить запчасти к заказу
	r.HandleFunc("/api/orderparts/{id:[0-9]+}", api.DeleteOrderParts).Methods("DELETE")       //json удалить запчасть с заказа
	r.HandleFunc("/api/orderworks/{id:[0-9]+}", api.GetOrderWorks).Methods("GET")             //json работы в заказе
	r.HandleFunc("/api/orderworks", api.PostOrderWorks).Methods("POST")                       //json добавить работу к заказу
	r.HandleFunc("/api/orderworks/{id:[0-9]+}", api.DeleteOrderWorks).Methods("DELETE")       //json удалить работу из заказа
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static"))))
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", r)
}
