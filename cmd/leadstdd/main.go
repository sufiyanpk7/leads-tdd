package main

import (
	"fmt"
	"log"
	"os"
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
	a := App{}
	log.Printf("Listening on %s...\n", addr)
	a.Initialize(os.Getenv("DATABASE_URL"))

	a.Run(addr)
}
