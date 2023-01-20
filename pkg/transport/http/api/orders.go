package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

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
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
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
	id, err := database.NewOrder(result)
	if err != nil || id == 0 {
		fmt.Println(err, "ошибка базы данных не удалось записать новый заказ")
		w.WriteHeader(500)
		return
	}

	var resul types.Order
	resul, err = database.ReadOrder(id)
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
func PutOrder(w http.ResponseWriter, r *http.Request) {
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
	err = database.ChangeOrder(res)
	if err != nil {
		fmt.Println(err, "ошибка базы данных изменения заказа")
		w.WriteHeader(500)
		return
	}
	var result types.Order
	result, err = database.ReadOrder(id)
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
func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	err = database.DelOrder(id)
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
