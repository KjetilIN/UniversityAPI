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

	if uniResponse.StatusCode != http.StatusOK {
		// The request was sucessfull but had no content
		http.Error(w, "UniApI: No result with found for '"+ search+"'", http.StatusNoContent)
		return
	}

	// Prepare empty list of structs to populate
	var uniStructs []UniStruct

	// Decode structs
	err := json.NewDecoder(uniResponse.Body).Decode(&uniStructs)
	if err != nil {
		http.Error(w, "Error during decoding. Happened on adding country info", http.StatusBadRequest)
		return
	}

	//Using the response from the Uni API and add the country info
	finalAPiResponse := addCountryInfoByName(w, uniStructs)

	//Return results 
	json.NewEncoder(w).Encode(finalAPiResponse)
}