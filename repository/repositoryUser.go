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

func (d DataBaseRepositoryUser) GetAllUser() ([]dtos.User, error) {
	lista := []dtos.User{}
	err := d.db.Select(&lista, "SELECT * from userap")
	if err != nil {
		return nil, err
	}
	return lista, nil
}

func (d DataBaseRepositoryUser) GetIDUser(id int) (dtos.User, error) {
	var user dtos.User
	err := d.db.Get(&user, "SELECT * from userap where ID=$1", id)
	if err != nil {
		return dtos.User{}, err
	}
	return user, nil
}

func (d DataBaseRepositoryUser) GetUsernameUser(username string) (dtos.User, error) {
	var user dtos.User
	err := d.db.Get(&user, "SELECT * from userap where username=$1", username)
	if err != nil {
		return dtos.User{}, err
	}
	return user, nil
}

func (d DataBaseRepositoryUser) DeleteUser(id int) error {
	_, err := d.db.Exec("DELETE from userap where id=$1", id)
	return err
}

func (d DataBaseRepositoryUser) PutUser(id int, a dtos.User) error {
	_, err := d.db.Exec("UPDATE userap SET username=$1, password=$2,secretary=$3,admin=$4 where id=$5", a.Username, a.Password, a.Secretary, a.Admin, id)
	return err
}
