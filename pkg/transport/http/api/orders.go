package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sabitvrustam/new/pkg/database/orders"
	"github.com/sabitvrustam/new/pkg/types"
)

type OrderAPI struct {
	db    *sql.DB
	order *orders.Order
}

func NewOrderAPI(db *sql.DB) *OrderAPI {
	return &OrderAPI{db: db, order: orders.NewOrder(db)}
}

func (a *OrderAPI) GetOrders(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	limit := vars["limit"]
	offset := vars["offset"]
	result, err := a.order.ReadOrders(limit, offset)
	if err != nil {
		fmt.Println(err, "не удалось считать данные ордеров колличество на странице"+limit+"offset"+offset)
		w.WriteHeader(500)
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

func (a *OrderAPI) GetOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	result, err := a.order.ReadOrder(id)
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

func (a *OrderAPI) PostOrder(w http.ResponseWriter, r *http.Request) {
	var result types.Id
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
	id, err := a.order.NewOrder1(result)
	if err != nil || id == 0 {
		fmt.Println(err, "ошибка базы данных не удалось записать новый заказ")
		w.WriteHeader(500)
		return
	}

	var resul types.Order
	resul, err = a.order.ReadOrder(id)
	if err != nil {
		fmt.Println(err, "не удалось считать данные ордера из базы данных по ид")
		w.WriteHeader(500)
		return
	}

	m, err := json.Marshal(resul)
	if err != nil {
		fmt.Println(err, "не удалось преобразовать данные в json")
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
		fmt.Println(err, "не удалось принять данные измененного ордера от пользователя")
		w.WriteHeader(408)
		return
	}
	var res types.Id
	err = json.Unmarshal(b, &res)
	if err != nil {
		fmt.Println(err, "ошибка unmarshal изменения заказа")
		w.WriteHeader(500)
		return
	}
	res.IdOrder = id
	err = a.order.ChangeOrder(res)
	if err != nil {
		fmt.Println(err, "ошибка базы данных изменения заказа")
		w.WriteHeader(500)
		return
	}
	var result types.Order
	result, err = a.order.ReadOrder(id)
	if err != nil {
		fmt.Println(err, "не удалось считать данные ордера из базы данных по ид")
		w.WriteHeader(500)
		return
	}
	m, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err, "не удалось преобразовать данные в json")
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
		fmt.Println(err)
	}
	err = a.order.DelOrder(id)
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
