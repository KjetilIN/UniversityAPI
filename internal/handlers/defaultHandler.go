package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"uniapi/internal/constants"
)

func DefaultHandler(w http.ResponseWriter, r *http.Request){
	//Header
	w.Header().Set("content-type", "application/json")

	//The default return message
	message := constants.ProjectInfo{Author: "10008", Version: "v1"}

	//Encoding json to response writer
	encoder := json.NewEncoder(w);
	err:= encoder.Encode(message)
	if err != nil{
		log.Println("Error on encoding default handler to response writer: " + err.Error())
		http.Error(w, "Error during encoding", http.StatusInternalServerError);
		return
	}
}
