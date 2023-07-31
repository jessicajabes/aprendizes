package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jessicajabes/aprendizes/handler"
	"github.com/jessicajabes/aprendizes/repository"
	"github.com/jessicajabes/aprendizes/service"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {

	db, err := sqlx.Connect("postgres", "user=postgres password=postgres dbname=aprendiz sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Conectado com sucesso!")

	myRouter := mux.NewRouter()
	dru := repository.NewDataBaseRepositoryUser(db)
	su := service.NewServiceUser(dru)
	hu := handler.NewHandlerUser(su)
	myRouter.HandleFunc("/user", hu.CreateUser).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":5555", myRouter))
}
