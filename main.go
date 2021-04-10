package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	//logger = log.With(logger, "ts", log.DefaultTimestampUTC, "loc", log.DefaultCaller)

	app := App{}

	err := godotenv.Load("fpl.env")

	// if err != nil {
		// log.Fatal(err)
	// }
	app.Initialize(
		os.Getenv("user"),
		os.Getenv("database"),
		os.Getenv("user"),
		)

	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Running on " + addr)
	app.Run(addr)
}