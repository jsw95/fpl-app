package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"time"
	"io/ioutil"
	"log"
	// "sync"
	"strconv"
)

type manager struct {
    Name string `json:"name"`
}

func main() {
	fmt.Println("h")

	// urls := []string{
	// 	"https://fantasy.premierleague.com/api/entry/1/",
	// 	"https://fantasy.premierleague.com/api/entry/2/",
	// 	"https://fantasy.premierleague.com/api/entry/3/",
	// }
	spaceClient := http.Client{
		Timeout: time.Second * 5, // Timeout after 2 seconds
	}
	baseUrl := "https://fantasy.premierleague.com/api/entry/"

	// httpClient = &http.Client(Timeout: 10 * time.Second)
	for i := 1; i < 1000; i++ {
		url := baseUrl + strconv.Itoa(i) + "/"
		// name := getManagerName(url)
		// fmt.Println(name)
		go getManagerName(url, spaceClient)


	}
	time.Sleep(5 * time.Second)
}

func getManagerName(url string, spaceClient http.Client) string {

	req, err := http.NewRequest(http.MethodGet, url, nil) 
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}
	
	if res.Body != nil {
		defer res.Body.Close()
	}
	
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	manager1 := manager{}
	jsonErr := json.Unmarshal(body, &manager1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	fmt.Println(manager1.Name)

	return manager1.Name
}