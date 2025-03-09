package main

import (
	"hexapi/domain"
	"hexapi/handler"
	"hexapi/service"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func main() {

	log.Println("Connecting to the database")
	db, err := sqlx.Open("mysql", "root@tcp(127.0.0.1)/test")
	if err != nil {
		panic(err)
	}

	// ph := handler.NewPersonHandler(service.NewPersonService(domain.NewPersonRepositoryStub()))
	ph := handler.NewPersonHandler(service.NewPersonService(domain.NewPersonRepositoryMysql(db)))

	r := mux.NewRouter()
	r.HandleFunc("/persons", ph.GetAllPersons).Methods(http.MethodGet)
	r.HandleFunc("/persons/{id:[0-9]+}", ph.GetPersonById).Methods(http.MethodGet)
	r.HandleFunc("/persons", ph.CreatePerson).Methods(http.MethodPost)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Starting server")
	log.Fatal(srv.ListenAndServe())
}
