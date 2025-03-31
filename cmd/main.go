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

	store, err := s.NewPostgressStore()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", store)

	server := api.NewAPIServer(address)

	if err := server.Start(); err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
