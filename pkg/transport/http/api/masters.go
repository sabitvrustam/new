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

func GetMasters(w http.ResponseWriter, r *http.Request) {
	result, err := database.ReadMasters()
	if err != nil {
		fmt.Println(err, "ошибка базы данных считывание мастеров")
		w.WriteHeader(500)
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

func PostMasters(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err, "ошибка приема данных нового мастера от пользователя")
		w.WriteHeader(404)
		return
	}
	var result types.Masters
	err = json.Unmarshal(b, &result)
	if err != nil {
		fmt.Println(err, "ошибка unmarshal новый масер")
		w.WriteHeader(500)
		return
	}
	result.Id, err = database.NewMaster(result)
	if err != nil || result.Id == 0 {
		fmt.Println(err, "ошибка базы данных сохранения нового пользователя")
		w.WriteHeader(500)
		return
	}
	m, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err, "не удалось преобразовать данные нового мастера в json")
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
	w.Write(m)
}
func PutMasters(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err, "не удалось принять данные изменения масера")
		w.WriteHeader(404)
		return
	}
	var result types.Masters
	err = json.Unmarshal(b, &result)
	if err != nil {
		fmt.Println(err, "ошибка unmarshal изменения мастера")
		w.WriteHeader(500)
		return
	}
	vars := mux.Vars(r)
	result.Id, err = strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	err = database.ChangMaster(result)
	if err != nil {
		fmt.Println(err, "ошибка базы данных изменение данных масера")
		w.WriteHeader(500)
		return
	}
	m, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err, "ошибка преобразования данных в json")
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
	w.Write(m)
}
func DeleteMasters(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		fmt.Println(err, "неправильный id")
		w.WriteHeader(404)
		return
	}
	err = database.DeleteMaster(id)
	if err != nil {
		fmt.Println(err, "ошибка базы данных удаление мастера")
		w.WriteHeader(500)
		return
	}
	m, err := json.Marshal(id)
	if err != nil {
		fmt.Println(err, "ошибка преобразования данных в json")
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
	w.Write(m)
}
