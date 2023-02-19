package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

// Create a reusable http.Client that is used by the uniinfo handler
var httpClient = &http.Client{
	Timeout: time.Second * 10, // Add a timeout to avoid hanging connections
}

//Function that setup the GET request and return error
func getFromUniAPI(searchWord string) (*http.Response, error) {
	// Create a new GET request
	req, err := http.NewRequest("GET", UNI_API_URL_PROD+"/search", nil)
	if err != nil {
		return nil, err
	}

	// Add the search query parameter to the URL
	// In this case add the name as parameter
	q := req.URL.Query()
	q.Add("name", searchWord)
	req.URL.RawQuery = q.Encode()

	// Send the request using the shared http.Client
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Return the response and any errors
	return resp, err
}


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