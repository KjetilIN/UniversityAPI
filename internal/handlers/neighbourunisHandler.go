package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"uniapi/internal/constants"
)

func NeighborUniHandler(w http.ResponseWriter, r *http.Request) {
	//Setting header content
	w.Header().Set("content-type", "application/json")

	// Parse the URL path
	path := strings.TrimSuffix(r.URL.Path, constants.NEIGHBOR_UNIS_PATH)
	pathParts := strings.Split(path, "/")

	// Check if path contains required variables
	if len(pathParts) != 6 {
		log.Println("Error on amount of parameters!")
		http.Error(w, "Invalid request path. Needs both country and middle: \n neighbourunis/{:country_name}/{:partial_or_complete_university_name}{?limit={:number}} ", http.StatusBadRequest)
		return
	}

	// Extract variables from the path
	countryName := pathParts[4]
	universityName := pathParts[5]

	// Check if limit is provided as a query parameter
	limitStr := r.URL.Query().Get("limit")
	limit, err := strconv.Atoi(limitStr)
	if(err != nil){
		// This might trow error because not a number or it doesn't exists. 
		// We set the variable as 0 for no limit, indicating no limit. 
		limit = 0
	}

	// The final response list 
	var neighborUnis []constants.UniversityInfo


	//Get the list of ISO codes for neighbor countries;
	codes, borderError := GetBorderCountry(countryName);
	if borderError != nil{
		log.Println("Error on get border country: " + err.Error())
		http.Error(w, "No border countries found", http.StatusNoContent)
		return 
	}

	for _, code := range codes{
		// 1. Get the country name by alpha code 
		currentCountryName, err := GetCountryFromAlphaCode(code);
		if err != nil{
			log.Println("Error on get Country from the alpha code method: " + err.Error())
			http.Error(w, "Invalid alpha code given: " + code, http.StatusBadRequest)
			return
		}
		// 2. Get all university that fit the country name and middle 

		var uniList []constants.UniversityInfo

		uniResponse, err := GetAllFromUniAPI(currentCountryName, universityName);
		if err != nil{
			log.Println("Error while trying to use middle and country name to do a GET request: ", err.Error())
			http.Error(w, "Invalid Request. See docs for usage.", http.StatusBadRequest)
			return 
		}

		var uniStruct []constants.UniStruct

		// Decode struct
		err = json.NewDecoder(uniResponse.Body).Decode(&uniStruct)
		if err != nil {
			log.Println("Error during decoding Uni API response to Uni Struct: " + err.Error())
			http.Error(w, "Error during decoding. Happened on adding country info", http.StatusInternalServerError)
			return
		}

		//3. Add the country information for each response 
		uniList = AddCountryInfoByName(w, uniStruct)

		// 3. Add them to the struct
		neighborUnis = append(neighborUnis, uniList...)

	}
	
	//Take only the limit, if there is a limit amount 
	if limit != 0{
		neighborUnis = neighborUnis[:limit]
	}

	//Return results 
	encoder:= json.NewEncoder(w)
	err = encoder.Encode(neighborUnis)
	if err != nil{
		log.Println("Error during encoding the Uni Information Struct from the Neighbor api endpoint: " + err.Error())
		http.Error(w, "Error during encoding", http.StatusInternalServerError)
	}
}