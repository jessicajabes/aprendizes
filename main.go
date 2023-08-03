package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jessicajabes/aprendizes/handler"
	"github.com/jessicajabes/aprendizes/repository"
	"github.com/jessicajabes/aprendizes/service"
	"github.com/jessicajabes/aprendizes/utils"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Middleware struct {
	sl *service.ServiceLogin
}

func (m Middleware) AuthRoute(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		valid, err := m.sl.Validate(token)
		if !valid || err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Not Authorized"))
			return
		}

		userID, err := m.sl.GetUserID(token)
		if !valid || err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Not Authorized"))
			return
		}

		userAdmin, err := m.sl.GetUserAdmin(token)
		if !valid || err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Not Authorized"))
			return
		}
		next(w, r.WithContext(utils.ContextWithUserAdmin((utils.ContextWithUserID(r.Context(), userID)), userAdmin)))
	}
}

func main() {

	db, err := sqlx.Connect("postgres", "user=postgres password=postgres dbname=aprendiz sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Successfully connected!")

	myRouter := mux.NewRouter()
	dru := repository.NewDataBaseRepositoryUser(db)
	sl := service.NewServiceLogin(dru)
	hl := handler.NewHandlerLogin(sl)

	myRouter.HandleFunc("/login", hl.Login).Methods(http.MethodPost)

	sr := service.NewServiceUser(dru)
	hu := handler.NewHandlerUser(sr)

	m := Middleware{
		sl: sl,
	}
	myRouter.HandleFunc("/user", m.AuthRoute(hu.CreateUser)).Methods(http.MethodPost)
	myRouter.HandleFunc("/user", m.AuthRoute(hu.GetAllUser)).Methods(http.MethodGet)
	myRouter.HandleFunc("/user/{id}", m.AuthRoute(hu.GetIDUser)).Methods(http.MethodGet)
	myRouter.HandleFunc("/user/{id}", m.AuthRoute(hu.DeleteUser)).Methods(http.MethodDelete)
	myRouter.HandleFunc("/user/{id}", m.AuthRoute(hu.PutUser)).Methods(http.MethodPut)

	log.Fatal(http.ListenAndServe(":5555", myRouter))
}
