package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/sabitvrustam/new/pkg/database/count"
	"github.com/sirupsen/logrus"
)

type CountAPI struct {
	db    *sql.DB
	count *count.Count
	log   *logrus.Logger
}

func NewCountAPI(db *sql.DB, log *logrus.Logger) *CountAPI {
	return &CountAPI{
		db:    db,
		count: count.NewCount(db, log),
		log:   log}
}
func (a *CountAPI) GetCount(w http.ResponseWriter, r *http.Request) {
	result, err := a.count.CountRead()
	if err != nil {
		a.log.Error(err, "ошибка базы данных считывание устройств")
		w.WriteHeader(500)
	}
	m, err := json.Marshal(result)
	if err != nil {
		a.log.Error(err, "не удалось преобразовать данные устройств в json")
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
	w.Write(m)
}
