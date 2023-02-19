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

	// Prepare empty list of structs to populate 
	var uniStructs []UniStuct

	// Decode structs
	err := json.NewDecoder(uniResponse.Body).Decode(&uniStructs)
	if err != nil {
		http.Error(w, "Error during decoding: "+err.Error(), http.StatusBadRequest)
		return
	}

	//The final response to the 
	var uniInfoResponse []UniversityInfo

	//Loop over each of the university and add the langauges
	var currentCountryInfo []CountryInfo
	var currentCountry string

	for _, uni := range uniStructs{

		if(uni.Country != currentCountry){
			// DO API REQUEST and set the new countryinfo stuct
			countryResponse, countryErr := getFromCountryApi(uni.Country)
			if countryErr != nil {
				http.Error(w, "ContryRepsonse error!", http.StatusBadRequest)
				return 
			}
			defer countryResponse.Body.Close() //Close body always 

			// Decode struct
			err := json.NewDecoder(countryResponse.Body).Decode(&currentCountryInfo)
			if err != nil {
				http.Error(w, "Error during decoding of country: "+err.Error(), http.StatusBadRequest)
				return
			}

			//Sucessfully decoded the struct, so we set the new country info
			currentCountry = uni.Country
		}

		//Build the New Struct
		var newUniInfo UniversityInfo = UniversityInfo{ UniStuct: uni, CountryInfo: currentCountryInfo[0]};

		//Add them into the response list 
		uniInfoResponse = append(uniInfoResponse, newUniInfo)
	}


	//Return results 
	json.NewEncoder(w).Encode(uniInfoResponse)
}