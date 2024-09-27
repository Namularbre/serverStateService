package main

import (
	"log"
	"net/http"
	"os"
	"serverStateService/handlers"
)

func main() {
	http.HandleFunc("/", handlers.GetServersHandler)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatalln("Error listening")
	}
}
