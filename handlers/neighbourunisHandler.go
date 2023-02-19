package handlers

import (
	"fmt"
	"net/http"
	"strings"
)


func getBorderCountry(country string) []string{
	return []string{}
}


func NeighbourUniHandler(w http.ResponseWriter, r *http.Request) {
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
	limit := r.URL.Query().Get("limit")

	// Use the variables as needed
	fmt.Fprintf(w, "Country Name: %s\nUniversity Name: %s\nLimit: %s", countryName, universityName, limit)
}