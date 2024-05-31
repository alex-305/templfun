package main

import (
	"log"

	"github.com/alex-305/templfun/handlers"
)

func main() {
	serv := handlers.CreateServer("localhost:2200")
	err := serv.Start()

	if err != nil {
		log.Fatal("Server failed to start")
	}
}
