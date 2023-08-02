package repository

import (
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

	_, err := d.db.Exec("INSERT into userap (username, password, secretary, admin) values ($1, $2, $3, $4)", u.Username, u.Password, u.Secretary, u.Admin)
	return err
}
