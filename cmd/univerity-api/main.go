package main

import (
	"log"
	"net/http"
	"os"
	"uniapi/internal/serverstats"
	"uniapi/internal/handlers"
	"uniapi/internal/mock"
	"uniapi/internal/constants"
)

//Main function to run the service
func main(){
	//Get the port from local or env 
	port := os.Getenv("PORT")
	if port == ""{
		log.Println("Port has not been set. Default: 8080")
		port = "8080"
	}

	//Starting server timer
	log.Println("Server timer starting.....")
	serverstats.InitServerTimer()


	//Handlers
	http.HandleFunc(constants.DIAG_PATH, handlers.DiagHandler)
	http.HandleFunc(constants.UNI_INFO_PATH, handlers.UniInfoHandler)
	http.HandleFunc(constants.NEIGHBOR_UNIS_PATH, handlers.NeighborUniHandler)

	//Mock handlers
	http.HandleFunc(constants.MOCK_COUNTRY_PATH, mock.CountryMockHandler)
	http.HandleFunc(constants.MOCK_UNI_PATH, mock.UniMockHandler)


	//Catch-all handler for undefined endpoints! (replaced the default handler)
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.Error(w, "Endpoint not found!\n See the docs for the defined endpoints!", http.StatusBadRequest)
    })


	//Start Server
	log.Println("Starting server on port " + port + "...")
	log.Fatal(http.ListenAndServe(":"+ port, nil))

}