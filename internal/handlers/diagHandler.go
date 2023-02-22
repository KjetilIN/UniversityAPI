package handlers

import (
	"encoding/json"
	"net/http"
	"uniapi/internal/serverstats"
)


func getStatusCode(link string) string{
	resp, _ := http.Get(link)
	return resp.Status
}

func DiagHandler(w http.ResponseWriter, r *http.Request){
	//Head Information
	w.Header().Set("content-type", "application/json")

	info := StatusInfo{
		UniApi: getStatusCode(UNI_API_URL_PROD),
		CountryApi: getStatusCode(COUNTRY_API_URL_PROD),
		Version: VERSION,
		Uptime: int(serverstats.Uptime().Seconds()),
	}
	
	//Encoding json
	encoder := json.NewEncoder(w);
	err:= encoder.Encode(info)

	//Handle error
	if err != nil{
		http.Error(w, "Error on output!", http.StatusInternalServerError);
	}


}