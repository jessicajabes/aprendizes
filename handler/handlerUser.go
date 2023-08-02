package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/jessicajabes/aprendizes/dtos"
)

type UserStorager interface {
	CreateUser(dtos.User) error
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
