package main

import (
	"fmt"
	"log"
	"net-cat/internal/server"
)

func main() {
	s := server.NewServer()

	err := s.Start()
	if err != nil {
		fmt.Println(err)
		log.Fatal("Error starting server")
	}
}
