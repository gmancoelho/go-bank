package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Start Go Kanban")

	server := newAPIServer(":8080")

	if err := server.Start(); err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
