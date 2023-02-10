package main

import (
	"log"
	"net/http"
	"os"
	"uniapi/handlers"
)

func main(){

	//Get the port from local or env 
	port := os.Getenv("PORT")
	if port == ""{
		log.Println("Port has not been set. Default: 8080")
		port = "8080"
	}
	//Handlers
	http.HandleFunc(handlers.DEFAULT_PATH, handlers.DefaultHandler)

	//Start Server
	log.Println("Starting server on port " + port + "...")
	log.Fatal(http.ListenAndServe(":"+ port, nil))

}