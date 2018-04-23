package main

import (
	"fmt"
	"log"
	"os"

<<<<<<< HEAD:cmd/leadstdd/main.go
	"github.com/zsgilber/leads-tdd-copy/pkg/api"
=======
	"github.com/zsgilber/leads-tdd/pkg/api"
>>>>>>> master:cmd/leadstdd/main.go
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
<<<<<<< HEAD:cmd/leadstdd/main.go
=======
	dburl := os.Getenv("DATABASE_URL")
	log.Printf("URL is %s\n", dburl)
>>>>>>> master:cmd/leadstdd/main.go
	log.Printf("Listening on %s...\n", addr)
	a.Initialize(os.Getenv("DATABASE_URL"))

	a.Run(addr)
}
