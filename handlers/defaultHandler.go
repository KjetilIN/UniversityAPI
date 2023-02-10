package handlers

import (
	"encoding/json"
	"net/http"
)

func DefaultHandler(w http.ResponseWriter, r *http.Request){
	//Header
	w.Header().Set("content-type", "application/json")

	message := ProjectInfo{Author: "Kjetil Indrehus", Version: "v1"}

	//Encoding json
	encoder := json.NewEncoder(w);
	err:= encoder.Encode(message)

	//Handle error
	if err != nil{
		http.Error(w, "Error on output!", http.StatusInternalServerError);
	}
}
