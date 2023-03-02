package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"
	"uniapi/internal/constants"
)

// Create a reusable http.Client that is used by the uni info handler
var httpClient = &http.Client{
	Timeout: time.Second * 10, // Add a timeout to avoid hanging connections
}

// Function that setup the GET request and returns the response and  error
func getUniversitiesWithName(searchWord string) (*http.Response, error) {
	// Building the url 
	URL := constants.UNI_API_URL_PROD + "/search?name=" + replaceSpaces(searchWord) 
	
	// Create a new GET request
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return nil, err
	}

	//Logging the request
	log.Println("GET with getUniversitiesWithName: " + URL)

	// Send the request using the shared http.Client
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Return the response and any errors
	return resp, err
}

// Function that does a GET request with the given country name and return response and error
func getCountriesFromCountryName(country string) (*http.Response, error) {
	// URL
	URL := constants.COUNTRY_API_NAME_URL_PROD + "/" + replaceSpaces(country)

	// Create a new GET request
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return nil, err
	}

	//Logging the request
	log.Println("GET with getCountriesFromCountryName: ", URL)

	// Send the request using the shared http.Client
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Return the response and any errors
	return resp, err
}

// Function gets the country name from the alpha code.
// Does a get request to the alpha code 
func getCountryNameFromAlphaCode(code string) (string, error) {
	// URL
	URL := constants.COUNTRY_API_ALPHA_URL_PROD + "/" + code

	// Create a new GET request
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return "", err
	}

	//Logging the request
	log.Println("GET country by Alpha code: ", URL)

	// Send the request using the shared http.Client
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}

	//Prepare to populate the border countries struct
	var countryNames []constants.CountryName

	decodeError := json.NewDecoder(resp.Body).Decode(&countryNames)
	if decodeError != nil {
		return "", err
	}
	
	//if no results after the decoder
	if(len(countryNames) < 1){
		return "", errors.New("204, no names to get")
	}

	// Return the name of the first response 
	return countryNames[0].Name.Common, nil
}

// Get the list of alpha codes from the body
func getBorderCountries(country string) ([]string, error) {
	// The url for the request
	URL := constants.COUNTRY_API_NAME_URL_PROD+"/"+country

	// Create a new GET request
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return nil, err
	}

	//logging the request
	log.Println("GET countries that are on the border: ", URL)

	// Send the request using the shared http.Client
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	//Prepare to populate the border countries struct and then decoding it 
	var borderCountries []constants.BorderCountries

	decodeError := json.NewDecoder(resp.Body).Decode(&borderCountries)
	if decodeError != nil {
		return nil, err
	}

	//if no results after the decoder
	if(len(borderCountries) < 1){
		return nil, errors.New("204, no border countries to get")
	}

	return borderCountries[0].Borders, nil
}

// Function that setup the GET request and return response and error
func getUniversitiesWithNameAndMiddle(country string, middle string) (*http.Response, error) {
	// URL
	URL := constants.UNI_API_URL_PROD + "/search?name=" + replaceSpaces(middle)  + "&country=" + country

	// Create a new GET request
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return nil, err
	}

	//Logging the request
	log.Println("GET country with country name and 'middle': " + URL)

	// Send the request using the shared http.Client
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Return the response and any errors
	return resp, err
}

// Takes the response from the UniApi and responsewriter. Returns list of UniversityInfo
func addCountryInfoToUniversities(w http.ResponseWriter, uniStructs []constants.UniStruct) []constants.UniversityInfo {

	// Sort the list of uni struct by country
	sort.Slice(uniStructs, func(i, j int) bool {
		return uniStructs[i].Country < uniStructs[j].Country
	})

	//The final response to the
	var uniInfoResponse []constants.UniversityInfo

	//Saving current country information to avoid multiple requests 
	var currentCountryInfo []constants.CountryInfo
	var currentCountry string

	//Loop over each of the university and add the languages
	for _, uni := range uniStructs {
		// Only do a new GET request if the University is in a different country. Only one get request
		if uni.Country != currentCountry {
			// DO API REQUEST and set the new country info struct
			countryResponse, countryErr := getCountriesFromCountryName(uni.Country)
			if countryErr != nil {
				log.Println("Error getting country by name method: ", countryErr.Error())
				http.Error(w, "Invalid request for " + uni.Country, http.StatusBadRequest)
				return nil
			}

			// Decode struct
			err := json.NewDecoder(countryResponse.Body).Decode(&currentCountryInfo)
			if err != nil {
				log.Println("Error during decoding of country: "+err.Error())
				http.Error(w, "Error during decoding", http.StatusInternalServerError)
				return nil
			}

			//Successfully decoded the struct, so we set the new country info
			currentCountry = uni.Country
		}

		//Build the New Struct
		var newUniInfo constants.UniversityInfo = constants.UniversityInfo{UniStruct: uni, CountryInfo: currentCountryInfo[0]}

		//Add them into the response list
		uniInfoResponse = append(uniInfoResponse, newUniInfo)
	}

	return uniInfoResponse

}

//Used to remove empty strings in a list
//Comes from an error on splitting an url. 
// /uni/v1/diag will turn into length of 5, because of this bug 
func removeEmptyStrings(stringList []string) []string {
    result := make([]string, 0, len(stringList)) // create a new slice to hold the non-empty strings
    for _, currentString := range stringList {
        if currentString != "" {
            result = append(result, currentString) // append the non-empty string to the new slice
        }
    }
    return result
}


//Check if the url is valid length.
//Uses responsewriter to return status if not, and returns false. 
//Takes the list of strings and the required length
func isOfValidLength(strList []string, required int,message string , w http.ResponseWriter) bool{
	//Remove empty strings
	strList = removeEmptyStrings(strList); 
	
	// Check if path contains required variables
	if len(strList) != required {
		log.Println("Error on amount of parameters! Should be ", required, ", was", len(strList))
		http.Error(w, "Invalid request path. \n"+ message + "\nCheck docs for use.", http.StatusBadRequest)
		return false
	}

	return true
}


//For endpoints where a GET request is the only allowed method, check if the request is a GET requests. 
func isCorrectRequestMethod(r *http.Request) bool{
	//Check if it is a GET request
	if(r.Method != http.MethodGet){
		log.Println("Wrong request type, tried ", r.Method)
		return false
	}

	return true
}

//Test for a function that takes a url and replace all spaces with %20
func replaceSpaces(url string) string {
	trimmed := strings.TrimSpace(url)
    return strings.ReplaceAll(trimmed, " ", "%20")
}
