package main

import (
	"log"
	"net/http"
)

func main() {
	router := InitializeRouter()

	log.Panic(http.ListenAndServe(":8080", router))
}
