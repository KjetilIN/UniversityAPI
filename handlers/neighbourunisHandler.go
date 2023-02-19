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
		limit = 0
	}

	//Get all border alpha codes 
	//borderCodes, err := getBorderCountry(countryName);
	if err != nil{
		http.Error(w, "No border countries found", http.StatusNoContent)
		return 
	}

	// The final response list 
	var neighbourunis []UniversityInfo


	//Get the list of ISO codes for neighbor contries;
	codes, _ := getBorderCountry(countryName);

	for _, code := range codes{
		// 1. Get the country name by alpha code 
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

		//3. Add the country information for each response 
		uniList = addCountryInfoByName(w, *uniResponse)

		// 3. Add them to the struct
		neighbourunis = append(neighbourunis, uniList...)

	}

	log.Println(limit)

	json.NewEncoder(w).Encode(neighbourunis)
	
}