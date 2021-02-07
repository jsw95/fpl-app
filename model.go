package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type player struct {
	Id        int32  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	IndexName string `json:"indexName"`

}

func (p *player) getPlayer(db *sql.DB) error {
	return db.QueryRow("SELECT * from public.players WHERE index_name=$1",
		p.IndexName).Scan(&p.Id, &p.FirstName, &p.LastName, &p.IndexName)
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello"))
	if err != nil {
		log.Fatalln("Fatal!")
	}
}

func (a *App) getPlayer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	indexName, _ := vars["IndexName"]
	log.Println("getPlayer triggered " + indexName)

	p := player{IndexName: indexName}

	err := p.getPlayer(a.DB)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Player not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, p)
}
