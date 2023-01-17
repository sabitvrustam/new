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
	result := database.ReadMasters()
	m, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err, "")
		w.WriteHeader(404)
	}
	w.WriteHeader(200)
	w.Write(m)
}

func PostMasters(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, _ := io.ReadAll(r.Body)
	var res types.Masters
	err := json.Unmarshal(b, &res)
	if err != nil {
		fmt.Println(err)
	}
	id, err := database.NewMaster(res)
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
func PutMasters(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, _ := io.ReadAll(r.Body)
	var res types.Masters
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
	err = database.ChangMaster(res)
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
func DeleteMasters(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	err = database.DeleteMaster(id)
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
