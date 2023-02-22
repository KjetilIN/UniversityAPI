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
	defer resp.Body.Close() // Close the body always

	// Return the response and any errors
	return resp, err
}

// Function that setup the GET request and return response and error
func getFromCountryFromName(country string) (*http.Response, error) {
	// URL
	url := COUNTRY_API_NAME_URL_PROD + "/" + country
	log.Println("Request on : ", url)

	// Create a new GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Send the request using the shared http.Client
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() // Closing the body

	// Return the response and any errors
	return resp, err
}

// Function gets the country name from the api
func getCountryFromAplhaCode(code string) (string, error) {
	// URL
	url := COUNTRY_API_ALPHA_URL_PROD + "/" + code
	log.Println("Request on : ", url)

	// Create a new GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	// Send the request using the shared http.Client
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close() //Closing body at the end

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
	// Create a new GET request
	req, err := http.NewRequest("GET", COUNTRY_API_NAME_URL_PROD+"/"+country, nil)
	if err != nil {
		return nil, err
	}

	//logging the request
	log.Println("Border Counties Req: ", req.URL)

	// Send the request using the shared http.Client
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() // Closing the body

	//Prepera to populate the border countries struct
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
	url := UNI_API_URL_PROD + "/search?name=" + middle + "&country=" + country
	log.Println("Request on : ", url)

	// Create a new GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Send the request using the shared http.Client
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() // Closing the body after request

	// Return the response and any errors
	return resp, err
}

// Takes the response from the UniApi and repsonsewriter. Returns list of UniversityInfo
func addCountryInfoByName(w http.ResponseWriter, resp http.Response) []UniversityInfo {
	// Prepare empty list of structs to populate
	var uniStructs []UniStuct

	// Decode structs
	err := json.NewDecoder(resp.Body).Decode(&uniStructs)
	if err != nil {
		http.Error(w, "Error during decoding. Happened on adding country info", http.StatusBadRequest)
		return nil
	}

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