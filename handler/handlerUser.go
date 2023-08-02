package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jessicajabes/aprendizes/dtos"
)

type UserStorager interface {
	CreateUser(dtos.User) error
	GetAllUser() ([]dtos.User, error)
	GetIDUser(id int) (dtos.User, error)
	DeleteUser(id int) error
	PutUser(id int, a dtos.User) error
}

type HandlerUser struct {
	service UserStorager
}

func NewHandlerUser(service UserStorager) *HandlerUser {
	return &HandlerUser{service: service}
}

func (h HandlerUser) CreateUser(w http.ResponseWriter, r *http.Request) {
	var u dtos.User
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	err = json.Unmarshal(data, &u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	err = h.service.CreateUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

}

func (h HandlerUser) GetAllUser(w http.ResponseWriter, r *http.Request) {
	lista, err := h.service.GetAllUser()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(lista)
}
func (h HandlerUser) GetIDUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	data, err := h.service.GetIDUser(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(data)

}

func (h HandlerUser) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	err = h.service.DeleteUser(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
}

func (h HandlerUser) PutUser(w http.ResponseWriter, r *http.Request) {
	var u dtos.User

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = json.Unmarshal(data, &u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = h.service.PutUser(id, u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
}
