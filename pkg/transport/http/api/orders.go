package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sabitvrustam/new/pkg/database"
	"github.com/sabitvrustam/new/pkg/types"
)

func GetOrders(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	limit := vars["limit"]
	offset := vars["offset"]
	result, err := database.ReadOrders(limit, offset)
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

func GetOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	result, err := database.ReadOrder(id)
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

func PostOrder(w http.ResponseWriter, r *http.Request) {
	var result types.Order
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
	id, err := database.NewOrder(result)
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
func PutOrder(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err, "не удалось принять данные измененного ордера от пользователя")
		w.WriteHeader(408)
		return
	}
	var res types.Order
	err = json.Unmarshal(b, &res)
	if err != nil {
		fmt.Println(err, "ошибка unmarshal изменения заказа")
		w.WriteHeader(500)
		return
	}
	id, err := database.NewOrder(res)
	if err != nil {
		fmt.Println(err, "ошибка базы данных изменения заказа")
		w.WriteHeader(500)
		return
	}
	if id == 0 {
		fmt.Println("введен некоректный номер ордера для изменения")
		w.WriteHeader(404)
		return
	}
	res.IdOrder = id
	m, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err, "не удалось преобразовать данные в json")
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
	w.Write(m)
}
