package main

import (
	"log"
	"os"
)

func main() {
	port := os.Getenv("ROUTER_PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("router service starting on :%s", port)
	// TODO: init config, db, redis, http server
}
