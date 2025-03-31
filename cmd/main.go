package main

import (
	"fmt"
	"log"

	"github.com/gmancoelho/go-bank/api"
	s "github.com/gmancoelho/go-bank/repository"
)

const address = ":8080"

func main() {
	fmt.Printf("Start Go Bank\n")

	store, err := s.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(address, store)

	if err := server.Start(); err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
