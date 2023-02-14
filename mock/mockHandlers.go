package mock

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// Function that is used for parsing a file 
func ParseFile(filename string) []byte {
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("File error: %v", err)
		os.Exit(1)
	}
	return file
}

//The mock handler for country 
func CountryMockHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
		case http.MethodGet:	
			log.Println("Stub handler get method: ")
			w.Header().Set("content-type", "application/json")
			output := ParseFile("./res/norway.json")
			fmt.Fprint(w, string(output))
			break

		default:
			http.Error(w, "Method not suported", http.StatusNotImplemented)
	}

}

// The uni mock handler 
func UniMockHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
		case http.MethodGet:
			log.Println("Stub handler get method: ")
			w.Header().Set("content-type", "application/json")
			output := ParseFile("./res/uni.json")
			fmt.Fprint(w, string(output))
			break

		default:
			http.Error(w, "Method not suported", http.StatusNotImplemented)
			break
	}

}
