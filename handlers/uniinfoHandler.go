package handlers

import (
	"encoding/json"
	"net/http"
)


func getFromUniAPI(searchWord string) *http.Response {
	resp, err := http.Get(UNI_API_URL)

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
	if uniResponse == nil{
		http.Error(w, "No result with = " + search, http.StatusNoContent);
		return;
	}


	// Prepare empty list of structs to populate
	uniStucts := []UniStuct{}

	// Decode structs
	err := json.NewDecoder(uniResponse.Body).Decode(&uniStucts)
	if err != nil {
		// Note: more often than not is this error due to client-side input, rather than server-side issues
		http.Error(w, "Error during decoding: "+err.Error(), http.StatusBadRequest)
		return
	}



	//Return results 
	encoder := json.NewEncoder(w);
	err = encoder.Encode(uniStucts)

	//Handle error
	if err != nil{
		http.Error(w, "Error on output!", http.StatusInternalServerError);
	}
}