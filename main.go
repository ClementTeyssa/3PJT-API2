package main

import (
	"log"
	"net/http"
)

func main() {
	router := InitializeRouter()
	log.Println("Rooter initialised")

	log.Panic(http.ListenAndServe(":8080", router))
}
