package repository

import (
	"fmt"

	"github.com/jessicajabes/aprendizes/dtos"
	"github.com/jmoiron/sqlx"
)

type DataBaseRepositoryUser struct {
	db *sqlx.DB
}

func NewDataBaseRepositoryUser(db *sqlx.DB) *DataBaseRepositoryUser {
	return &DataBaseRepositoryUser{db: db}
}

func (d DataBaseRepositoryUser) CreateUser(u dtos.User) error {
	fmt.Println(u.Username)
	fmt.Println(u.Password)
	fmt.Println(u.Secretary)
	_, err := d.db.Exec("INSERT into userap (username, password, secretary) values ($1, $2, $3)", u.Username, u.Password, u.Secretary)
	return err
}
