package main

import (
	"log"

	"github.com/elliot40404/icache-echo/internal/server"
)

func main() {
	log.Println("Starting server...")
	server.Run()
}
