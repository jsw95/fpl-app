package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"text/template"
)

type player struct {
	Id        int32  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	IndexName string `json:"indexName"`

}

type ShowPlayerPage struct {
	PageTitle string
	Players []player
	TeamPlayers []player
}
type LoginPage struct {
	PageTitle string
}

func (p *player) getPlayer(db *sql.DB) error {
	return db.QueryRow("SELECT * from public.players WHERE index_name=$1",
		p.IndexName).Scan(&p.Id, &p.FirstName, &p.LastName, &p.IndexName)
}


// func getPlayersFromTeamID(teamId string) []player {

// 	baseUrl := "https://fantasy.premierleague.com/api/entry/"
// 	url := baseUrl + strconv.Itoa(teamId) + "/"



// }

func  (a *App) ShowPlayers(w http.ResponseWriter, r *http.Request){

	data := ShowPlayerPage{
		PageTitle: "Here are some players",
		Players: []player{},
	}
	indexName := r.FormValue("player_name")

	if indexName != "" {
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

		data.Players = append(data.Players, p)
	}


	tmpl, err := template.ParseFiles("templates/showPlayers.html")

	if err != nil {
		log.Fatalln(err)
	}

	_ = tmpl.Execute(w, data)




}

func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello"))
	if err != nil {
		log.Fatalln("Fatal!")
	}
}

func (a *App) AuthHandler(w http.ResponseWriter, r *http.Request) {
	data := LoginPage{
		PageTitle: "Login Here",
	}
	tmpl, err := template.ParseFiles("templates/auth.html")

	if err != nil {
		log.Fatalln(err)
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	getSessionCookies(username, password)

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Fatalln(err)
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


