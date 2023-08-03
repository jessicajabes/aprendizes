package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/jessicajabes/aprendizes/dtos"
)

type ServiceLogin interface {
	Login(u dtos.User) (string, error)
}

type HandlerLogin struct {
	service ServiceLogin
}

func NewHandlerLogin(service ServiceLogin) *HandlerLogin {
	return &HandlerLogin{service: service}
}

func (h HandlerLogin) Login(w http.ResponseWriter, r *http.Request) {
	var u dtos.User
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = json.Unmarshal(data, &u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	tokenString, err := h.service.Login(u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(dtos.LoginResponse{AccessToken: tokenString})

}
