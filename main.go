package main

import (
	// "github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	//logger = log.With(logger, "ts", log.DefaultTimestampUTC, "loc", log.DefaultCaller)
	
	a := App{}

	// err := godotenv.Load("fpl.env")

	// if err != nil {
		// log.Fatal(err)
	// }
	a.Initialize(
		os.Getenv("user"),
		os.Getenv("database"),
		os.Getenv("user"),
		)

	port := "8080"
	log.Println("Running on http://localhost:" + port)
	a.Run(":" + port)
}