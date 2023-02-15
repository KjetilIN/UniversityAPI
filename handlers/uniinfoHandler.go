package handlers

import (
	"fmt"
	"net/http"
)


func UniInfoHandler(w http.ResponseWriter, r *http.Request){
	//Setting header content
	w.Header().Set("content-type", "text")
	// Get all parameteres from the url
	keys:= r.URL.Query()

	//Search keyword
	search:=keys.Get("search")
	
	fmt.Fprint(w,search)


}