package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"net/http"
	"log"
)

type App struct {
	Router *mux.Router
	DB *sql.DB
}


func (a *App) Initialize(user, dbname, password string) {

	a.Router = mux.NewRouter()
	a.initializeRoutes()


}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/", HomePageHandler)
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8010", a.Router))
}



