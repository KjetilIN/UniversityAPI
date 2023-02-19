package handlers

import (
	"encoding/json"
	"net/http"
)


func getFromUniAPI(searchWord string) *http.Response {
	resp, err := http.Get(UNI_API_URL_PROD + "/search?name="+searchWord)

	if err != nil{
		return nil
	}
	return resp;
}



func UniInfoHandler(w http.ResponseWriter, r *http.Request){
	//Setting header content
	w.Header().Set("content-type", "application/json")
	// Get all parameteres from the url
	keys:= r.URL.Query()

	//Search keyword
	search:=keys.Get("search")

	//Doing a GET request to the UNI API
	uniResponse := getFromUniAPI(search);
	defer uniResponse.Body.Close()

	if uniResponse.StatusCode != http.StatusOK {
		http.Error(w, "No result with = "+ search, uniResponse.StatusCode)
		return
	}


	// Prepare empty list of structs to populate 
	var uniStructs []UniStuct

	// Decode structs
	err := json.NewDecoder(uniResponse.Body).Decode(&uniStructs)
	if err != nil {
		http.Error(w, "Error during decoding: "+err.Error(), http.StatusBadRequest)
		return
	}
	//Return results 
	json.NewEncoder(w).Encode(uniStructs)
}