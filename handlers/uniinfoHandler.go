package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
)



func UniInfoHandler(w http.ResponseWriter, r *http.Request){
	//Setting header content
	w.Header().Set("content-type", "application/json")


	//Get the search word from the path of the url 
	search := strings.TrimPrefix(r.URL.Path, UNI_INFO_PATH)

	//Doing a GET request to the UNI API
	uniResponse, uniError := getFromUniAPI(search);
	if uniError != nil {
		http.Error(w, "UniResponse error!", http.StatusBadRequest)
		return 
	}
	defer uniResponse.Body.Close() //Close body always 

	if uniResponse.StatusCode != http.StatusOK {
		// The request was sucessfull but had no content
		http.Error(w, "UniApI: No result with found for '"+ search+"'", http.StatusNoContent)
		return
	}

	//Using the response from the Uni API and add the country info
	finalAPiResponse := addCountryInfoByName(w,*uniResponse)

	//Return results 
	json.NewEncoder(w).Encode(finalAPiResponse)
}