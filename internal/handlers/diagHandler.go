package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"uniapi/internal/constants"
	"uniapi/internal/serverstats"
)

// Uses the http package to do a GET request on the given link. Returns the status
func getStatusCode(link string) string{
	resp, _ := http.Get(link)
	return resp.Status
}

//Checks if the diag url is of valid length
func isDiagUrlValid(url string, w http.ResponseWriter) bool {
	//Splitting the string
	words := strings.Split(url, "/")

	//Removing empty strings created by the string.split
	words = removeEmptyStrings(words) 
	if(len(words) != 3){
		log.Println("Bad request. Url not long enough. Should be 3, was ", len(words))
		http.Error(w, "Bad request. Longer than required", http.StatusBadRequest)
		return false
	}
	return true
}

func DiagHandler(w http.ResponseWriter, r *http.Request){
	//Head Information
	w.Header().Set("content-type", "application/json")

	//check the url 
	urlStatus := isDiagUrlValid(r.URL.Path, w);
	if(!urlStatus){
		return
	}

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