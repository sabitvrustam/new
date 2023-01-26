package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sabitvrustam/new/pkg/database/parts"
	"github.com/sabitvrustam/new/pkg/types"
)

type PartAPI struct {
	db   *sql.DB
	part *parts.Part
}

func NewPartAPI(db *sql.DB) *PartAPI {
	return &PartAPI{db: db, part: parts.NewPart(db)}
}

func (a *PartAPI) GetParts(w http.ResponseWriter, r *http.Request) {
	result := a.part.ReadParts()
	m, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err, "")
		w.WriteHeader(404)
	}
	w.WriteHeader(200)
	w.Write(m)
}

func (a *PartAPI) PostParts(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, _ := io.ReadAll(r.Body)
	var res types.Part
	err := json.Unmarshal(b, &res)
	if err != nil {
		fmt.Println(err)
	}
	id, err := a.part.NewPart1(res)
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
func (a *PartAPI) PutPart(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, _ := io.ReadAll(r.Body)
	var res types.Part
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
	err = a.part.ChangePart(res)
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
func (a *PartAPI) DeletePart(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	err = a.part.DeletePart(id)
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
