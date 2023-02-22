package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"
	"time"
)

// Create a reusable http.Client that is used by the uniinfo handler
var httpClient = &http.Client{
	Timeout: time.Second * 10, // Add a timeout to avoid hanging connections
}

// Function that setup the GET request and return error
func getFromUniAPI(searchWord string) (*http.Response, error) {
	// Building the url 
	URL := UNI_API_URL_PROD + "/search?name=" + searchWord
	
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
	URL := COUNTRY_API_NAME_URL_PROD + "/" + country

	// Create a new GET request
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return nil, err
	}

	//Loggin the request
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
func getCountryFromAplhaCode(code string) (string, error) {
	// URL
	URL := COUNTRY_API_ALPHA_URL_PROD + "/" + code

	// Create a new GET request
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return "", err
	}

	//Logging the request
	log.Println("GET country by Aplha code: ", URL)

	// Send the request using the shared http.Client
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}

	//Prepera to populate the border countries struct
	var countryNames []CountryName

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
	URL := COUNTRY_API_NAME_URL_PROD+"/"+country

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

	//Prepera to populate the border countries struct and then decoding it 
	var borderCountrys []BorderCountries

	decodeError := json.NewDecoder(resp.Body).Decode(&borderCountrys)
	if decodeError != nil {
		return nil, err
	}

	return borderCountrys[0].Borders, nil
}

// Function that setup the GET request and return response and error
func getAllFromUniAPI(country string, middle string) (*http.Response, error) {
	// URL
	URL := UNI_API_URL_PROD + "/search?name=" + middle + "&country=" + country

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

// Takes the response from the UniApi and repsonsewriter. Returns list of UniversityInfo
func addCountryInfoByName(w http.ResponseWriter, uniStructs []UniStuct) []UniversityInfo {

	// Sort the list of unistruct by country
	sort.Slice(uniStructs, func(i, j int) bool {
		return uniStructs[i].Country < uniStructs[j].Country
	})

	//The final response to the
	var uniInfoResponse []UniversityInfo

	//Loop over each of the university and add the langauges
	var currentCountryInfo []CountryInfo
	var currentCountry string

	for _, uni := range uniStructs {
		// Only do a new GET request if the University is in a diffrent contry. Only one get request
		if uni.Country != currentCountry {
			// DO API REQUEST and set the new countryinfo stuct
			countryResponse, countryErr := getFromCountryFromName(uni.Country)
			if countryErr != nil {
				http.Error(w, "ContryRepsonse error!", http.StatusBadRequest)
				return nil
			}

			// Decode struct
			err := json.NewDecoder(countryResponse.Body).Decode(&currentCountryInfo)
			if err != nil {
				http.Error(w, "Error during decoding of country: "+err.Error(), http.StatusBadRequest)
				return nil
			}

			//Sucessfully decoded the struct, so we set the new country info
			currentCountry = uni.Country
		}

		//Build the New Struct
		var newUniInfo UniversityInfo = UniversityInfo{UniStuct: uni, CountryInfo: currentCountryInfo[0]}

		//Add them into the response list
		uniInfoResponse = append(uniInfoResponse, newUniInfo)
	}

	return uniInfoResponse

}