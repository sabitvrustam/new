package http

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sabitvrustam/new/pkg/transport/http/api"
	"github.com/sirupsen/logrus"
)

func StartHandler(db *sql.DB, log *logrus.Logger) {

	t := NewTemplates(log)
	r := mux.NewRouter()
	deviceAPI := api.NewDeviceAPI(db, log)
	masterAPI := api.NewMasterAPI(db, log)
	orderAPI := api.NewOrderAPI(db, log)
	userAPI := api.NewUsersAPI(db, log)
	partAPI := api.NewPartAPI(db, log)
	workAPI := api.NewWork(db, log)
	statusAPI := api.NewStatusAPI(db, log)
	countAPI := api.NewCountAPI(db, log)

	r.HandleFunc("/", t.indexPage)
	r.HandleFunc("/order/new", t.newOrderPage)
	r.HandleFunc("/order/status", t.statusOrderPage)
	r.HandleFunc("/order/change", t.makeChangesOrder)
	r.HandleFunc("/order", t.ordersPage)
	r.HandleFunc("/parts", t.parts)
	r.HandleFunc("/works", t.works)

	r.HandleFunc("/api/count", countAPI.GetCount).Methods("get")                                  //колличество заказов
	r.HandleFunc("/api/users", userAPI.GetUsers).Methods("GET")                                   //json список клиентов
	r.HandleFunc("/api/users/search/{l_name:[А-Яа-яЁё]+}", userAPI.GetUsersSearch).Methods("GET") //json поиск клиента
	r.HandleFunc("/api/users/{id:[0-9]+}", userAPI.GetUser).Methods("GET")                        //json статус клиента
	r.HandleFunc("/api/users", userAPI.PostUser).Methods("POST")                                  //json новый клиент
	r.HandleFunc("/api/users/{id:[0-9]+}", userAPI.PutUser).Methods("PUT")                        //json изменить данные клиента
	r.HandleFunc("/api/users/{id:[0-9]+}", userAPI.DeleteUser).Methods("DELETE")                  //json удалить клиента
	r.HandleFunc("/api/device", deviceAPI.GetDevices).Methods("GET")                              //json список устройств
	r.HandleFunc("/api/device/search/{sn}", deviceAPI.GetDevicesSearch).Methods("GET")            //json поиск устройства
	r.HandleFunc("/api/device/{id:[0-9]+}", deviceAPI.GetDevice).Methods("GET")                   //json статус устройства
	r.HandleFunc("/api/device", deviceAPI.PostDevice).Methods("POST")                             //json новое устройство
	r.HandleFunc("/api/device/{id:[0-9]+}", deviceAPI.PutDevice).Methods("PUT")                   //json изменение устройства
	r.HandleFunc("/api/device/{id:[0-9]+}", deviceAPI.DeleteDevice).Methods("DELETE")             //json удаление устройства
	r.HandleFunc("/api/order/{limit:[0-9]+}/{offset:[0-9]+}", orderAPI.GetOrders).Methods("get")  //json список заказов
	r.HandleFunc("/api/order", orderAPI.PostOrder).Methods("POST")                                //json новый заказ
	r.HandleFunc("/api/order/{id:[0-9]+}", orderAPI.GetOrder).Methods("GET")                      //json статус заказа
	r.HandleFunc("/api/order/{id:[0-9]+}", orderAPI.PutOrder).Methods("PUT")                      //json изменить заказ
	r.HandleFunc("/api/order/{id:[0-9]+}", orderAPI.DeleteOrder).Methods("DELETE")                //json удалить заказ
	r.HandleFunc("/api/masters", masterAPI.GetMasters).Methods("GET")                             //json список мастеров
	r.HandleFunc("/api/masters", masterAPI.PostMasters).Methods("POST")                           //json новый мастер
	r.HandleFunc("/api/masters/{id:[0-9]+}", masterAPI.PutMasters).Methods("PUT")                 //json изменить мастера
	r.HandleFunc("/api/masters/{id:[0-9]+}", masterAPI.DeleteMasters).Methods("DELETE")           //json удалить мастера
	r.HandleFunc("/api/parts", partAPI.GetParts).Methods("GET")                                   //json список запчастей
	r.HandleFunc("/api/parts", partAPI.PostParts).Methods("POST")                                 //json новая запчасть
	r.HandleFunc("/api/parts/{id:[0-9]+}", partAPI.PutPart).Methods("PUT")                        //json изменить запчасть
	r.HandleFunc("/api/parts/{id:[0-9]+}", partAPI.DeletePart).Methods("DELETE")                  //json удалить запчасть
	r.HandleFunc("/api/works", workAPI.GetWorks).Methods("GET")                                   //json список работ
	r.HandleFunc("/api/works", workAPI.PostWork).Methods("POST")                                  //json новая работа
	r.HandleFunc("/api/works/{id:[0-9]+}", workAPI.PutWork).Methods("PUT")                        //json изменить работу
	r.HandleFunc("/api/works/{id:[0-9]+}", workAPI.DeleteWork).Methods("DELETE")                  // json удалить работу
	r.HandleFunc("/api/orderparts/{id:[0-9]+}", orderAPI.GetOrderParts).Methods("GET")            //json запчасти в заказе
	r.HandleFunc("/api/orderparts", orderAPI.PostOrderParts).Methods("POST")                      //json добавить запчасти к заказу
	r.HandleFunc("/api/orderparts/{id:[0-9]+}", orderAPI.DeleteOrderParts).Methods("DELETE")      //json удалить запчасть с заказа
	r.HandleFunc("/api/orderworks/{id:[0-9]+}", orderAPI.GetOrderWorks).Methods("GET")            //json работы в заказе
	r.HandleFunc("/api/orderworks", orderAPI.PostOrderWorks).Methods("POST")                      //json добавить работу к заказу
	r.HandleFunc("/api/orderworks/{id:[0-9]+}", orderAPI.DeleteOrderWorks).Methods("DELETE")      //json удалить работу из заказа
	r.HandleFunc("/api/status", statusAPI.GetStatus).Methods("get")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static"))))
	log.Info("Локальный сервер запущен порт 8080")
	http.ListenAndServe(":8080", r)
}
