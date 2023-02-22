package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func NeighbourUniHandler(w http.ResponseWriter, r *http.Request) {
	//Setting header content
	w.Header().Set("content-type", "application/json")

	// Parse the URL path
	path := strings.TrimSuffix(r.URL.Path, NEIGHBOUR_UNIS_PATH)
	pathParts := strings.Split(path, "/")

	// Check if path contains required variables
	if len(pathParts) < 6 {
		http.Error(w, "Invalid request path. Needs both country and middle: \nneighbourunis/{:country_name}/{:partial_or_complete_university_name}{?limit={:number}} ", http.StatusBadRequest)
		return
	}

	// Extract variables from the path
	countryName := pathParts[4]
	universityName := pathParts[5]

	// Check if limit is provided as a query parameter
	limitStr := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(limitStr)
	if(err != nil){
		// We set the variable as 0 for no limit 
		limit = 0
	}

	// The final response list 
	var neighbourunis []UniversityInfo


	//Get the list of ISO codes for neighbor contries;
	codes, borderError := getBorderCountry(countryName);
	if borderError != nil{
		http.Error(w, "No border countries found", http.StatusNoContent)
		return 
	}

	for _, code := range codes{
		// 1. Get the country name by alpha code 
		log.Println("Getting info from " + code)
		currentCountryName, err := getCountryFromAplhaCode(code);
		if err != nil{
			http.Error(w, "Not correct alpha code for; " + code, http.StatusBadRequest)
			return
		}
		// 2. Get all university that fit the country name and middle 

		var uniList []UniversityInfo

		uniResponse, err := getAllFromUniAPI(currentCountryName, universityName);
		if err != nil{
			http.Error(w, "Error on get All from Uni Api", http.StatusBadRequest)
			return 
		}

		var uniStructs []UniStuct

		// Decode structs
		err = json.NewDecoder(uniResponse.Body).Decode(&uniStructs)
		if err != nil {
			http.Error(w, "Error during decoding. Happened on adding country info", http.StatusBadRequest)
			return
		}

		//3. Add the country information for each response 
		uniList = addCountryInfoByName(w, uniStructs)

		// 3. Add them to the struct
		neighbourunis = append(neighbourunis, uniList...)

	}
	
	//Take only the limit, if there is a limit amount 
	if limit != 0{
		neighbourunis = neighbourunis[:limit]
	}

	//Encode the result 
	json.NewEncoder(w).Encode(neighbourunis)
}