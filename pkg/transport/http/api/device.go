package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sabitvrustam/new/pkg/database"
	"github.com/sabitvrustam/new/pkg/types"
)

type DeviceAPI struct {
	db     *sql.DB
	device *database.Device
}

func NewDeviceAPI(db *sql.DB) *DeviceAPI {
	return &DeviceAPI{db: db, device: database.NewDevice(db)}
}

func (a *DeviceAPI) GetDevices(w http.ResponseWriter, r *http.Request) {
	result, err := a.device.ReadDevices(0, "")
	if err != nil {
		fmt.Println(err, "ошибка базы данных считывание устройств")
		w.WriteHeader(500)
	}
	m, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err, "не удалось преобразовать данные устройств в json")
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
	w.Write(m)
}
func (a *DeviceAPI) GetDevicesSearch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sn := vars["sn"]
	result, err := a.device.ReadDevices(0, sn)
	if err != nil {
		fmt.Println(err, "ошибка базы данных считывание устройств")
		w.WriteHeader(500)
	}
	m, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err, "не удалось преобразовать данные устройств в json")
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
	w.Write(m)

}
func (a DeviceAPI) GetDevice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	result, err := a.device.ReadDevices(id, "")
	if err != nil {
		fmt.Println(err, "ошибка базы данных считывание устройств")
		w.WriteHeader(500)
	}
	m, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err, "не удалось преобразовать данные устройств в json")
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
	w.Write(m)
}

func (a *DeviceAPI) PostDevice(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err, "ошибка приема данных нового устройства от пользователя")
		w.WriteHeader(404)
		return
	}
	var result types.Device
	err = json.Unmarshal(b, &result)
	if err != nil {
		fmt.Println(err, "ошибка unmarshal нового устройства")
		w.WriteHeader(500)
		return
	}
	result.Id, err = a.device.NewDevice1(result)
	if err != nil || result.Id == 0 {
		fmt.Println(err, "ошибка базы данных сохранения нового устройства")
		w.WriteHeader(500)
		return
	}
	m, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err, "не удалось преобразовать данные нового устройства в json")
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(200)
	w.Write(m)
}
func (a *DeviceAPI) PutDevice(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err, "не удалось принять данные изменения устройства")
		w.WriteHeader(404)
		return
	}
	var result types.Device
	err = json.Unmarshal(b, &result)
	if err != nil {
		fmt.Println(err, "ошибка unmarshal изменения устройства")
		w.WriteHeader(500)
		return
	}
	vars := mux.Vars(r)
	result.Id, err = strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	err = a.device.ChangDevice(result)
	if err != nil {
		fmt.Println(err, "ошибка базы данных изменение данных устройства")
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
func (a *DeviceAPI) DeleteDevice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		fmt.Println(err, "неправильный id")
		w.WriteHeader(404)
		return
	}
	err = a.device.DelDevice(id)
	if err != nil {
		fmt.Println(err, "ошибка базы данных удаление устройства")
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
