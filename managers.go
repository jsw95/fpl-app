package main

import (
	"encoding/json"
	// "fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

)

type manager struct {
	Name string `json:"name"`
}

func addManagersToDb() {

	spaceClient := http.Client{
		Timeout: time.Second * 5, // Timeout after 2 seconds
	}
	baseUrl := "https://fantasy.premierleague.com/api/entry/"

	var wg sync.WaitGroup
	batchSize := 750
	numBatches := 10
	for batch := 0; batch < numBatches; batch++ {

		for i := 1; i < batchSize; i++ {
			wg.Add(1)
			userId := i + batch * batchSize
			url := baseUrl + strconv.Itoa(userId) + "/"
			go getManagerName(url, spaceClient, &wg)

		}
		wg.Wait()
	}
	// time.Sleep(5 * time.Second)
}

func getManagerName(url string, spaceClient http.Client, wg *sync.WaitGroup) string {
	defer wg.Done()
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
	
	return manager1.Name
}
