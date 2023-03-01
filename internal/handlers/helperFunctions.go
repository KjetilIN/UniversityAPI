package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"time"
	"uniapi/internal/constants"
)

// Create a reusable http.Client that is used by the uni info handler
var httpClient = &http.Client{
	Timeout: time.Second * 10, // Add a timeout to avoid hanging connections
}

// Function that setup the GET request and return error
func getFromUniAPI(searchWord string) (*http.Response, error) {
	// Building the url 
	URL := constants.UNI_API_URL_PROD + "/search?name=" + searchWord
	
	// Create a new GET request
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return nil, err
	}

	//Logging the request
	log.Println("GET from UniApi: " + URL)

	// Send the request using the shared http.Client
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Return the response and any errors
	return resp, err
}

// Function that setup the GET request and return response and error
func getFromCountryFromName(country string) (*http.Response, error) {
	// URL
	URL := constants.COUNTRY_API_NAME_URL_PROD + "/" + country

	// Create a new GET request
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return nil, err
	}

	//Logging the request
	log.Println("GET country name: ", URL)

	// Send the request using the shared http.Client
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Return the response and any errors
	return resp, err
}

// Function gets the country name from the api
func getCountryFromAlphaCode(code string) (string, error) {
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

	// Return the response and any errors
	return countryNames[0].Name.Common, err
}

// Get the list of alpha codes from the body
func getBorderCountry(country string) ([]string, error) {
	// The url for the request
	URL := constants.COUNTRY_API_NAME_URL_PROD+"/"+country

	// Create a new GET request
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return nil, err
	}

	//logging the request
	log.Println("GET countries that are on the border : ", URL)

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

	return borderCountries[0].Borders, nil
}

// Function that setup the GET request and return response and error
func getAllFromUniAPI(country string, middle string) (*http.Response, error) {
	// URL
	URL := constants.UNI_API_URL_PROD + "/search?name=" + middle + "&country=" + country

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
func addCountryInfoByName(w http.ResponseWriter, uniStructs []constants.UniStruct) []constants.UniversityInfo {

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
			countryResponse, countryErr := getFromCountryFromName(uni.Country)
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
//Comes from an error on splitting an url. /uni/v1/diag will turn into length of 5
func removeEmptyStrings(strs []string) []string {
    result := make([]string, 0, len(strs)) // create a new slice to hold the non-empty strings
    for _, str := range strs {
        if str != "" {
            result = append(result, str) // append the non-empty string to the new slice
        }
    }
    return result
}


//Check if the url is valid length.
//Uses responsewriter to return status if not, and returns false. 
//Takes the list of strings and the required length
func isValidLength(strList []string, required int, w http.ResponseWriter) bool{

	strList = removeEmptyStrings(strList); //Remove empty strings
	
	// Check if path contains required variables
	if len(strList) != required {
		log.Println("Error on amount of parameters! Should be ", required, ", was", len(strList))
		http.Error(w, "Invalid request path. Either too long or short. Check docs for use.", http.StatusBadRequest)
		return false
	}

	return true
}