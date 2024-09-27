package main

import (
	"log"
	"net/http"
	"os"
	"serverStateService/handlers"
)

func main() {
	http.HandleFunc("/", handlers.GetServersHandler)

	completeAddress := os.Getenv("ADDRESS") + ":" + os.Getenv("PORT")

	log.Println(completeAddress)

	err := http.ListenAndServe(completeAddress, nil)
	if err != nil {
		log.Fatalln("Error listening")
	}
}
