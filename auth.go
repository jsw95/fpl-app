package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func getSessionCookies(username, password string) {
	authUrl := "https://users.premierleague.com/accounts/login/"

	
	postBody, _ := json.Marshal(map[string]string{
		"password": password,
		"login": username,
		"redirect_uri": "https://fantasy.premierleague.com/a/login",
		"app": "plfpl-web",
	 })
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(authUrl, "application/json", responseBody)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	sb := string(body)

	log.Println(sb[0])
} 