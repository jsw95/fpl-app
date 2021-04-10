package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"os"
	"net/http"
)

type App struct {
	Router *mux.Router
	DB *sql.DB
}


func (a *App) Initialize(user, dbname, password string) {
	var connectionString string
	if os.Getenv("DATABASE_URL") != "" {
		connectionString = os.Getenv("DATABASE_URL")
	} else {
		connectionString = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
	}
	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter() 
	a.initializeRoutes()


}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/", HomePageHandler)
	a.Router.HandleFunc("/player_name/{IndexName}", a.getPlayer)
	a.Router.HandleFunc("/show_players", a.ShowPlayers)
	a.Router.HandleFunc("/login", a.AuthHandler)
}

func (a *App) Run(addr string) {
	log.Println("Trying to connect on " + addr)
	log.Fatal(http.ListenAndServe(addr, a.Router))
}



