package service

import (
	"github.com/jessicajabes/aprendizes/dtos"
)

type UserStorager interface {
	CreateUser(dtos.User) error
}

type ServiceUser struct {
	repository UserStorager
}

func NewServiceUser(repository UserStorager) *ServiceUser {
	return &ServiceUser{repository: repository}
}

func (s ServiceUser) CreateUser(user dtos.User) error {
	return s.repository.CreateUser(user)
}
