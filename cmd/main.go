package main

import (
	"log"

	"github.com/khalednadam/ecomgo/cmd/http"
)

func main() {
	server := http.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
