package main

import (
	"log"
	"net/http"
	"os"
)

func main(){

	//Get the port from local or env 
	port := os.Getenv("PORT")
	defaultPort := 8080;
	if port == ""{
		log.Println("Port has not been set. Default: ", defaultPort)
	}
	//Handlers

	//Start Server
	log.Println("Starting server on port " + port + "...")
	log.Fatal(http.ListenAndServe(":"+port, nil))

}