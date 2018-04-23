package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zsgilber/leads-tdd/pkg/api"
)

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}

func main() {
	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}
	a := api.App{}
	dburl := os.Getenv("DATABASE_URL")
	log.Printf("URL is %s\n", dburl)
	log.Printf("Listening on %s...\n", addr)
	a.Initialize(os.Getenv("DATABASE_URL"))

	a.Run(addr)
}
