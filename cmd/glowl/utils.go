package main

import (
	"Glowl/internal/common/api"
	"log"
)

func serverRun() {
	srv, err := api.NewServer("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal(err)
	}
	if err := srv.Start(); err != nil {
		log.Fatal(err)
	}
}
