package mock

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// Function that is used for parsing a file 
func parseFile(filename string) []byte {
	file, err := os.ReadFile(filename)
	if err != nil {
		// Handles errors like this, because the mock response is not going to change,
		log.Printf("File error: " + err.Error())
		os.Exit(1)
	}
	return file
}

//The mock handler for country 
func CountryMockHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet:	
			log.Println("GET request to the Country mock endpoint")
			w.Header().Set("content-type", "application/json")
			output := parseFile("./internal/res/norway.json")
			fmt.Fprint(w, string(output))
			break

		default:
			http.Error(w, "Method not supported", http.StatusNotImplemented)
	}

}
// The uni mock handler 
func UniMockHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet:
			log.Println("GET request to the Uni mock endpoint")
			w.Header().Set("content-type", "application/json")
			output := parseFile("./internal/res/uni.json")
			fmt.Fprint(w, string(output))
			break

		default:
			http.Error(w, "Method not supported", http.StatusNotImplemented)
			break
	}

}
