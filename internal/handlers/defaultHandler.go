package handlers

import (
	"encoding/json"
	"net/http"
	"uniapi/internal/constants"
)

func DefaultHandler(w http.ResponseWriter, r *http.Request){
	//Header
	w.Header().Set("content-type", "application/json")

	message := constants.ProjectInfo{Author: "10008", Version: "v1"}

	//Encoding json
	encoder := json.NewEncoder(w);
	err:= encoder.Encode(message)

	//Handle error
	if err != nil{
		http.Error(w, "Error on output!", http.StatusInternalServerError);
	}
}
