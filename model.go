package main

import (
	"database/sql"
	"log"
	"net/http"
)

type player struct {
	id int32
	firstName string
	lastName string
	indexName string

}

func (p *player) getPlayer(db *sql.DB) {
	_ := return db.QueryRow("SELECT * from players WHERE index_name=$1", p.indexName).Scan(&p.id, &p.firstName, &p.lastName, &p.indexName)
}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello"))
	if err != nil {
		log.Fatalln("Fatal!")
	}
}
