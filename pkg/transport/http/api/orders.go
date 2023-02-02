package api

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sabitvrustam/new/pkg/database/orders"
	"github.com/sabitvrustam/new/pkg/types"
	"github.com/sirupsen/logrus"
)

type OrderAPI struct {
	db    *sql.DB
	log   *logrus.Logger
	order *orders.Order
}

func NewOrderAPI(db *sql.DB, log *logrus.Logger) *OrderAPI {
	return &OrderAPI{
		db:    db,
		log:   log,
		order: orders.NewOrder(db, log)}
}

func (a *OrderAPI) GetOrders(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	limit, err := strconv.ParseUint(vars["limit"], 10, 64)
	if err != nil {
		a.log.Error(err)
	}
	offset, err := strconv.ParseUint(vars["offset"], 10, 64)
	if err != nil {
		a.log.Error(err)
	}
	result, err := a.order.ReadOrders(limit, offset)
	if err != nil {
		a.log.Error(err, "не удалось считать данные ордеров колличество на странице")
		w.WriteHeader(500)
		return
	}
	marshalResult, err := json.Marshal(result)
	if err != nil {
		a.log.Error(err, "не удалось преобразовать данные Ордера в json")
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
	w.Write(marshalResult)

}

func (a *OrderAPI) GetOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	result, err := a.order.GetOrderByID(id)
	if err != nil {
		a.log.Error(err, "не удалось считать данные ордера из базы данных по ид")
		w.WriteHeader(500)
		return
	}

	marshalResult, err := json.Marshal(result)
	if err != nil {
		a.log.Error(err, "не удалось преобразовать данные Ордера в json")
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
	w.Write(marshalResult)

}

func (a *OrderAPI) PostOrder(w http.ResponseWriter, r *http.Request) {
	var result types.Id
	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		a.log.Error(err, "не удалось принять данные нового ордера от пользователя")
		w.WriteHeader(408)
		return
	}
	err = json.Unmarshal(b, &result)
	if err != nil {
		a.log.Error(err, "ошибка unmarshal нового заказа")
		w.WriteHeader(500)
		return
	}
	id, err := a.order.NewOrder1(result)
	if err != nil || id == 0 {
		a.log.Error(err, "ошибка базы данных не удалось записать новый заказ")
		w.WriteHeader(500)
		return
	}

	var resul *types.Order
	resul, err = a.order.GetOrderByID(id)
	if err != nil {
		a.log.Error(err, "Не удалось считать данные ордера из базы данных по ид")
		w.WriteHeader(500)
		return
	}

	m, err := json.Marshal(resul)
	if err != nil {
		a.log.Error(err, "не удалось преобразовать данные в json")
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
	w.Write(m)
}
func (a *OrderAPI) PutOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		a.log.Error(err, "не удалось принять данные измененного ордера от пользователя")
		w.WriteHeader(408)
		return
	}
	var res types.Id
	err = json.Unmarshal(b, &res)
	if err != nil {
		a.log.Error(err, "ошибка unmarshal изменения заказа")
		w.WriteHeader(500)
		return
	}
	res.IdOrder = id
	err = a.order.ChangeOrder(res)
	if err != nil {
		a.log.Error(err, "ошибка базы данных изменения заказа")
		w.WriteHeader(500)
		return
	}
	var result *types.Order
	result, err = a.order.GetOrderByID(id)
	if err != nil {
		a.log.Error(err, "не удалось считать данные ордера из базы данных по ид")
		w.WriteHeader(500)
		return
	}
	m, err := json.Marshal(result)
	if err != nil {
		a.log.Error(err, "не удалось преобразовать данные в json")
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
	w.Write(m)
}
func (a *OrderAPI) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		a.log.Error(err)
	}
	err = a.order.DelOrder(id)
	if err != nil {
		a.log.Error(err)
	}

	m, err := json.Marshal(id)
	if err != nil {
		a.log.Error(err, "")
	}
	w.WriteHeader(200)
	w.Write(m)
}
