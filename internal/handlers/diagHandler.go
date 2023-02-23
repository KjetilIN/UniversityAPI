package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"uniapi/internal/constants"
	"uniapi/internal/serverstats"
)

// Uses the http package to do a GET request on the given link. Returns the status
func getStatusCode(link string) string{
	resp, _ := http.Get(link)
	return resp.Status
}

func DiagHandler(w http.ResponseWriter, r *http.Request){
	//Head Information
	w.Header().Set("content-type", "application/json")
	
	//Information for the client 
	info := constants.StatusInfo{
		UniApi: getStatusCode(constants.UNI_API_URL_PROD),
		CountryApi: getStatusCode(constants.COUNTRY_API_URL_PROD),
		Version: constants.VERSION,
		Uptime: int(serverstats.GetUptime().Seconds()),
	}
	
	//Encoding json
	encoder := json.NewEncoder(w);
	err:= encoder.Encode(info)

	//Handle error
	if err != nil{
		log.Println("Error on decoding diag struct to response writer: " + err.Error())
		http.Error(w, "Error on encoding", http.StatusInternalServerError);
		return
	}
}