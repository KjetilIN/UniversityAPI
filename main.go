package main

import (
	"log"
	"net/http"
	"os"
	"uniapi/handlers"
	"uniapi/mock"
	"uniapi/serverstats"
)


func main(){

	//Get the port from local or env 
	port := os.Getenv("PORT")
	if port == ""{
		log.Println("Port has not been set. Default: 8080")
		port = "8080"
	}

	//Starting server timer
	log.Println("Server timer starting.....")
	serverstats.Init()


	//Handlers
	http.HandleFunc(handlers.DEFAULT_PATH, handlers.DefaultHandler)
	http.HandleFunc(handlers.DIAG_PATH, handlers.DiagHandler)
	http.HandleFunc(handlers.UNI_INFO_PATH, handlers.UniInfoHandler)
	http.HandleFunc(handlers.NEIGHBOUR_UNIS_PATH, handlers.NeighbourUniHandler)

	//Mock handlers
	http.HandleFunc(handlers.MOCK_COUNTRY_PATH, mock.CountryMockHandler)
	http.HandleFunc(handlers.MOCK_UNI_PATH, mock.UniMockHandler)

	//Start Server
	log.Println("Starting server on port " + port + "...")
	log.Fatal(http.ListenAndServe(":"+ port, nil))

}