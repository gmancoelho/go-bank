package main

import (
	"fmt"

	"github.com/gmancoelho/go-bank/api"
)

const address = ":8080"

func main() {
	fmt.Printf("Start Go Bank\n")

	server := api.NewAPIServer(address)

	if err := server.Start(); err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
