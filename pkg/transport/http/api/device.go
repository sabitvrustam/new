package api

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sabitvrustam/new/pkg/database/devices"
	"github.com/sabitvrustam/new/pkg/types"
	"github.com/sirupsen/logrus"
)

type DeviceAPI struct {
	db     *sql.DB
	device *devices.Device
	log    *logrus.Logger
}

func NewDeviceAPI(db *sql.DB, log *logrus.Logger) *DeviceAPI {
	return &DeviceAPI{
		db:     db,
		device: devices.NewDevice(db, log),
		log:    log}
}

func (a *DeviceAPI) GetDevices(w http.ResponseWriter, r *http.Request) {
	result, err := a.device.DevicesRead(0, "")
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
func (a *DeviceAPI) GetDevicesSearch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sn := vars["sn"]
	result, err := a.device.DevicesRead(0, sn)
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
func (a DeviceAPI) GetDevice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		a.log.Error(err, "jfjh")
	}
	result, err := a.device.DevicesRead(id, "")
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

func (a *DeviceAPI) PostDevice(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		a.log.Error(err, "ошибка приема данных нового устройства от пользователя")
		w.WriteHeader(404)
		return
	}
	var result types.Device
	err = json.Unmarshal(b, &result)
	if err != nil {
		a.log.Error(err, "ошибка unmarshal нового устройства")
		w.WriteHeader(500)
		return
	}
	result.Id, err = a.device.DeviceWrite(result)
	if err != nil || result.Id == 0 {
		a.log.Error(err, "ошибка базы данных сохранения нового устройства")
		w.WriteHeader(500)
		return
	}
	m, err := json.Marshal(result)
	if err != nil {
		a.log.Error(err, "не удалось преобразовать данные нового устройства в json")
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
		a.log.Error(err, "не удалось принять данные изменения устройства")
		w.WriteHeader(404)
		return
	}
	var result types.Device
	err = json.Unmarshal(b, &result)
	if err != nil {
		a.log.Error(err, "ошибка unmarshal изменения устройства")
		w.WriteHeader(500)
		return
	}
	vars := mux.Vars(r)
	result.Id, err = strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		a.log.Error()
	}
	err = a.device.DeviceCange(result)
	if err != nil {
		a.log.Error(err, "ошибка базы данных изменение данных устройства")
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
func (a *DeviceAPI) DeleteDevice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		a.log.Error(err, "неправильный id")
		w.WriteHeader(404)
		return
	}
	err = a.device.DeviceDelete(id)
	if err != nil {
		a.log.Error(err, "ошибка базы данных удаление устройства")
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
