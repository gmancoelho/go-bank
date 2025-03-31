package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Start Go Bank\n")

	server := newAPIServer(":8080")

	if err := server.Start(); err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
