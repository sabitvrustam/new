package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sabitvrustam/new/pkg/database/works"
	"github.com/sabitvrustam/new/pkg/types"
)

type WorkAPI struct {
	db   *sql.DB
	work *works.Work
}

func NewWork(db *sql.DB) *WorkAPI {
	return &WorkAPI{db: db, work: works.NewWork(db)}
}

func (a *WorkAPI) GetWorks(w http.ResponseWriter, r *http.Request) {
	result := a.work.ReadWoeks()
	m, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err, "")
		w.WriteHeader(404)
	}
	w.WriteHeader(200)
	w.Write(m)
}

func (a *WorkAPI) PostWork(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, _ := io.ReadAll(r.Body)
	var res types.Work
	err := json.Unmarshal(b, &res)
	if err != nil {
		fmt.Println(err)
	}
	id, err := a.work.WriteWork(res)
	if err != nil {
		fmt.Println(err)
	}
	res.Id = id
	m, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err, "")
	}
	w.WriteHeader(200)
	w.Write(m)
}
func (a *WorkAPI) PutWork(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, _ := io.ReadAll(r.Body)
	var res types.Work
	err := json.Unmarshal(b, &res)
	if err != nil {
		fmt.Println(err)
	}
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	res.Id = id
	err = a.work.ChangeWork(res)
	if err != nil {
		fmt.Println(err)
	}

	m, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err, "")
	}
	w.WriteHeader(200)
	w.Write(m)

}

func (a *WorkAPI) DeleteWork(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	err = a.work.DeleteWork(id)
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
