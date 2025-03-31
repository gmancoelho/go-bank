package main

import (
	"fmt"
    "github.com/gmancoelho/go-bank/api"
)

func main() {
	fmt.Printf("Start Go Bank\n")

	server := newAPIServer(":8080")

	if err := server.Start(); err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
