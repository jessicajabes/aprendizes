package service

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jessicajabes/aprendizes/dtos"
)

type UserGetter interface {
	GetUsernameUser(u string) (dtos.User, error)
}

type ServiceLogin struct {
	repository UserGetter
	secret     string
}

func NewServiceLogin(repository UserGetter) *ServiceLogin {
	return &ServiceLogin{repository: repository}
}

func (s ServiceLogin) Login(u dtos.User) (string, error) {
	data, err := s.repository.GetUsernameUser(u.Username)
	if err != nil {
		return "", err
	}

	if u.Password != data.Password {
		return "", errors.New("username or password does not exist")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    fmt.Sprintf("%d", data.ID),
		"user_admin": fmt.Sprintf("%t", data.Admin),
	})
	return token.SignedString([]byte(s.secret))

}

func (s ServiceLogin) Validate(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secret), nil
	})
	if err != nil {
		return false, err
	}
	return token.Valid, nil
}

func (s ServiceLogin) GetUserID(tokenString string) (int, error) {
	var claims jwt.MapClaims
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secret), nil
	})
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(claims["user_id"].(string))
}

func (s ServiceLogin) GetUserAdmin(tokenString string) (bool, error) {
	var claims jwt.MapClaims
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secret), nil
	})
	if err != nil {
		return false, err
	}
	return strconv.ParseBool(claims["user_admin"].(string))
}
