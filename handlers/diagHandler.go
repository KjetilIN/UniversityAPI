package handlers

import "net/http"



func DiagHandler(w http.ResponseWriter, r *http.Request){
	//Head Information
	w.Header().Set("content-type", "application/json")



}