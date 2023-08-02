package service

import (
	"github.com/jessicajabes/aprendizes/dtos"
)

type UserStorager interface {
	CreateUser(dtos.User) error
	GetAllUser() ([]dtos.User, error)
	GetIDUser(id int) (dtos.User, error)
	//GetUsernameUser(username string) (dtos.User, error)
	DeleteUser(id int) error
	PutUser(id int, a dtos.User) error
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
func (s ServiceUser) GetAllUser() ([]dtos.User, error) {
	return s.repository.GetAllUser()
}
func (s ServiceUser) GetIDUser(id int) (dtos.User, error) {
	return s.repository.GetIDUser(id)
}
func (s ServiceUser) DeleteUser(id int) error {
	return s.repository.DeleteUser(id)
}
func (s ServiceUser) PutUser(id int, a dtos.User) error {
	return s.repository.PutUser(id, a)
}
