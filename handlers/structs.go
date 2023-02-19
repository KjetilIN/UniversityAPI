package handlers

//Default handlets return struct
type ProjectInfo struct{
	Author string `json:"author"`
	Version string `json:"version"`
}

// Diag handler return struct 
type StatusInfo struct{
	UniApi string `json:"universitiesapi"`
	CountryApi string `json:"countriesapi"`
	Version string `json:"version"`
	Uptime int `json:"uptime"`
}	

//Stuct with the 


// Struct used by encoder to get the information from the university api
type UniStuct struct{
	Name string `json:"name"`
	Country string `json:"country"`
	Isocode string `json:"alpha_two_code"`
	Webpages []string `json:"web_pages"`
}

// University information structs used for the endpoints 
type UniverityInfo struct {
	UniStuct
	Languages []string `json:"languages"`
	MapLink string `json:"map"`
}