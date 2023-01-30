package api

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sabitvrustam/new/pkg/database/masters"
	"github.com/sabitvrustam/new/pkg/types"
	"github.com/sirupsen/logrus"
)

type MasterAPI struct {
	db     *sql.DB
	master *masters.Master
	log    *logrus.Logger
}

func NewMasterAPI(db *sql.DB, log *logrus.Logger) *MasterAPI {
	return &MasterAPI{
		db:     db,
		master: masters.NewMaster(db, log),
		log:    log}
}

func (a *MasterAPI) GetMasters(w http.ResponseWriter, r *http.Request) {
	result, err := a.master.MastersRead()
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

func (a *MasterAPI) PostMasters(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		a.log.Error(err, "ошибка приема данных нового мастера от пользователя")
		w.WriteHeader(404)
		return
	}
	var result types.Master
	err = json.Unmarshal(b, &result)
	if err != nil {
		a.log.Error(err, "ошибка unmarshal новый масер")
		w.WriteHeader(500)
		return
	}
	result.Id, err = a.master.MastersWrite(result)
	if err != nil || result.Id == 0 {
		a.log.Error(err, "ошибка базы данных сохранения нового пользователя")
		w.WriteHeader(500)
		return
	}
	m, err := json.Marshal(result)
	if err != nil {
		a.log.Error(err, "не удалось преобразовать данные нового мастера в json")
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
	w.Write(m)
}

func (a *MasterAPI) PutMasters(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		a.log.Error(err, "не удалось принять данные изменения масера")
		w.WriteHeader(404)
		return
	}
	var result types.Master
	err = json.Unmarshal(b, &result)
	if err != nil {
		a.log.Error(err, "ошибка unmarshal изменения мастера")
		w.WriteHeader(500)
		return
	}
	vars := mux.Vars(r)
	result.Id, err = strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		a.log.Error(err)
	}
	err = a.master.MastersChange(result)
	if err != nil {
		a.log.Error(err, "ошибка базы данных изменение данных масера")
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

func (a *MasterAPI) DeleteMasters(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		a.log.Error(err, "неправильный id")
		w.WriteHeader(404)
		return
	}
	err = a.master.MastersDelete(id)
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
