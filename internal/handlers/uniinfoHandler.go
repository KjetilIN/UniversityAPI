package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"log"
	"uniapi/internal/constants"
)


func UniInfoHandler(w http.ResponseWriter, r *http.Request){
	//Setting header content
	w.Header().Set("content-type", "application/json")

	//Check if it is a GET request
	requestStatus := isCorrectRequestMethod(r);
	if(!requestStatus){
		http.Error(w, "Bad request. Only get is available for the uniinfo endpoint.", http.StatusBadRequest)
		return
	}

	//Get the search word from the path of the url 
	search := strings.TrimPrefix(r.URL.Path, constants.UNI_INFO_PATH)

	//Check if valid length, error is handled by the method
	isValid :=isOfValidLength(strings.Split(r.URL.Path, "/"),4,"Please add a search keyword(s).",w);
	if(!isValid){
		//Only return because the method handles the errors. 
		return
	}

	//Doing a GET request to the UNI API
	uniResponse, uniError := getUniversitiesWithName(search);
	if uniError != nil {
		log.Println("Error on get request to uni api: " + uniError.Error())
		http.Error(w, "Invalid request for " + search, http.StatusBadRequest)
		return 
	}
	defer uniResponse.Body.Close()

	// Prepare empty list of structs to populate
	var uniStructs []constants.UniStruct

	// Decode structs
	err := json.NewDecoder(uniResponse.Body).Decode(&uniStructs)
	if err != nil {
		log.Println("Error during decoding Uni API response to Uni Struct: " + err.Error())
		http.Error(w, "Error during decoding", http.StatusInternalServerError)
		return
	}else if (len(uniStructs) == 0){
		log.Println("No results for following url: " + constants.UNI_API_URL_PROD + "/search?name=" + search)
		http.Error(w, "No results for keyword: " + search, http.StatusBadRequest)
		return 
	}

	//Using the response from the Uni API and add the country info
	finalAPiResponse := addCountryInfoToUniversities(w, uniStructs)

	//Return results 
	encoder:= json.NewEncoder(w)
	err = encoder.Encode(finalAPiResponse)
	if err != nil{
		log.Println("Error during encoding the Uni Information Struct: " + err.Error())
		http.Error(w, "Error during encoding", http.StatusInternalServerError)
	}
}