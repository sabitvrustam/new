package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sabitvrustam/new/pkg/types"
)

func (a *OrderAPI) GetOrderParts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		a.log.Error(err, "неправильный id Order parts")
		w.WriteHeader(404)
		return
	}
	result, err := a.order.ReadOrderParts(nil, id)
	if err != nil {
		a.log.Error(err, "ошибка базы данных считывание мастеров")
		w.WriteHeader(500)
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

func (a *OrderAPI) PostOrderParts(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		a.log.Error(err, "ошибка приема данных нового мастера от пользователя")
		w.WriteHeader(404)
		return
	}
	var result types.OrderParts
	err = json.Unmarshal(b, &result)
	if err != nil {
		a.log.Error(err, "ошибка unmarshal новый масер")
		w.WriteHeader(500)
		return
	}
	id, err := a.order.NewOrderParts(result)
	if err != nil || id == 0 {
		a.log.Error(err, "ошибка базы данных сохранения нового пользователя")
		w.WriteHeader(500)
		return
	}
	resul, err := a.order.ReadOrderParts(&id, 0)
	if err != nil {
		a.log.Error(err, "ошибка базы данных считывание мастеров")
		w.WriteHeader(500)
	}

	m, err := json.Marshal(resul)
	if err != nil {
		a.log.Error(err, "не удалось преобразовать данные нового мастера в json")
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
	w.Write(m)
}
func (a *OrderAPI) DeleteOrderParts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		a.log.Error(err, "неправильный id")
		w.WriteHeader(404)
		return
	}
	err = a.order.DelOrderParts(id)
	if err != nil {
		a.log.Error(err, "ошибка базы данных удаление мастера")
		w.WriteHeader(500)
		return
	}
	m, err := json.Marshal(id)
	if err != nil {
		a.log.Error(err, "ошибка преобразования данных в json")
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
	w.Write(m)
}
