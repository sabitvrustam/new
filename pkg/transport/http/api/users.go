package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sabitvrustam/new/pkg/database/users"
	"github.com/sabitvrustam/new/pkg/types"
)

type UserAPI struct {
	db   *sql.DB
	user *users.Users
}

func NewUsersAPI(db *sql.DB) *UserAPI {
	return &UserAPI{db: db, user: users.NewUser(db)}
}

func (a *UserAPI) GetUsers(w http.ResponseWriter, r *http.Request) {
	result, err := a.user.ReadUsers()
	if err != nil {
		fmt.Println(err, "ошибка базы данных считывание пользователей")
		w.WriteHeader(500)
	}
	m, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err, "не удалось преобразовать данные пользователей в json")
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
	w.Write(m)
}
func (a *UserAPI) GetUsersSearch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lName := vars["l_name"]
	result, err := a.user.ReadUsersSearch(lName)
	if err != nil {
		fmt.Println(err, "ошибка базы данных считывание пользователей")
		w.WriteHeader(500)
	}
	m, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err, "не удалось преобразовать данные пользователей в json")
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
	w.Write(m)

}
func (a *UserAPI) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	result, err := a.user.ReadUser(id)
	if err != nil {
		fmt.Println(err, "ошибка базы данных считывание пользователя")
		w.WriteHeader(500)
	}
	m, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err, "не удалось преобразовать данные пользователя в json")
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
	w.Write(m)
}

func (a *UserAPI) PostUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err, "ошибка приема данных нового клиента от пользователя")
		w.WriteHeader(404)
		return
	}
	var result types.User
	err = json.Unmarshal(b, &result)
	if err != nil {
		fmt.Println(err, "ошибка unmarshal новый клиент")
		w.WriteHeader(500)
		return
	}
	result.Id, err = a.user.NewUser1(result)
	if err != nil || result.Id == 0 {
		fmt.Println(err, "ошибка базы данных сохранения нового пользователя")
		w.WriteHeader(500)
		return
	}
	m, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err, "не удалось преобразовать данные нового пользователя в json")
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
	w.Write(m)
}
func (a *UserAPI) PutUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err, "не удалось принять данные изменения клиента")
		w.WriteHeader(404)
		return
	}
	var result types.User
	err = json.Unmarshal(b, &result)
	if err != nil {
		fmt.Println(err, "ошибка unmarshal изменения клиента")
		w.WriteHeader(500)
		return
	}
	vars := mux.Vars(r)
	result.Id, err = strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	err = a.user.ChangUser(result)
	if err != nil {
		fmt.Println(err, "ошибка базы данных изменение данных клиента")
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
func (a *UserAPI) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		fmt.Println(err, "неправильный id")
		w.WriteHeader(404)
		return
	}
	err = a.user.DelUser(id)
	if err != nil {
		fmt.Println(err, "ошибка базы данных удаление клиента")
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
