package main

import (
	"log"
	"os"
)

func main() {
	port := os.Getenv("CONTEXT_BRIDGE_PORT")
	if port == "" {
		port = "8081"
	}
	log.Printf("context-bridge service starting on :%s", port)
	// TODO: init config, redis, postgres, http server
}
