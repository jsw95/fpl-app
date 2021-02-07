package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	a := App{}

	err := godotenv.Load("fpl.env")

	if err != nil {
		log.Fatalf("error")
	}
	a.Initialize(
		os.Getenv("user"),
		os.Getenv("database"),
		os.Getenv("user"),
		)

	a.Run(":8080")

}