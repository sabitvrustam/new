package api

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sabitvrustam/new/pkg/database/users"
	"github.com/sabitvrustam/new/pkg/types"
	"github.com/sirupsen/logrus"
)

type UserAPI struct {
	db   *sql.DB
	log  *logrus.Logger
	user *users.Users
}

func NewUsersAPI(db *sql.DB, log *logrus.Logger) *UserAPI {
	return &UserAPI{
		db:   db,
		log:  log,
		user: users.NewUser(db, log)}
}

func (a *UserAPI) GetUsers(w http.ResponseWriter, r *http.Request) {
	result, err := a.user.ReadUsers(nil, nil)
	if err != nil {
		a.log.Error(err, "ошибка базы данных считывание пользователей")
		w.WriteHeader(500)
	}
	m, err := json.Marshal(result)
	if err != nil {
		a.log.Error(err, "не удалось преобразовать данные пользователей в json")
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
	w.Write(m)
}

func (a *UserAPI) GetUsersSearch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lName := vars["l_name"]
	result, err := a.user.ReadUsers(&lName, nil)
	if err != nil {
		a.log.Error(err, "ошибка базы данных считывание пользователей")
		w.WriteHeader(500)
	}
	m, err := json.Marshal(result)
	if err != nil {
		a.log.Error(err, "не удалось преобразовать данные пользователей в json")
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
		a.log.Error(err)
	}
	result, err := a.user.ReadUsers(nil, &id)
	if err != nil {
		a.log.Error(err, "ошибка базы данных считывание пользователя")
		w.WriteHeader(500)
	}
	m, err := json.Marshal(result)
	if err != nil {
		a.log.Error(err, "не удалось преобразовать данные пользователя в json")
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
		a.log.Error(err, "ошибка приема данных нового клиента от пользователя")
		w.WriteHeader(404)
		return
	}
	var result types.User
	err = json.Unmarshal(b, &result)
	if err != nil {
		a.log.Error(err, "ошибка unmarshal новый клиент")
		w.WriteHeader(500)
		return
	}
	result.Id, err = a.user.NewUser1(result)
	if err != nil || result.Id == 0 {
		a.log.Error(err, "ошибка базы данных сохранения нового пользователя")
		w.WriteHeader(500)
		return
	}
	m, err := json.Marshal(result)
	if err != nil {
		a.log.Error(err, "не удалось преобразовать данные нового пользователя в json")
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
		a.log.Error(err, "не удалось принять данные изменения клиента")
		w.WriteHeader(404)
		return
	}
	var result types.User
	err = json.Unmarshal(b, &result)
	if err != nil {
		a.log.Error(err, "ошибка unmarshal изменения клиента")
		w.WriteHeader(500)
		return
	}
	vars := mux.Vars(r)
	result.Id, err = strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		a.log.Error(err)
	}
	err = a.user.ChangUser(result)
	if err != nil {
		a.log.Error(err, "ошибка базы данных изменение данных клиента")
		w.WriteHeader(500)
		return
	}
	m, err := json.Marshal(result)
	if err != nil {
		a.log.Error(err, "ошибка преобразования данных в json")
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
		a.log.Error(err, "неправильный id")
		w.WriteHeader(404)
		return
	}
	err = a.user.DelUser(id)
	if err != nil {
		a.log.Error(err, "ошибка базы данных удаление клиента")
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
