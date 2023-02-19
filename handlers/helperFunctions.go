package handlers

import (
	"encoding/json"
	"log"
	"net/http"
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


//Function that setup the GET request and return response and error 
func getFromCountryApi(country string) (*http.Response, error) {
	// URL
	url := COUNTRY_API_URL_PROD + "/" + country
	log.Println("Request on : ", url)

	// Create a new GET request
	req, err := http.NewRequest("GET", url , nil)
	if err != nil {
		return nil, err
	}

	// Send the request using the shared http.Client
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Return the response and any errors
	return resp, err
}

//Get the list of alpha codes from the body
func getBorderCountry(country string) ([]string, error){
	// Create a new GET request
	req, err := http.NewRequest("GET", COUNTRY_API_URL_PROD +"/" + country, nil)
	if err != nil {
		return nil, err
	}

	log.Println("Border Counties Req: ", req.URL)

	// Send the request using the shared http.Client
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	//Prepera to populate the border countries struct
	var borderCountrys []BorderCountries

	decodeError := json.NewDecoder(req.Body).Decode(&borderCountrys);
	if decodeError != nil{
		return nil, err
	}

	return borderCountrys[0].Borders, nil
}