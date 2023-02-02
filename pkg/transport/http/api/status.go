package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/sabitvrustam/new/pkg/database/status"
	"github.com/sirupsen/logrus"
)

type StatusAPI struct {
	db     *sql.DB
	log    *logrus.Logger
	status *status.Status
}

func NewStatusAPI(db *sql.DB, log *logrus.Logger) *StatusAPI {
	return &StatusAPI{
		db:     db,
		log:    log,
		status: status.NewStatus(db, log)}
}

func (a *StatusAPI) GetStatus(w http.ResponseWriter, r *http.Request) {
	result := a.status.ReadStatus()
	m, err := json.Marshal(result)
	if err != nil {
		a.log.Error(err, "")
		w.WriteHeader(404)
	}
	w.WriteHeader(200)
	w.Write(m)
}
